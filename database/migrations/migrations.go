package migrations

import (
	"github.com/joseluissanchez77/GoTesisApiArduino/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB){
	db.AutoMigrate(models.SensorData{})
}