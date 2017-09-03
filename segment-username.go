// Copyright 2016-2017 The psh Authors. All rights reserved.

package psh

import "bytes"

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

// Compile renders the segment results.
func (s *SegmentUsername) Compile() []byte {
	var b bytes.Buffer
	b.Write(SetBackground(236))
	b.Write([]byte(" \\u "))
	return b.Bytes()
}
