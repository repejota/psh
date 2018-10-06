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
	_, _ = exec.Command("git", "config", "--global", "user.email", "repejota@gmail.com").Output()
	_, _ = exec.Command("git", "config", "--global", "user.name", "Raül Pérez").Output()
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
	_, _ = exec.Command("git", "config", "--global", "user.email", "repejota@gmail.com").Output()
	_, _ = exec.Command("git", "config", "--global", "user.name", "Raül Pérez").Output()
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
