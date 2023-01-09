package server

import (
	"fmt"

	"github.com/leonardogazio/golang-account-crud/server/routes"
	"github.com/leonardogazio/golang-account-crud/utils"
)

// Start ...
func Start() error {
	addr := fmt.Sprintf(":%s", utils.GetEnv("SERVER_HTTP_PORT", "80"))
	utils.Logger.Sugar().Infof("HTTP Server listening on: %s", addr)

	return routes.Setup().Run(addr)
}
