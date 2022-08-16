package models

import (
	"time"

	"gorm.io/gorm"
)
type CatalogData struct{
	ID          uint            `json:"id" gorm:"primaryKey"`
	ParentId	int64			`json:"parent_id" gorm:"string;default:NULL"`
	Description string 				`json:"description"`
	Keywork string 				`json:"keywork"`
	Status	string			`json:"status" gorm:"string;default:activo"`
	CreatedAt time.Time 		`json:"created"`
	UpdatedAt time.Time 		`json:"updated"`
	DeletedAt gorm.DeletedAt	`gorm:"index" json:"deleted"`
}