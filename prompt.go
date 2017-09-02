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

// String implements Stringer interface and allows the type to describe itself
// as an string result.
func (p *Prompt) String() string {
	var b bytes.Buffer
	for _, s := range p.segments {
		b.WriteString(s.String())
	}
	return b.String()
}
