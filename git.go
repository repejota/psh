package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func getGitBranch() string {
	out, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	if err != nil {
		return ""
	}
	return strings.Trim(string(out), "\n")
}

func getGitChanges() string {
	out, err := exec.Command("git", "status", "-s", "--porcelain").Output()
	if err != nil {
		return ""
	}
	counter := 0
	s := bufio.NewScanner(bytes.NewReader(out))
	for s.Scan() {
		counter++
	}
	return strconv.Itoa(counter)
}

func getGitPartial() string {
	return fmt.Sprintf("%s:%s", getGitChanges(), getGitBranch())
}
