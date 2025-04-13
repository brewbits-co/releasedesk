package fields

// Extendable represents a structure that supports additional metadata fields.
type Extendable struct {
	// Metadata is a key-value store for additional attributes or properties
	// that extend the base entity.
	Metadata map[string]string
}
