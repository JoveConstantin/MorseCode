package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	service "github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func MainHandle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if req.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.WriteHeader(http.StatusOK)
	http.ServeFile(w, req, "../index.html")
}

func UploadHandle(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	if req.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := req.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusInternalServerError)
		return
	}
	file, header, err := req.FormFile("myFile")
	if err != nil {
		http.Error(w, "Error retrieving file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		return
	}
	fileStrings := string(fileBytes)
	convertData, err := service.MorseOrTextConverter(fileStrings)
	outputFilename := time.Now().UTC().Format("20060102_150405") + filepath.Ext(header.Filename)
	outputFile, err := os.Create(outputFilename)
	if err != nil {
		http.Error(w, "Failed to create file", http.StatusInternalServerError)
		return
	}
	defer outputFile.Close()

	if _, err := outputFile.WriteString(convertData); err != nil {
		http.Error(w, "Failed to write result to file", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = fmt.Fprintf(w, "Converted data:\n%s\n\nOriginal text:\n%s", convertData, fileStrings)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}
