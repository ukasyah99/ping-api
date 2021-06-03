package main

import (
	"github.com/joho/godotenv"
	"github.com/ukasyah99/ping-api/routes"

	_ "github.com/ukasyah99/ping-api/docs"
)

// @title Ping API
// @version 1.0
// @description A simple ping API.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name The Unlicense
// @license.url http://unlicense.org

// @host localhost:8002

func main() {
	godotenv.Load()
	routes.Run()
}
