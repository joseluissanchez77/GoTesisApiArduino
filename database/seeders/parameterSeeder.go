package seeders

import (
	"gorm.io/gorm"
	"github.com/joseluissanchez77/GoTesisApiArduino/database/fakers"
)

type Seeder struct {
	Seeder interface{}
}


func RegisterSeeders(db *gorm.DB) []Seeder {
	return []Seeder{
		{Seeder: fakers.ParameterDataFaker(db)},
	}
}

func DBSeed(db *gorm.DB) error {
	for _, seeder := range RegisterSeeders(db) {
		err := db.Debug().Create(seeder.Seeder).Error
		if err != nil {
			return err
		}
	}

	return nil
}



// var parameterDatas = []models.ParameterData{
// 	models.ParameterData{
// 		Sensor: "Motor",
// 		WeekdayEs:    "Domingo",
// 		WeekdayEn: "Domingo",
// 		Weekday: "Do",
// 		HourInitial : "17:00",
// 		HourEnd : "17:10",

// 	},
// 	models.ParameterData{
// 		Sensor: "Motor",
// 		WeekdayEs:    "Domingo",
// 		WeekdayEn: "Domingo",
// 		Weekday: "Do",
// 		HourInitial : "17:00",
// 		HourEnd : "17:10",
// 	},


// }

// func All() []seed.Seed {
// 	return []seed.Seed{}
// }

