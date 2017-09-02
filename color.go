package psh

import "fmt"

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
