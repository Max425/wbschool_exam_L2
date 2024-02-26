package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindAnagrams(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected map[string][]string
	}{
		{
			name:  "SimpleCase",
			input: []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"},
			expected: map[string][]string{
				"акптя":  {"пятак", "пятка", "тяпка"},
				"иклост": {"листок", "слиток", "столик"},
			},
		},
		{
			name:     "NoAnagrams",
			input:    []string{"apple", "banana", "cherry"},
			expected: map[string][]string{},
		},
		{
			name:  "MixedCase",
			input: []string{"Кот", "ток", "кто", "КТО", "отк"},
			expected: map[string][]string{
				"кот": {"Кот", "ток", "кто", "КТО", "отк"},
			},
		},
		{
			name:     "OneWordSets",
			input:    []string{"hello", "world", "program"},
			expected: map[string][]string{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := FindAnagrams(test.input)
			assert.Equal(t, test.expected, result, "Unexpected result")
		})
	}
}
