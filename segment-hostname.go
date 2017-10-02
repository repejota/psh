// Copyright 2016-2017 The psh Authors. All rights reserved.

package psh

import (
	"bytes"
	"fmt"
)

const (
	// SegmentHostnameBackground is the background color to use
	SegmentHostnameBackground = 238 // #444444
)

// SegmentHostname implements the hostname partial of the prompt.
//
// It renders the current hostname up to the first '.'
//
// TODO:
// * Add support for long/conmplete hostname.
type SegmentHostname struct {
	Data []byte
}

// NewSegmentHostname creates an instace of SegmentHostname type.
func NewSegmentHostname() *SegmentHostname {
	segment := &SegmentHostname{}
	return segment
}

// Compile collects the data for this segment.
func (s *SegmentHostname) Compile() {
	s.Data = []byte("\\h")
}

// Render renders the segment results.
func (s *SegmentHostname) Render() []byte {
	var b bytes.Buffer
	if len(s.Data) != 0 {
		b.Write(SetBackground(SegmentHostnameBackground))
		fmt.Fprint(&b, " ")
		b.Write(s.Data)
		fmt.Fprint(&b, " ")
	}
	return b.Bytes()
}
