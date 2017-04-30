package prompt

// Prompt ...
type Prompt struct {
	options Options
}

// NewPrompt ...
func NewPrompt() (p Prompt) {
	p.options = NewOptions()
	return p
}
