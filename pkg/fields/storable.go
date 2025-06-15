package fields

import "fmt"

// Storable represents the attributes required to store a file.
type Storable struct {
	// Filename is the name of the file, including its extension.
	Filename string `xorm:"not null"`
	// MimeType is the MIME type of the file (e.g., "application/json", "image/png").
	MimeType string `xorm:"varchar(100) not null"`
	// Size is the size of the file in bytes.
	Size FileSize `xorm:"not null"`
	// Path is the storage location where the file is saved.
	Path string `xorm:"not null"`
}

type FileSize int64

// String converts a FileSize to a human-readable string like "1K", "234M", "2G", etc.
func (s FileSize) String() string {
	const (
		kilobyte = 1024
		megabyte = kilobyte * 1024
		gigabyte = megabyte * 1024
		terabyte = gigabyte * 1024
	)

	switch {
	case s >= terabyte:
		return fmt.Sprintf("%.2fT", float64(s)/float64(terabyte))
	case s >= gigabyte:
		return fmt.Sprintf("%.2fG", float64(s)/float64(gigabyte))
	case s >= megabyte:
		return fmt.Sprintf("%.2fM", float64(s)/float64(megabyte))
	case s >= kilobyte:
		return fmt.Sprintf("%.2fK", float64(s)/float64(kilobyte))
	default:
		return fmt.Sprintf("%dB", s)
	}
}
