package server

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func router() *gin.Engine {
	router := gin.Default()
	router.SetTrustedProxies(nil)

	router.Use(static.ServeRoot("/", "public"))

	api := router.Group("/api")
	{
		api.GET("/ping", ping)
	}
	return router
}
