package routes

import (
	"github.com/gin-gonic/gin"
	hndOperationType "github.com/leonardogazio/golang-account-crud/server/routes/handlers/operationtype"
)

// Setup ...
func Setup() *gin.Engine {
	engine := gin.Default()
	root := engine.Group("/v1")
	{
		root.GET("/operation_types", hndOperationType.Read)
		root.GET("/operation_types/:id", hndOperationType.Read)
		root.POST("/operation_types", hndOperationType.Create)
		root.PUT("/operation_types", hndOperationType.Update)
		root.DELETE("/operation_types/:id", hndOperationType.Delete)
	}
	return engine
}
