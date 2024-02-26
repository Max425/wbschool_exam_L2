package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatch(t *testing.T) {
	tests := []struct {
		name           string
		flags          Flags
		input          string
		expectedResult bool
	}{
		{
			name: "Match with IgnoreCase enabled",
			flags: Flags{
				Pattern:    "example",
				Fixed:      false,
				IgnoreCase: true,
				Invert:     false,
			},
			input:          "This is an Example text",
			expectedResult: true,
		},
		{
			name: "Match with IgnoreCase disabled",
			flags: Flags{
				Pattern:    "example",
				Fixed:      false,
				IgnoreCase: false,
				Invert:     false,
			},
			input:          "This is an example text",
			expectedResult: true,
		},
		{
			name: "Exact match with Fixed enabled",
			flags: Flags{
				Pattern:    "example",
				Fixed:      true,
				IgnoreCase: false,
				Invert:     false,
			},
			input:          "This is an example text",
			expectedResult: true,
		},
		{
			name: "Inverted match",
			flags: Flags{
				Pattern:    "example",
				Fixed:      false,
				IgnoreCase: false,
				Invert:     true,
			},
			input:          "This is an example text",
			expectedResult: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := match(&test.flags, test.input)
			assert.Equal(t, test.expectedResult, result, "Unexpected match result")
		})
	}
}
