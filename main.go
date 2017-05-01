package main

import (
	"fmt"

	"github.com/repejota/psh/command"
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

	var rootCmd = command.RunCmd
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(command.ColorTestCmd)
	rootCmd.Execute()
}
