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
