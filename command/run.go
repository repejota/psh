package command

import (
	"fmt"

	"github.com/repejota/psh/prompt"
	"github.com/spf13/cobra"
)

// RunCmd ...
var RunCmd = &cobra.Command{
	Hidden: true,
	Use:    "",
	Short:  "Run the prompt generator",
	Long:   `Run the prompt generator and prints the final string.`,
	Run:    runFunc,
}

func runFunc(cmd *cobra.Command, args []string) {
	p := prompt.NewPrompt()
	p.Build()
	result := p.Render()
	fmt.Printf(result)
}
