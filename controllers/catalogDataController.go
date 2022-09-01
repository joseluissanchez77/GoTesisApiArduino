package controllers

import (
	// "fmt"
	"net/http"
	"github.com/joseluissanchez77/GoTesisApiArduino/models"
	"github.com/joseluissanchez77/GoTesisApiArduino/database"
	"github.com/gin-gonic/gin"
)


func ShowAllCatalogData(c *gin.Context){

	
	db:= database.GetDatabase()

	var catalogDataRpt models.CatalogData

	/* obtener el paramatro del request todos
	c.Request.URL.Query() */
	//aceder al request

	
	errosKey := db.Where("keywork = ?", c.Query("keywork")).Find(&catalogDataRpt).Error 
	if errosKey != nil{
		c.JSON(400, gin.H{
			"error" : "No se encontro data del keywork: "+errosKey.Error(),
		})
		return
	}
	
	// var id_cat = catalogDataRpt.ID

	var catalogDataModel []models.CatalogData
	erros := db.Where("parent_id = ?", catalogDataRpt.ID).Find(&catalogDataModel).Error 
	if erros != nil{
		c.JSON(400, gin.H{
			"error" : "No se encontro data del parent Id: "+erros.Error(),
		})
		return
	}


	c.JSON(http.StatusOK, catalogDataModel)

}
