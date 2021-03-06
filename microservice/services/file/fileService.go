package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type fileService struct{}

// NewFileService func
func NewFileService() FileServiceInterface {
	return &fileService{}
}

// Write func
func (s fileService) Write(message []byte, directory string) error {
	data := map[string]interface{}{}
	decodeErr := json.Unmarshal(message, &data)

	if decodeErr != nil {
		return decodeErr
	}

	writeFileErr := writeFile(data, directory)

	if writeFileErr != nil {
		return writeFileErr
	}

	return nil
}

// Exist func
func (s fileService) Exist(path string) (bool, error) {
	_, err := os.Stat(path)

	if err != nil {
		return false, err
	}

	return true, nil
}

func writeFile(data map[string]interface{}, directory string) error {
	file, fileErr := json.MarshalIndent(data, "", " ")

	if fileErr != nil {
		return fileErr
	}

	filename := data["ID"]
	path := fmt.Sprintf("%s/%s.json", directory, filename)
	writeErr := ioutil.WriteFile(path, file, 0644)

	if writeErr != nil {
		return writeErr
	}

	return nil
}
