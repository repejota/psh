package main

import "testing"

func TestReset(t *testing.T) {
	var prompt Prompt
	prompt.reset()
	if prompt.Prompt != "$(tput sgr0)" {
		t.Error("Expected '$(tput sgr0)', but got ", prompt.Prompt)
	}
}

func TestSetBg(t *testing.T) {
	var prompt Prompt
	prompt.setBg(1)
	if prompt.Prompt != "$(tput setab 1)" {
		t.Error("Expected '$(tput setab 1)', but got ", prompt.Prompt)
	}
}

func TestSetFg(t *testing.T) {
	var prompt Prompt
	prompt.setFg(1)
	if prompt.Prompt != "$(tput setaf 1)" {
		t.Error("Expected '$(tput setaf 1)', but got ", prompt.Prompt)
	}
}

func TestSetColor(t *testing.T) {
	var prompt Prompt
	prompt.setColor(1, 2)
	if prompt.Prompt != "$(tput setaf 1)$(tput setab 2)" {
		t.Error("Expected '$(tput setaf 1)$(tput setab 2)', but got ", prompt.Prompt)
	}
}
