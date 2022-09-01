package routes
import (
	// "net/http"
	"github.com/joseluissanchez77/GoTesisApiArduino/controllers"
	"github.com/gin-gonic/gin"
	// "github.com/gin-contrib/cors"
	// "github.com/itsjamie/gin-cors"
	// "time"
	// "github.com/rs/cors"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine{


	main := router.Group("api/v1")
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
			catalogData.GET("/", controllers.ShowAllCatalogData)
			// catalogData.POST("/", controllers.CreateCatalogData)
			// catalogData.PUT("/", controllers.UpdateCatalogData)
			// catalogData.DELETE("/:id", controllers.DeleteCatalogData)
		}
	}

	// main.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"https://myxml.in"},
	// 	AllowMethods:     []string{"PUT", "PATCH", "GET","POST"},
	// 	AllowHeaders:     []string{"Origin"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	AllowOriginFunc: func(origin string) bool {
	// 		return origin == "https://github.com"
	// 	},
	// 	MaxAge: 12 * time.Hour,
	// }))


	// main.Use(gin.Logger(), gin.Recovery())

	// main.Use(cors.New(cors.Config{
	// 	AllowAllOrigins: true,
	// }))

	return router;
}
