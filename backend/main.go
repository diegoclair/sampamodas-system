package main

import (
	"os"

	"github.com/diegoclair/go_utils-lib/logger"
	"github.com/diegoclair/sampamodas-system/backend/data"
	"github.com/diegoclair/sampamodas-system/backend/server"

	"github.com/diegoclair/sampamodas-system/backend/service"
)

func main() {
	logger.Info("Reading the initial configs...")

	db, err := data.Connect()
	if err != nil {
		panic(err)
	}
	svc := service.New(db)
	server := server.InitServer(svc)
	logger.Info("About to start the application...")

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	if err := server.Run(":" + port); err != nil {
		panic(err)
	}
}
