package handlers

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("../index.html")
	if err != nil {
		log.Printf("Ошибка: %v", err)
		http.Error(w, "Ошибка", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Printf("Ошибка: %v", err)
		http.Error(w, "Ошибка", http.StatusInternalServerError)
	}
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(10 << 20)

	file, header, err := r.FormFile("myFile")
	if err != nil {
		log.Fatal(err)
		http.Error(w, "ошибка при получении файла", http.StatusBadRequest)
		return
	}

	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Ошибка", http.StatusInternalServerError)
		return
	}

	converted, err := service.Convert(string(fileBytes))
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Ошибка", http.StatusInternalServerError)
		return
	}

	fileName := fmt.Sprintf("result_%s%s",
		time.Now().UTC().Format("20060102_150405"),
		filepath.Ext(header.Filename))

	err = os.WriteFile(fileName, []byte(converted), 0644)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Ошибка", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	pattern :=
		`<!DOCTYPE html>
  <html lang="ru"><head>
  <meta charset="utf-8" />
  <title>Ответ</title>
  </head>
<body>%s</body></html>`

	result := fmt.Sprintf(pattern, converted)
	w.Write([]byte(result))
}
