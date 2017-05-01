package partial

import (
	"bufio"
	"bytes"
	"os"
	"os/exec"
	"strings"

	"github.com/repejota/psh/environment"
)

// GitPartial ...
type GitPartial struct {
}

// NewGitPartial ...
func NewGitPartial() (gp GitPartial) {
	return gp
}

// Run ...
func (gp GitPartial) Run(e *environment.Environment) {
	if existsPath(".git") {
		e.Git.IsAGitRepo = true
		e.Git.Branch = getGitBranch()
		e.Git.Changes = getGitChanges()
	}
}

// existsPath returns whether the given file or directory exists or not
func existsPath(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

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
