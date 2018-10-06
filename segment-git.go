// Copyright 2016-2018, psh project Authors.
//
// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with this
// work for additional information regarding copyright ownership.  The ASF
// licenses this file to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.  See the
// License for the specific language governing permissions and limitations
// under the License.

package psh

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

const (
	// SegmentGitBackground is the background color to use
	SegmentGitBackground = 148 // #afd700

	// SegmentGitForeground is the foreground color to use
	SegmentGitForeground = 0 // #000000
)

// SegmentGit implements the git SCM partial of the prompt.
//
// It renders the current branch if the folder contains a repository.
//
// TODO:
// * git rev-parse --show-toplevel : shows root of the repo
// * git status --porcelain -b : shows conmplete status of the repo
type SegmentGit struct {
	Data []byte
}

// NewSegmentGit creates an instace of SegmentGit type.
func NewSegmentGit() *SegmentGit {
	segment := &SegmentGit{}
	return segment
}

// Compile collects the data for this segment.
func (s *SegmentGit) Compile() {
	currentBranch := getCurrentBranch()
	s.Data = []byte(currentBranch)
}

// Render renders the segment results.
func (s *SegmentGit) Render() []byte {
	var b bytes.Buffer
	if len(s.Data) != 0 {
		b.Write(SetBackground(148))
		fmt.Fprint(&b, " ")
		b.Write(SetForeground(0))
		b.Write(s.Data)
		b.Write(ResetForeground())
		fmt.Fprint(&b, " ")
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
