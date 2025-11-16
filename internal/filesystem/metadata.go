package filesystem

import (
	"time"
)

type Metadata struct {
	Path      string    `json:"path"`
	Name      string    `json:"name"`
	Extension string    `json:"extension,omitempty"`
	Size      int64     `json:"size,omitempty"`
	IsFile    bool      `json:"is_file"`
	Owner     string    `json:"owner"` // UID of the owner
	Group     string    `json:"group"` // GID of the group
	UpdatedAt time.Time `json:"updated_at"`
}

func GetMetadata(path string) (Metadata, error) {
	return GetLocalMetadata(path)
}
