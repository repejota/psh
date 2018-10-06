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
	"testing"
)

func TestSegmentVirtualEnvCompile(t *testing.T) {
	err := os.Setenv("VIRTUAL_ENV", "/home/user/venv")
	if err != nil {
		t.Fatalf("Can't set the required enviornment.")
	}
	expected := `venv`
	segment := NewSegmentVirtualEnv()
	segment.Compile()
	if string(segment.Data) != expected {
		t.Fatalf("Compiled data expected to be %q but got %q", expected, string(segment.Data))
	}
}

func TestSegmentVirtualEnvRender(t *testing.T) {
	err := os.Setenv("VIRTUAL_ENV", "/home/user/venv")
	if err != nil {
		t.Fatalf("Can't set the required enviornment.")
	}
	expected := `\[\e[48;5;22m\]  venv `
	segment := NewSegmentVirtualEnv()
	segment.Compile()
	out := segment.Render()
	if string(out) != expected {
		t.Fatalf("Rendered segment expected to be %q but got %q", expected, string(out))
	}
}
