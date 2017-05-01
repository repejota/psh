package prompt

import (
	"bytes"
	"html/template"
	"regexp"

	"github.com/repejota/psh/environment"
	"github.com/repejota/psh/partial"
	"github.com/repejota/psh/prompt/helpers"
	"github.com/repejota/psh/prompt/themes"
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
	var result bytes.Buffer
	tmpl := template.New("prompt").Funcs(helpers.TemplateHelpers)
	re := regexp.MustCompile(`\r?\n`)
	templateSource := re.ReplaceAllString(themes.DefaultTemplate, "")
	templateSource = templateSource + " "
	tmpl, err := tmpl.Parse(templateSource)
	if err != nil {
		panic(err)
	}
	tmpl.Execute(&result, p.Environment)
	return result.String()
}
