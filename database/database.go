package database

import (
	"github.com/joseluissanchez77/GoTesisApiArduino/database/migrations"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var db *gorm.DB

func StartDB(){
	// str := "host=localhost port=5432 user=postgres dbname=DB_ARDUINOTESIS sslmode=disable password=root"
	str := "host=108.175.14.214 port=5432 user=postgres dbname=DB_ARDUINOTESIS sslmode=disable password=j0s3_20221"

	database, err := gorm.Open(postgres.Open(str), &gorm.Config{})

	if err != nil{
		log.Fatal("errorrrr : ",err)
	}

	db = database

	config, _ := db.DB()

	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)

	migrations.RunMigrations(db)
}


func GetDatabase() *gorm.DB{
	return db
}