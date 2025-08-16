package server

import (
	"go1fl-sprint6-final/internal/handlers"
	"log"
	"net/http"
	"time"
)

type Server struct {
	logger    *log.Logger
	server    *http.Server
	indexPath string
}

func New(logger *log.Logger) *Server {
	router := http.NewServeMux()
	indexPath := "../index.html"
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, indexPath)
	})
	router.HandleFunc("/upload", handlers.UploadHandler)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		logger:    logger,
		server:    srv,
		indexPath: indexPath,
	}
}

func generateResultPage(filename, original, converted string) string {
	return `<!DOCTYPE html>
<html>
<head>
	<title>Conversion Result</title>
	<style>
		body { font-family: Arial, sans-serif; margin: 20px; }
		.result { background: #f5f5f5; padding: 15px; margin: 10px 0; border-radius: 5px; }
	</style>
</head>
<body>
	<h1>Conversion Result</h1>
	<div>
		<h3>Original content (` + filename + `):</h3>
		<div class="result">` + original + `</div>
	</div>
	<div>
		<h3>Converted content:</h3>
		<div class="result">` + converted + `</div>
	</div>
	<a href="/">Upload another file</a>
</body>
</html>`
}

func (s *Server) Start() error {
	s.logger.Printf("Server starting on %s", s.server.Addr)
	return s.server.ListenAndServe()
}
