package main

import (
	"log"
	"os"

	"github.com/miffi/backend/internal/server"
)

func main() {
	CORSHeader := os.Getenv("CORS_HEADER")
	if CORSHeader == "" {
		log.Fatal("CORS_HEADER must be set to the frontend's URL")
	}

	server := server.New(CORSHeader)
	server.Run()
}
