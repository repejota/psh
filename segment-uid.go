package psh

// SegmentUID implements the UID part of the prompt.
//
// It renders the character '#' if the effective UID of the current user is 0,
// otherwise renders the character '$'.
type SegmentUID struct {
}

// NewSegmentUID creates an instace of SegmentUID type.
func NewSegmentUID() *SegmentUID {
	segment := &SegmentUID{}
	return segment
}

// String implements Stringer interface and renders the type as string.
func (s *SegmentUID) String() string {
	return "\\$"
}
