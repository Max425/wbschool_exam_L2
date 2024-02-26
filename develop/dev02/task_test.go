package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnpackString(t *testing.T) {
	tests := []struct {
		input          string
		expected       string
		expectedErrMsg string
	}{
		{"a4bc2d5e", "aaaabccddddde", ""},
		{"abcd", "abcd", ""},
		{"", "", ""},
		{"qwe\\4\\5", "qwe45", ""},
		{"qwe\\45", "qwe44444", ""},
		{"qwe\\\\5", "qwe\\\\\\\\\\", ""},
		{"45", "", "некорректная строка"},
	}

	for _, test := range tests {
		actual, err := UnpackString(test.input)
		if test.expectedErrMsg != "" {
			assert.NotNil(t, err)
			assert.Equal(t, test.expectedErrMsg, err.Error())
		} else {
			assert.Nil(t, err)
			assert.Equal(t, test.expected, actual)
		}
	}
}
