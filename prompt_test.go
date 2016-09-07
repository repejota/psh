package main

import "testing"

func TestPrompt(t *testing.T) {
	var prompt Prompt
	if prompt.Prompt != "" {
		t.Error("Expected empty string, but got ", prompt.Prompt)
	}
}
