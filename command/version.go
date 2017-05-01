package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// Version is the current version of psh following the server specs
	// filled automatically every time psh is built.
	Version string

	// Build is the current git commit and branch information filled
	// automatically every time psh is built in the following format:
	// <branch>-<gitcommit>
	Build string
)

// VersionCmd ...
var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of psh",
	Long:  `Print the current version and build information for psh.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s-%s\n", Version, Build)
	},
}
