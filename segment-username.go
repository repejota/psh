// Copyright 2016-2017 The psh Authors. All rights reserved.

package psh

import "bytes"

// SegmentUsername implements the username partial of the prompt.
//
// It renders the current username.
type SegmentUsername struct {
	Data string
}

// NewSegmentUsername creates an instace of SegmentUsername type.
func NewSegmentUsername() *SegmentUsername {
	segment := &SegmentUsername{}
	return segment
}

// Compile ...
func (s *SegmentUsername) Compile() {
	s.Data = "\\u"
}

// Render renders the segment results.
func (s *SegmentUsername) Render() []byte {
	var b bytes.Buffer
	if s.Data != "" {
		b.Write(SetBackground(236))
		b.Write([]byte(" "))
		b.Write([]byte(s.Data))
		b.Write([]byte(" "))
	}
	return b.Bytes()
}
