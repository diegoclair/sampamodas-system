package main

import (
	"os"

	"github.com/diegoclair/go_utils-lib/logger"
	"github.com/diegoclair/sampamodas-system/backend/domain/service"
	"github.com/diegoclair/sampamodas-system/backend/infra/data"
	"github.com/diegoclair/sampamodas-system/backend/server"
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
		port = "5000"
	}

	if err := server.Start(":" + port); err != nil {
		panic(err)
	}
}
