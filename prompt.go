package main

import (
	"fmt"
	"strconv"
)

// Prompt type
type Prompt struct {
	prompt string
}

func (p *Prompt) append(partial string) {
	p.prompt = fmt.Sprintf("%s%s", p.prompt, partial)
}

func (p *Prompt) prepend(partial string) {
	p.prompt = fmt.Sprintf("%s%s", partial, p.prompt)
}

func (p *Prompt) reset() {
	p.prompt = fmt.Sprintf("%s%s", p.prompt, "$(tput sgr0)")
}

func (p *Prompt) setColor(fg int, bg int) {
	foreground := "$(tput setaf " + strconv.Itoa(fg) + ")"
	background := "$(tput setab " + strconv.Itoa(bg) + ")"
	p.prompt = fmt.Sprintf("%s%s%s", p.prompt, foreground, background)
}

func (p *Prompt) setBg(bg int) {
	background := "$(tput setab " + strconv.Itoa(bg) + ")"
	p.prompt = fmt.Sprintf("%s%s", p.prompt, background)
}

func (p *Prompt) setFg(fg int) {
	foreground := "$(tput setaf " + strconv.Itoa(fg) + ")"
	p.prompt = fmt.Sprintf("%s%s", p.prompt, foreground)
}

// Build the prompt
func (p *Prompt) getPrompt() string {
	p.reset()
	p.setBg(244)
	p.append(JOBS)
	p.setColor(244, 240)
	p.append(" ")

	p.reset()
	p.setBg(240)
	p.append(PATH)
	p.append(" ")

	p.reset()
	p.setColor(240, 236)
	p.append(" ")

	p.reset()
	p.setBg(236)
	p.append(getGitPartial())
	p.append(" ")
	p.reset()
	p.setFg(236)
	p.append("")

	p.reset()
	p.append(" ")

	return p.prompt
}
