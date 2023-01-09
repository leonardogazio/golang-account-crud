package main

import (
	"github.com/leonardogazio/golang-account-crud/db"
	"github.com/leonardogazio/golang-account-crud/server"
	"github.com/leonardogazio/golang-account-crud/utils"
	"go.uber.org/zap"
)

func init() {
	utils.NewLogger()
}

func main() {
	if _, err := db.GetConnection(); err != nil {
		utils.Logger.Fatal("Database Connection Error", zap.Error(err))
	}

	if err := server.Start(); err != nil {
		utils.Logger.Fatal("Server Initialize Error", zap.Error(err))
	}
}
