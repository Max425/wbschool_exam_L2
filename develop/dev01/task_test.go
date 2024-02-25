package main

import "testing"

func TestNtpServerIsCorrect(t *testing.T) {
	err := getTime("Error text")
	if err == nil {
		t.Error(err)
	}
}
