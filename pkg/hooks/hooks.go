package hooks

// Hooks defines a set of lifecycle methods that can be implemented by structs
// to perform actions before or after certain events, such as saving or updating data.
type Hooks interface {
	// BeforeSave is called before saving the record (either create or update).
	BeforeSave() error
	// BeforeCreate is called before a new record is created.
	BeforeCreate() error
	// AfterCreate is called after a new record has been successfully created.
	AfterCreate() error
	// BeforeUpdate is called before an existing record is updated.
	BeforeUpdate() error
	// AfterUpdate is called after an existing record has been successfully updated.
	AfterUpdate() error
	// AfterSave is called after the record has been saved (either create or update).
	AfterSave() error
}

// BaseHooks provides default no-op implementations of the Hooks interface methods.
// It can be embedded in other structs to avoid having to implement every hook manually.
type BaseHooks struct{}

// BeforeCreate is a no-op implementation of the BeforeCreate hook.
// Override this method in the embedding struct to add custom behavior before creating a record.
func (b *BaseHooks) BeforeCreate() error {
	return nil
}

// AfterCreate is a no-op implementation of the AfterCreate hook.
// Override this method in the embedding struct to add custom behavior after creating a record.
func (b *BaseHooks) AfterCreate() error {
	return nil
}

// BeforeUpdate is a no-op implementation of the BeforeUpdate hook.
// Override this method in the embedding struct to add custom behavior before updating a record.
func (b *BaseHooks) BeforeUpdate() error {
	return nil
}

// AfterUpdate is a no-op implementation of the AfterUpdate hook.
// Override this method in the embedding struct to add custom behavior after updating a record.
func (b *BaseHooks) AfterUpdate() error {
	return nil
}
