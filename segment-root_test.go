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

import "testing"

func TestSegmentRootCompile(t *testing.T) {
	expected := `\$`
	segment := NewSegmentRoot()
	segment.Compile()
	if string(segment.Data) != expected {
		t.Fatalf("Compiled data expected to be %q but got %q", expected, string(segment.Data))
	}
}

func TestSegmentRootRender(t *testing.T) {
	expected := `\[\e[48;5;236m\] \$ `
	segment := NewSegmentRoot()
	segment.Compile()
	out := segment.Render()
	if string(out) != expected {
		t.Fatalf("Rendered segment expected to be %q but got %q", expected, string(out))
	}
}
