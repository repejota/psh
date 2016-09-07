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
	jobsPartialPtr := flag.Bool("jobs-partial", true, "Enable or disable jobs partial.")
	pathPartialPtr := flag.Bool("path-partial", true, "Enable or disable path partial.")
	gitPartialPtr := flag.Bool("git-partial", true, "Enable or disable git partial.")
	flag.Parse()

	if *versionPtr {
		fmt.Println(Version)
		os.Exit(0)
	}

	var o Options
	if *jobsPartialPtr {
		o.JobsPartial = true
	} else {
		o.JobsPartial = false
	}
	if *pathPartialPtr {
		o.PathPartial = true
	} else {
		o.PathPartial = false
	}
	if *gitPartialPtr {
		o.GitPartial = true
	} else {
		o.GitPartial = false
	}

	var p Prompt
	prompt := p.getPrompt(o)
	fmt.Printf(prompt)
}
