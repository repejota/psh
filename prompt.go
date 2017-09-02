package psh

import (
	"bytes"
)

// Prompt is the type that defines a shell prompt.
type Prompt struct {
	segments []Segment
}

// NewPrompt creates a new Prompt type instance.
func NewPrompt() *Prompt {
	p := &Prompt{}
	p.segments = make([]Segment, 0)
	return p
}

// AddSegment appends a segment to the prompt.
func (p *Prompt) AddSegment(s Segment) error {
	p.segments = append(p.segments, s)
	return nil
}

// Render compiles all the segments of the prompt and concatenate its results.
func (p *Prompt) Render() []byte {
	var b bytes.Buffer
	// Render all segments
	for _, s := range p.segments {
		b.Write(s.Render())
	}
	// Reset foreground and background colors
	b.Write(ResetFgAndBg())
	return b.Bytes()
}
