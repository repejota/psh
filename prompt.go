// Copyright 2016-2017 The psh Authors. All rights reserved.

package psh

import (
	"errors"
	"strings"
	"sync"
)

// Prompt type that defines the shell prompt structure.
type Prompt struct {
	mu       sync.Mutex
	keys     []string
	segments map[string]Segment
}

// NewPrompt creates a new prompt and returns a pointer to its instance.
func NewPrompt(segments string) (*Prompt, error) {
	keys := strings.Split(segments, ",")
	prompt := &Prompt{
		keys:     keys,
		segments: make(map[string]Segment, len(keys)),
	}
	return prompt, nil
}

// Compile traverses all available segments, and in paralel add them to the
// prompt. It also process each segment to get all the data needed to render.
func (p *Prompt) Compile() error {
	var wg sync.WaitGroup
	wg.Add(len(p.keys))
	for _, v := range p.keys {
		go p.addSegment(&wg, v)
	}
	wg.Wait()
	return nil
}

// Render ...
func (p *Prompt) Render() (string, error) {
	return "", nil
}

// addSegment adds a segment instance by its name to the prompt.
func (p *Prompt) addSegment(wg *sync.WaitGroup, key string) error {
	defer wg.Done()
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
	default:
		return errors.New("segment unknown")
	}
	p.mu.Lock()
	p.segments[key] = segment
	p.mu.Unlock()
	return nil
}

/*
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
