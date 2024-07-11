package main

import (
	"net/http"

	"github.com/aukdev/ruwan_go_10_07_2024/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	var server = gin.New()
	// main route
	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "hi",
		})
	})

	var apiRoute = server.Group("/api/v1")
	routes.RootRoute(apiRoute)
	println("server")
	server.Run()
}
