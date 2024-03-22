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

// CreateSocialMedia godoc
//
//	@Summary		Create Social Media
//	@Description	Create Social Mediaby json Social Media
//	@Tags			socialmedia
//	@Accept			json
//	@Produce		json
//	@Param			socialmedia	body		structs.AddSocialMedia	true	"Show List of SocialMedia"
//	@Success		201			{object}	structs.AddSocialMediaResult
//	@Security		ApiKeyAuth
//	@Router			/socialmedias [post]
func CreateSocialMedia(ctx *gin.Context) {
	db := database.GetDB()
	var socialMedia structs.SocialMedia
	ctx.ShouldBind(&socialMedia)
	socialMedia.UserID = ctx.MustGet("id").(uint)
	if err := db.Create(&socialMedia).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	result := structs.AddSocialMediaResult{
		Id:               socialMedia.ID,
		Name:             socialMedia.Name,
		Social_media_url: socialMedia.Social_media_url,
		UserID:           socialMedia.UserID,
		CreatedAt:        socialMedia.CreatedAt,
	}
	ctx.JSON(http.StatusCreated, result)
	return
}

// SocialMedia godoc
//
//	@Summary		Show List of Social Media
//	@Description	Show List of Social Media by json Social Media
//	@Tags			socialmedia
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	structs.GetSocialMediaResult
//	@Security		ApiKeyAuth
//	@Router			/socialmedias [get]
func GetSocialMedia(ctx *gin.Context) {
	db := database.GetDB()
	type SocialMedia struct {
		structs.SocialMedia
		User structs.User
	}

	var socialMedia []SocialMedia
	db.Session(&gorm.Session{SkipHooks: true}).Preload("User").Find(&socialMedia)

	var result []structs.GetSocialMediaResult

	for _, v := range socialMedia {
		dest := structs.GetSocialMediaResult{
			Id:               v.ID,
			Name:             v.Name,
			Social_media_url: v.Social_media_url,
			UserID:           v.UserID,
			CreatedAt:        v.CreatedAt,
			UpdatedAt:        v.UpdatedAt,
		}
		dest.User.Id = v.User.ID
		dest.User.Email = v.User.Email
		dest.User.Username = v.User.Username
		result = append(result, dest)
	}
	ctx.JSON(http.StatusOK, result)
}

// UpdateSocialMedia godoc
//
//	@Summary		Update an Social Media
//	@Description	Update by json Social Media
//	@Tags			socialmedia
//	@Accept			json
//	@Produce		json
//	@Param			socialmediaId	path		int							true	"Social Media ID"
//	@Param			photo			body		structs.UpdateSocialMedia	true	"Update Social Media"
//	@Success		200				{object}	structs.UpdateSocialMediaResult
//	@Security		ApiKeyAuth
//	@Router			/socialmedias/{socialmediaId} [put]
func UpdateSocialMedia(ctx *gin.Context) {
	socialMediaId := ctx.Param("socialMediaId")
	db := database.GetDB()
	var socialMedia structs.SocialMedia
	ctx.ShouldBind(&socialMedia)
	err := db.Model(&socialMedia).Clauses(clause.Returning{Columns: []clause.Column{{Name: "id"}, {Name: "name"}, {Name: "social_media_url"}, {Name: "user_id"}, {Name: "updated_at"}}}).Where("id = ?", socialMediaId).Updates(socialMedia).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	result := structs.UpdateSocialMediaResult{
		Id:               socialMedia.ID,
		Name:             socialMedia.Name,
		Social_media_url: socialMedia.Social_media_url,
		UserID:           socialMedia.UserID,
		UpdatedAt:        socialMedia.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, result)
}

// DeleteSocialMedia godoc
//
//	@Summary		Delete an socialmedia
//	@Description	Delete by socialmedia ID
//	@Tags			socialmedia
//	@Accept			json
//	@Produce		json
//	@Param			socialmediaId	path		int	true	"socialmedia ID"	Format(int64)
//	@Success		200				{object}	structs.DeleteSocialMediaResult
//	@Security		ApiKeyAuth
//	@Router			/socialmedias/{socialmediaId} [delete]
func DeleteSocialMedia(ctx *gin.Context) {
	socialMediaId, _ := strconv.Atoi(ctx.Param("socialMediaId"))
	db := database.GetDB()
	var socialMedia structs.SocialMedia
	socialMedia.ID = uint(socialMediaId)
	err := db.Delete(&socialMedia).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your social media has been successfully deleted",
	})
}
