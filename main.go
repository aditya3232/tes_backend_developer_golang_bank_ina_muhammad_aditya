package main

import (
	"fmt"

	"github.com/aditya3232/tes_backend_developer_golang_bank_ina_muhammad_aditya/config"
	"github.com/aditya3232/tes_backend_developer_golang_bank_ina_muhammad_aditya/connection"
	"github.com/aditya3232/tes_backend_developer_golang_bank_ina_muhammad_aditya/helper"
	"github.com/aditya3232/tes_backend_developer_golang_bank_ina_muhammad_aditya/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// panic recovery
	defer helper.RecoverPanic()

	router := gin.Default()
	if config.CONFIG.DEBUG == 0 {
		gin.SetMode(gin.ReleaseMode)
	}

	routes.Initialize(router)
	router.Run(fmt.Sprintf("%s:%s", config.CONFIG.APP_HOST, config.CONFIG.APP_PORT))

	defer connection.Close()
}
