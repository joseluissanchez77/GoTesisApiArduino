package models

import (
	"time"

	"gorm.io/gorm"
)
type SensorData struct{
	ID          uint            `json:"id" gorm:"primaryKey"`
	Description string 			`json:"description"`
	Celsius float32 			`json:"celsius"`
	Fahrenheit float32 			`json:"fahrenheit"`
	WaterLevel float32 			`json:"water_level"`
	Ph float32 					`json:"ph"`
	Nutrition int64 			`gorm:"type:integer;default:1" json:"nutrition"`
	CreatedAt time.Time 		`json:"created"`
	UpdatedAt time.Time 		`json:"updated"`
	DeletedAt gorm.DeletedAt	`gorm:"index" json:"deleted"`
	
}