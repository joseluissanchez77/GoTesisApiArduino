package controllers

import (
	"github.com/joseluissanchez77/GoTesisApiArduino/models"
	"github.com/joseluissanchez77/GoTesisApiArduino/database"
	"github.com/gin-gonic/gin"
)


func CreateFoodParameterData(c *gin.Context){

	
	db := database.GetDatabase()

	var parameterData models.ParameterData
	// log.Println("---->>>", &parameterData)
	err := c.ShouldBindJSON(&parameterData)
	
	if err != nil{
		c.JSON(400, gin.H{
			"error" : "No se puede enlazar json: "+err.Error(),
		})
		return
	}

	//crear
	err = db.Create(&parameterData).Error 

	if err != nil{
		c.JSON(400, gin.H{
			"error" : "Error al guardar datos de alimentacion: "+err.Error(),
		})
		return
	}


	c.JSON(200, parameterData)
}