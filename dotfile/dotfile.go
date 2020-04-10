package dotfile

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/getshway/shway/sidekick/sidekick"
)

var (
	ignoreFiles = []string{
		".ds_store",
		"brewfile",
	}
	ignoreDirs = []string{
		".git",
	}
	includeFilePattern = regexp.MustCompile(`^\.{0,1}[a-zA-Z_]+$`)
)

// Paths stores dotfile path
type Paths struct {
	Shells []string
	Brews  []string
}

// FindAllFiles gets all dotfile path
func FindAllFiles(root string) (result Paths, err error) {
	result.Shells = []string{}
	result.Brews = []string{}
	err = filepath.Walk(root, func(fp string, f os.FileInfo, err error) error {
		if fp != root {
			nl := strings.ToLower(f.Name())
			if f.IsDir() && sidekick.InStrings(nl, ignoreDirs) {
				return filepath.SkipDir
			}
			if !f.IsDir() {
				switch true {
				case sidekick.InStrings(nl, []string{"homebrew.yml", "homebrew.yaml"}):
					// get homebrew files
					result.Brews = append(result.Brews, fp)
				case includeFilePattern.MatchString(nl) && !sidekick.InStrings(nl, ignoreFiles):
					// get shell configs
					result.Shells = append(result.Shells, fp)
				}
			}
		}
		return nil
	})
	if err != nil {
		result = Paths{}
	}
	return
}

// FindAllBrewFiles gets all brew file path
func FindAllBrewFiles(root string) (result []string, err error) {
	result = []string{}
	err = filepath.Walk(root, func(fp string, f os.FileInfo, err error) error {
		nl := strings.ToLower(f.Name())
		if fp != root {
			if f.IsDir() && sidekick.InStrings(nl, ignoreDirs) {
				return filepath.SkipDir
			}
			if !f.IsDir() && sidekick.InStrings(nl, []string{"homebrew.yml", "homebrew.yaml"}) {
				result = append(result, fp)
			}
		}
		return nil
	})
	return
}

// FindAllShellFiles gets all dotfile path
func FindAllShellFiles(root string) (result []string, err error) {
	result = []string{}
	err = filepath.Walk(root, func(fp string, f os.FileInfo, err error) error {
		nl := strings.ToLower(f.Name())
		if fp != root {
			if f.IsDir() && sidekick.InStrings(nl, ignoreDirs) {
				return filepath.SkipDir
			}
			if !f.IsDir() && includeFilePattern.MatchString(nl) && !sidekick.InStrings(nl, ignoreFiles) {
				result = append(result, fp)
			}
		}
		return nil
	})
	return
}
