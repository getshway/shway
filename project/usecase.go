package project

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/getshway/shway/domain"
	"github.com/getshway/shway/dotfile"
	"github.com/getshway/shway/sidekick/files"
	"github.com/getshway/shway/sidekick/git"
	"github.com/getshway/shway/sidekick/homebrew"

	"github.com/pkg/errors"
)

const tmpl = `{{- range . }}
source {{ . }}
{{- end }}
`

// List project directory
func List() (result []string, err error) {
	home, err := domain.Home()
	if err != nil {
		return nil, errors.Wrap(err, "fail to get shway's home directory path")
	}
	pd := filepath.Join(home, domain.ShwayProjectDir)
	if files.IsNotExist(pd) {
		return nil, errors.New(fmt.Sprintf("directory is not found : %s", home))
	}
	fs, err := ioutil.ReadDir(pd)
	if err != nil {
		return nil, errors.Wrapf(err, "fail to search directory : %s", home)
	}
	result = []string{"default"}
	for _, f := range fs {
		if f.IsDir() {
			result = append(result, f.Name())
		}
	}
	return
}

// Load shell configs
func Load(name string) (result string, err error) {
	// get shway home directory path
	home, err := domain.Home()
	if err != nil {
		err = errors.Wrap(err, "fail to get shway's home directory path")
		return
	}

	pd := getProjectDirPath(home, name)
	if files.IsNotExist(pd) {
		if err = os.Mkdir(pd, 0700); err != nil {
			err = errors.Wrapf(err, "not found project directory and fail to create a directory : %s", pd)
			return
		}
	}

	sources, err := dotfile.FindAllShellFiles(pd)
	if err != nil {
		err = errors.Wrap(err, "fail to load shway's project directories")
		return
	}

	tpl := template.Must(template.New("load").Parse(tmpl))
	var out bytes.Buffer
	err = tpl.Execute(&out, sources)
	return out.String(), err
}

// Set ...
func Set(name, rp string) (err error) {
	// get shway home directory path
	home, err := domain.Home()
	if err != nil {
		return errors.Wrap(err, "fail to get shway's home directory path")
	}

	pd := getProjectDirPath(home, name)
	if files.IsNotExist(pd) {
		if err = os.Mkdir(pd, 0700); err != nil {
			return errors.Wrapf(err, "not found project directory and fail to create a directory : %s", pd)
		}
	}

	if _, err = git.Clone(pd, rp); err != nil {
		return
	}

	bs, err := dotfile.FindAllBrewFiles(pd)
	if err != nil {
		return
	}
	// run homebrew
	for _, bw := range bs {
		if err = homebrew.Yaml(bw); err != nil {
			return
		}
	}

	return
}

// Update ...
func Update(name string) (err error) {
	// get shway home directory path
	home, err := domain.Home()
	if err != nil {
		return errors.Wrap(err, "fail to get shway's home directory path")
	}

	pd := getProjectDirPath(home, name)
	if files.IsNotExist(pd) {
		if err = os.Mkdir(pd, 0700); err != nil {
			return errors.Wrapf(err, "not found project directory and fail to create a directory : %s", pd)
		}
	}

	// update a directory
	if files.IsExist(filepath.Join(pd, ".git")) {
		if _, err = git.Pull(pd); err != nil {
			return errors.Wrap(err, "fail to update")
		}
	}

	bs, err := dotfile.FindAllBrewFiles(pd)
	if err != nil {
		return
	}
	// run homebrew
	for _, bw := range bs {
		if err = homebrew.Yaml(bw); err != nil {
			return
		}
	}

	return
}

func getProjectDirPath(home, name string) (p string) {
	if name == "default" {
		return filepath.Join(home, domain.ShwayDefaultDir)
	}
	return filepath.Join(home, domain.ShwayProjectDir, name)
}
