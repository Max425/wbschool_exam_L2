package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCut(t *testing.T) {
	tests := []struct {
		name     string
		flags    Flags
		input    string
		expected string
		notSkip  bool
	}{
		{
			name:     "Test with -f flag",
			flags:    Flags{Fields: "1,3", Delimiter: ",", Separated: false},
			input:    "apple,orange,banana",
			expected: "apple,banana",
			notSkip:  true,
		},
		{
			name:     "Test without -f flag",
			flags:    Flags{Delimiter: ",", Separated: false},
			input:    "apple,orange,banana",
			expected: "apple,orange,banana",
			notSkip:  true,
		},
		{
			name:     "Test with -s flag, line with delimiter",
			flags:    Flags{Delimiter: ",", Separated: true},
			input:    "apple,orange,banana",
			expected: "apple,orange,banana",
			notSkip:  true,
		},
		{
			name:     "Test with -s flag, line without delimiter",
			flags:    Flags{Delimiter: " ", Separated: true},
			input:    "apple,orange",
			expected: "",
			notSkip:  false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, notSkip := Cut(&test.flags, test.input)
			assert.Equal(t, test.expected, result)
			assert.Equal(t, test.notSkip, notSkip)
		})
	}
}
