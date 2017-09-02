package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/repejota/psh"
)

var (
	// Version is the latest release number.
	//
	// This number is the latest tag from the git repository.
	Version string

	// Build is the lastest build string.
	//
	// This string is the branch name and the commit hash (short format)
	Build string
)

func main() {
	// Define CLI flags
	versionPtr := flag.Bool("version",
		false,
		"Show version information")
	colortestPtr := flag.Bool("colortest",
		false,
		"Show available colors on the terminal.")
	backgroundtestPtr := flag.Bool("backgroundtest",
		false,
		"Show available backgrounds on the terminal.")

	flag.Parse()

	if *versionPtr {
		showVersion()
	}

	if *colortestPtr {
		colorTest()
	}

	if *backgroundtestPtr {
		backgroundTest()
	}

	// Create a prompt instance
	prompt := psh.NewPrompt()

	// Add segment root to the prompt
	sroot := psh.NewSegmentRoot()
	err := prompt.AddSegment(sroot)
	if err != nil {
		log.Fatal("Can't add segment root to the prompt", err)
	}

	res := prompt.Render()
	// Prints rendered prompt to stdout
	fmt.Printf("%s ", res)
}

// showVersion prints the current version information.
func showVersion() {
	fmt.Println("psh : Version", Version, "Build", Build)
	os.Exit(0)
}

// colorTest executes a foreground color test.
func colorTest() {
	escape := "\x1b"
	reset := escape + "[0m"
	for c := 0; c < 256; c++ {
		color := fmt.Sprintf("%s[38;5;%sm", escape, strconv.Itoa(c))
		fmt.Printf("Color %d: '%s%s%s'%s", c, color, "test", reset, "\n")
	}
	os.Exit(0)
}

// backgroundTest executes a backgtound color test.
func backgroundTest() {
	escape := "\x1b"
	reset := escape + "[0m"
	for c := 0; c < 256; c++ {
		color := fmt.Sprintf("%s[48;5;%sm", escape, strconv.Itoa(c))
		fmt.Printf("Color %d: '%s%s%s'%s", c, color, "test", reset, "\n")
	}
	os.Exit(0)
}
