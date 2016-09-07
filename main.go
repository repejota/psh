package main

import "fmt"

func main() {
	var p Prompt
	prompt := p.getPrompt()
	fmt.Printf(prompt)
}
