package themes

// DefaultTemplate ...
var DefaultTemplate = `
{{color 0 108}}  {{.Shell.PromptEscapeJobs}} {{color 108 8}}
{{color 21 8}} {{.Shell.PromptEscapeCurentWorkingDirectory}} {{if .Git.IsAGitRepo}}{{color 8 19}}{{else}}{{color 8 0}}{{end}}
{{if .Git.IsAGitRepo}}{{color 21 19}}  {{.Git.Changes}} {{.Git.Branch}} {{color 19 0}}{{end}}
{{reset}}
`
