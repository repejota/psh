package psh

import (
	"fmt"
	"os"
	"strconv"
)

const (
	// colorTemplate sets the foreground and background on the terminal.
	colorTemplate = "\\[\\e%s\\]"
)

// ResetFgAndBg reset the foreground and the background on the terminal to
// the default values.
func ResetFgAndBg() []byte {
	escapeSequence := fmt.Sprintf(colorTemplate, "[0m")
	return []byte(escapeSequence)
}

// ColorTest executes a foreground color test.
func ColorTest() {
	escape := "\x1b"
	reset := escape + "[0m"
	for c := 0; c < 256; c++ {
		color := fmt.Sprintf("%s[38;5;%sm", escape, strconv.Itoa(c))
		fmt.Printf("Color %d: '%s%s%s'%s", c, color, "test", reset, "\n")
	}
	os.Exit(0)
}

// BackgroundTest executes a backgtound color test.
func BackgroundTest() {
	escape := "\x1b"
	reset := escape + "[0m"
	for c := 0; c < 256; c++ {
		color := fmt.Sprintf("%s[48;5;%sm", escape, strconv.Itoa(c))
		fmt.Printf("Color %d: '%s%s%s'%s", c, color, "test", reset, "\n")
	}
	os.Exit(0)
}
