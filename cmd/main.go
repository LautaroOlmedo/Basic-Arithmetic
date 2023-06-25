package main

import (
	"Arithmetic/internal/adapters/app/api"
	"Arithmetic/internal/adapters/core/arithmetic"
	"Arithmetic/internal/adapters/framework/left/grpc"
	"Arithmetic/internal/adapters/framework/right/db"
	"Arithmetic/internal/ports"

	"log"
	"os"
)

func main() {
	var err error

	// PORTS
	var dbAdapter ports.DbPort
	var core ports.ArithmeticPort
	var appAdapter ports.APIPort
	var gRPCadapter ports.GRPCPort

	// ENVIRONMENT VARIABLES
	dbDriver := os.Getenv("DB_DRIVER")
	dbSourceName := os.Getenv("DS_NAME")

	dbAdapter, err = db.NewAdapter(dbDriver, dbSourceName)
	if err != nil {
		log.Fatalf("Failed to initiate database connection: %v", err)
	}

	defer dbAdapter.CloseDBConnection()

	core = arithmetic.NewAdapter()
	appAdapter = api.NewAdapter(dbAdapter, core)
	gRPCadapter = grpc.NewAdapter(appAdapter)

	gRPCadapter.Run()

}
