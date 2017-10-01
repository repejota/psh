// Copyright 2016-2017 The psh Authors. All rights reserved.
package psh

import "testing"

func TestSegmentHostnameCompile(t *testing.T) {
	expected := "\\h"
	segment := NewSegmentHostname()
	segment.Compile()
	if string(segment.Data) != expected {
		t.Fatalf("Compiled data expected to be %s but got %s", expected, segment.Data)
	}
}
