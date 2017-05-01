package prompt

import (
	"bytes"
	"fmt"
	"html/template"
	"strconv"

	"github.com/repejota/psh/environment"
	"github.com/repejota/psh/partial"
)

// Prompt ...
type Prompt struct {
	Options     Options
	Environment environment.Environment
}

// NewPrompt ...
func NewPrompt() (p Prompt) {
	p.Options = NewOptions()
	p.Environment = environment.NewEnvironment()
	return p
}

// Build ...
func (p *Prompt) Build() {
	sp := partial.NewShellPartial()
	sp.Run(&p.Environment)
	gp := partial.NewGitPartial()
	gp.Run(&p.Environment)
}

// Render ...
func (p *Prompt) Render() string {
	escape := "\x1b"
	funcMap := template.FuncMap{
		"color": func(fg int, bg int) string {
			foreground := escape + "[38;5;" + strconv.Itoa(fg) + "m"
			background := escape + "[48;5;" + strconv.Itoa(bg) + "m"
			return fmt.Sprintf("%s%s", foreground, background)
		},
		"fg": func(c int) string {
			foreground := escape + "[38;5;" + strconv.Itoa(c) + "m"
			return foreground
		},
		"bg": func(c int) string {
			background := escape + "[48;5;" + strconv.Itoa(c) + "m"
			return background
		},
		"reset": func() string {
			reset := escape + "[0m"
			return reset
		},
	}
	var result bytes.Buffer
	tmpl := template.New("prompt").Funcs(funcMap)
	templateSource := ""
	templateSource = templateSource + "{{color 0 108}}  {{.Shell.PromptEscapeJobs}} {{color 108 8}}{{reset}}"
	templateSource = templateSource + "{{color 21 8}} {{.Shell.PromptEscapeCurentWorkingDirectory}} {{color 8 19}}{{reset}}"
	templateSource = templateSource + "{{if .Git.IsAGitRepo}}{{color 21 19}}  {{.Git.Changes}} {{.Git.Branch}}{{end}} {{color 19 0}}{{reset}}"
	templateSource = templateSource + " "
	tmpl, err := tmpl.Parse(templateSource)
	if err != nil {
		panic(err)
	}
	tmpl.Execute(&result, p.Environment)
	return result.String()
}
