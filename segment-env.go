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
	envName := ""
	envName, ok := os.LookupEnv("ENV")
	if ok {
		envName = os.Getenv("ENV")
	}
	s.Data = []byte(envName)
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
