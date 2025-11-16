package store

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type FileStorage struct{}

var FileStore = &FileStorage{}

func (fs *FileStorage) Init() error {
	storeDir := os.Getenv("FILE_STORE_PATH")
	if storeDir == "" {
		var err error
		storeDir, err = os.Getwd()
		if err != nil {
			return err
		}
	}

	// Ensure directory exists
	return os.MkdirAll(storeDir, 0755)
}

func (fs *FileStorage) Write(userId string, data interface{}) error {
	storeDir := os.Getenv("FILE_STORE_PATH")
	if storeDir == "" {
		dir, err := os.Getwd()
		if err != nil {
			return err
		}
		storeDir = dir
	}

	// Ensure directory exists
	if err := os.MkdirAll(storeDir, 0755); err != nil {
		return err
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	filePath := filepath.Join(storeDir, userId+".json")
	return os.WriteFile(filePath, jsonData, 0644)
}

func (fs *FileStorage) Read(userId string) (interface{}, error) {
	storeDir := os.Getenv("FILE_STORE_PATH")
	if storeDir == "" {
		dir, err := os.Getwd()
		if err != nil {
			return nil, err
		}
		storeDir = dir
	}

	filePath := filepath.Join(storeDir, userId+".json")
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var result interface{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
