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
	// str := "host=108.175.14.214 port=5432 user=postgres dbname=DB_ARDUINOTESIS sslmode=disable password=j0s3_20221"
	str := "host=ec2-100-26-39-41.compute-1.amazonaws.com port=5432 user=dbpdxyvcdpudic dbname=d3959ibfssf0v3 password=232ce7d7bbcada6854014e9de1b1dd384cc33e4a7d22fe54c06056d74d310a36"

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