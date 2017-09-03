// Copyright 2016-2017 The psh Authors. All rights reserved.

package psh

import (
	"strings"
)

// Prompt is the type that defines a shell prompt.
type Prompt struct {
	segmentKeys []string
	segments    map[string]Segment
}

// NewPrompt creates a new Prompt type instance.
func NewPrompt(segmentsFlag string) *Prompt {
	p := &Prompt{}
	p.segments = make(map[string]Segment)
	p.segmentKeys = p.parseSegmentsFlag(segmentsFlag)
	return p
}

// parseSegmentsFlag ...
func (p *Prompt) parseSegmentsFlag(segmentsFlag string) []string {
	// Get the segments the prompt will use
	segmentKeys := strings.Split(segmentsFlag, ",")
	// Clean list
	for key, segmentName := range segmentKeys {
		if segmentName == "" {
			segmentKeys = segmentKeys[:key+copy(segmentKeys[key:], segmentKeys[key+1:])]
		}
	}
	return segmentKeys
}

// Compile ...
func (p *Prompt) Compile() error {
	return nil
}

// Render ...
func (p *Prompt) Render() (string, error) {
	return "", nil
}

/*
// AddSegment appends a segment to the prompt.
func (p *Prompt) AddSegment(key string) error {
	switch key {
	case "root":
		p.segments[key] = NewSegmentRoot()
	case "username":
		p.segments[key] = NewSegmentUsername()
	case "hostname":
		p.segments[key] = NewSegmentHostname()
	case "cwd":
		p.segments[key] = NewSegmentCWD()
	default:
		return errors.New("segment unknown")
	}
	return nil
}


// Render compiles all the segments on the prompt and concatenates its results.
//
// Receives the list of segments to render them in the correct order.
func (p *Prompt) Render(segmentsList []string) ([]byte, error) {
	var b bytes.Buffer
	// Render all segments
	for _, segmentKey := range segmentsList {
		segmentInstance := p.segments[segmentKey]
		b.Write(segmentInstance.Compile())
	}
	// Reset foreground and background colors
	b.Write(ResetFgBg())
	b.WriteRune(' ')
	return b.Bytes(), nil
}
*/
