package main

import "fmt"

// append Appends a partial string to the end of the prompt
func (p *Prompt) append(partial string) {
	p.Prompt = fmt.Sprintf("%s%s", p.Prompt, partial)
}

// prepend Prepends a partial string to the start of the prompt
func (p *Prompt) prepend(partial string) {
	p.Prompt = fmt.Sprintf("%s%s", partial, p.Prompt)
}
