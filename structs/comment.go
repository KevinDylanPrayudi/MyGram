package structs

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model

	UserID  uint
	PhotoID uint   `form:"photo_id" json:"photo_id" valid:"required,type(uint)"`
	Message string `form:"message" json:"message" valid:"required,type(string)"`
}

type AddComment struct {
	PhotoID uint   `json:"photo_id" example:"1" format:"int64"`
	Message string `json:"message" example:"message"`
}

type AddCommentResult struct {
	Id        uint      `json:"id" example:"1" format:"int64"`
	Message   string    `json:"message" example:"message"`
	PhotoID   uint      `json:"photo_id" example:"1" format:"int64"`
	UserID    uint      `json:"user_id" example:"1" format:"int64"`
	CreatedAt time.Time `json:"created_at" example:"date"`
}

type GetCommentResult struct {
	Id        uint      `json:"id" example:"1" format:"int64"`
	Message   string    `json:"message" example:"message"`
	PhotoID   uint      `json:"photo_id" example:"1" format:"int64"`
	UserID    uint      `json:"user_id" example:"1" format:"int64"`
	CreatedAt time.Time `json:"created_at" example:"date"`
	UpdatedAt time.Time `json:"updated_at" example:"date"`
	User      struct {
		Id       uint   `json:"id" example:"1" format:"int64"`
		Email    string `json:"email" example:"testing@gmail.com"`
		Username string `json:"username" example:"testing"`
	}
}

type UpdateComment struct {
	Message string `json:"message" example:"update message"`
}

type UpdateCommentResult struct {
	Id        uint      `json:"id" example:"1" format:"int64"`
	Message   string    `json:"message" example:"update message"`
	PhotoID   uint      `json:"photo_id" example:"1" format:"int64"`
	UserID    uint      `json:"user_id" example:"1" format:"int64"`
	UpdatedAt time.Time `json:"updated_at" example:"date"`
}

type DeleteCommentResult struct {
	Message string `json:"message" example:"Your comment has been successfully deleted"`
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(c)
	if err != nil {
		return err
	}
	return
}

func (c *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	if govalidator.IsNull(c.Message) {
		return errors.New("U hafta input message field")
	}
	return
}
