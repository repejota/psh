package main

import "testing"

func TestExistPath(t *testing.T) {
	ep := existsPath(".git")
	if !ep {
		t.Error("Expected true, got ", ep)
	}
}

func TestDoNotExistPath(t *testing.T) {
	ep := existsPath("foobar")
	if ep {
		t.Error("Expected false, got ", ep)
	}
}
