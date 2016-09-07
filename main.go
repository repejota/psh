package main

import (
	"flag"
	"fmt"
	"os"
)

// Version is the version number
var Version string

func main() {
	versionPtr := flag.Bool("version", false, "Show version number.")
	flag.Parse()

	if *versionPtr {
		fmt.Println(Version)
		os.Exit(0)
	}

	var p Prompt
	prompt := p.getPrompt()
	fmt.Printf(prompt)
}
