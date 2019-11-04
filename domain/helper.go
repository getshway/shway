package domain

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/mitchellh/go-homedir"
)

// IsAvailableOS checks if OS is allowed to run by shway
func IsAvailableOS() (result bool) {
	return (runtime.GOOS == "darwin")
}

// Home get shway's home directory
func Home() (result string, err error) {
	if result = os.Getenv("SHWAY_HOME"); result != "" {
		return
	}
	h, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	return filepath.Join(h, ".shway"), nil
}
