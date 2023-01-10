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

// @BasePath /v1
//
// @Summary Create operation type
// @Description Create a new operation type
// @Tags example
// @Accept  json
// @Produce  json
// @Param request body models.SwaggerCreateTemplate true "Create operation type payload"
// @Success 201 "Successfully Created!"
// @Router /operation_types [post]
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

// @BasePath /v1
//
// @Summary Get operation types
// @Schemes
// @Description Get operation type list
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {object} models.Response{message=string,data=[]models.OperationType}
// @Failure 404
// @Router /operation_types [get]
func Read(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	response := controllers.GetOperationTypes(id)
	context.JSON(response.StatusCode, response)
}

// ReadByID ...

// @BasePath /v1
//
// @Summary Get operation type
// @Schemes
// @Description Get operation type by ID
// @Tags example
// @Accept json
// @Produce json
// @Param id	path	int		true	"Operation Type ID"
// @Success 200 {object} models.Response{message=string,data=models.OperationType}
// @Failure 404
// @Router /operation_types/{id} [get]
func ReadByID(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	response := controllers.GetOperationTypes(id)
	context.JSON(response.StatusCode, response)
}

// Update ...

// @BasePath /v1
//
// @Summary Update operation type
// @Description Update a new operation type
// @Tags example
// @Accept  json
// @Produce  json
// @Param request body models.SwaggerUpdateTemplate true "Update operation type payload"
// @Success 200 "Successfully Updated!"
// @Router /operation_types [put]
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

// @BasePath /v1
//
// @Summary Delete operation type
// @Description Delete a new operation type
// @Tags example
// @Accept  json
// @Produce  json
// @Param id	path	int		true	"Operation Type ID"
// @Success 200 "Successfully Deleted!"
// @Router /operation_types/{id} [delete]
func Delete(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	response := controllers.DeleteOperationType(&models.OperationType{ID: int64(id)})
	context.JSON(response.StatusCode, response)
}
