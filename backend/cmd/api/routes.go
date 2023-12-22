package main

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func (app *application) router() *gin.Engine {
	router := gin.Default()

	router.Use(static.ServeRoot("/", "public"))

	api := router.Group("/api")
	{
		api.GET("/ping", app.ping)
	}
	return router
}
