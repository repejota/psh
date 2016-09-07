package main

// Prompt type
type Prompt struct {
	Options Options
	Prompt  string
}

// getPrompt Builds the prompt
func (p *Prompt) getPrompt(o Options) string {
	if o.JobsPartial {
		p.PartialJobs()
	}
	if o.PathPartial {
		p.PartialPath()
	}
	if o.GitPartial {
		p.PartialGit()
	}
	p.reset()
	p.append(" ")
	return p.Prompt
}
