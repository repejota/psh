package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func getGitBranch() string {
	out, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	if err != nil {
		return ""
	}
	return strings.Trim(string(out), "\n")
}

func getGitChanges() int {
	out, err := exec.Command("git", "status", "-s", "--porcelain").Output()
	if err != nil {
		return -1
	}
	changes := 0
	s := bufio.NewScanner(bytes.NewReader(out))
	for s.Scan() {
		changes++
	}
	return changes
}

func getGitPartial() string {
	partial := ""
	if existsPath(".git") {
		changes := getGitChanges()
		if changes > 0 {
			partial = fmt.Sprintf("%d  ", changes)
		}
		partial = fmt.Sprintf("%s%s", partial, getGitBranch())
	}
	return partial
}
