package storage

import "path/filepath"

// ConvertChecksumToPath generates a storage path for a file based on its checksum with a two-character subdirectory.
// Example: If the StorageDir is "/data/storage" and the checksum is "abcdef123456",
// the resulting path will be "/data/storage/ab/abcdef123456".
func ConvertChecksumToPath(checksum string) string {
	subFolder := checksum[:2]
	filePath := filepath.Join(StorageDir, subFolder, checksum)
	return filePath
}
