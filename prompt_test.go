// Copyright 2016-2017 The psh Authors. All rights reserved.
package psh

import "testing"

func TestGetSegmentsList(t *testing.T) {
	var tests = []struct {
		input    string
		expected []string
	}{
		{"", []string{}},
		{"a,b", []string{"a", "b"}},
		{"a,,c", []string{"a", "c"}},
	}
	for _, tt := range tests {
		output := getSegmentsList(tt.input)
		if len(output) != len(tt.expected) {
			t.Fatalf("Expected to have %d elements but got %d", len(tt.expected), len(output))
		}
	}
}
