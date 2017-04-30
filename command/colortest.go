package command

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// ColorTestCmd ...
func ColorTestCmd(cmd *cobra.Command, args []string) {
	escape := "\x1b"
	reset := escape + "[0m"

	// Colors
	disableColors, _ := cmd.Flags().GetBool("disable-colors")
	if !disableColors {
		color := ""
		for c := 0; c < 256; c++ {
			color = escape + "[38;5;" + strconv.Itoa(c) + "m"
			fmt.Printf("Color %d: '%s%s%s'%s", c, color, "test", reset, "\n")
		}
	}

	// Backgrounds
	disableBackgrounds, _ := cmd.Flags().GetBool("disable-backgrounds")
	if !disableBackgrounds {
		background := ""
		for b := 0; b < 256; b++ {
			background = escape + "[48;5;" + strconv.Itoa(b) + "m"
			fmt.Printf("Background %d: '%s%s%s'%s", b, background, "test", reset, "\n")
		}
	}
}
