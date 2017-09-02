// Copyright 2016-2017 The psh Authors. All rights reserved.

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

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
	segmentsPtr := flag.String("segments",
		psh.DefaultSegments,
		"Segments to show on the propmt.")

	flag.Parse()

	if *versionPtr {
		showVersion()
	}

	if *colortestPtr {
		psh.ColorTest()
	}

	if *backgroundtestPtr {
		psh.BackgroundTest()
	}

	segmentsList := strings.Split(*segmentsPtr, ",")

	// Create a prompt instance
	prompt := psh.NewPrompt(segmentsList)

	// Add segments to the propmt
	for _, sname := range segmentsList {
		err := prompt.AddSegment(sname)
		if err != nil {
			log.Fatalf("Can't add segment %s to the prompt: %s", sname, err)
		}
	}

	// Render prompt
	res := prompt.Render()

	fmt.Printf("%s", res)
}

// showVersion prints the current version information.
func showVersion() {
	fmt.Println("psh : Version", Version, "Build", Build)
	os.Exit(0)
}
