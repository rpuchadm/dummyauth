package main

import (
	"myproject/aplicaciones"
	"myproject/ini"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	router := gin.Default()
	router.Any("/aplicaciones", aplicaciones.RouteAplicaciones)
	router.GET("/init", ini.RouteInit)
	router.GET("/drop", ini.RouteDrop)

	router.Run(":8080")
}
