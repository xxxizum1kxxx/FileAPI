package filesystem

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func IsFilePathLegal(path string) bool {
	// Check if the path is empty
	if path == "" {
		return false
	}
	// Check if the path contains path traversal attack (like ../ or ..\)
	if strings.Contains(path, "..") {
		return false
	}
	// Clean the path and check if it is valid
	cleanPath := filepath.Clean(path)
	// Prevent path traversal attack (the cleaned path should not contain ..)
	if strings.Contains(cleanPath, "..") {
		return false
	}
	// Ensure the cleaned path is not empty
	if cleanPath == "" || cleanPath == "." {
		return false
	}
	return true
}

func IsExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true
}

// SecureJoin joins base and rel into an absolute path under base,
// rejecting any traversal that would escape the base directory.
func SecureJoin(base string, rel string) (string, error) {
	if base == "" {
		return "", errors.New("base path is empty")
	}
	// Disallow absolute relative part
	if filepath.IsAbs(rel) {
		return "", errors.New("relative path must not be absolute")
	}
	// Clean components
	baseAbs, err := filepath.Abs(base)
	if err != nil {
		return "", err
	}
	relClean := filepath.Clean(rel)
	joined := filepath.Join(baseAbs, relClean)
	joinedAbs, err := filepath.Abs(joined)
	if err != nil {
		return "", err
	}
	// Ensure joined path is within base
	relToBase, err := filepath.Rel(baseAbs, joinedAbs)
	if err != nil {
		return "", err
	}
	if relToBase == ".." || strings.HasPrefix(relToBase, "..") {
		return "", errors.New("path escapes base directory")
	}
	return joinedAbs, nil
}
