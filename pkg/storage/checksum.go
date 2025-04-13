package storage

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"io"
	"mime/multipart"
)

// CalculateChecksums computes Md5, SHA-256, and SHA-512 checksums for the input file using a multi-writer.
// It returns the hex-encoded Md5, SHA-256, and SHA-512 hashes, along with any error encountered during processing.
func CalculateChecksums(file multipart.File) (string, string, string, error) {
	md5Hash := md5.New()
	sha256Hash := sha256.New()
	sha512Hash := sha512.New()
	multiWriter := io.MultiWriter(md5Hash, sha256Hash, sha512Hash)

	if _, err := io.Copy(multiWriter, file); err != nil {
		return "", "", "", err
	}

	md5Sum := hex.EncodeToString(md5Hash.Sum(nil))
	sha256Sum := hex.EncodeToString(sha256Hash.Sum(nil))
	sha512Sum := hex.EncodeToString(sha512Hash.Sum(nil))

	return md5Sum, sha256Sum, sha512Sum, nil
}
