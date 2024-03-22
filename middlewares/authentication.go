package middlewares

import (
	"final-assignment/helpers"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		getToken := strings.Split(ctx.GetHeader("Authorization"), " ")
		if len(getToken) < 2 {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "please set Header Authorization",
			})
			return
		}
		token := getToken[1]
		userId, err := helpers.ParseToken(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "Your Token Isn't Valid",
			})
		}
		ctx.Set("id", userId)
		ctx.Next()
	}
}
