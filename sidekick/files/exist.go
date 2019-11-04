package files

import "os"

// IsExist returns a boolean indicating whether the error is known to report that a file or directory exists.
func IsExist(p string) bool {
	return !IsNotExist(p)
}

// IsNotExist returns a boolean indicating whether the error is known to report that a file or directory does not exist.
func IsNotExist(p string) bool {
	_, err := os.Stat(p)
	return os.IsNotExist(err)
}
