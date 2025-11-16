package httpserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/xxxizum1kxxx/fileapi/internal/config"
)

func StartServer(config *config.Config) {
	router := http.NewServeMux()
	if err := RegisterRoutes(router); err != nil {
		log.Fatal(err)
	}
	if err := RegisterFileRoutes(router, config.Filesystem.Path); err != nil {
		log.Fatal(err)
	}
	if err := RegisterMetadataRoutes(router, config.Filesystem.Path); err != nil {
		log.Fatal(err)
	}
	if err := RegisterDirectoryRoutes(router, config.Filesystem.Path); err != nil {
		log.Fatal(err)
	}

	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.Server.Port), router); err != nil {
		log.Fatal(err)
	}
}
