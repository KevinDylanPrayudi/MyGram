package controllers

import (
	"final-assignment/database"
	"final-assignment/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// CreateComment godoc
//
//	@Summary		Create comment
//	@Description	Create by json comment
//	@Tags			comment
//	@Accept			json
//	@Produce		json
//	@Param			comment	body		structs.AddComment	true	"Show List of Comment"
//	@Success		201		{object}	structs.AddCommentResult
//	@Security		ApiKeyAuth
//	@Router			/comment [post]
func CreateComment(ctx *gin.Context) {
	db := database.GetDB()
	var comment structs.Comment
	ctx.ShouldBind(&comment)
	comment.UserID = ctx.MustGet("id").(uint)
	if err := db.Create(&comment).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	result := structs.AddCommentResult{
		Id:        comment.ID,
		Message:   comment.Message,
		PhotoID:   comment.PhotoID,
		UserID:    comment.UserID,
		CreatedAt: comment.CreatedAt,
	}
	ctx.JSON(http.StatusCreated, result)
	return
}

// GetComment godoc
//
//	@Summary		Show List of Comment
//	@Description	Show List of Comment by json Comment
//	@Tags			comment
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	structs.GetCommentResult
//	@Security		ApiKeyAuth
//	@Router			/comment [get]
func GetComment(ctx *gin.Context) {
	db := database.GetDB()
	type Comment struct {
		structs.Comment
		User  structs.User
		Photo structs.Photo
	}

	var comment []Comment
	db.Session(&gorm.Session{SkipHooks: true}).Preload("User").Preload("Photo").Find(&comment)

	var result []structs.GetCommentResult

	for _, v := range comment {
		dest := structs.GetCommentResult{
			Id:        v.ID,
			Message:   v.Message,
			PhotoID:   v.PhotoID,
			UserID:    v.UserID,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		dest.User.Id = v.User.ID
		dest.User.Email = v.User.Email
		dest.User.Username = v.User.Username
		result = append(result, dest)
	}
	ctx.JSON(http.StatusOK, result)
}

// UpdateComment godoc
//
//	@Summary		Update an comment
//	@Description	Update by json comment
//	@Tags			comment
//	@Accept			json
//	@Produce		json
//	@Param			commentId	path		int						true	"Comment ID"
//	@Param			comment		body		structs.UpdateComment	true	"Update Comment"
//	@Success		200			{object}	structs.UpdateCommentResult
//	@Security		ApiKeyAuth
//	@Router			/comment/{commentId} [put]
func UpdateComment(ctx *gin.Context) {
	commentId := ctx.Param("commentId")
	db := database.GetDB()
	var comment structs.Comment
	ctx.ShouldBind(&comment)
	err := db.Model(&comment).Clauses(clause.Returning{Columns: []clause.Column{{Name: "id"}, {Name: "message"}, {Name: "photo_id"}, {Name: "user_id"}, {Name: "updated_at"}}}).Select("message").Where("id = ?", commentId).Updates(comment).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	result := structs.UpdateCommentResult{
		Id:        comment.ID,
		Message:   comment.Message,
		PhotoID:   comment.PhotoID,
		UserID:    comment.UserID,
		UpdatedAt: comment.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, result)
}

// DeleteComment godoc
//
//	@Summary		Delete an comment
//	@Description	Delete by comment ID
//	@Tags			comment
//	@Accept			json
//	@Produce		json
//	@Param			commentId	path		int	true	"Comment ID"	Format(int64)
//	@Success		200			{object}	structs.DeleteCommentResult
//	@Security		ApiKeyAuth
//	@Router			/comment/{commentId} [delete]
func DeleteComment(ctx *gin.Context) {
	commentId, _ := strconv.Atoi(ctx.Param("commentId"))
	db := database.GetDB()
	var comment structs.Comment
	comment.ID = uint(commentId)
	err := db.Delete(&comment).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your comment has been successfully deleted",
	})
}
