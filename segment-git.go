// Copyright 2016-2017 The psh Authors. All rights reserved.

package psh

import (
	"bytes"
	"os/exec"
	"strings"
)

// git rev-parse --show-toplevel : shows root of the repo
// git status --porcelain -b : shows conmplete status of the repo

// SegmentGit implements the git SCM partial of the prompt.
//
// It renders the current branch if the folder contains a repository.
type SegmentGit struct {
	Data []byte
}

// NewSegmentGit creates an instace of SegmentHostname type.
func NewSegmentGit() *SegmentGit {
	segment := &SegmentGit{}
	return segment
}

// Compile ...
func (s *SegmentGit) Compile() {
	currentBranch := getCurrentBranch()
	s.Data = []byte(currentBranch)
}

// Render renders the segment results.
func (s *SegmentGit) Render() []byte {
	var b bytes.Buffer
	if len(s.Data) != 0 {
		b.Write(SetBackground(202))
		b.Write([]byte(" "))
		b.Write(s.Data)
		b.Write([]byte(" "))
	}
	return b.Bytes()
}

func getCurrentBranch() string {
	command := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	commandOut, err := command.Output()
	if err != nil {
		return ""
	}
	currentBranch := strings.Trim(string(commandOut), "\n")
	return currentBranch
}
