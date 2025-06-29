package fields

// Hashable represents data used to verify the integrity of a file.
type Hashable struct {
	// Md5 is the Md5 checksum for verifying the file's integrity.
	Md5 string `xorm:"varchar(32) not null"`
	// Sha256 is the SHA-256 checksum for verifying the file's integrity.
	Sha256 string `xorm:"varchar(64) not null"`
	// Sha512 is the SHA-512 checksum for verifying the file's integrity.
	Sha512 string `xorm:"varchar(128) not null"`
}
