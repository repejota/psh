// Copyright 2016-2017 The psh Authors. All rights reserved.

package psh

import (
	"bytes"
	"strings"
)

// Prompt type that defines the shell prompt structure.
type Prompt struct {
	keys     []string
	segments map[string]Segment
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
func (p *Prompt) Compile() error {
	for _, v := range p.keys {
		p.addSegment(v)
	}
	return nil
}

// Render compiles all the segments on the prompt and concatenates its results.
//
// Receives the list of segments to render them in the correct order.
func (p *Prompt) Render() ([]byte, error) {
	var b bytes.Buffer
	// Render all segments
	for _, v := range p.keys {
		segmentInstance := p.segments[v]
		b.Write(segmentInstance.Render())
	}
	// Reset foreground and background colors
	b.Write(ResetFgBg())
	b.WriteRune(' ')
	return b.Bytes(), nil
}

// addSegment adds a segment instance by its name to the prompt.
// TODO : Should be done in paralel
func (p *Prompt) addSegment(key string) {
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
	default:
		segment = NewSegmentUnknown()
	}
	segment.Compile()
	p.segments[key] = segment
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
