package routers

import (
	"final-assignment/controllers"
	_ "final-assignment/customValidations"
	_ "final-assignment/docs"
	"final-assignment/middlewares"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() {
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.POST("/users/register", controllers.Register)
	router.POST("/users/login", controllers.Login)
	authorized := router.Group("/")

	authorized.Use(middlewares.Authentication())
	{
		authorized.PUT("/users/:userId", middlewares.Authorization("user"), controllers.UpdateUser)
		authorized.DELETE("/users/:userId", middlewares.Authorization("user"), controllers.DeleteUser)
		authorized.GET("/photo", controllers.GetPhoto)
		authorized.POST("/photo", controllers.CreatePhoto)
		authorized.PUT("/photo/:photoId", middlewares.Authorization("photo"), controllers.UpdatePhoto)
		authorized.DELETE("/photo/:photoId", middlewares.Authorization("photo"), controllers.DeletePhoto)
		authorized.GET("/comment", controllers.GetComment)
		authorized.POST("/comment", controllers.CreateComment)
		authorized.PUT("/comment/:commentId", middlewares.Authorization("comment"), controllers.UpdateComment)
		authorized.DELETE("/comment/:commentId", middlewares.Authorization("comment"), controllers.DeleteComment)
		authorized.GET("/socialmedias", controllers.GetSocialMedia)
		authorized.POST("/socialmedias", controllers.CreateSocialMedia)
		authorized.PUT("/socialmedias/:socialMediaId", middlewares.Authorization("socialmedia"), controllers.UpdateSocialMedia)
		authorized.DELETE("/socialmedias/:socialMediaId", middlewares.Authorization("socialmedia"), controllers.DeleteSocialMedia)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")
}
