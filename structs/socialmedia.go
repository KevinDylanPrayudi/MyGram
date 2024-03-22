package structs

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	gorm.Model

	Name             string `form:"name" json:"name" valid:"required,type(string)"`
	Social_media_url string `form:"social_media_url" json:"social_media_url" valid:"required,type(string)"`
	UserID           uint
}

type AddSocialMedia struct {
	Name             string `json:"name" example:"testing"`
	Social_media_url string `json:"social_media_url" example:"testing"`
}

type AddSocialMediaResult struct {
	Id               uint      `json:"id" example:"1" type:"int64"`
	Name             string    `json:"name" example:"testing"`
	Social_media_url string    `json:"social_media_url" example:"testing"`
	UserID           uint      `json:"user_id" example:"1" type:"int64"`
	CreatedAt        time.Time `json:"created_at" example:"date"`
}

type GetSocialMediaResult struct {
	Id               uint      `json:"id" example:"1" type:"int64"`
	Name             string    `json:"name" example:"testing"`
	Social_media_url string    `json:"social_media_url" example:"testing"`
	UserID           uint      `json:"user_id" example:"1" type:"int64"`
	CreatedAt        time.Time `json:"created_at" example:"date"`
	UpdatedAt        time.Time `json:"updated_at" example:"date"`
	User             struct {
		Id       uint   `json:"id" example:"1" format:"int64"`
		Email    string `json:"email" example:"testing@gmail.com"`
		Username string `json:"username" example:"testing"`
	}
}

type UpdateSocialMedia struct {
	Name             string `json:"name" example:"update testing"`
	Social_media_url string `json:"social_media_url" example:"update testing"`
}

type UpdateSocialMediaResult struct {
	Id               uint      `json:"id" example:"1" type:"int64"`
	Name             string    `json:"name" example:"update testing"`
	Social_media_url string    `json:"social_media_url" example:"update testing"`
	UserID           uint      `json:"user_id" example:"1" type:"int64"`
	UpdatedAt        time.Time `json:"updated_at" example:"date"`
}

type DeleteSocialMediaResult struct {
	Message string `json:"message" example:"Your social media has been successfully deleted"`
}

func (sm *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(sm)
	if err != nil {
		return err
	}
	return
}

func (p *SocialMedia) BeforeUpdate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(p)
	if err != nil {
		return err
	}
	return
}
