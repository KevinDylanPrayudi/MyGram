package structs

import (
	"errors"
	"final-assignment/helpers"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Username    string `gorm:"uniqueIndex" form:"username" json:"username" valid:"required,isUnique(username)~username is already in use,type(string)"`
	Email       string `gorm:"uniqueIndex" form:"email" json:"email" valid:"required,isUnique(email)~email is already in use,type(string),email"`
	Password    string `form:"password" json:"password" valid:"required,minstringlength(6),type(string)" example:"testing@gmail.com"`
	Age         int    `form:"age" json:"age" valid:"required,type(int),minAge~U hafta older than 8 yo for signing up"`
	Photos      []Photo
	Comments    []Comment
	SocialMedia []SocialMedia
}

type AddUser struct {
	Age      int    `json:"age" example:"9" format:"int64"`
	Email    string `json:"email" example:"testing@gmail.com"`
	Username string `json:"username" example:"testing"`
	Password string `json:"password" example:"123456"`
}

type AddUserResult struct {
	Age      int    `json:"age" example:"9" format:"int64"`
	Email    string `json:"email" example:"testing@gmail.com"`
	Id       uint   `json:"id" example:"1" format:"int64"`
	Username string `json:"username" example:"testing"`
}

type LoginUser struct {
	Email    string `json:"email" example:"testing@gmail.com"`
	Password string `json:"password" example:"123456"`
}

type LoginUserResult struct {
	Token string `json:"token" example:"result of generated token"`
}

type UpdateUser struct {
	Email    string `json:"email" example:"testing1@gmail.com"`
	Username string `json:"username" example:"testing1"`
}

type UpdateUserResult struct {
	AddUserResult
	UpdatedAt time.Time `json:"updated_at" example:"date"`
}

type DeleteUserResult struct {
	Message string `json:"message" example:"Your account has been successfully deleted"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(u)
	if err != nil {
		return err
	}
	u.Password = helpers.HashAndSalt([]byte(u.Password))
	return
}

func (u *User) AfterFind(tx *gorm.DB) error {
	plainPwd := tx.Statement.Context.Value("plainPwd").(string)
	if !helpers.ComparePasswords(u.Password, []byte(plainPwd)) {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	userId := tx.Statement.Context.Value("userId").(uint)
	if govalidator.IsNull(u.Username) {
		return errors.New("Your username hafta be filled")
	}
	if govalidator.IsNull(u.Email) {
		return errors.New("Your email hafta be filled")
	}
	if tx.Session(&gorm.Session{SkipHooks: true}).First(&User{}, userId).RowsAffected == 0 {
		return errors.New("The User's ID isn't found")
	}
	return
}
