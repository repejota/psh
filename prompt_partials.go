package main

// PartialJobs renders the jobs partial
func (p *Prompt) PartialJobs() {
	p.reset()
	p.setColor(p.Options.JobsPartialForeground, p.Options.JobsPartialBackground)
	p.append("  ")
	p.append(BashEscapeJobs)
	p.append(" ")
	p.setColor(p.Options.JobsPartialBackground, p.Options.PathPartialBackground)
	p.append(p.Options.PartialSuffix)
	p.append(" ")
}

// PartialPath renders the current path partial
func (p *Prompt) PartialPath() {
	p.reset()
	p.setColor(p.Options.PathPartialForeground, p.Options.PathPartialBackground)
	p.append(BashEscapePath)
	p.append(" ")
	p.setColor(p.Options.PathPartialBackground, p.Options.GitPartialBackground)
	p.append(p.Options.PartialSuffix)
	p.append(" ")
}

// PartialGit renders the git partial
func (p *Prompt) PartialGit() {
	if !existsPath(".git") {
		p.append("\\[$(tput cub 2)\\]")
		p.reset()
		p.setFg(p.Options.PathPartialBackground)
		p.append(p.Options.PartialSuffix)
		return
	}
	p.reset()
	p.setColor(p.Options.GitPartialForeground, p.Options.GitPartialBackground)
	p.append(getGitPartial())
	p.append(" ")
	p.reset()
	p.setFg(p.Options.GitPartialBackground)
	p.append(p.Options.PartialSuffix)
}
