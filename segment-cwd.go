// Copyright 2016-2017 The psh Authors. All rights reserved.

package psh

import "bytes"

// SegmentCWD implements the current working directory partial of the prompt.
//
// It renders the current working directory path.
type SegmentCWD struct {
	Data []byte
}

// NewSegmentCWD creates an instace of SegmentCWD type.
func NewSegmentCWD() *SegmentCWD {
	segment := &SegmentCWD{}
	return segment
}

// Compile ...
func (s *SegmentCWD) Compile() {
	s.Data = []byte("\\w")
}

// Render renders the segment results.
func (s *SegmentCWD) Render() []byte {
	var b bytes.Buffer
	if len(s.Data) != 0 {
		b.Write(SetBackground(238))
		b.Write([]byte(" "))
		b.Write(s.Data)
		b.Write([]byte(" "))
	}
	return b.Bytes()
}
