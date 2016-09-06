package main

// Prompt type
type Prompt struct {
	prompt string
}

// Build the prompt
func (p *Prompt) getPrompt() string {
	prompt := "\\u@\\H:\\w\\$ "
	p.prompt = prompt
	return p.prompt
}
