package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {

	logger := log.New(os.Stdout, "SERVER: ", log.LstdFlags|log.Lshortfile)
	srv := server.New(logger)

	err := srv.Start()
	if err != nil && err != http.ErrServerClosed {
		logger.Fatal(err)
	}
}
