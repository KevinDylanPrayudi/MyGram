package customValidations

import (
	"final-assignment/database"
	"final-assignment/structs"
	"log"
	"regexp"
	"strconv"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

func init() {
	govalidator.ParamTagMap["isUnique"] = govalidator.ParamValidator(isUnique)
	govalidator.ParamTagRegexMap["isUnique"] = regexp.MustCompile("^isUnique\\((\\w+)\\)$")
	govalidator.TagMap["minAge"] = govalidator.Validator(minAge)
}

func isUnique(v string, params ...string) bool {
	var user structs.User
	db := database.GetDB()
	if params[0] == "username" {
		return db.Session(&gorm.Session{SkipHooks: true}).First(&user, "username = ?", v).RowsAffected == 0
	} else if params[0] == "email" {
		return db.Session(&gorm.Session{SkipHooks: true}).First(&user, "email = ?", v).RowsAffected == 0
	}
	return false
}

func minAge(v string) bool {
	vInt, err := strconv.Atoi(v)
	if err != nil {
		log.Panic(err)
	}
	return vInt > 8
}
