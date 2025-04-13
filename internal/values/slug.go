package values

import (
	"strings"
	"unicode"
)

type Slug string

// Format converts the Slug text into a URL-friendly slug.
// It converts to lowercase, replaces spaces with hyphens, and removes non-alphanumeric characters.
func (s *Slug) Format() {
	// Convert the name to lowercase
	text := strings.ToLower(string(*s))

	// Create a builder to efficiently build the slug
	var slug strings.Builder

	// Iterate through each character and process
	for _, r := range text {
		// If the character is alphanumeric or a space, include it
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == ' ' {
			slug.WriteRune(r)
		} else {
			// Replace non-alphanumeric characters with nothing (remove them)
			continue
		}
	}

	// Replace spaces with hyphens
	*s = Slug(strings.ReplaceAll(slug.String(), " ", "-"))
}
