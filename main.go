package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

var (
	// Version is the version number
	Version string

	// Build is the build hash
	Build string
)

func colorTest() {
	bg := 0
	reset := "$(tput sgr0)"
	for fg := 0; fg < 16; fg++ {
		foreground := "$(tput setaf " + strconv.Itoa(fg) + ")"
		background := "$(tput setab " + strconv.Itoa(bg) + ")"
		fmt.Printf("echo -e \"%s%s fg:%d bg:%d - %s%s\"", foreground, background, fg, bg, "sample ", reset)
	}
	os.Exit(0)
}

func commandVersion() {
	fmt.Printf("%s-%s\n", Version, Build)
	os.Exit(0)
}

func main() {
	jobsPartialPtr := flag.Bool("jobs-partial", true, "Enable or disable jobs partial.")
	pathPartialPtr := flag.Bool("path-partial", true, "Enable or disable path partial.")
	gitPartialPtr := flag.Bool("git-partial", true, "Enable or disable git partial.")

	colorTestPtr := flag.Bool("colortest", false, "Execute a color test.")

	versionPtr := flag.Bool("version", false, "Show version number.")

	flag.Parse()

	if *versionPtr {
		commandVersion()
	}

	if *colorTestPtr {
		colorTest()
	}

	options := NewOptions()
	options.JobsPartial = *jobsPartialPtr
	options.PathPartial = *pathPartialPtr
	options.GitPartial = *gitPartialPtr

	// Build prompt
	prompt := NewPrompt(options)
	prompt.Build()

	// Render prompt
	var result string
	result = prompt.Render()

	fmt.Printf(result)
}
