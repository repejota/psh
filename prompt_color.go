package main

import (
	"fmt"
	"strconv"
)

// reset Resets color, background and attributes on the terminal
func (p *Prompt) reset() {
	p.Prompt = fmt.Sprintf("%s%s", p.Prompt, "$(tput sgr0)")
}

// setColor Sets the color and the background of the current cursor
func (p *Prompt) setColor(fg int, bg int) {
	foreground := "$(tput setaf " + strconv.Itoa(fg) + ")"
	background := "$(tput setab " + strconv.Itoa(bg) + ")"
	p.Prompt = fmt.Sprintf("%s%s%s", p.Prompt, foreground, background)
}

// setBg sets the background of the current cursor
func (p *Prompt) setBg(bg int) {
	background := "$(tput setab " + strconv.Itoa(bg) + ")"
	p.Prompt = fmt.Sprintf("%s%s", p.Prompt, background)
}

// setFg sets the foreground color of the current cursor
func (p *Prompt) setFg(fg int) {
	foreground := "$(tput setaf " + strconv.Itoa(fg) + ")"
	p.Prompt = fmt.Sprintf("%s%s", p.Prompt, foreground)
}
