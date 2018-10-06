// Copyright 2016-2017 The psh Authors. All rights reserved.

package psh

import (
	"bytes"
	"fmt"
	"os"
)

const (
	// SegmentEnvBackground is the background color to use
	SegmentEnvBackground = 166
)

// SegmentEnv implements the python virtual env partial of the prompt.
//
// Renders the current virtalenv python environment name if it is available.
type SegmentEnv struct {
	Data []byte
}

// NewSegmentEnv creates an instace of SegmentEnv type.
func NewSegmentEnv() *SegmentEnv {
	segment := &SegmentEnv{}
	return segment
}

// Compile collects the data for this segment.
func (s *SegmentEnv) Compile() {
	s.Data = []byte("")
	EnvName, ok := os.LookupEnv("ENV")
	if ok {
		EnvName = os.Getenv("ENV")
		s.Data = []byte(EnvName)
	}
}

// Render renders the segment results.
func (s *SegmentEnv) Render() []byte {
	var b bytes.Buffer
	if len(s.Data) > 0 {
		b.Write(SetBackground(SegmentEnvBackground))
		fmt.Fprint(&b, " ")
		fmt.Fprint(&b, " ")
		b.Write(s.Data)
		fmt.Fprint(&b, " ")
	}
	return b.Bytes()
}
