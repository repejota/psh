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
	jobsPartialBackgroundPtr := flag.Int("jobs-partial-bg", -1, "Background color for the jobs partial.")
	pathPartialPtr := flag.Bool("path-partial", true, "Enable or disable path partial.")
	gitPartialPtr := flag.Bool("git-partial", true, "Enable or disable git partial.")
	flag.Parse()

	if *versionPtr {
		fmt.Println(Version)
		os.Exit(0)
	}

	var options Options

	options.setDefaults()

	options.JobsPartial = *jobsPartialPtr
	if *jobsPartialBackgroundPtr != -1 {
		options.JobsPartialBackground = *jobsPartialBackgroundPtr
	}
	options.JobsPartialBackground = *jobsPartialBackgroundPtr

	options.PathPartial = *pathPartialPtr

	options.GitPartial = *gitPartialPtr

	var prompt Prompt
	prompt.Options = options
	fmt.Printf(prompt.getPrompt())
}
