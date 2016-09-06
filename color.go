package main

import "fmt"

const (
	Reset = "$(tput sgr0)"
)

func resetFormat(prompt string) string {
	p := fmt.Sprintf("%s%s", prompt, Reset)
	return p
}

func fgColor(prompt string, color int) string {
	p := fmt.Sprintf("$(tput setaf %d)%s", color, prompt)
	return p
}
