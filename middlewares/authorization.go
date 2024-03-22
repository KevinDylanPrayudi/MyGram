package middlewares

import (
	"final-assignment/database"
	"final-assignment/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Authorization(param string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		contextUserID := ctx.MustGet("id")
		switch param {
		case "user":
			{
				paramUserID, err := strconv.Atoi(ctx.Param("userId"))
				if err != nil {
					ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
						"message": "Your Param isn't valid",
					})
					return
				}
				if uint(paramUserID) != contextUserID {
					ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
						"message": "U aren't authorization",
					})
					return
				}
			}
		case "photo":
			{
				paramPhotoID, err := strconv.Atoi(ctx.Param("photoId"))
				if err != nil {
					ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
						"message": "Your Param isn't valid",
					})
					return
				}
				db := database.GetDB()
				if err = db.First(&structs.Photo{}, "id = ? AND user_id = ?", paramPhotoID, contextUserID).Error; err != nil {
					ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
						"message": "U aren't authorization",
					})
					return
				}
			}
		case "comment":
			{
				paramCommentID, err := strconv.Atoi(ctx.Param("commentId"))
				if err != nil {
					ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
						"message": "Your Param isn't valid",
					})
					return
				}
				db := database.GetDB()
				if err = db.First(&structs.Comment{}, "id = ? AND user_id = ?", paramCommentID, contextUserID).Error; err != nil {
					ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
						"message": "U aren't authorization",
					})
					return
				}
			}
		case "socialmedia":
			{
				paramSocialMediaId, err := strconv.Atoi(ctx.Param("socialMediaId"))
				if err != nil {
					ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
						"message": "Your Param isn't valid",
					})
					return
				}
				db := database.GetDB()
				if err = db.First(&structs.SocialMedia{}, "id = ? AND user_id = ?", paramSocialMediaId, contextUserID).Error; err != nil {
					ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
						"message": "U aren't authorization",
					})
					return
				}
			}
		}
		ctx.Next()
	}
}
