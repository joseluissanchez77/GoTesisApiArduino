package models

import (
	"time"

	"gorm.io/gorm"
)
type SensorData struct{
	ID          uint            `json:"id" gorm:"primaryKey"`
	Description string 			`json:"description"`
	Celsius float32 			`json:"celsius" gorm:"type:decimal(10,2)"`
	Fahrenheit float32 			`json:"fahrenheit" gorm:"type:decimal(10,2)"`
	WaterLevel float32 			`json:"water_level" gorm:"type:decimal(10,2)"`
	Ph float32 					`json:"ph" gorm:"type:decimal(10,2)"`
	Nutrition string 			`gorm:"type:string;default:apagado" json:"nutrition"`
	WaterPump string 			`json:"water_pump" gorm:"type:string;default:apagado"`,
	DescripcionTurbidez string 				`json:"descripcion_turbidez"`,
	Turbidez		float32			`json:"turbidez" gorm:"type:decimal(10,2);default:0"`,
	TimeAndDateLocal	string	`json:"time_and_date_local"`
	CreatedAt time.Time 		`json:"created"`
	UpdatedAt time.Time 		`json:"updated"`
	DeletedAt gorm.DeletedAt	`gorm:"index" json:"deleted"`

}