package utils

import (
	"fmt"
	"time"
)

// FormatTime takes a time.Time and returns a human-readable
// relative time description (e.g., "3 hours ago") or a formatted date if the date
// is older than 2 days.
func FormatTime(date time.Time) string {
	now := time.Now()
	diff := now.Sub(date)

	days := int(diff.Hours() / 24)

	if days > 2 {
		// Return the formatted date for timestamps older than 2 days
		return date.Format("January 2, 2006")
	} else if days > 0 {
		if days == 1 {
			return "1 day ago"
		}
		return fmt.Sprintf("%d days ago", days)
	}

	hours := int(diff.Hours())
	if hours > 0 {
		if hours == 1 {
			return "1 hour ago"
		}
		return fmt.Sprintf("%d hours ago", hours)
	}

	minutes := int(diff.Minutes())
	if minutes > 0 {
		if minutes == 1 {
			return "1 minute ago"
		}
		return fmt.Sprintf("%d minutes ago", minutes)
	}

	seconds := int(diff.Seconds())
	if seconds <= 5 {
		return "just now"
	}
	return fmt.Sprintf("%d seconds ago", seconds)
}
