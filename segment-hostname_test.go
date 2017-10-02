// Copyright 2016-2017 The psh Authors. All rights reserved.
package psh

import "testing"

func TestSegmentHostnameCompile(t *testing.T) {
	expected := `\h`
	segment := NewSegmentHostname()
	segment.Compile()
	if string(segment.Data) != expected {
		t.Fatalf("Compiled data expected to be %s but got %s", expected, string(segment.Data))
	}
}

func TestSegmentHostnameRender(t *testing.T) {
	expected := `\[\e[48;5;238m\] \h `
	segment := NewSegmentHostname()
	segment.Compile()
	out := segment.Render()
	if string(out) != expected {
		t.Fatalf("Rendered segment expected to be %q but got %q", expected, string(out))
	}
}
