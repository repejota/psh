package main

import "testing"

func TestAppend(t *testing.T) {
	var prompt Prompt
	prompt.append("foo")
	if prompt.Prompt != "foo" {
		t.Error("Expected empty string, but got ", prompt.Prompt)
	}
}

func TestPrepend(t *testing.T) {
	var prompt Prompt
	prompt.prepend("bar")
	if prompt.Prompt != "bar" {
		t.Error("Expected empty string, but got ", prompt.Prompt)
	}
}
