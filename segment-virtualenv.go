// Copyright 2016-2017 The psh Authors. All rights reserved.

package psh

import (
	"bytes"
	"fmt"
	"os"
	"path"
)

const (
	// SegmentVirtualEnvBackground is the background color to use
	SegmentVirtualEnvBackground = 22 // #585858
)

// SegmentVirtualEnv implements the python virtual env partial of the prompt.
//
// Renders the current virtalenv python environment name if it is available.
type SegmentVirtualEnv struct {
	Data []byte
}

// NewSegmentVirtualEnv creates an instace of SegmentVirtualEnv type.
func NewSegmentVirtualEnv() *SegmentVirtualEnv {
	segment := &SegmentVirtualEnv{}
	return segment
}

// Compile collects the data for this segment.
func (s *SegmentVirtualEnv) Compile() {
	s.Data = []byte("")
	virtualEnvPath, ok := os.LookupEnv("VIRTUAL_ENV")
	if ok {
		virtualEnvName := path.Base(virtualEnvPath)
		s.Data = []byte(virtualEnvName)
	}
}

// Render renders the segment results.
func (s *SegmentVirtualEnv) Render() []byte {
	var b bytes.Buffer
	if len(s.Data) > 0 {
		b.Write(SetBackground(SegmentVirtualEnvBackground))
		fmt.Fprint(&b, " ")
		fmt.Fprint(&b, " ")
		b.Write(s.Data)
		fmt.Fprint(&b, " ")
	}
	return b.Bytes()
}
