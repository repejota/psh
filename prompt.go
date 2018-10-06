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
	"strings"
	"sync"
)

// Prompt type that defines the shell prompt structure.
type Prompt struct {
	keys         []string
	segmentsLock sync.Mutex
	segments     map[string]Segment
}

// NewPrompt creates a new prompt and returns a pointer to its instance.
func NewPrompt(segments string) *Prompt {
	keys := getSegmentsList(segments)
	prompt := &Prompt{
		keys:     keys,
		segments: make(map[string]Segment),
	}
	return prompt
}

// Compile traverses all available segments. It also process each segment to
// get all the data needed to render.
func (p *Prompt) Compile() {
	var wg sync.WaitGroup
	wg.Add(len(p.keys))
	for _, v := range p.keys {
		go func(v string) {
			defer wg.Done()
			segment := p.getSegment(v)
			segment.Compile()
			p.segmentsLock.Lock()
			p.segments[v] = segment
			p.segmentsLock.Unlock()
		}(v)
	}
	wg.Wait()
}

// Render compiles all the segments on the prompt and concatenates its results.
//
// Receives the list of segments to render them in the correct order.
func (p *Prompt) Render() []byte {
	var b bytes.Buffer
	// Render all segments
	for _, v := range p.keys {
		segmentInstance := p.segments[v]
		b.Write(segmentInstance.Render())
	}
	// Reset foreground and background colors
	b.Write(ResetFgBg())
	b.WriteRune(' ')
	return b.Bytes()
}

// getSegment gets a segment instance by its name to the prompt.
func (p *Prompt) getSegment(key string) Segment {
	var segment Segment
	switch key {
	case "root":
		segment = NewSegmentRoot()
	case "username":
		segment = NewSegmentUsername()
	case "hostname":
		segment = NewSegmentHostname()
	case "cwd":
		segment = NewSegmentCWD()
	case "git":
		segment = NewSegmentGit()
	case "virtualenv":
		segment = NewSegmentVirtualEnv()
	case "env":
		segment = NewSegmentEnv()
	default:
		segment = NewSegmentUnknown()
	}
	return segment
}

func getSegmentsList(segments string) []string {
	keys := strings.Split(segments, ",")
	for i, v := range keys {
		if v == "" {
			keys = append(keys[:i], keys[i+1:]...)
		}
	}
	return keys
}
