package controllers

import (
	"errors"
	"final-assignment/database"
	"final-assignment/helpers"
	"final-assignment/structs"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// CreatePhoto godoc
//
//	@Summary		Create Photo
//	@Description	Create by json photo
//	@Tags			photo
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			title	formData	string	true	"title of photo"
//	@Param			caption	formData	string	false	"caption of photo"
//	@Param			file	formData	file	true	"account image"
//	@Success		201		{object}	structs.AddPhotoResult
//	@Security		ApiKeyAuth
//	@Router			/photo [post]
func CreatePhoto(ctx *gin.Context) {
	db := database.GetDB()
	var photo structs.Photo
	ctx.ShouldBind(&photo)
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	filename := file.Filename
	extension := filepath.Ext(filename)
	extension = strings.TrimPrefix(extension, ".")
	if extension != "jpg" && extension != "png" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": errors.New("You must upload a JPEG or PNG file"),
		})
		return
	}
	photoUrl, err := helpers.Upload(file)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	photo.Photo_url = photoUrl
	photo.UserID = ctx.MustGet("id").(uint)
	if err := db.Create(&photo).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	result := structs.AddPhotoResult{
		Id:        photo.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		Photo_url: photo.Photo_url,
		UserID:    photo.UserID,
		CreatedAt: photo.CreatedAt,
	}
	ctx.JSON(http.StatusCreated, result)
	return
}

// GetPhoto godoc
//
//	@Summary		Show List of Photo
//	@Description	Show List of Photo by json photo
//	@Tags			photo
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	structs.GetPhotoResult
//	@Security		ApiKeyAuth
//	@Router			/photo [get]
func GetPhoto(ctx *gin.Context) {
	db := database.GetDB()
	type Photo struct {
		structs.Photo
		User structs.User
	}

	var photo []Photo
	db.Session(&gorm.Session{SkipHooks: true}).Preload("User").Find(&photo)

	var result []structs.GetPhotoResult

	for _, v := range photo {
		dest := structs.GetPhotoResult{
			Id:        v.ID,
			Title:     v.Title,
			Caption:   v.Caption,
			Photo_url: v.Photo_url,
			UserID:    v.UserID,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		dest.User.Email = v.User.Email
		dest.User.Username = v.User.Username
		result = append(result, dest)
	}
	ctx.JSON(http.StatusOK, result)
}

// UpdatePhoto godoc
//
//	@Summary		Update an photo
//	@Description	Update by json photo
//	@Tags			photo
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			photoId	path		int					true	"Photo ID"
//	@Param			title	formData	string	true	"title of photo"
//	@Param			caption	formData	string	true	"caption of photo"
//	@Param			file	formData	file	true	"account image"
//	@Success		200		{object}	structs.UpdatePhotoResult
//	@Security		ApiKeyAuth
//	@Router			/photo/{photoId} [put]
func UpdatePhoto(ctx *gin.Context) {
	photoId := ctx.Param("photoId")
	db := database.GetDB()
	var photo structs.Photo
	ctx.ShouldBind(&photo)
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	filename := file.Filename
	extension := filepath.Ext(filename)
	extension = strings.TrimPrefix(extension, ".")
	if extension != "jpg" && extension != "png" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": errors.New("You must upload a JPEG or PNG file"),
		})
		return
	}
	photoUrl, err := helpers.Upload(file)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	photo.Photo_url = photoUrl
	err = db.Model(&photo).Clauses(clause.Returning{Columns: []clause.Column{{Name: "id"}, {Name: "title"}, {Name: "caption"}, {Name: "photo_url"}, {Name: "user_id"}, {Name: "updated_at"}}}).Where("id = ?", photoId).Updates(photo).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	result := structs.UpdatePhotoResult{
		Id:        photo.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		Photo_url: photo.Photo_url,
		UserID:    photo.UserID,
		UpdatedAt: photo.UpdatedAt,
	}
	ctx.JSON(http.StatusOK, result)
}

// DeletePhoto godoc
//
//	@Summary		Delete an photo
//	@Description	Delete by photo ID
//	@Tags			photo
//	@Accept			json
//	@Produce		json
//	@Param			photoId	path		int	true	"photo ID"	Format(int64)
//	@Success		200		{object}	structs.DeletePhotoResult
//	@Security		ApiKeyAuth
//	@Router			/photo/{photoId} [delete]
func DeletePhoto(ctx *gin.Context) {
	photoId, _ := strconv.Atoi(ctx.Param("photoId"))
	db := database.GetDB()
	var photo structs.Photo
	photo.ID = uint(photoId)
	err := db.Select(clause.Associations).Delete(&photo).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your photo has been successfully deleted",
	})
}
