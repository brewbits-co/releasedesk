package build

import (
	"errors"
	"github.com/brewbits-co/releasedesk/internal/domains/build"
	"github.com/brewbits-co/releasedesk/internal/values"
	"github.com/brewbits-co/releasedesk/pkg/storage"
	"mime/multipart"
	"path/filepath"
)

func (s *service) UploadBuild(slug values.Slug, os values.OS, info build.BasicInfo, files map[values.Architecture]*multipart.FileHeader, Metadata map[string]string) (build.Build, error) {
	for _, header := range files {
		validFileType := isValidFileType(os, header.Filename)
		if !validFileType {
			return build.Build{}, errors.New("invalid file type")
		}
	}

	platformEntity, err := s.platformRepo.GetByAppSlugAndOS(slug, os)
	if err != nil {
		return build.Build{}, err
	}

	info.PlatformID = platformEntity.ID
	buildEntity := build.NewBuild(info)
	buildEntity.Metadata = Metadata

	for arch, header := range files {
		file, err := header.Open()
		if err != nil {
			return build.Build{}, errors.New("failed to open file")
		}
		// TODO: Fix this defer
		defer file.Close()

		md5Sum, sha256Sum, sha512Sum, err := storage.CalculateChecksums(file)
		if err != nil {
			return build.Build{}, err
		}

		filePath := storage.ConvertChecksumToPath(sha256Sum)
		err = storage.SaveFile(filePath, file)
		if err != nil {
			return build.Build{}, err
		}

		artifact := build.NewArtifact(header.Filename, header.Size, arch, md5Sum, sha256Sum, sha512Sum)

		buildEntity.Artifacts = append(buildEntity.Artifacts, artifact)
	}

	err = s.buildRepo.Save(&buildEntity)
	if err != nil {
		// TODO: Delete file
		return build.Build{}, err
	}

	return buildEntity, err
}

func isValidFileType(os values.OS, fileName string) bool {
	allowedExtensions := map[values.OS][]string{
		values.Windows: {".appx", ".appxbundle", ".appxupload", ".msix", ".msixbundle", ".msixupload", ".exe", ".zip", ".msi"},
		values.MacOS:   {".zip", ".platform.zip", ".dmg", ".pkg"},
		values.Linux:   {".AppImage", ".deb", ".rpm", ".tgz", ".gz", ".snap", ".flatpak"},
		values.Android: {".apk", ".aab"},
		values.IOS:     {".ipa"},
	}

	ext := filepath.Ext(fileName)
	for _, allowedExt := range allowedExtensions[os] {
		if ext == allowedExt {
			return true
		}
	}
	return false
}
