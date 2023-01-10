package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/leonardogazio/golang-account-crud/docs"
	hndOperationType "github.com/leonardogazio/golang-account-crud/server/routes/handlers/operationtype"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Setup ...
func Setup() *gin.Engine {
	engine := gin.Default()
	docs.SwaggerInfo.BasePath = "/v1"
	root := engine.Group("/v1")
	{
		root.GET("/operation_types", hndOperationType.Read)
		root.GET("/operation_types/:id", hndOperationType.ReadByID)
		root.POST("/operation_types", hndOperationType.Create)
		root.PUT("/operation_types", hndOperationType.Update)
		root.DELETE("/operation_types/:id", hndOperationType.Delete)

		// Swagger route.
		engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, ginSwagger.DefaultModelsExpandDepth(-1)))
	}
	return engine
}
