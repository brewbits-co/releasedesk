package values

// Role is an enum type for user roles represented as integers.
type Role int

const (
	// Admin have full access to all products and administrative tasks.
	Admin Role = iota + 1
	// Manager have full access to all products.
	Manager // 2
	// Developer have limited access to selected products.
	Developer
	// Tester have limited access to selected products.
	Tester
	// Viewer have read access to selected products.
	Viewer // 3
)
