package models

import (
	"time"

	"gorm.io/gorm"
)
type ParameterData struct{
	ID          uint            `json:"id" gorm:"primaryKey"`
	Sensor string 				`json:"sensor"`
	WeekdayEs string 			`json:"weekdayEs"`
	WeekdayEn string 			`json:"weekdayEn"`
	Weekday string 				`gorm:"type:enum('Sun', 'Mon', 'Tue','Wed','Thu','Fri','Sat')" json:"weekday"`
	Hour string 				`json:"hour"`
	CreatedAt time.Time 		`json:"created"`
	UpdatedAt time.Time 		`json:"updated"`
	DeletedAt gorm.DeletedAt	`gorm:"index" json:"deleted"`
}