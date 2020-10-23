package sheepdog

import "os"

// Get the absolute path to sheepdog's home directory or fallback to the relative path.
func HomeDir() string {
	path, err := os.UserHomeDir()
	if err != nil {
		return "~/"
	}
	return path + "/.sheepdog"
}

// Create the Sheepdog home directory
func CreateHomeDir() error {
	return os.Mkdir(HomeDir(), 0755)
}