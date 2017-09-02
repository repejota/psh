// Copyright 2016-2017 The psh Authors. All rights reserved.

package psh

import "bytes"

// SegmentHostname implements the hostname partial of the prompt.
//
// It renders the current hostname.
type SegmentHostname struct {
}

// NewSegmentHostname creates an instace of SegmentHostname type.
func NewSegmentHostname() *SegmentHostname {
	segment := &SegmentHostname{}
	return segment
}

// Render renders the segment results.
func (s *SegmentHostname) Render() []byte {
	var b bytes.Buffer
	b.Write(SetBackground(237))
	b.Write([]byte(" \\h "))
	return b.Bytes()
}
