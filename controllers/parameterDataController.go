package controllers

import (
	"net/http"
	"github.com/joseluissanchez77/GoTesisApiArduino/models"
	"github.com/joseluissanchez77/GoTesisApiArduino/database"
	"github.com/gin-gonic/gin"
)


func CreateFoodParameterData(c *gin.Context){

	
	db := database.GetDatabase()

	//aceder al request
	parameterData:=  models.ParameterData{}
	c.Bind(&parameterData)

	var diaEs string
	var diaEn string
	keyworkRpt := parameterData.Weekday
	
	switch keyworkRpt {
		case "Sun":
			diaEs = "Domingo"
			diaEn = "Sunday"
		case "Mon":
			diaEs = "Lunes"
			diaEn = "Monday"
		case "Tue":
			diaEs = "Martes"
			diaEn = "Tuesday"
		case "Wed":
			diaEs = "Miércoles"
			diaEn = "Wednesday"
		case "Thu":
			diaEs = "Jueves"
			diaEn = "Thursday"
		case "Fri":
			diaEs = "Viernes"
			diaEn = "Friday"
		case "Sat":
			diaEs = "Sabado"
			diaEn = "Saturday"
	}


	parameterDataModf := models.ParameterData{
		Keywork : "ALI",
		Sensor : "Servo Alimentacion",
		InitialRangeOff	: 0,
		EndRangeOff	: 0,
		WeekdayEs : diaEs,
		WeekdayEn : diaEn,
		Weekday :parameterData.Weekday,
		HourInitial : parameterData.HourInitial,
		HourEnd :parameterData.HourEnd,
	}

	//crear
	err := db.Create(&parameterDataModf).Error 

	if err != nil{
		c.JSON(400, gin.H{
			"error" : "Error al guardar datos de alimentacion: "+err.Error(),
		})
		return
	}


	c.JSON(http.StatusOK, parameterDataModf)
}



func CreateWaterPumpsParameterData(c *gin.Context){

	
	db := database.GetDatabase()

	//aceder al request
	parameterData:=  models.ParameterData{}
	c.Bind(&parameterData)

	parameterDataModf := models.ParameterData{
		Keywork : "BOM",
		Sensor : "Bomba Agua",
		InitialRangeOff	: parameterData.InitialRangeOff,
		EndRangeOff	: parameterData.EndRangeOff,
		WeekdayEs : "",
		WeekdayEn : "",
		Weekday :"",
		HourInitial : "",
		HourEnd :"",
	}

	//crear
	err := db.Create(&parameterDataModf).Error 

	if err != nil{
		c.JSON(400, gin.H{
			"error" : "Error al guardar datos de bomba de agua: "+err.Error(),
		})
		return
	}


	c.JSON(http.StatusOK, parameterDataModf)
}

