package operationtype

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/leonardogazio/golang-account-crud/db/controllers"
	"github.com/leonardogazio/golang-account-crud/models"
	"github.com/leonardogazio/golang-account-crud/utils/validation"
)

// Create ...
func Create(context *gin.Context) {
	var m models.OperationType

	if err := context.ShouldBindJSON(&m); err != nil {
		context.JSON(http.StatusBadRequest, models.NewResponse(-1, "Invalid Params", validation.GenerateValidationResponse(err)))
		return
	}

	response := controllers.CreateOperationType(&m)
	context.JSON(response.StatusCode, response)
}

// Read ...
func Read(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	response := controllers.GetOperationTypes(id)
	context.JSON(response.StatusCode, response)
}

// Update ...
func Update(context *gin.Context) {
	var m models.OperationType

	if err := context.ShouldBindJSON(&m); err != nil {
		context.JSON(http.StatusBadRequest, models.NewResponse(-1, "Invalid Params", validation.GenerateValidationResponse(err)))
		return
	}

	response := controllers.UpdateOperationType(&m)
	context.JSON(response.StatusCode, response)
}

// Delete ...
func Delete(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	response := controllers.DeleteOperationType(&models.OperationType{ID: int64(id)})
	context.JSON(response.StatusCode, response)
}
