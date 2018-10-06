// Copyright 2016-2018, psh project Authors.
//
// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with this
// work for additional information regarding copyright ownership.  The ASF
// licenses this file to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.  See the
// License for the specific language governing permissions and limitations
// under the License.

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
