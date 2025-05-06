package values

// VersionFormat defines the versioning format used in an application.
// It is used for automatically suggesting new release versions and verifying compliance
// with the defined versioning rules.
type VersionFormat string

const (
	// CustomFormat is a versioning format that allows you to fit application-specific needs.
	CustomFormat VersionFormat = "CustomFormat"
	// SemVer is a 3-component number in the format of MAJOR.MINOR.PATCH.
	SemVer VersionFormat = "SemVer"
	// CalVer is a versioning convention based on your application's release calendar.
	CalVer VersionFormat = "CalVer"
)
