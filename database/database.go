package database

import (
	"final-assignment/structs"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func PostGresDB() {
	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=postgres password=123 dbname=assignment-final port=5432 sslmode=disable",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		log.Panic("Postgress is halted caused by", err)
	}
	db.AutoMigrate(&structs.User{}, &structs.Photo{}, &structs.Comment{}, &structs.SocialMedia{})
	db.Exec("TRUNCATE users, photos, comments, social_media;")
	db.Exec("SELECT setval(pg_get_serial_sequence('users', 'id'), COALESCE((SELECT MAX(id) + 1 FROM users), 1), false);")
	db.Exec("SELECT setval(pg_get_serial_sequence('photos', 'id'), COALESCE((SELECT MAX(id) + 1 FROM photos), 1), false);")
	db.Exec("SELECT setval(pg_get_serial_sequence('comments', 'id'), COALESCE((SELECT MAX(id) + 1 FROM comments), 1), false);")
	db.Exec("SELECT setval(pg_get_serial_sequence('social_media', 'id'), COALESCE((SELECT MAX(id) + 1 FROM social_media), 1), false);")
}

func GetDB() *gorm.DB {
	return db
}
