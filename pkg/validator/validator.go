package validator

// Validator defines a set of methods for validating the state of a struct
// before certain operations are performed.
type Validator interface {
	// IsValid checks if the struct's fields adhere to the predefined business rules.
	// Returns an error if any validation rule is violated.
	IsValid() error
}

// BaseValidator provides default no-op implementations for the Validator interface methods.
// It can be embedded in other structs to provide default validation behavior.
type BaseValidator struct{}

// IsValid is a no-op implementation that always returns nil.
// Override this method in the embedding struct to implement custom validation logic.
func (b *BaseValidator) IsValid() error {
	return nil
}
