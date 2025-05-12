package release

import (
	"github.com/brewbits-co/releasedesk/internal/domains/release"
	"github.com/brewbits-co/releasedesk/pkg/schemas"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
	"strings"
)

// HandleSaveReleaseNotes handles the submission of release notes and changelogs
func (c *releaseController) HandleSaveReleaseNotes(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.PlainText(w, r, err.Error())
		return
	}

	// Get the release ID from the URL
	releaseIDStr := chi.URLParam(r, "id")
	releaseID, err := strconv.Atoi(releaseIDStr)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, schemas.NewErrorResponse("Invalid release ID", nil))
		return
	}

	// Get the release notes text from the form
	releaseNotesText := r.Form.Get("release-notes")

	// Process changelog entries
	var changelogs []release.Changelog

	// Process all changelog entries
	for key, values := range r.Form {
		if len(values) == 0 || values[0] == "" {
			continue
		}

		// Check for existing entries (format: <type>_<id>=<text>)
		for _, changeType := range []string{"Changed", "Added", "Fixed", "Removed", "Security", "Deprecated"} {
			if strings.HasPrefix(key, changeType+"_") {
				// Extract the changelog ID
				parts := strings.SplitN(key, "_", 2)
				if len(parts) != 2 {
					continue
				}

				changelogIDStr := parts[1]
				changelogID, err := strconv.Atoi(changelogIDStr)
				if err != nil {
					continue
				}

				// Create the changelog entry
				changelog := release.Changelog{
					ID:         changelogID,
					ReleaseID:  releaseID,
					Text:       values[0],
					ChangeType: release.ChangeType(changeType),
				}
				changelogs = append(changelogs, changelog)
				break
			}
		}

		// Check for new entries (format: <type>=<text>)
		switch key {
		case "Changed", "Added", "Fixed", "Removed", "Security", "Deprecated":
			changelog := release.Changelog{
				ReleaseID:  releaseID,
				Text:       values[0],
				ChangeType: release.ChangeType(strings.Title(key)),
			}
			changelogs = append(changelogs, changelog)
		}
	}

	// Save the release notes and changelogs
	_, err = c.service.SaveReleaseNotes(releaseID, releaseNotesText, changelogs)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, schemas.NewErrorResponse("Failed to save release notes", []string{err.Error()}))
		return
	}

	// Redirect to the release notes page
	http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
}
