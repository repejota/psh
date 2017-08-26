package command

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// ColorTestCmd ...
var ColorTestCmd = &cobra.Command{
	Use:   "colortest",
	Short: "Run the color test",
	Long:  `Run a local color test to show all available color and backgrounds on this terminal.`,
	Run:   colorTestFunc,
}

func colorTestFunc(cmd *cobra.Command, args []string) {
	escape := "\x1b"
	reset := escape + "[0m"

	// Colors
	color := ""
	for c := 0; c < 256; c++ {
		color = escape + "[38;5;" + strconv.Itoa(c) + "m"
		fmt.Printf("Color %d: '%s%s%s'%s", c, color, "test", reset, "\n")
	}

	// Backgrounds
	background := ""
	for b := 0; b < 256; b++ {
		background = escape + "[48;5;" + strconv.Itoa(b) + "m"
		fmt.Printf("Background %d: '%s%s%s'%s", b, background, "test", reset, "\n")
	}
}
