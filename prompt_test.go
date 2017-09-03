// Copyright 2016-2017 The psh Authors. All rights reserved.
package psh

import (
	"testing"
)

func TestNewPrompt(t *testing.T) {
	prompt := NewPrompt()
	if len(prompt.segments) != 0 {
		t.Errorf("Prompt segments length should be 0 but got %d", len(prompt.segments))
	}
}

func TestAddSegmentUnkown(t *testing.T) {
	prompt := NewPrompt()
	err := prompt.AddSegment("invalidSegmentKey")
	if err.Error() != "segment unknown" {
		t.Errorf("Error should be 'segment unknown' but got %s", err)
	}
}

func TestRender(t *testing.T) {
	prompt := NewPrompt()
	segmentList := make([]string, 0)
	result, err := prompt.Render(segmentList)
	if err != nil {
		t.Errorf("Error should be nil but got %s", err)
	}
	if result == nil {
		t.Errorf("Result should not be nil but got %s", result)
	}
}
