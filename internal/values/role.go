package values

// Role is an enum type for user roles represented as integers.
type Role int

const (
	// Admin has full access to all applications and administrative tasks.
	Admin Role = iota + 1
	// Manager has full access to all applications.
	Manager // 2
	// Developer have limited access to selected applications.
	Developer
	// Tester has limited access to selected applications.
	Tester
	// Viewer has read access to selected applications.
	Viewer // 3
)
