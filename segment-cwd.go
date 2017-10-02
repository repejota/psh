// Copyright 2016-2017 The psh Authors. All rights reserved.

package psh

import (
	"bytes"
	"fmt"
)

const (
	// SegmentCWDBackground is the background color to use
	SegmentCWDBackground = 237 // #3a3a3a
)

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

// Compile collects the data for this segment.
func (s *SegmentCWD) Compile() {
	s.Data = []byte("\\w")
}

// Render renders the segment results.
func (s *SegmentCWD) Render() []byte {
	var b bytes.Buffer
	if len(s.Data) != 0 {
		b.Write(SetBackground(SegmentCWDBackground))
		fmt.Fprint(&b, " ")
		b.Write(s.Data)
		fmt.Fprint(&b, " ")
	}
	return b.Bytes()
}
