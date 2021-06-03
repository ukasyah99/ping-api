package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ukasyah99/ping-api/routes/ping"
	"github.com/ukasyah99/ping-api/routes/swagger"
)

var r = gin.Default()

func Run() {
	getRoutes()
	r.Run()
}

func getRoutes() {
	ping.AddRoutes(r)
	swagger.AddRoutes(r)
}
