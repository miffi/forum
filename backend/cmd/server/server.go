package main

import (
	"os"

	"github.com/miffi/backend/internal/server"
)

func main() {
	CORSHeader := os.Getenv("CORS_HEADER")

	server := server.New(CORSHeader)
	server.Run()
}
