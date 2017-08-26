package partial

import (
	"os"

	"github.com/repejota/psh/environment"
)

// ShellPartial ...
type ShellPartial struct {
}

// NewShellPartial ...
func NewShellPartial() (sp ShellPartial) {
	return sp
}

// Run ...
func (sp ShellPartial) Run(e *environment.Environment) {
	e.Shell.ShellEnv = os.Getenv("SHELL")
	e.Shell.PromptEscapeUsername = "\\u"
	e.Shell.PromptEscapeHostname = "\\h"
	e.Shell.PromptEscapeFullHostName = "\\H"
	e.Shell.PromptEscapeCurentWorkingDirectory = "\\w"
	e.Shell.PromptEscapeBaseWorkingDirectory = "\\W"
	e.Shell.PromptEscapeJobs = "\\j"
}
