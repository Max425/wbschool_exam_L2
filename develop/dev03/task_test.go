package main

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	tests := []struct {
		name     string
		flags    *Flags
		input    []string
		expected []string
	}{
		{
			name:     "SortAlphabetically",
			flags:    &Flags{},
			input:    []string{"banana", "apple", "cherry"},
			expected: []string{"apple", "banana", "cherry"},
		},
		{
			name:     "SortNumerically",
			flags:    &Flags{N: true},
			input:    []string{"10", "5", "20"},
			expected: []string{"5", "10", "20"},
		},
		{
			name:     "SortReverse",
			flags:    &Flags{R: true},
			input:    []string{"banana", "apple", "cherry"},
			expected: []string{"cherry", "banana", "apple"},
		},
		{
			name:     "SortUnique",
			flags:    &Flags{U: true},
			input:    []string{"banana", "apple", "banana", "cherry"},
			expected: []string{"apple", "banana", "cherry"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Sort(test.flags, test.input)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}
}
