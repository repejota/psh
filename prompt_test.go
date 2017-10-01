// Copyright 2016-2017 The psh Authors. All rights reserved.
package psh

import (
	"testing"
)

func TestGetSegmentsList(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{"", 0},
		{"a,b", 2},
		{"a,,c", 2},
	}
	for _, tt := range tests {
		output := getSegmentsList(tt.input)
		if len(output) != tt.expected {
			t.Fatalf("Expected to have %d elements but got %d", tt.expected, len(output))
		}
	}
}

func TestNewPrompt(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{"", 0},
		{"a,b", 2},
		{"a,,c", 2},
	}
	for _, tt := range tests {
		prompt := NewPrompt(tt.input)
		if len(prompt.keys) != tt.expected {
			t.Fatalf("Prompt expected to have %d keys but got %d", tt.expected, len(prompt.keys))
		}
	}
}
