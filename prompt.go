package main

import (
	"fmt"
	"strconv"
)

// Prompt type
type Prompt struct {
	prompt string
}

// append Appends a partial string to the end of the prompt
func (p *Prompt) append(partial string) {
	p.prompt = fmt.Sprintf("%s%s", p.prompt, partial)
}

// prepend Prepends a partial string to the start of the prompt
func (p *Prompt) prepend(partial string) {
	p.prompt = fmt.Sprintf("%s%s", partial, p.prompt)
}

// reset Resets color, background and attributes on the terminal
func (p *Prompt) reset() {
	p.prompt = fmt.Sprintf("%s%s", p.prompt, "$(tput sgr0)")
}

// setColor Sets the color and the background of the current cursor
func (p *Prompt) setColor(fg int, bg int) {
	foreground := "$(tput setaf " + strconv.Itoa(fg) + ")"
	background := "$(tput setab " + strconv.Itoa(bg) + ")"
	p.prompt = fmt.Sprintf("%s%s%s", p.prompt, foreground, background)
}

// setBg sets the background of the current cursor
func (p *Prompt) setBg(bg int) {
	background := "$(tput setab " + strconv.Itoa(bg) + ")"
	p.prompt = fmt.Sprintf("%s%s", p.prompt, background)
}

// setFg sets the foreground color of the current cursor
func (p *Prompt) setFg(fg int) {
	foreground := "$(tput setaf " + strconv.Itoa(fg) + ")"
	p.prompt = fmt.Sprintf("%s%s", p.prompt, foreground)
}

// getPrompt Builds the prompt
func (p *Prompt) getPrompt() string {
	p.reset()
	p.setBg(244)
	p.append(BashEscapeJobs)
	p.setColor(244, 240)
	p.append(" ")

	p.reset()
	p.setBg(240)
	p.append(BashEscapePath)
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
