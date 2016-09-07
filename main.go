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
	jobsPartialForegroundPtr := flag.Int("jobs-partial-fg", -1, "Foreground color for the jobs partial.")
	pathPartialPtr := flag.Bool("path-partial", true, "Enable or disable path partial.")
	pathPartialBackgroundPtr := flag.Int("path-partial-bg", -1, "Background color for the path partial.")
	pathPartialForegroundPtr := flag.Int("path-partial-fg", -1, "Foreground color for the path partial.")
	gitPartialPtr := flag.Bool("git-partial", true, "Enable or disable git partial.")
	gitPartialBackgroundPtr := flag.Int("git-partial-bg", -1, "Background color for the git partial.")
	gitPartialForegroundPtr := flag.Int("git-partial-fg", -1, "Foreground color for the git partial.")
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
	if *jobsPartialForegroundPtr != -1 {
		options.JobsPartialForeground = *jobsPartialForegroundPtr
	}

	options.PathPartial = *pathPartialPtr
	if *pathPartialBackgroundPtr != -1 {
		options.PathPartialBackground = *pathPartialBackgroundPtr
	}
	if *pathPartialForegroundPtr != -1 {
		options.PathPartialForeground = *pathPartialForegroundPtr
	}

	options.GitPartial = *gitPartialPtr
	if *gitPartialBackgroundPtr != -1 {
		options.GitPartialBackground = *gitPartialBackgroundPtr
	}
	if *gitPartialForegroundPtr != -1 {
		options.GitPartialForeground = *gitPartialForegroundPtr
	}

	var prompt Prompt
	prompt.Options = options
	fmt.Printf(prompt.getPrompt())
}
