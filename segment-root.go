package psh

// SegmentRoot implements the UID part of the prompt.
//
// It renders the character '#' if the effective UID of the current user is 0,
// otherwise renders the character '$'.
type SegmentRoot struct {
}

// NewSegmentRoot creates an instace of SegmentRoot type.
func NewSegmentRoot() *SegmentRoot {
	segment := &SegmentRoot{}
	return segment
}

// Render renders the segment resaults.
func (s *SegmentRoot) Render() []byte {
	return []byte("\\$")
}
