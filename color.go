package main

import "fmt"

func resetFormat(prompt string) string {
	p := fmt.Sprintf("%s%s", prompt, "$(tput sgr0)")
	return p
}

func fgColor(prompt string, color int) string {
	p := fmt.Sprintf("$(tput setaf %d)%s", color, prompt)
	return p
}

func bgColor(prompt string, color int) string {
	p := fmt.Sprintf("$(tput setab %d)%s", color, prompt)
	return p
}
