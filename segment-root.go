// Copyright 2016-2017 The psh Authors. All rights reserved.

package psh

import (
	"bytes"
	"fmt"
)

const (
	// SegmentRootBackground is the background color to use
	SegmentRootBackground = 236 // #303030
)

// SegmentRoot implements the root partial of the prompt.
//
// It renders the special character string "\$".
// In bash this results to character '#' if the effective UID is 0,
// otherwise '$'.
type SegmentRoot struct {
	Data []byte
}

// NewSegmentRoot creates an instace of SegmentRoot type.
func NewSegmentRoot() *SegmentRoot {
	segment := &SegmentRoot{}
	return segment
}

// Compile ...
func (s *SegmentRoot) Compile() {
	s.Data = []byte("\\$")
}

// Render renders the segment results.
func (s *SegmentRoot) Render() []byte {
	var b bytes.Buffer
	b.Write(SetBackground(SegmentRootBackground))
	fmt.Fprint(&b, " ")
	b.Write(s.Data)
	fmt.Fprint(&b, " ")
	return b.Bytes()
}
