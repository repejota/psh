package main

// Options type
type Options struct {
	JobsPartial           bool
	JobsPartialBackground int
	PathPartial           bool
	GitPartial            bool
}

func (o *Options) setDefaults() {
	o.JobsPartial = true
	o.JobsPartialBackground = 244
	o.PathPartial = true
	o.GitPartial = true
}
