package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
