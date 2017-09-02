// Copyright 2016-2017 The psh Authors. All rights reserved.

package psh

// SegmentUsername implements the username partial of the prompt.
//
// It renders the current username.
type SegmentUsername struct {
}

// NewSegmentUsername creates an instace of SegmentUsername type.
func NewSegmentUsername() *SegmentUsername {
	segment := &SegmentUsername{}
	return segment
}

// Render renders the segment results.
func (s *SegmentUsername) Render() []byte {
	return []byte("\\u ")
}
