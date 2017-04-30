package main

import (
	"github.com/repejota/psh/partial"
)

// Prompt type
type Prompt struct {
	Options  Options
	Prompt   string
	Partials map[string]partial.Partial
}

// NewPrompt creates a new prompt type.
func NewPrompt(o Options) (p Prompt) {
	p.Options = o
	p.Partials = make(map[string]partial.Partial)
	p.Prompt = ""
	return p
}

// Build Builds the prompt.
func (p *Prompt) Build() {
	if p.Options.JobsPartial {
		jobspartial := partial.NewJobsPartial()
		jobspartial.Update()
		p.Partials["jobs"] = jobspartial
	}
	if p.Options.PathPartial {
		pathpartial := partial.NewPathPartial()
		pathpartial.Update()
		p.Partials["path"] = pathpartial
	}
	if p.Options.GitPartial {
		gitpartial := partial.NewGitPartial()
		gitpartial.Update()
		p.Partials["git"] = gitpartial
	}
}

// Render renders prompt to final string.
func (p *Prompt) Render() string {
	if p.Partials["jobs"] != nil {
		p.append(p.Partials["jobs"].Render())
		p.PartialJobs()
	}
	if p.Partials["path"] != nil {
		p.append(p.Partials["path"].Render())
		p.PartialPath()
	}
	if p.Partials["git"] != nil {
		p.append(p.Partials["git"].Render())
		p.PartialGit()
	}
	p.reset()
	p.append(" ")
	return p.Prompt
}
