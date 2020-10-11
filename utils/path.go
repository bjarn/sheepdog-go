package utils

import "os"

// Get the absolute path to the user's home directory or fallback to the relative path.
func UserHomeDir() string {
	path, err := os.UserHomeDir()
	if err != nil {
		return "~/"
	}
	return path
}

// Get the absolute path to sheepdog's home directory or fallback to the relative path.
func SheepdogHomeDir() string {
	path, err := os.UserHomeDir()
	if err != nil {
		return "~/"
	}
	return path + "/.sheepdog"
}