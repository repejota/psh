package themes

// DefaultTemplate ...
var DefaultTemplate = `
{{/* Jobs Partial */}}
{{color 0 108}}  {{.Shell.PromptEscapeJobs}} {{color 108 8}}

{{/* Path Partial */}}
{{color 21 8}} {{.Shell.PromptEscapeCurentWorkingDirectory}} {{if .Git.IsAGitRepo}}{{color 8 19}}{{else}}{{color 8 0}}{{end}}

{{/* Git Partial */}}
{{if .Git.IsAGitRepo}}
{{color 21 19}}  {{if .Git.Changes}}{{.Git.Changes}} {{end}}{{.Git.Branch}} {{color 19 0}}
{{end}}
{{reset}}
`
