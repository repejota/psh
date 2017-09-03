// Copyright 2016-2017 The psh Authors. All rights reserved.

package psh

import "bytes"

// SegmentCWD implements the cueewnr working directory partial of the prompt.
//
// It renders the current working directory path.
type SegmentCWD struct {
}

// NewSegmentCWD creates an instace of SegmentCWD type.
func NewSegmentCWD() *SegmentCWD {
	segment := &SegmentCWD{}
	return segment
}

// Compile renders the segment results.
func (s *SegmentCWD) Compile() []byte {
	var b bytes.Buffer
	b.Write(SetBackground(238))
	b.Write([]byte(" \\w "))
	return b.Bytes()
}
