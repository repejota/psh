// Copyright 2016-2017 The psh Authors. All rights reserved.

package psh

import (
	"fmt"
	"os"
)

const (
	// colorTemplate ...
	colorTemplate = "\\[\\e%s\\]"
)

// ResetFgBg resets foreground and background to default values.
func ResetFgBg() []byte {
	escapeSequence := fmt.Sprintf(colorTemplate, "[0m")
	return []byte(escapeSequence)
}

// ResetForeground reset foreground to default value.
func ResetForeground() []byte {
	escapeSequence := fmt.Sprintf(colorTemplate, "[39m")
	return []byte(escapeSequence)
}

// ResetBackground reset background to default value.
func ResetBackground() []byte {
	escapeSequence := fmt.Sprintf(colorTemplate, "[49m")
	return []byte(escapeSequence)
}

// SetBackground sets background color on the terminal.
func SetBackground(background uint8) []byte {
	bg := fmt.Sprintf("[48;5;%dm", background)
	escapeSequence := fmt.Sprintf(colorTemplate, bg)
	return []byte(escapeSequence)
}

// SetForeground sets foreground color on the terminal.
func SetForeground(foreground uint8) []byte {
	fg := fmt.Sprintf("[38;5;%dm", foreground)
	escapeSequence := fmt.Sprintf(colorTemplate, fg)
	return []byte(escapeSequence)
}

// ColorTest executes a foreground color test.
func ColorTest() {
	escape := "\x1b"
	for c := 0; c < 256; c++ {
		color := fmt.Sprintf("%s[38;5;%dm", escape, c)
		fmt.Printf("%s%d\n", color, c)
	}
	fmt.Printf("%s[0m", escape)
	os.Exit(0)
}

// BackgroundTest executes a backgtound color test.
func BackgroundTest() {
	escape := "\x1b"
	for c := 0; c < 256; c++ {
		background := fmt.Sprintf("%s[48;5;%dm", escape, c)
		fmt.Printf("%s%d\n", background, c)
	}
	fmt.Printf("%s[0m", escape)
	os.Exit(0)
}
