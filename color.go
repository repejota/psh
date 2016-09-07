package main

import "fmt"

func fgColor(prompt string, color int) string {
	p := fmt.Sprintf("$(tput setaf %d)%s", color, prompt)
	return p
}

func bgColor(prompt string, color int) string {
	p := fmt.Sprintf("$(tput setab %d)%s", color, prompt)
	return p
}
