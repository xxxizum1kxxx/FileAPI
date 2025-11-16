package handler

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/xxxizum1kxxx/fileapi/internal/filesystem"
)

// ResolvePath extracts the relative path from r.URL.Path by trimming the given routePrefix,
// optionally allows empty relative path, performs URL-decoding and SecureJoin with basePath.
// Returns the absolute filesystem path or an error message with appropriate HTTP status code.
func ResolvePath(r *http.Request, basePath string, routePrefix string, allowEmpty bool) (string, int, string) {
	relativePath := strings.TrimPrefix(r.URL.Path, routePrefix)

	if relativePath == "" {
		if allowEmpty {
			// root listing: use basePath as target directory
			return strings.TrimSuffix(basePath, "/"), http.StatusOK, ""
		}
		return "", http.StatusBadRequest, "Path is required"
	}

	decoded, err := url.PathUnescape(relativePath)
	if err != nil {
		return "", http.StatusBadRequest, "Invalid path encoding"
	}

	absPath, err := filesystem.SecureJoin(strings.TrimSuffix(basePath, "/"), decoded)
	if err != nil {
		return "", http.StatusBadRequest, "Invalid path"
	}
	return absPath, http.StatusOK, ""
}
