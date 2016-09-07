package main

func (p *Prompt) PartialJobs() {
	p.reset()
	p.setBg(244)
	p.append(BashEscapeJobs)
	p.setColor(244, 240)
	p.append(" ")
}

func (p *Prompt) PartialPath() {
	p.reset()
	p.setBg(240)
	p.append(BashEscapePath)
	p.append(" ")
	p.reset()
	p.setColor(240, 236)
	p.append(" ")
}

func (p *Prompt) PartialGit() {
	p.reset()
	p.setBg(236)
	p.append(getGitPartial())
	p.append(" ")
	p.reset()
	p.setFg(236)
	p.append("")
}
