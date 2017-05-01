package environment

// ShellEnvironment ...
type ShellEnvironment struct {
	ShellEnv                           string
	PromptEscapeUsername               string
	PromptEscapeHostname               string
	PromptEscapeFullHostName           string
	PromptEscapeCurentWorkingDirectory string
	PromptEscapeBaseWorkingDirectory   string
	PromptEscapeJobs                   string
}

// GitEnvironment ...
type GitEnvironment struct {
	IsAGitRepo bool
	Branch     string
	Changes    int
}

// Environment ...
type Environment struct {
	Shell ShellEnvironment
	Git   GitEnvironment
}

// NewEnvironment ...
func NewEnvironment() (e Environment) {
	e.Git.IsAGitRepo = false
	e.Git.Changes = 0
	return e
}
