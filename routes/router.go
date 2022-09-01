package routes
import (


	"github.com/joseluissanchez77/GoTesisApiArduino/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine{



	main := router.Group("/api/v1")
	{
		sensorData := main.Group("sensor-data")
		{
			sensorData.GET("/:id", controllers.ShowSensorData)
			sensorData.GET("/last-record", controllers.LastRecordSensorData)
			sensorData.GET("/", controllers.ShowAllSensorData)
			sensorData.POST("/", controllers.CreateSensorData)
			sensorData.PUT("/", controllers.UpdateSensorData)
			sensorData.DELETE("/:id", controllers.DeleteSensorData)
		}

		

		parameterData := main.Group("parameter-data")
		{
			// parameterData.GET("/:id", controllers.ShowParameterData)
			// parameterData.GET("/last-record", controllers.LastRecordparameterData)
			parameterData.GET("/parameter", controllers.ShowAllparameterData)
			parameterData.POST("/food", controllers.CreateFoodParameterData)
			parameterData.PUT("/food", controllers.UpdateFoodParameterData)
			parameterData.POST("/water-pumps", controllers.CreateWaterPumpsParameterData)
			parameterData.PUT("/water-pumps", controllers.UpdateWaterPumpsParameterData)
			// parameterData.PUT("/", controllers.UpdateparameterData)
			// parameterData.DELETE("/:id", controllers.DeleteparameterData)
		}

		catalogData := main.Group("catalog-data")
		{
			// catalogData.GET("/:id", controllers.ShowCatalogData)
			// catalogData.GET("/last-record", controllers.LastRecordCatalogData)
			catalogData.GET("", controllers.ShowAllCatalogData)
			// catalogData.POST("/", controllers.CreateCatalogData)
			// catalogData.PUT("/", controllers.UpdateCatalogData)
			// catalogData.DELETE("/:id", controllers.DeleteCatalogData)
		}
	}

	return router;
}

