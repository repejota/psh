package main

import (
	"fmt"

	"github.com/repejota/psh/command"
	"github.com/repejota/psh/prompt"
	"github.com/spf13/cobra"
)

var (
	// Version ...
	Version string

	// Build ...
	Build string
)

func main() {
	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of psh",
		Long:  `Print the current version and build information for psh.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%s-%s\n", Version, Build)
		},
	}

	var colortestCmd = &cobra.Command{
		Use:   "colortest",
		Short: "Run the color test",
		Long:  `Run a local color test to show all available color and backgrounds on this terminal.`,
		Run:   command.ColorTestCmd,
	}
	colortestCmd.Flags().Bool("disable-colors", false, "do not execute foreground color test")
	colortestCmd.Flags().Bool("disable-backgrounds", false, "do not execute background color test")

	var runCmd = &cobra.Command{
		Use:   "run",
		Short: "Run the prompt generator",
		Long:  `Run the prompt generator and prints the final string.`,
		Run: func(cmd *cobra.Command, args []string) {
			var result string
			_ = prompt.NewPrompt()
			fmt.Printf(result)
		},
	}

	var rootCmd = runCmd
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(colortestCmd)
	rootCmd.Execute()
}
