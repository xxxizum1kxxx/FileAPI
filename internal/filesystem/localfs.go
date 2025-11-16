package filesystem

import (
	"bytes"
	"os"
	"path/filepath"
	"strconv"
	"syscall"
)

func GetLocalMetadata(path string) (Metadata, error) {
	path = filepath.Clean(path)

	info, err := os.Stat(path)
	if err != nil {
		return Metadata{}, err
	}

	isFile := !info.IsDir()

	// default value, prevent Sys() or type assertion failure
	var uid, gid string

	if sys := info.Sys(); sys != nil {
		if stat, ok := sys.(*syscall.Stat_t); ok {
			uid = strconv.FormatUint(uint64(stat.Uid), 10)
			gid = strconv.FormatUint(uint64(stat.Gid), 10)
		}
	}

	m := Metadata{
		Path:      path,
		Name:      filepath.Base(path),
		IsFile:    isFile,
		UpdatedAt: info.ModTime(),
		Owner:     uid,
		Group:     gid,
	}

	if isFile {
		m.Size = info.Size()
		m.Extension = filepath.Ext(path)
	}

	return m, nil
}

func GetLocalFile(path string) (bytes.Buffer, error) {
	file, err := os.Open(path)
	if err != nil {
		return bytes.Buffer{}, err
	}
	defer file.Close()
	buf := bytes.NewBuffer(make([]byte, 0))
	_, err = buf.ReadFrom(file)
	if err != nil {
		return bytes.Buffer{}, err
	}
	return *buf, nil
}
