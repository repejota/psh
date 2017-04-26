package main

// Prompt type
type Prompt struct {
	Options Options
	Prompt  string
}

// NewPrompt creates a new prompt type.
func NewPrompt(o Options) (p Prompt) {
	p.Options = o
	return p
}

// BuildPrompt Builds the prompt.
func (p *Prompt) BuildPrompt() {
	if p.Options.JobsPartial {
		p.PartialJobs()
	}
	if p.Options.PathPartial {
		p.PartialPath()
	}
	if p.Options.GitPartial {
		p.PartialGit()
	}
}

// RenderPrompt renders prompt to final string.
func (p *Prompt) RenderPrompt() string {
	p.reset()
	p.append(" ")
	return p.Prompt
}
