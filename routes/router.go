package routes
import (
	"github.com/joseluissanchez77/GoTesisApiArduino/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine{

	main := router.Group("api/v1")
	{
		sensorData := main.Group("sensor-data")
		{
			sensorData.GET("/:id", controllers.ShowSensorData)
			sensorData.GET("/", controllers.ShowAllSensorData)
			sensorData.POST("/", controllers.CreateSensorData)
			sensorData.PUT("/", controllers.UpdateSensorData)
			sensorData.DELETE("/:id", controllers.DeleteSensorData)
			sensorData.GET("/last-record", controllers.LastRecordSensorData)
		}
	}

	return router;
}

