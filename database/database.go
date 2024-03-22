package database

import (
	"final-assignment/structs"
	"log"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
	user = os.Getenv("PGUSER")
	password = os.Getenv("PGPASSWORD")
	dbname = os.Getenv("PGDATABASE")
	port = os.Getenv("PGPORT")
	host = os.Getenv("PGHOST")
)
func PostGresDB() {
	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port),
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
