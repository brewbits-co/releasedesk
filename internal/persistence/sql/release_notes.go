package sql

import (
	"database/sql"
	"github.com/brewbits-co/releasedesk/internal/domains/release"
	"xorm.io/xorm"
)

// NewReleaseNotesRepository is the constructor for releaseNotesRepository
func NewReleaseNotesRepository(engine *xorm.Engine) release.ReleaseNotesRepository {
	return &releaseNotesRepository{engine: engine}
}

// releaseNotesRepository is the implementation of release.ReleaseNotesRepository
type releaseNotesRepository struct {
	engine *xorm.Engine
}

// Save persists a ReleaseNotes entity and its Changelogs to the database in a single transaction
func (r *releaseNotesRepository) Save(releaseNotes *release.ReleaseNotes) error {
	// Start a transaction
	session := r.engine.NewSession()
	defer session.Close()

	err := session.Begin()
	if err != nil {
		return err
	}

	// Defer a rollback in case anything fails
	defer func() {
		if err != nil {
			session.Rollback()
		}
	}()

	// Update the ReleaseNotes text in the Releases table
	if _, err = session.Table("release").
		Where("id = ?", releaseNotes.ReleaseID).
		Update(map[string]interface{}{
			"release_notes": releaseNotes.Text,
		}); err != nil {
		return err
	}

	// Save all changelogs
	for i := range releaseNotes.Changelogs {
		changelog := &releaseNotes.Changelogs[i]
		changelog.ReleaseID = releaseNotes.ReleaseID

		if changelog.ID > 0 {
			// Update existing changelog
			if _, err = session.Table("changelog").
				Where("id = ? AND release_id = ?", changelog.ID, changelog.ReleaseID).
				Update(map[string]interface{}{
					"text":        changelog.Text,
					"change_type": changelog.ChangeType,
				}); err != nil {
				return err
			}
		} else {
			// Insert new changelog
			if _, err = session.Table("changelog").
				Insert(changelog); err != nil {
				return err
			}
		}
	}

	// Commit the transaction
	return session.Commit()
}

// FindByReleaseID retrieves a ReleaseNotes entity by its ReleaseID along with its Changelogs
func (r *releaseNotesRepository) GetByReleaseID(releaseID int) (release.ReleaseNotes, error) {
	var releaseNotes release.ReleaseNotes

	// Get release notes text
	has, err := r.engine.Table("release").
		Select("id as release_id, release_notes as text").
		Where("id = ?", releaseID).
		Get(&releaseNotes)
	if err != nil {
		return release.ReleaseNotes{}, err
	}
	if !has {
		return release.ReleaseNotes{}, sql.ErrNoRows
	}

	// Get changelogs
	changelogs, err := r.FindChangelogsByReleaseID(releaseID)
	if err != nil {
		return release.ReleaseNotes{}, err
	}

	releaseNotes.Changelogs = changelogs
	return releaseNotes, nil
}

// FindChangelogsByReleaseID retrieves all Changelog entities for a specific ReleaseID
func (r *releaseNotesRepository) FindChangelogsByReleaseID(releaseID int) ([]release.Changelog, error) {
	session := r.engine.Where("release_id = ?", releaseID).OrderBy("ID")
	var changelogs []release.Changelog
	err := session.Find(&changelogs)
	if err != nil {
		return nil, err
	}
	return changelogs, nil
}
