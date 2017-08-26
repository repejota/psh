package main

import (
	"github.com/repejota/psh/command"
)

func main() {
	var rootCmd = command.RunCmd
	rootCmd.AddCommand(command.VersionCmd)
	rootCmd.AddCommand(command.ColorTestCmd)
	rootCmd.Execute()
}
