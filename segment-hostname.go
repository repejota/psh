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
