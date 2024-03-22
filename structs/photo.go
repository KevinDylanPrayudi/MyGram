package structs

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model

	Title     string `form:"title" json:"title" valid:"required,type(string)"`
	Caption   string `form:"caption" json:"caption"`
	Photo_url string
	UserID    uint
	Comments  []Comment
}

type AddPhoto struct {
	Title     string `json:"title" example:"title"`
	Caption   string `json:"caption" example:"caption"`
	Photo_url string `json:"photo_url" example:"photo url"`
}

type AddPhotoResult struct {
	Id        uint      `json:"id" example:"1" format:"int64"`
	Title     string    `json:"title" example:"title"`
	Caption   string    `json:"caption" example:"caption"`
	Photo_url string    `json:"photo_url" example:"photo url"`
	UserID    uint      `json:"user_id" example:"1" format:"int64"`
	CreatedAt time.Time `json:"created_at" example:"date"`
}

type GetPhotoResult struct {
	Id        uint      `json:"id" example:"1" format:"int64"`
	Title     string    `json:"title" example:"title"`
	Caption   string    `json:"caption" example:"caption"`
	Photo_url string    `json:"photo_url" example:"photo_url"`
	UserID    uint      `json:"user_id" example:"1" format:"int64"`
	CreatedAt time.Time `json:"created_at" example:"date"`
	UpdatedAt time.Time `json:"updated_at" example:"date"`
	User      struct {
		Email    string `json:"email" example:"testing@gmail.com"`
		Username string `json:"username" example:"testing"`
	}
}

type UpdatePhoto struct {
	Title     string `json:"title" example:"update title"`
	Caption   string `json:"caption" example:"update caption"`
	Photo_url string `json:"photo_url" example:"update photo url"`
}

type UpdatePhotoResult struct {
	Id        uint      `json:"id" example:"1" format:"int64"`
	Title     string    `json:"title" example:"update title"`
	Caption   string    `json:"caption" example:"update caption"`
	Photo_url string    `json:"photo_url" example:"update photo url"`
	UserID    uint      `json:"user_id" example:"1" format:"int64"`
	UpdatedAt time.Time `json:"updated_at" example:"date"`
}

type DeletePhotoResult struct {
	Message string `json:"message" example:"Your photo has been successfully deleted"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(p)
	if err != nil {
		return err
	}
	return
}

func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(p)
	if err != nil {
		return err
	}
	return
}
