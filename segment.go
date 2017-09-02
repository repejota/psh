// Copyright 2016-2017 The psh Authors. All rights reserved.

package psh

// Segment is a partial or portion of the final prompt.
type Segment interface {
	Render() []byte
}
