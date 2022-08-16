package migrations

import (
	// "github.com/joseluissanchez77/GoTesisApiArduino/database/seeders"
	"github.com/joseluissanchez77/GoTesisApiArduino/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB){
	db.AutoMigrate(models.SensorData{})
	db.AutoMigrate(models.ParameterData{})
	db.AutoMigrate(models.CatalogData{})
	
	// seeders.DBSeed(db)
}