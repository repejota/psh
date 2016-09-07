package main

import "testing"

func TestExistPath(t *testing.T) {
	ep := existsPath(".git")
	if !ep {
		t.Error("Expected true, got ", ep)
	}

	ep = existsPath("foobar")
	if ep {
		t.Error("Expected false, got ", ep)
	}
}
