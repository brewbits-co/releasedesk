package values

// Role is an enum type for user roles represented as integers.
type Role int

const (
	// Admin have full access to all applications and administrative tasks.
	Admin Role = iota + 1
	// Manager have full access to all applications.
	Manager // 2
	// Developer have limited access to selected applications.
	Developer
	// Tester have limited access to selected applications.
	Tester
	// Viewer have read access to selected applications.
	Viewer // 3
)
