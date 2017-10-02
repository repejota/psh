// Copyright 2016-2017 The psh Authors. All rights reserved.

package psh

import (
	"bytes"
	"fmt"
)

const (
	// SegmentUsernameBackground is the background color to use
	SegmentUsernameBackground = 240 // #585858
)

// SegmentUsername implements the username partial of the prompt.
//
// Renders the current username.
//
// TODO:
// * If the user is root (uid=0) use a different background.
type SegmentUsername struct {
	Data []byte
}

// NewSegmentUsername creates an instace of SegmentUsername type.
func NewSegmentUsername() *SegmentUsername {
	segment := &SegmentUsername{}
	return segment
}

// Compile collects the data for this segment.
func (s *SegmentUsername) Compile() {
	s.Data = []byte("\\u")
}

// Render renders the segment results.
func (s *SegmentUsername) Render() []byte {
	var b bytes.Buffer
	b.Write(SetBackground(SegmentUsernameBackground))
	fmt.Fprint(&b, " ")
	b.Write(s.Data)
	fmt.Fprint(&b, " ")
	return b.Bytes()
}
