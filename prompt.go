package main

import "fmt"

// Prompt type
type Prompt struct {
	prompt string
}

// Build the prompt
func (p *Prompt) getPrompt() string {
	p.prompt = fmt.Sprintf("%s:%s%s ",
		resetFormat(fgColor(USER_HOSTNAME, 240)),
		PATH,
		PROMPT,
	)
	p.prompt = resetFormat(p.prompt)
	return p.prompt
}
