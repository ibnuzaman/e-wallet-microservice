package main

import (
	"ewallet-framework/cmd"
	"ewallet-framework/helpers"
)

func main() {

	// Load Config
	helpers.SetupConfig()

	// Load Logger
	helpers.SetupLogger()

	// Load MySQL
	helpers.SetupMySQL()

	// Load GRPC
	go cmd.ServerGRPC()

	// Load HTTP
	cmd.ServerHttp()
}
