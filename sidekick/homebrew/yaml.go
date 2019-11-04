package homebrew

import (
	"fmt"
	"io/ioutil"

	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"
)

// Bundle homebrew request
type Bundle struct {
	Taps     []string  `yaml:"taps"`
	Formulae []Formula `yaml:"formulae"`
	Casks    []Cask    `yaml:"casks"`
}

// Formula is defined a homebrew formula
type Formula struct {
	Name string `yaml:"name"`
}

// Cask is defined a homebrew cask
type Cask struct {
	Name string `yaml:"name"`
}

// Yaml loads yaml file and update and install packages managed by homebrew and homebrew-cask
//    Notice: This function does not upgrade
func Yaml(fp string) (err error) {
	// load yaml file
	bs, err := ioutil.ReadFile(fp)
	if err != nil {
		return errors.Wrap(err, "fail to load homebrew bundle file")
	}
	var b Bundle
	if err = yaml.Unmarshal(bs, &b); err != nil {
		return errors.Wrap(err, "fail to parse homebrew bundle file")
	}

	if err = Update(); err != nil {
		return errors.Wrap(err, "fail to update homebrew before install package")
	}

	if len(b.Taps) > 0 {
		for _, tp := range b.Taps {
			if err = Tap(tp); err != nil {
				return errors.Wrapf(err, fmt.Sprintf("fail to load homebrew tap : %s", tp))
			}
		}
	}

	for _, f := range b.Formulae {
		if err = Install(f.Name); err != nil {
			return errors.Wrapf(err, fmt.Sprintf("fail to install package : %s", f.Name))
		}
	}

	for _, c := range b.Casks {
		if err = InstallCask(c.Name); err != nil {
			return errors.Wrapf(err, fmt.Sprintf("fail to install package (cask): %s", c.Name))
		}
	}

	if err = Cleanup(); err != nil {
		fmt.Println("[warn] fail to clean up homebrew after install")
	}
	return
}
