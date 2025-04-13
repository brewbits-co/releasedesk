package build

import (
	"github.com/brewbits-co/releasedesk/internal/values"
	"github.com/brewbits-co/releasedesk/pkg/fields"
	"github.com/brewbits-co/releasedesk/pkg/storage"
)

func NewArtifact(filename string, size int64, arch values.Architecture, md5 string, sha256 string, sha512 string) Artifact {
	filePath := storage.ConvertChecksumToPath(sha256)

	return Artifact{
		Storable: fields.Storable{
			Filename: filename,
			MimeType: "application/vnd.android.package-archive",
			Size:     fields.FileSize(size),
			Path:     filePath,
		},
		Hashable: fields.Hashable{
			Md5:    md5,
			Sha256: sha256,
			Sha512: sha512,
		},
		ID:           0,
		BuildID:      0,
		Architecture: arch,
	}
}

type Artifact struct {
	fields.Storable
	fields.Hashable
	// ID is the unique identifier of an Artifact.
	ID int `db:"ID"`
	// BuildID is the identifier of the Build that this Artifact belongs.
	BuildID int `db:"BuildID"`
	// Architecture represents the system architecture for which the artifact is built.
	Architecture values.Architecture `db:"Architecture"`
}
