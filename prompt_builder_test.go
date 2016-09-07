package main

import "testing"

func TestAppend(t *testing.T) {
	var prompt Prompt
	prompt.Prompt = "foobar"
	prompt.append("suffix")
	if prompt.Prompt != "foobarsuffix" {
		t.Error("Expected 'foobarsuffix' string, but got ", prompt.Prompt)
	}
}

func TestPrepend(t *testing.T) {
	var prompt Prompt
	prompt.Prompt = "foobar"
	prompt.prepend("prefix")
	if prompt.Prompt != "prefixfoobar" {
		t.Error("Expected 'prefixfoobar' string, but got ", prompt.Prompt)
	}
}
