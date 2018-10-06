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
	// SegmentRootBackground is the background color to use
	SegmentRootBackground = 236 // #303030
)

// SegmentRoot implements the root partial of the prompt.
//
// It renders the special character string "\$".
// In bash this results to character '#' if the effective UID is 0,
// otherwise '$'.
//
// TODO:
// * Change background color if the last command exit status != 0
type SegmentRoot struct {
	Data []byte
}

// NewSegmentRoot creates an instace of SegmentRoot type.
func NewSegmentRoot() *SegmentRoot {
	segment := &SegmentRoot{}
	return segment
}

// Compile collects the data for this segment.
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
