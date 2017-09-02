// Copyright 2016-2017 The psh Authors. All rights reserved.

package psh

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

// Render renders the segment results.
func (s *SegmentCWD) Render() []byte {
	return []byte("\\w ")
}
