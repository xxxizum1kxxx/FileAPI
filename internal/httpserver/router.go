package httpserver

import (
	"fmt"
	"net/http"

	"github.com/xxxizum1kxxx/fileapi/internal/httpserver/handler"
)

func RegisterRoutes(router *http.ServeMux) error {
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	return nil
}

func RegisterMetadataRoutes(router *http.ServeMux, basePath string) error {
	router.HandleFunc("/api/v1/metadata/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			// Get metadata of a file or directory
			metadataPath, status, errmsg := handler.ResolvePath(r, basePath, "/api/v1/metadata/", false)
			if errmsg != "" {
				http.Error(w, errmsg, status)
				return
			}
			handler.MetadataGet(w, r, metadataPath)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
	})
	return nil
}

func RegisterFileRoutes(router *http.ServeMux, basePath string) error {
	router.HandleFunc("/api/v1/files/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			// Get a file
			filePath, status, errmsg := handler.ResolvePath(r, basePath, "/api/v1/files/", false)
			if errmsg != "" {
				http.Error(w, errmsg, status)
				return
			}
			handler.FileGet(w, r, filePath)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
	})
	return nil
}

func RegisterDirectoryRoutes(router *http.ServeMux, basePath string) error {
	router.HandleFunc("/api/v1/directory/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			// Get a directory
			directoryPath, status, errmsg := handler.ResolvePath(r, basePath, "/api/v1/directory/", true)
			if errmsg != "" {
				http.Error(w, errmsg, status)
				return
			}
			handler.DirGet(w, r, directoryPath)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
	})
	return nil
}
