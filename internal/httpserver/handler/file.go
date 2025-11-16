package handler

import (
	"net/http"
	"path/filepath"

	"github.com/xxxizum1kxxx/fileapi/internal/filesystem"
)

func FileGet(w http.ResponseWriter, r *http.Request, path string) {
	if !filesystem.IsFilePathLegal(path) {
		http.Error(w, "Invalid path", http.StatusBadRequest)
		return
	}
	file, err := filesystem.GetFile(path)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	// Set the Content-Type based on the file extension
	ext := filepath.Ext(path)
	contentType := getContentType(ext)
	w.Header().Set("Content-Type", contentType)

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(file.Bytes()); err != nil {
		// Best-effort: client may have gone away; avoid leaking internal errors
		return
	}
}

func getContentType(ext string) string {
	contentTypes := map[string]string{
		".html": "text/html",
		".css":  "text/css",
		".js":   "application/javascript",
		".json": "application/json",
		".xml":  "application/xml",
		".png":  "image/png",
		".jpg":  "image/jpeg",
		".jpeg": "image/jpeg",
		".gif":  "image/gif",
		".svg":  "image/svg+xml",
		".pdf":  "application/pdf",
		".txt":  "text/plain",
		".csv":  "text/csv",
	}
	if ct, ok := contentTypes[ext]; ok {
		return ct
	}
	return "application/octet-stream"
}
