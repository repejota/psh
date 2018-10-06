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
