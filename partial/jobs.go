package partial

// Jobs ...
type Jobs struct {
	BashEscapeSequence string
}

// NewJobsPartial ...
func NewJobsPartial() (jobs Jobs) {
	return jobs
}

// Update ...
func (j Jobs) Update() {
	j.BashEscapeSequence = "\\j"
}

// Render ...
func (j Jobs) Render() string {
	return ""
}
