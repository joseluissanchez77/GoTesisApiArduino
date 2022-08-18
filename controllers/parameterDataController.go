package controllers

import (
	// "strconv"
	"net/http"
	"github.com/joseluissanchez77/GoTesisApiArduino/models"
	"github.com/joseluissanchez77/GoTesisApiArduino/database"
	"github.com/gin-gonic/gin"
)

/*
guardar alimentacion
*/
func CreateFoodParameterData(c *gin.Context){

	
	db := database.GetDatabase()

	//aceder al request
	parameterData:=  models.ParameterData{}
	c.Bind(&parameterData)

	//validar rango de hora
	if parameterData.HourEnd <= parameterData.HourInitial  {
		c.JSON(http.StatusBadRequest , gin.H{"response": "Hora de fin debe ser mayor que la hora de incio.", "dataExists": nil})
		return
	}

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


/**
guardar nvel de bomba
*/
func CreateWaterPumpsParameterData(c *gin.Context){

	
	db := database.GetDatabase()

	
	//validar si ya existe al menos uno
	var countPumps = models.ParameterData{Weekday: "BOM"}
	db.First(&countPumps)

	if countPumps.ID!=0 {
		c.JSON(http.StatusBadRequest , gin.H{"response": "Ya existe un dato guardado, solo se debe actualizar", "dataExists": countPumps})
		return
	}


	//aceder al request
	parameterData:=  models.ParameterData{}
	c.Bind(&parameterData)

	//validar rango de hora
	if parameterData.EndRangeOff < parameterData.InitialRangeOff  {
		c.JSON(http.StatusBadRequest , gin.H{"response": "Rango de fin debe ser mayor quel rengo de incio.", "dataExists": nil})
		return
	}

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

/*
Bomba de agua
*/
func UpdateWaterPumpsParameterData(c *gin.Context){

	db := database.GetDatabase()

	//aceder al request
	parameterData:=  models.ParameterData{}
	c.Bind(&parameterData)

	//validar rango de hora
	if parameterData.EndRangeOff < parameterData.InitialRangeOff  {
		c.JSON(http.StatusBadRequest , gin.H{"response": "Rango de fin debe ser mayor quel rengo de incio.", "dataExists": nil})
		return
	}

	//validar si  existe 
	var countId = models.ParameterData{ID: parameterData.ID}
	errorVal := db.Where("keywork = 'BOM'").First(&countId).Error

	// if countId.ID!=0 {
	if errorVal != nil {
		c.JSON(http.StatusBadRequest , gin.H{"response": "No existe registro con el Id indicado", "dataExists": nil})
		return
	}

	parameterDataModf := models.ParameterData{
		ID	: parameterData.ID,
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
	err := db.Save(&parameterDataModf).Error 

	if err != nil{
		c.JSON(400, gin.H{
			"error" : "Error al ACTUALIZAR datos de bomba de agua: "+err.Error(),
		})
		return
	}


	c.JSON(200, parameterData)
}

/**
actaulizar ALimentacion
*/
func UpdateFoodParameterData(c *gin.Context){

	
	db := database.GetDatabase()

	//aceder al request
	parameterData:=  models.ParameterData{}
	c.Bind(&parameterData)

	//validar rango de hora
	if parameterData.HourEnd <= parameterData.HourInitial  {
		c.JSON(http.StatusBadRequest , gin.H{"response": "Hora de fin debe ser mayor que la hora de incio.", "dataExists": nil})
		return
	}

	//validar si  existe r
	var countId = models.ParameterData{ID: parameterData.ID}
	errV := db.Where("keywork = 'ALI'").First(&countId).Error

	if errV != nil {
		c.JSON(http.StatusBadRequest , gin.H{"response": "No existe registro con el Id indicado", "dataExists": nil})
		return
	}

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
		ID : parameterData.ID,
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
	err := db.Save(&parameterDataModf).Error 

	if err != nil{
		c.JSON(400, gin.H{
			"error" : "Error al actualizar datos de alimentacion: "+err.Error(),
		})
		return
	}


	c.JSON(http.StatusOK, parameterDataModf)
}



func ShowAllparameterData(c *gin.Context){

	//recibe el id del request - por default llega en string
	keywork := c.Query("keywork")

	db:= database.GetDatabase()

	var parameterData []models.ParameterData


	if keywork == "ALI" {	
		
		err := db.Where("keywork = 'ALI'").Order("id desc").Find(&parameterData).Error

		if err != nil{
			c.JSON(400, gin.H{"error" : "Error al obtener lista de parametros: "+err.Error()})
			return
		}

		c.JSON(200, parameterData)

		return
	}else if  keywork == "BOM" {
		err := db.Where("keywork = 'BOM'").Order("id desc").Find(&parameterData).Error

		if err != nil{
			c.JSON(400, gin.H{"error" : "Error al obtener lista de parametros: "+err.Error()})
			return
		}

		c.JSON(200, parameterData)

		return

	}

	err := db.Order("id desc").Find(&parameterData).Error

	if err != nil{
		c.JSON(400, gin.H{"error" : "Error al obtener lista de datos sensor: "+err.Error()})
		return
	}

	c.JSON(200, parameterData)

}
