package fakers

import (
	"github.com/joseluissanchez77/GoTesisApiArduino/models"
	"time"

	"gorm.io/gorm"
)

func ParameterDataFaker(db *gorm.DB) *models.ParameterData {
	return &models.ParameterData{
		Sensor: "Motor",
		WeekdayEs:    "Domingo",
		WeekdayEn: "Domingo",
		Weekday: "Do",
		HourInitial : "17:00",
		HourEnd : "17:10",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     gorm.DeletedAt{},
	}
}
