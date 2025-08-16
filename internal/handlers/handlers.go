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

	"go1fl-sprint6-final/internal/service"
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
	if r.Method != http.MethodPost {
		http.Error(w, "Ошибка", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		log.Printf("Ошибка: %v", err)
		http.Error(w, "Ошибка", http.StatusInternalServerError)
		return
	}

	file, header, err := r.FormFile("myFile")
	if err != nil {
		log.Printf("Ошибка: %v", err)
		http.Error(w, "Ошибка", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		log.Printf("Ошибка: %v", err)
		http.Error(w, "Ошибка", http.StatusInternalServerError)
		return
	}

	converted, err := service.Convert(string(fileBytes))
	if err != nil {
		log.Printf("Ошибка: %v", err)
		http.Error(w, "Ошибка", http.StatusInternalServerError)
		return
	}

	fileName := fmt.Sprintf("result_%s%s",
		time.Now().UTC().Format("20060102_150405"),
		filepath.Ext(header.Filename))

	err = os.WriteFile(fileName, []byte(converted), 0644)
	if err != nil {
		log.Printf("Ошибка: %v", err)
		http.Error(w, "Ошибка", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(converted))
}
