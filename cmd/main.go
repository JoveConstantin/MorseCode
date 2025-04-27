package main

import (
	"log"
	"os"

	server "github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "MyServer:", log.LstdFlags)
	server := server.MyRouter(logger)
	logger.Println("Port:", server.Server.Addr)
	if err := server.Server.ListenAndServe(); err != nil {
		logger.Fatal("Server failed: ", err)
	}

}
