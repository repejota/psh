// Copyright 2016-2017 The psh Authors. All rights reserved.
package psh

import "testing"

func TestSegmentCWDCompile(t *testing.T) {
	expected := `\w`
	segment := NewSegmentCWD()
	segment.Compile()
	if string(segment.Data) != expected {
		t.Fatalf("Compiled data expected to be %q but got %q", expected, string(segment.Data))
	}
}

func TestSegmentCWDRender(t *testing.T) {
	expected := `\[\e[48;5;237m\] \w `
	segment := NewSegmentCWD()
	segment.Compile()
	out := segment.Render()
	if string(out) != expected {
		t.Fatalf("Rendered segment expected to be %q but got %q", expected, string(out))
	}
}
