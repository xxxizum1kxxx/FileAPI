package filesystem

import (
	"bytes"
	"os"
)

func GetFile(path string) (bytes.Buffer, error) {
	filebytes, err := ReadLocalFile(path)
	if err != nil {
		return bytes.Buffer{}, err
	}
	return filebytes, nil
}

func GetDirectory(path string) ([]Metadata, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	metadataList := []Metadata{}
	for _, file := range files {
		metadata, err := GetMetadata(path + "/" + file.Name())
		if err != nil {
			return nil, err
		}
		metadataList = append(metadataList, metadata)
	}
	return metadataList, nil
}
