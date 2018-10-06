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
	"fmt"
)

const (
	// colorTemplate is an escaped template to apply color to text using ANSI
	// escape sequences.
	colorTemplate = `\[\e%s\]`
)

// ResetFgBg resets foreground and background to default values.
func ResetFgBg() []byte {
	escapeSequence := fmt.Sprintf(colorTemplate, "[0m")
	return []byte(escapeSequence)
}

// ResetForeground reset foreground to default value.
func ResetForeground() []byte {
	escapeSequence := fmt.Sprintf(colorTemplate, "[39m")
	return []byte(escapeSequence)
}

// ResetBackground reset background to default value.
func ResetBackground() []byte {
	escapeSequence := fmt.Sprintf(colorTemplate, "[49m")
	return []byte(escapeSequence)
}

// SetBackground sets background color on the terminal
func SetBackground(background uint8) []byte {
	bg := fmt.Sprintf("[48;5;%dm", background)
	escapeSequence := fmt.Sprintf(colorTemplate, bg)
	return []byte(escapeSequence)
}

// SetForeground sets foreground color on the terminal.
func SetForeground(foreground uint8) []byte {
	fg := fmt.Sprintf("[38;5;%dm", foreground)
	escapeSequence := fmt.Sprintf(colorTemplate, fg)
	return []byte(escapeSequence)
}

// SetTrueColorBackground sets background color on the terminal using true
// color format.
func SetTrueColorBackground(r, g, b int) []byte {
	fg := fmt.Sprintf("[48;2;%d;%d;%dm", r, g, b)
	escapeSequence := fmt.Sprintf(colorTemplate, fg)
	return []byte(escapeSequence)
}

// SetTrueColorForeground sets foreground color on the terminal using true
// color format.
func SetTrueColorForeground(r, g, b int) []byte {
	fg := fmt.Sprintf("[38;2;%d;%d;%dm", r, g, b)
	escapeSequence := fmt.Sprintf(colorTemplate, fg)
	return []byte(escapeSequence)
}
