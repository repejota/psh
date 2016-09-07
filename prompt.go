package main

// Prompt type
type Prompt struct {
	Options Options
	Prompt  string
}

// getPrompt Builds the prompt
func (p *Prompt) getPrompt() string {
	if p.Options.JobsPartial {
		p.PartialJobs()
	}
	if p.Options.PathPartial {
		p.PartialPath()
	}
	if p.Options.GitPartial {
		p.PartialGit()
	}
	p.reset()
	p.append(" ")
	return p.Prompt
}
