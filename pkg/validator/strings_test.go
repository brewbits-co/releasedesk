package validator

import (
	"testing"
)

func TestIsEmail(t *testing.T) {
	tests := []struct {
		email    string
		expected bool
	}{
		{"john.doe@example.com", true},
		{"invalid.email", false},
		{"missing@domain", false},
		{"text <hello@world.com>", false},
		{"email@example.com (Joe Smith)", false},
		{"", false},
	}

	for _, test := range tests {
		result := IsEmail(test.email)
		if result != test.expected {
			t.Errorf("For email %s, expected %t but got %t", test.email, test.expected, result)
		}
	}
}

func TestIsPasswordStrong(t *testing.T) {
	tests := []struct {
		password string
		expected bool
	}{
		{"StrongP@ssw0rd", true},
		{"weak", false},
		{"NoSymbol123ButNumber", true},
		{"OnlyText", false},
		{"lowerc@ase1", false},
		{"UPPERC@", false},
		{"", false},
	}

	for _, test := range tests {
		result := IsPasswordStrong(test.password)
		if result != test.expected {
			t.Errorf("For password %s, expected %t but got %t", test.password, test.expected, result)
		}
	}
}

func TestIsAnyEmpty(t *testing.T) {
	tests := []struct {
		values   []string
		expected bool
	}{
		{[]string{"", "hello", "world"}, true},
		{[]string{"foo", "bar", "baz"}, false},
		{[]string{""}, true},
		{[]string{}, false},
	}

	for _, test := range tests {
		result := IsAnyEmpty(test.values...)
		if result != test.expected {
			t.Errorf("For values %v, expected %t but got %t", test.values, test.expected, result)
		}
	}
}
