package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
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
	versionPtr := flag.Bool("version", false, "Show version information")
	colortestPtr := flag.Bool("colortest", false, "Show colors available on the terminal.")
	backgroundtestPtr := flag.Bool("backgroundtest", false, "Show backgrounds available on the terminal.")
	flag.Parse()

	// Version command
	if *versionPtr {
		fmt.Println("psh : Version", Version, "Build", Build)
		os.Exit(0)
	}

	// Color test command
	if *colortestPtr {
		escape := "\x1b"
		reset := escape + "[0m"
		for c := 0; c < 256; c++ {
			color := fmt.Sprintf("%s[38;5;%sm", escape, strconv.Itoa(c))
			fmt.Printf("Color %d: '%s%s%s'%s", c, color, "test", reset, "\n")
		}
		os.Exit(0)
	}

	// Background test command
	if *backgroundtestPtr {
		escape := "\x1b"
		reset := escape + "[0m"
		for c := 0; c < 256; c++ {
			color := fmt.Sprintf("%s[48;5;%sm", escape, strconv.Itoa(c))
			fmt.Printf("Color %d: '%s%s%s'%s", c, color, "test", reset, "\n")
		}
	}

	fmt.Println("$ ")
}
