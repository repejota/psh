// Copyright 2016-2017 The psh Authors. All rights reserved.

package psh

import (
	"bytes"
)

// SegmentRoot implements the root partial of the prompt.
//
// It renders the character '#' if the effective UID of the current user is 0,
// otherwise renders the character '$'.
type SegmentRoot struct {
	Data []byte
}

// NewSegmentRoot creates an instace of SegmentRoot type.
func NewSegmentRoot() *SegmentRoot {
	segment := &SegmentRoot{}
	return segment
}

// Compile ...
func (s *SegmentRoot) Compile() {
	s.Data = []byte("\\$")
}

// Render renders the segment results.
func (s *SegmentRoot) Render() []byte {
	var b bytes.Buffer
	if len(s.Data) != 0 {
		b.Write(SetBackground(236))
		b.Write([]byte(" "))
		b.Write(s.Data)
		b.Write([]byte(" "))
	}
	return b.Bytes()
}
