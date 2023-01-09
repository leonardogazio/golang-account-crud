package handlers

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/leonardogazio/golang-account-crud/db"
	"github.com/leonardogazio/golang-account-crud/server/routes"
	"github.com/leonardogazio/golang-account-crud/utils"
)

// TestEngine ...
var TestEngine *gin.Engine

func Init(m *testing.M) {
	utils.NewLogger()
	TestEngine = routes.Setup()
	db.InitMocked = true
	db.GetConnection() //database connection is mocked in this case, error check not needed
	os.Exit(m.Run())
}
