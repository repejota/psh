package main

import "testing"

func TestDefaultOptions(t *testing.T) {
	var options Options

	options.setDefaults()

	if !options.JobsPartial {
		t.Error("Expected true but found ", options.JobsPartial)
	}

	if !options.PathPartial {
		t.Error("Expected true but found ", options.PathPartial)
	}

	if !options.GitPartial {
		t.Error("Expected true but found ", options.GitPartial)
	}
}
