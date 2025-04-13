package values

type Architecture string

const (
	// NoArch represents a file without a specified architecture
	NoArch Architecture = ""
	// X86 represents the 32-bit x86 architecture.
	X86 Architecture = "x86"
	// X64 represents the 64-bit x86 architecture (also known as amd64 or x86_64).
	X64 Architecture = "x64"
	// ARM64 represents the 64-bit ARM architecture (e.g., Apple Silicon, modern ARM-based devices).
	ARM64 Architecture = "ARM64"
	// ARM represents the 32-bit ARM architecture (common in older ARM devices).
	ARM Architecture = "ARM"
)
