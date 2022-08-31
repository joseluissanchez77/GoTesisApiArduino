package controllers

import (
	// "bytes"
	"net/http"
	"fmt"
	"time"
	// "log"
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


// func Ternary[T any](condition bool, If, Else T) T {
//     if condition {
//         return If
//     }
//     return Else
// }
// fmt.Println(Ternary(1 < 2, "yes", "no")) // yes

func CreateSensorData(c *gin.Context){

	db := database.GetDatabase()

	loc, erro := time.LoadLocation("America/Guayaquil")

    if erro != nil {
        fmt.Println(erro)
    }


	day  := time.Now().UTC().In(loc).Format("Mon")
	date := time.Now().UTC().In(loc).Format("01-02-2006")
	currentHour := time.Now().UTC().In(loc).Format("15:04:05")
	fmt.Println(day, date, currentHour)

	var parameterData models.ParameterData

	// erroP = db./* Select([]string{"hour_initial", "hour_end"}). */First(&parameterData, newid).Error  Tue
	// erros := db.Where("Weekday LIKE ?", "%Sat%").Find(&parameterData).Error 
	erros := db.Where("Weekday = ? and keywork = ?",day , "ALI").Find(&parameterData).Error 
	
	if erros != nil{
		c.JSON(400, gin.H{
			"error" : "No se encontro parametros: "+erros.Error(),
		})
		return
	}

	//aceder al request
	sensorData:=  models.SensorData{}
	c.Bind(&sensorData)

	// c.JSON(http.StatusOK, gin.H{"response": sensorData.Description, "para": parameterData.HourInitial})


	// var sensordata models.SensorData
	// // log.Println("---->>>", &sensordata)
	// err := c.ShouldBindJSON(&sensordata)

	var nut string
	if !(parameterData.HourInitial >= currentHour)  && !(currentHour <= parameterData.HourEnd)  {
		nut = "apagado"
	} else {
		nut = "alimentando"
	}

	var waterPump string
	var parameterDataBomb models.ParameterData
	errosBom := db.Where("keywork = ?", "BOM").Find(&parameterDataBomb).Error 
	if errosBom != nil{
		c.JSON(400, gin.H{
			"error" : "No se encontro parametros: "+errosBom.Error(),
		})
		return
	}
	
	if (parameterDataBomb.InitialRangeOff >= sensorData.WaterLevel) && (sensorData.WaterLevel <= parameterDataBomb.EndRangeOff)  {
		waterPump = "apagado"//apagado
	} else {
		waterPump = "encendido"//encendido
	}

	
	sensordataModf := models.SensorData{
		Description : 		sensorData.Description,
		Celsius  	: 		sensorData.Celsius,
		Fahrenheit 	: 		sensorData.Fahrenheit,
		WaterLevel 	: 		sensorData.WaterLevel,
		Ph			: 		sensorData.Ph,
		Nutrition	: 		nut ,
		WaterPump	:		waterPump,
		TimeAndDateLocal :  time.Now().UTC().In(loc).Format("01-02-2006 15:04:05"),
		DescripcionTurbidez : sensorData.DescripcionTurbidez,
		Turbidez	: 		sensorData.Turbidez,
	}

	// err := c.ShouldBindJSON(&sensordataModf)
	
	// if err != nil{
	// 	c.JSON(400, gin.H{
	// 		"error" : "No se puede enlazar json: "+err.Error(),
	// 	})
	// 	return
	// }

	//crear
	err := db.Create(&sensordataModf).Error 

	if err != nil{
		c.JSON(400, gin.H{
			"error" : "Error al guardar datos del sensor: "+err.Error(),
		})
		return
	}


	c.JSON(http.StatusOK, sensordataModf)


	// description := c.PostForm("description")

	// c.JSON(http.StatusOK, gin.H{"response": c.PostForm("description")})
	// c.JSON(200, description)
	
	/* db := database.GetDatabase()

	rqt := models.SensorData{
		Description : 		"test365",
		Celsius  	: 		1,
		Fahrenheit 	: 		2,
		WaterLevel 	: 		3,
		Ph			: 		4,
		Nutrition	: 		1,
	}

	err := db.Create(&rqt)
	
	if err != nil{
		log.Println(err)
	}

	c.JSON(200, rqt.ID) */

	// // loc, erro := time.LoadLocation("America/Guayaquil")

    // // if erro != nil {
    // //     fmt.Println(erro)
    // // }

	// // var parameterData models.ParameterData
	// // newid := 1
	// // // db := database.GetDatabase()
	// // erro = db./* Select([]string{"hour_initial", "hour_end"}). */First(&parameterData, newid).Error 

	// // if erro != nil{
	// // 	c.JSON(400, gin.H{
	// // 		"error" : "No se encontro datos del sensor: "+erro.Error(),
	// // 	})
	// // 	return
	// // }
	// // c.JSON(http.StatusOK, time.Now().UTC().In(loc).Format("Mon 01-02-2006 15:04:05"))

	/* buf := make([]byte, 1024)
	num, _ := c.Request.Body.Read(buf)
	reqBody := string(buf[0:num])
	 */

	// // // // db := database.GetDatabase()

	// // // // var sensordata models.SensorData
	// // // // // log.Println("---->>>", &sensordata)
	// // // // err := c.ShouldBindJSON(&sensordata)
	
	// // // // if err != nil{
	// // // // 	c.JSON(400, gin.H{
	// // // // 		"error" : "No se puede enlazar json: "+err.Error(),
	// // // // 	})
	// // // // 	return
	// // // // }

	// // // // //crear
	// // // // err = db.Create(&sensordata).Error 

	// // // // if err != nil{
	// // // // 	c.JSON(400, gin.H{
	// // // // 		"error" : "Error al guardar datos del sensor: "+err.Error(),
	// // // // 	})
	// // // // 	return
	// // // // }


	// // // // c.JSON(200, sensordata)
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