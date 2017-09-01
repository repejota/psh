package psh

import "fmt"

// Prompt is the type that defines a shell prompt
type Prompt struct {
}

// NewPrompt creates a new Prompt type instance
func NewPrompt() Prompt {
	p := Prompt{}
	return p
}

// String implements Stringer interface and allows the type to describe itself as an string result.
func (p Prompt) String() string {
	return fmt.Sprintf("$")
}
