// Copyright 2016-2017 The psh Authors. All rights reserved.
package psh

import (
	"os"
	"os/exec"
	"testing"
)

func TestSegmentGitCompile(t *testing.T) {
	_ = os.Chdir("/tmp")
	expected := ``
	segment := NewSegmentGit()
	segment.Compile()
	if string(segment.Data) != expected {
		t.Fatalf("Compiled data expected to be %q but got %q", expected, string(segment.Data))
	}
}

func TestSegmentGitRender(t *testing.T) {
	_ = os.Chdir("/tmp")
	expected := ``
	segment := NewSegmentGit()
	segment.Compile()
	out := segment.Render()
	if string(out) != expected {
		t.Fatalf("Rendered segment expected to be %q but got %q", expected, string(out))
	}
}

func TestSegmentGitCompileBranch(t *testing.T) {
	_ = os.Mkdir("/tmp/repo", 0755)
	_ = os.Chdir("/tmp/repo")
	_, _ = exec.Command("git", "config", "user.email", "you@example.com").Output()
	_, _ = exec.Command("git", "config", "user.name", "Your Name").Output()
	_, _ = exec.Command("git", "init").Output()
	_, _ = exec.Command("touch", "README.md").Output()
	_, _ = exec.Command("git", "add", ".").Output()
	_, _ = exec.Command("git", "commit", "-am", ".").Output()
	defer func() { _ = os.RemoveAll("/tmp/repo") }()

	expected := `master`
	segment := NewSegmentGit()
	segment.Compile()
	if string(segment.Data) != expected {
		t.Fatalf("Compiled data expected to be %q but got %q", expected, string(segment.Data))
	}
}

func TestSegmentGitRenderBranch(t *testing.T) {
	_ = os.Mkdir("/tmp/repo", 0755)
	_ = os.Chdir("/tmp/repo")
	_, _ = exec.Command("git", "config", "user.email", "you@example.com").Output()
	_, _ = exec.Command("git", "config", "user.name", "Your Name").Output()
	_, _ = exec.Command("git", "init").Output()
	_, _ = exec.Command("touch", "README.md").Output()
	_, _ = exec.Command("git", "add", ".").Output()
	_, _ = exec.Command("git", "commit", "-am", ".").Output()
	defer func() { _ = os.RemoveAll("/tmp/repo") }()

	expected := `\[\e[48;5;148m\] \[\e[38;5;0m\]master\[\e[39m\] `
	segment := NewSegmentGit()
	segment.Compile()
	out := segment.Render()
	if string(out) != expected {
		t.Fatalf("Rendered segment expected to be %q but got %q", expected, string(out))
	}
}
