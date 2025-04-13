package build

import (
	"errors"
	"github.com/brewbits-co/releasedesk/internal/domains/build"
	"github.com/brewbits-co/releasedesk/internal/values"
	"github.com/brewbits-co/releasedesk/pkg/storage"
	"mime/multipart"
	"path/filepath"
)

func (s *service) UploadBuild(slug values.Slug, platform values.Platform, info build.BasicInfo, files map[values.Architecture]*multipart.FileHeader, Metadata map[string]string) (build.Build, error) {
	for _, header := range files {
		validFileType := isValidFileType(platform, header.Filename)
		if !validFileType {
			return build.Build{}, errors.New("invalid file type")
		}
	}

	appEntity, err := s.appRepo.GetByProductSlugAndPlatform(slug, platform)
	if err != nil {
		return build.Build{}, err
	}

	info.AppID = appEntity.ID
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

func isValidFileType(platform values.Platform, fileName string) bool {
	allowedExtensions := map[values.Platform][]string{
		values.Windows: {".appx", ".appxbundle", ".appxupload", ".msix", ".msixbundle", ".msixupload", ".exe", ".zip", ".msi"},
		values.MacOS:   {".zip", ".app.zip", ".dmg", ".pkg"},
		values.Linux:   {".AppImage", ".deb", ".rpm", ".tgz", ".gz", ".snap", ".flatpak"},
		values.Android: {".apk", ".aab"},
		values.IOS:     {".ipa"},
		values.Other:   {".zip"},
	}

	ext := filepath.Ext(fileName)
	for _, allowedExt := range allowedExtensions[platform] {
		if ext == allowedExt {
			return true
		}
	}
	return false
}
