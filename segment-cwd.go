// Copyright 2016-2017 The psh Authors. All rights reserved.

package psh

import "bytes"

// SegmentCWD implements the cueewnr working directory partial of the prompt.
//
// It renders the current working directory path.
type SegmentCWD struct {
	Data string
}

// NewSegmentCWD creates an instace of SegmentCWD type.
func NewSegmentCWD() *SegmentCWD {
	segment := &SegmentCWD{}
	return segment
}

// Compile ...
func (s *SegmentCWD) Compile() {
	s.Data = "\\w"
}

// Render renders the segment results.
func (s *SegmentCWD) Render() []byte {
	var b bytes.Buffer
	if s.Data != "" {
		b.Write(SetBackground(238))
		b.Write([]byte(" "))
		b.Write([]byte(s.Data))
		b.Write([]byte(" "))
	}
	return b.Bytes()
}
