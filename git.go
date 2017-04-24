package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// getGitBranch returns the current git branch name
func getGitBranch() string {
	out, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	if err != nil {
		return ""
	}
	return strings.Trim(string(out), "\n")
}

// getGitChanges counts the number of changes on this repo
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

// getGitPartial builds Git partial string
func getGitPartial() string {
	partial := ""
	changes := getGitChanges()
	partial = fmt.Sprintf("%s%s", partial, " ")
	if changes > 0 {
		partial = fmt.Sprintf("%s%d ", partial, changes)
	}
	partial = fmt.Sprintf("%s%s", partial, getGitBranch())
	return partial
}
