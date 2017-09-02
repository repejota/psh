// Copyright 2016-2017 The psh Authors. All rights reserved.

package psh

import (
	"bytes"
	"errors"
)

// Prompt is the type that defines a shell prompt.
type Prompt struct {
	segments map[string]Segment
}

// NewPrompt creates a new Prompt type instance.
func NewPrompt(segmentsList []string) *Prompt {
	p := &Prompt{}
	p.segments = make(map[string]Segment, len(segmentsList))
	return p
}

// AddSegment appends a segment to the prompt.
func (p *Prompt) AddSegment(key string) error {
	switch key {
	case "root":
		p.segments[key] = NewSegmentRoot()
		return nil
	case "username":
		p.segments[key] = NewSegmentUsername()
		return nil
	case "hostname":
		p.segments[key] = NewSegmentHostname()
		return nil
	case "cwd":
		p.segments[key] = NewSegmentCWD()
		return nil
	default:
		return errors.New("segment unknown")
	}
}

// Render compiles all the segments of the prompt and concatenate its results.
//
// Receives the list of segments to render them in the correct order.
func (p *Prompt) Render(segmentsList []string) ([]byte, error) {
	var b bytes.Buffer
	// Render all segments
	for _, segmentKey := range segmentsList {
		segment := p.segments[segmentKey]
		b.Write(segment.Render())
	}
	// Reset foreground and background colors
	b.Write(ResetFgBg())
	b.WriteRune(' ')
	return b.Bytes(), nil
}
