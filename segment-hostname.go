// Copyright 2016-2017 The psh Authors. All rights reserved.

package psh

import "bytes"

// SegmentHostname implements the hostname partial of the prompt.
//
// It renders the current hostname.
type SegmentHostname struct {
	Data string
}

// NewSegmentHostname creates an instace of SegmentHostname type.
func NewSegmentHostname() *SegmentHostname {
	segment := &SegmentHostname{}
	return segment
}

// Compile ...
func (s *SegmentHostname) Compile() {
	s.Data = "\\h"
}

// Render renders the segment results.
func (s *SegmentHostname) Render() []byte {
	var b bytes.Buffer
	if s.Data != "" {
		b.Write(SetBackground(237))
		b.Write([]byte(" "))
		b.Write([]byte(s.Data))
		b.Write([]byte(" "))
	}
	return b.Bytes()
}
