package main

import (
	"log"

	"github.com/xxxizum1kxxx/fileapi/internal/config"
	"github.com/xxxizum1kxxx/fileapi/internal/httpserver"
)

func main() {
	config, err := config.LoadConfig("configs/config.json")
	if err != nil {
		log.Fatal(err)
	}
	httpserver.StartServer(config)
}
