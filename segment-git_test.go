// Copyright 2016-2017 The psh Authors. All rights reserved.
package psh

import (
	"os"
	"os/exec"
	"testing"
)

func TestSegmentGitCompile(t *testing.T) {
	os.Chdir("/tmp")
	expected := ``
	segment := NewSegmentGit()
	segment.Compile()
	if string(segment.Data) != expected {
		t.Fatalf("Compiled data expected to be %q but got %q", expected, string(segment.Data))
	}
}

func TestSegmentGitRender(t *testing.T) {
	os.Chdir("/tmp")
	expected := ``
	segment := NewSegmentGit()
	segment.Compile()
	out := segment.Render()
	if string(out) != expected {
		t.Fatalf("Rendered segment expected to be %q but got %q", expected, string(out))
	}
}

func TestSegmentGitCompileBranch(t *testing.T) {
	os.Mkdir("/tmp/repo", 0755)
	os.Chdir("/tmp/repo")
	exec.Command("git", "init").Output()
	exec.Command("touch", "README.md").Output()
	exec.Command("git", "add", ".").Output()
	exec.Command("git", "commit", "-am", ".").Output()
	defer os.RemoveAll("/tmp/repo")

	expected := `master`
	segment := NewSegmentGit()
	segment.Compile()
	if string(segment.Data) != expected {
		t.Fatalf("Compiled data expected to be %q but got %q", expected, string(segment.Data))
	}
}

func TestSegmentGitRenderBranch(t *testing.T) {
	os.Mkdir("/tmp/repo", 0755)
	os.Chdir("/tmp/repo")
	exec.Command("git", "init").Output()
	exec.Command("touch", "README.md").Output()
	exec.Command("git", "add", ".").Output()
	exec.Command("git", "commit", "-am", ".").Output()
	defer os.RemoveAll("/tmp/repo")

	expected := `\[\e[48;5;148m\] \[\e[38;5;0m\]master\[\e[39m\] `
	segment := NewSegmentGit()
	segment.Compile()
	out := segment.Render()
	if string(out) != expected {
		t.Fatalf("Rendered segment expected to be %q but got %q", expected, string(out))
	}
}
