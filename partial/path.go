package partial

// Path ...
type Path struct {
	BashEscapeSequence string
}

// NewPathPartial ...
func NewPathPartial() (path Path) {
	return path
}

// Update ...
func (p Path) Update() {
	p.BashEscapeSequence = "\\w"
}

// Render ...
func (p Path) Render() string {
	return ""
}
