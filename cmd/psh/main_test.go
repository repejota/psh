// Copyright 2016-2017 The psh Authors. All rights reserved.

package main

import "testing"

func TestShowVersion(t *testing.T) {
	Version = "0.0.12"
	Build = "develop-16ddeed"
	expectedVersionInfo := "psh : Version 0.0.12 Build develop-16ddeed"
	versionInfo := showVersion()
	if versionInfo != expectedVersionInfo {
		t.Errorf("Expected '%s' but got %s", expectedVersionInfo, versionInfo)
	}
}
