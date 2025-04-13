package storage

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

const StorageDir = "_data/storage"

// SaveFile saves a multipart file to the specified file path on the filesystem, ensuring the target directory exists.
// Returns an error if directory creation, file operations, or writing fails.
func SaveFile(filePath string, file multipart.File) error {
	// Ensure the folder exists
	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return err
	}

	// Save the file to the storage folder
	if _, err := file.Seek(0, io.SeekStart); err != nil {
		return err
	}
	storedFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer storedFile.Close()

	if _, err := io.Copy(storedFile, file); err != nil {
		return err
	}
	return nil
}
