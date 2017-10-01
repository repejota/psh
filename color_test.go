// Copyright 2016-2017 The psh Authors. All rights reserved.

package psh

import "testing"

func TestResetFgBg(t *testing.T) {
	expected := `\[\e[0m\]`
	output := ResetFgBg()
	if string(output) != expected {
		t.Fatalf("Expected %s but got %s", expected, string(output))
	}
}

func TestResetForegound(t *testing.T) {
	expected := `\[\e[39m\]`
	output := ResetForeground()
	if string(output) != expected {
		t.Fatalf("Expected %s but got %s", expected, string(output))
	}
}

func TestResetBackground(t *testing.T) {
	expected := `\[\e[49m\]`
	output := ResetBackground()
	if string(output) != expected {
		t.Fatalf("Expected %s but got %s", expected, string(output))
	}
}
