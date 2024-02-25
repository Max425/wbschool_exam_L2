package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnpacked(t *testing.T) {
	testers := map[string]string{
		"a4bc2d5e":  "aaaabccddddde",
		"abcd":      "abcd",
		"":          "",
		"qwe\\4\\5": "qwe45",
		"qwe\\45":   "qwe44444",
		"qwe\\\\5":  "qwe\\\\\\\\\\",
	}
	for k, v := range testers {
		actual, _ := Unpacked(k)
		if actual != v {
			t.Error("Excpected:", v, "Actual:", actual)
		}
	}
}

func TestUnpackedError(t *testing.T) {
	tester := "45432"
	_, err := Unpacked(tester)
	assert.NotNil(t, err)
}
