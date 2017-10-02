// Copyright 2016-2017 The psh Authors. All rights reserved.
package psh

import "testing"

func TestSegmentUsernameCompile(t *testing.T) {
	expected := `\u`
	segment := NewSegmentUsername()
	segment.Compile()
	if string(segment.Data) != expected {
		t.Fatalf("Compiled data expected to be %s but got %s", expected, string(segment.Data))
	}
}

func TestSegmentUsernameRender(t *testing.T) {
	expected := `\[\e[48;5;240m\] \u `
	segment := NewSegmentUsername()
	segment.Compile()
	out := segment.Render()
	if string(out) != expected {
		t.Fatalf("Rendered segment expected to be %q but got %q", expected, string(out))
	}
}
