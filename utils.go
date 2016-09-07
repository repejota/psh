package main

import "os"

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
