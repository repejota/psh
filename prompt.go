package main

import "fmt"

// Prompt type
type Prompt struct {
	prompt string
}

// Build the prompt
func (p *Prompt) getPrompt() string {
	p.prompt = fmt.Sprintf("%s%s%s ",
		bgColor(JOBS, 239),
		//bgColor(USER_HOSTNAME, 240),
		bgColor(PATH, 238),
		resetFormat(PROMPT),
	)
	return resetFormat(p.prompt)
}
