package sql

import (
	"database/sql"
	"github.com/brewbits-co/releasedesk/internal/domains/release"
	"github.com/jmoiron/sqlx"
)

// NewReleaseNotesRepository is the constructor for releaseNotesRepository
func NewReleaseNotesRepository(db *sqlx.DB) release.ReleaseNotesRepository {
	return &releaseNotesRepository{db: db}
}

// releaseNotesRepository is the implementation of release.ReleaseNotesRepository
type releaseNotesRepository struct {
	db *sqlx.DB
}

// Save persists a ReleaseNotes entity and its Changelogs to the database in a single transaction
func (r *releaseNotesRepository) Save(releaseNotes *release.ReleaseNotes) error {
	// Start a transaction
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}

	// Defer a rollback in case anything fails
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Update the ReleaseNotes text in the Releases table
	q := `UPDATE Releases SET ReleaseNotes = :Text WHERE ID = :ReleaseID`
	_, err = tx.NamedExec(q, releaseNotes)
	if err != nil {
		return err
	}

	// Save all changelogs
	for i := range releaseNotes.Changelogs {
		changelog := &releaseNotes.Changelogs[i]
		changelog.ReleaseID = releaseNotes.ReleaseID

		if changelog.ID > 0 {
			// Update existing changelog
			q := `UPDATE Changelogs SET Text = :Text, ChangeType = :ChangeType 
				WHERE ID = :ID AND ReleaseID = :ReleaseID`
			_, err = tx.NamedExec(q, changelog)
		} else {
			// Insert new changelog
			q := `INSERT INTO Changelogs (ReleaseID, Text, ChangeType) 
				VALUES (:ReleaseID, :Text, :ChangeType)`
			var exec sql.Result
			exec, err = tx.NamedExec(q, changelog)
			if err != nil {
				return err
			}

			insertId, _ := exec.LastInsertId()
			changelog.ID = int(insertId)
		}

		if err != nil {
			return err
		}
	}

	// Commit the transaction
	return tx.Commit()
}

// FindByReleaseID retrieves a ReleaseNotes entity by its ReleaseID along with its Changelogs
func (r *releaseNotesRepository) FindByReleaseID(releaseID int) (release.ReleaseNotes, error) {
	var releaseNotes release.ReleaseNotes

	// Get release notes text
	q := `SELECT ID, ReleaseNotes as Text FROM Releases WHERE ID = $1 LIMIT 1`
	err := r.db.QueryRowx(q, releaseID).Scan(&releaseNotes.ReleaseID, &releaseNotes.Text)
	if err != nil {
		return release.ReleaseNotes{}, err
	}

	releaseNotes.ReleaseID = releaseID

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
	q := `SELECT ID, ReleaseID, Text, ChangeType FROM Changelogs WHERE ReleaseID = $1 ORDER BY ID`
	rows, err := r.db.Queryx(q, releaseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var changelogs []release.Changelog

	for rows.Next() {
		var changelog release.Changelog
		if err := rows.StructScan(&changelog); err != nil {
			return nil, err
		}
		changelogs = append(changelogs, changelog)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return changelogs, nil
}
