package models

import (
	"time"

	"gorm.io/gorm"
)
type ParameterData struct{
	ID          uint            `json:"id" gorm:"primaryKey"`
	Sensor string 				`json:"sensor"`
	WeekdayEs string 			`json:"weekday_es"`
	WeekdayEn string 			`json:"weekday_en"`
	// Weekday string 				`gorm:"type:enum('Sun', 'Mon', 'Tue','Wed','Thu','Fri','Sat')" json:"weekday"`
	Weekday string 				`json:"weekday"`
	HourInitial string 			`json:"hour_initial"`
	HourEnd string 				`json:"hour_end"`
	CreatedAt time.Time 		`json:"created"`
	UpdatedAt time.Time 		`json:"updated"`
	DeletedAt gorm.DeletedAt	`gorm:"index" json:"deleted"`
}