package files

import (
	"os"
)

// MakeDirIfNotExists creates directory if doesn' t exist yet
func MakeDirIfNotExists(p string) error {
	if IsExist(p) {
		return nil
	}
	return os.Mkdir(p, 0755)
}
