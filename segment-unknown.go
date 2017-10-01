// Copyright 2016-2017 The psh Authors. All rights reserved.

package psh

// SegmentUnknown implements the not-found segment partial of the prompt.
//
// It renders an empty string.
type SegmentUnknown struct {
	Data []byte
}

// NewSegmentUnknown creates an instace of SegmentUnknown type.
func NewSegmentUnknown() *SegmentUnknown {
	segment := &SegmentUnknown{}
	return segment
}

// Compile ...
func (s *SegmentUnknown) Compile() {
	s.Data = nil
}

// Render renders the segment results.
func (s *SegmentUnknown) Render() []byte {
	return nil
}
