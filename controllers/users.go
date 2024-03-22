package controllers

import (
	"context"
	"final-assignment/database"
	"final-assignment/helpers"
	"final-assignment/structs"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm/clause"
)

// AddUsers godoc
//
//	@Summary		Add an users
//	@Description	add by json users
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			photo	body		structs.AddUser	true	"Add user"
//	@Success		201		{object}	structs.AddUserResult
//	@Router			/users/register [post]
func Register(ctx *gin.Context) {
	db := database.GetDB()
	var user structs.User
	ctx.ShouldBind(&user)
	if err := db.Create(&user).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	result := structs.AddUserResult{
		Age:      user.Age,
		Email:    user.Email,
		Id:       user.ID,
		Username: user.Username,
	}
	ctx.JSON(http.StatusCreated, result)
	return
}

// LoginUsers godoc
//
//	@Summary		Login an users
//	@Description	Login by json users
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			login	body		structs.LoginUser	true	"Login user"
//	@Success		200		{object}	structs.LoginUserResult
//	@Router			/users/login [post]
func Login(ctx *gin.Context) {
	db := database.GetDB()
	var user structs.User
	ctx.ShouldBind(&user)
	pwdCtx := context.WithValue(context.Background(), "plainPwd", user.Password)
	err := db.WithContext(pwdCtx).First(&user, "email = ?", user.Email).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}
	result, err := helpers.GenerateJWT(user.Email, user.ID)
	if err != nil {
		log.Println(err)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"token": result,
	})
}

// UpdateUsers godoc
//
//	@Summary		Update an users
//	@Description	Update by json users
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			userId	path		int					true	"User ID"
//	@Param			user	body		structs.UpdateUser	true	"Update user"
//	@Success		200		{object}	structs.UpdateUserResult
//	@Security		ApiKeyAuth
//	@Router			/users/{userId} [put]
func UpdateUser(ctx *gin.Context) {
	userId := ctx.Param("userId")
	db := database.GetDB()
	var user structs.User
	ctx.ShouldBind(&user)
	userIdCtx := context.WithValue(context.Background(), "userId", ctx.MustGet("id"))
	err := db.WithContext(userIdCtx).Clauses(clause.Returning{Columns: []clause.Column{{Name: "id"}, {Name: "email"}, {Name: "username"}, {Name: "age"}, {Name: "updated_at"}}}).Model(&user).Select("username", "email").Where("id = ?", userId).Updates(user).Error
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			switch pgErr.Code {
			case "23505":
				{
					if err.(*pgconn.PgError).ConstraintName == "idx_users_username" {
						ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
							"message": "The username is already exist",
						})
						return
					} else {
						ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
							"message": "The email is already exist",
						})
						return
					}
				}
			default:
				{
					ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
						"message": err.Error(),
					})
					return
				}
			}
		} else {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		}
	}
	result := structs.UpdateUserResult{
		UpdatedAt: user.UpdatedAt,
	}
	result.AddUserResult.Age = user.Age
	result.AddUserResult.Email = user.Email
	result.AddUserResult.Id = user.ID
	result.AddUserResult.Username = user.Username

	ctx.JSON(http.StatusOK, result)
}

// DeleteUser godoc
//
//	@Summary		Delete an User
//	@Description	Delete by User ID
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			userId	path		int	true	"User ID"	Format(int64)
//	@Success		200		{object}	structs.DeleteUserResult
//	@Security		ApiKeyAuth
//	@Router			/users/{userId} [delete]
func DeleteUser(ctx *gin.Context) {
	userId, _ := strconv.Atoi(ctx.Param("userId"))
	db := database.GetDB()
	var user structs.User
	user.ID = uint(userId)
	err := db.Select(clause.Associations).Delete(&user).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your account has been successfully deleted",
	})
}
