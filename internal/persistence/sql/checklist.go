package sql

import (
	"github.com/brewbits-co/releasedesk/internal/domains/release"
	"xorm.io/xorm"
)

// NewChecklistRepository is the constructor for checklistRepository
func NewChecklistRepository(engine *xorm.Engine) release.ChecklistRepository {
	return &checklistRepository{engine: engine}
}

// checklistRepository is the implementation of release.ChecklistRepository
type checklistRepository struct {
	engine *xorm.Engine
}

// FindByReleaseID retrieves all ChecklistItem entities for a specific ReleaseID
func (r *checklistRepository) FindByReleaseID(releaseID int) ([]release.ChecklistItem, error) {
	var items []release.ChecklistItem
	err := r.engine.Where("release_id = ?", releaseID).OrderBy("\"order\"").Find(&items)
	if err != nil {
		return nil, err
	}
	return items, nil
}

// Save persists a new ChecklistItem entity to the database
func (r *checklistRepository) Save(item *release.ChecklistItem) error {
	_, err := r.engine.Insert(item)
	return err
}

// Update updates an existing ChecklistItem entity in the database
func (r *checklistRepository) Update(item *release.ChecklistItem) error {
	_, err := r.engine.ID(item.ID).Update(item)
	return err
}

// Delete removes a ChecklistItem entity from the database by its ID
func (r *checklistRepository) Delete(itemID int) error {
	_, err := r.engine.Delete(&release.ChecklistItem{ID: itemID})
	return err
}
