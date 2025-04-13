package build

import (
	"github.com/brewbits-co/releasedesk/internal/domains/build"
	"github.com/brewbits-co/releasedesk/internal/values"
	"github.com/brewbits-co/releasedesk/pkg/fields"
	"github.com/brewbits-co/releasedesk/pkg/storage"
	"github.com/brewbits-co/releasedesk/pkg/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func (c *buildController) HandleBuildUpload(w http.ResponseWriter, r *http.Request) {
	// Parse the incoming multipart form
	r.ParseMultipartForm(10 << 20) // Limit upload size to ~10MB

	slug := chi.URLParam(r, "slug")
	platform := chi.URLParam(r, "platform")

	var buildUploadRequest = struct {
		build.BasicInfo
		fields.Extendable
	}{}

	err := utils.NewDecoder().Decode(&buildUploadRequest, r.Form)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.PlainText(w, r, err.Error())
		return
	}

	// Allow multiple file upload for certain platforms than need different architectures (ex: Windows)
	// The field name should by default File and for specify an architecture File_x86, File_x64, etc.
	var files = make(map[values.Architecture]*multipart.FileHeader)

	if fhs := r.MultipartForm.File["File"]; len(fhs) > 0 {
		fileHeader := fhs[0]
		files[values.NoArch] = fileHeader
	}
	if fhs := r.MultipartForm.File["File_x86"]; len(fhs) > 0 {
		fileHeader := fhs[0]
		files[values.X86] = fileHeader
	}
	if fhs := r.MultipartForm.File["File_x64"]; len(fhs) > 0 {
		fileHeader := fhs[0]
		files[values.X64] = fileHeader
	}
	if fhs := r.MultipartForm.File["File_ARM64"]; len(fhs) > 0 {
		fileHeader := fhs[0]
		files[values.ARM64] = fileHeader
	}
	if fhs := r.MultipartForm.File["File_ARM"]; len(fhs) > 0 {
		fileHeader := fhs[0]
		files[values.ARM] = fileHeader
	}

	if len(files) == 0 {
		render.Status(r, http.StatusBadRequest)
		return
	}

	_, err = c.service.UploadBuild(
		values.Slug(slug),
		values.Platform(platform),
		buildUploadRequest.BasicInfo,
		files,
		buildUploadRequest.Metadata,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	render.Status(r, http.StatusCreated)
}

func (c *buildController) HandleArtifactDownload(w http.ResponseWriter, r *http.Request) {
	checksum := chi.URLParam(r, "checksum")
	if len(checksum) < 2 {
		http.Error(w, "Invalid checksum", http.StatusBadRequest)
		return
	}

	// Determine file path based on checksum
	subFolder := checksum[:2]
	filePath := filepath.Join(storage.StorageDir, subFolder, checksum)

	// Open the file for reading
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	// Serve the file as a download
	w.Header().Set("Content-Disposition", `attachment; filename="AnyDesk.exe"`)
	w.Header().Set("Content-Type", "application/octet-stream")
	if _, err := io.Copy(w, file); err != nil {
		http.Error(w, "Error serving file", http.StatusInternalServerError)
		return
	}
}
