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
	"testing"
)

func TestGetSegmentsList(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{"", 0},
		{"a,b", 2},
		{"a,,c", 2},
	}
	for _, tt := range tests {
		output := getSegmentsList(tt.input)
		if len(output) != tt.expected {
			t.Fatalf("Expected to have %d elements but got %d", tt.expected, len(output))
		}
	}
}

func TestNewPrompt(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{"", 0},
		{"a,b", 2},
		{"a,,c", 2},
	}
	for _, tt := range tests {
		prompt := NewPrompt(tt.input)
		if len(prompt.keys) != tt.expected {
			t.Fatalf("Prompt expected to have %d keys but got %d", tt.expected, len(prompt.keys))
		}
	}
}
