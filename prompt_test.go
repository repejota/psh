package main

import "testing"

func TestGetPrompt(t *testing.T) {
	prompt := getPrompt()
	if prompt != "\\u@\\H:\\w\\$ " {
		t.Error("\\u@\\H:\\w\\$ ", prompt)
	}
}
