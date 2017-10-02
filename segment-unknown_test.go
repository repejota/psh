// Copyright 2016-2017 The psh Authors. All rights reserved.
package psh

import "testing"

func TestSegmentUnknown(t *testing.T) {
	expected := ""
	segment := NewSegmentUnknown()
	segment.Compile()
	if string(segment.Data) != expected {
		t.Fatalf("Compiled data expected to be %q but got %q", expected, segment.Data)
	}

	out := segment.Render()
	if string(out) != expected {
		t.Fatalf("Compiled data expected to be %q but got %q", expected, segment.Data)
	}
}
