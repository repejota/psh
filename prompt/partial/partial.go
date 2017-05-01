package partial

import (
	"github.com/repejota/psh/environment"
)

// Partial ...
type Partial interface {
	// Run ...
	Run(e *environment.Environment)
}
