package server

import (
	"github.com/gin-gonic/gin"
)

type Server interface {
	Run()
}

type server struct {
	router *gin.Engine
}

func New(CORSHeader string) Server {
	r := router(CORSHeader)
	return &server{router: r}
}

func (server *server) Run() {
	server.router.Run()
}
