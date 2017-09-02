package psh

// Segment is a partial or portion of the final prompt
type Segment interface {
	String() string
}
