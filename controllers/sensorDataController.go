package controllers

import (
	"github.com/joseluissanchez77/GoTesisApiArduino/models"
	"github.com/joseluissanchez77/GoTesisApiArduino/database"
	"strconv"
	"github.com/gin-gonic/gin"
)

func ShowSensorData(c *gin.Context){
	//recibe el id del request - por default llega en string
	id := c.Param("id")

	newid, err := strconv.Atoi(id)

	if err != nil{
		c.JSON(400, gin.H{
			"error" : "ID no es entero (int)",
		})

		return
	}

	db := database.GetDatabase()

	var sensordata models.SensorData
	err = db.First(&sensordata, newid).Error 

	if err != nil{
		c.JSON(400, gin.H{
			"error" : "No se encontro datos del sensor: "+err.Error(),
		})
		return
	}


	c.JSON(200, sensordata)
}


func CreateSensorData(c *gin.Context){

	db := database.GetDatabase()

	var sensordata models.SensorData

	err := c.ShouldBindJSON(&sensordata)

	if err != nil{
		c.JSON(400, gin.H{
			"error" : "No se puede enlazar json: "+err.Error(),
		})
		return
	}

	//crear
	err = db.Create(&sensordata).Error 

	if err != nil{
		c.JSON(400, gin.H{
			"error" : "Error al guardar datos del sensor: "+err.Error(),
		})
		return
	}


	c.JSON(200, sensordata)
}


func ShowAllSensorData(c *gin.Context){

	db:= database.GetDatabase()

	var sensordata []models.SensorData

	err := db.Order("id desc").Find(&sensordata).Error

	if err != nil{
		c.JSON(400, gin.H{
			"error" : "Error al obtener lista de datos sensor: "+err.Error(),
		})
		return
	}


	c.JSON(200, sensordata)
}


func UpdateSensorData(c *gin.Context){

	db := database.GetDatabase()

	var sensordata models.SensorData

	err := c.ShouldBindJSON(&sensordata)

	if err != nil{
		c.JSON(400, gin.H{
			"error" : "No se puede enlazar json: "+err.Error(),
		})
		return
	}

	//crear
	err = db.Save(&sensordata).Error 

	if err != nil{
		c.JSON(400, gin.H{
			"error" : "Error al ACTUALIZAR datos del sensor: "+err.Error(),
		})
		return
	}


	c.JSON(200, sensordata)
}


func DeleteSensorData(c *gin.Context){
	//recibe el id del request - por default llega en string
	id := c.Param("id")

	newid, err := strconv.Atoi(id)

	if err != nil{
		c.JSON(400, gin.H{
			"error" : "ID no es entero (int)",
		})

		return
	}

	db := database.GetDatabase()

	
	err = db.Delete(&models.SensorData{}, newid).Error 

	if err != nil{
		c.JSON(400, gin.H{
			"error" : "No se UDO BORRAR datos del sensor: "+err.Error(),
		})
		return
	}


	c.Status(204)
}



func LastRecordSensorData(c *gin.Context){

	db:= database.GetDatabase()

	var sensordata models.SensorData

	err := db.Order("id desc").Limit( 1 ).Find(&sensordata).Error

	if err != nil{
		c.JSON(400, gin.H{
			"error" : "Error al obtener lista de datos sensor: "+err.Error(),
		})
		return
	}


	c.JSON(200, sensordata)
}