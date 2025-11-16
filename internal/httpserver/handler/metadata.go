package handler

import (
	"encoding/json"
	"net/http"

	"github.com/xxxizum1kxxx/fileapi/internal/filesystem"
)

func MetadataGet(w http.ResponseWriter, r *http.Request, path string) {
	if !filesystem.IsFilePathLegal(path) {
		http.Error(w, "Invalid path", http.StatusBadRequest)
		return
	}
	metadata, err := filesystem.GetMetadata(path)
	if err != nil {
		http.Error(w, "Metadata not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(metadata); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
