// Copyright 2016-2017 The psh Authors. All rights reserved.

package main

import (
	"flag"
	"fmt"
	"log"
	"os"

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

	versionFlag        bool
	colortestFlag      bool
	backgroundtestFlag bool
	segmentsFlag       string
)

func init() {
	flag.BoolVar(&versionFlag, "version", false, "Show version information")
	flag.BoolVar(&colortestFlag, "colortest", false, "Show available colors on the terminazl.")
	flag.BoolVar(&backgroundtestFlag, "backgroundtest", false, "Show available backgrounds on the terminal.")
	flag.StringVar(&segmentsFlag, "segments", psh.DefaultSegments, "Segments to show on the propmt.")
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("psh: ")

	flag.Parse()
	if versionFlag {
		versionInfo := showVersion()
		fmt.Println(versionInfo)
		os.Exit(0)
	}
	if colortestFlag {
		ColorTest()
		os.Exit(0)
	}
	if backgroundtestFlag {
		BackgroundTest()
		os.Exit(0)
	}

	err := doPSH()
	if err != nil {
		log.Fatal(err)
	}
}

// doPSH builds the prompt decoupled from main function to easy tests.
func doPSH() error {
	// Create a prompt instance
	prompt := psh.NewPrompt(segmentsFlag)

	// Compile prompt (all of its segments)
	err := prompt.Compile()
	if err != nil {
		return err
	}

	// Render prompt
	result, err := prompt.Render()
	if err != nil {
		return err
	}

	fmt.Printf("%s", result)
	return nil
}

// showVersion prints the current version information.
func showVersion() string {
	versionInfo := fmt.Sprintf("psh : Version %s Build %s", Version, Build)
	return versionInfo
}

// ColorTest executes a foreground color test.
func ColorTest() {
	escape := "\x1b"
	for c := 0; c < 256; c++ {
		color := fmt.Sprintf("%s[38;5;%dm", escape, c)
		fmt.Printf("%s%d\n", color, c)
	}
	fmt.Printf("%s[0m", escape)
}

// BackgroundTest executes a backgtound color test.
func BackgroundTest() {
	escape := "\x1b"
	for c := 0; c < 256; c++ {
		background := fmt.Sprintf("%s[48;5;%dm", escape, c)
		fmt.Printf("%s%d\n", background, c)
	}
	fmt.Printf("%s[0m", escape)
}
