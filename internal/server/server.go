package server

import (
	"log"
	"net/http"
	"time"

	handler "github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type MyServ struct {
	Logger *log.Logger
	Server *http.Server
}

func MyRouter(logger *log.Logger) *MyServ {

	myRouter := http.NewServeMux()
	myRouter.HandleFunc("/", handler.MainHandle)
	myRouter.HandleFunc("/upload", handler.UploadHandle)

	myServ := &MyServ{
		Logger: logger,
		Server: &http.Server{
			Addr:    ":8080",
			Handler: myRouter,
			// Добавьте остальные параметры:
			ErrorLog:     logger,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  15 * time.Second,
		},
	}

	return myServ
}
