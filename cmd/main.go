package main

import (
	"log"
	"net/http"
	"os"

	"go1fl-sprint6-final/internal/server"
)

func main() {

	logger := log.New(os.Stdout, "SERVER: ", log.LstdFlags|log.Lshortfile)
	srv := server.New(logger)

	err := srv.Start()
	if err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Ошибка: %v", err)
	}
}
