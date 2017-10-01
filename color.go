// Copyright 2016-2017 The psh Authors. All rights reserved.

package psh

import (
	"fmt"
)

const (
	// colorTemplate is an escaped template to apply color to text using ANSI
	// escape sequences.
	colorTemplate = `\[\e%s\]`
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
