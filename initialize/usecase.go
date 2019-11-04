package initialize

import (
	"bytes"
	"os"
	"path/filepath"
	"text/template"

	"github.com/getshway/shway/domain"
	"github.com/getshway/shway/dotfile"
	"github.com/pkg/errors"
)

const tmpl = `shway() {
	case "$1" in
	init | load)
		source <( {{ .Executable }} $@ ) || {{ .Executable }} $@
		;;
	set | update)
		{{ .Executable }} $@ && source <( {{ .Executable }} load $2 ) || {{ .Executable }} $@
		;;
	*)
		{{ .Executable }} $@
		;;
	esac
}

_shway() {
	IFS=' ' read -A reply <<< "help set update load list init"
}
compctl -K _shway shway

{{- range .Sources }}
source {{ . }}
{{- end }}
`

// Run initialize shway
func Run() (result string, err error) {
	// check OS
	if !domain.IsAvailableOS() {
		err = domain.ErrNotAvailableOS
		return
	}
	// get shway home directory path
	home, err := domain.Home()
	if err != nil {
		err = errors.Wrap(err, "fail to get shway's home directory path")
		return
	}

	// make directories
	dfp := filepath.Join(home, domain.ShwayDefaultDir)
	if err = os.MkdirAll(dfp, 0700); err != nil {
		err = errors.Wrap(err, "fail to create shway's defualt directory")
		return
	}
	pfp := filepath.Join(home, domain.ShwayProjectDir)
	if err = os.MkdirAll(pfp, 0700); err != nil {
		err = errors.Wrap(err, "fail to create shway's projects directory")
		return
	}

	sources := []string{}
	fs, err := dotfile.FindAllShellFiles(dfp)
	if err != nil {
		err = errors.Wrap(err, "fail to load shway's project directories")
		return
	}
	sources = append(sources, fs...)

	fs, err = dotfile.FindAllShellFiles(pfp)
	if err != nil {
		err = errors.Wrap(err, "fail to load shway's project directories")
		return
	}
	sources = append(sources, fs...)

	executable, err := os.Executable()
	if err != nil {
		return
	}

	tpl := template.Must(template.New("init").Parse(tmpl))
	var out bytes.Buffer
	err = tpl.Execute(&out, struct {
		Executable string
		Sources    []string
	}{
		Executable: executable,
		Sources:    sources,
	})
	return out.String(), err
}
