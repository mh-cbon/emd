package std

import (
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/mh-cbon/emd/emd"
)

// Register standard helpers to the generator.
func Register(g *emd.Generator) error {

	g.AddFunc("file", func(f string) (string, error) {
		s, err := ioutil.ReadFile(f)
		ext := filepath.Ext(f)
		ext = strings.TrimPrefix(ext, ".")
		res := `
__> ` + f + `__
` + "```" + ext + `
` + strings.TrimSpace(string(s)) + `
` + "```"
		return res, err
	})

	g.AddFunc("cli", func(bin string, args ...string) (string, error) {
		cmd := exec.Command(bin, args...)
		out, err := cmd.CombinedOutput()
		fbin := filepath.Base(bin)
		res := `
__$ ` + fbin + ` ` + strings.Join(args, " ") + `__
` + "```sh" + `
` + strings.TrimSpace(string(out)) + `
` + "```"
		return res, err
	})

	g.AddTemplate(`{{define "gh/releases" -}}
Check the [release page](https://github.com/{{.User}}/{{.Name}}/releases)!
{{- end}}`)

	g.AddTemplate(`{{define "badge/travis" -}}
[![travis Status](https://travis-ci.org/{{.User}}/{{.Name}}.svg?branch={{.Branch}})](https://travis-ci.org/{{.User}}/{{.Name}})
{{- end}}`)

	g.AddTemplate(`{{define "badge/appveyor" -}}
[![appveyor Status](https://ci.appveyor.com/api/projects/status/{{.ProviderName}}/{{.User}}/{{.Name}}?branch={{.Branch}}&svg=true)](https://ci.appveyor.com/project/{{.User}}/{{.Name}})
{{- end}}`)

	g.AddTemplate(`{{define "choco/install" -}}
` + "```sh" + `
choco install {{.Name}}
` + "```" + `
{{- end}}`)

	g.AddTemplate(`{{define "linux/gh_src_repo" -}}
` + "```sh" + `
wget -O - https://raw.githubusercontent.com/mh-cbon/latest/master/source.sh \
| GH={{.User}}/{{.Name}} sh -xe
# or
curl -L https://raw.githubusercontent.com/mh-cbon/latest/master/source.sh \
| GH={{.User}}/{{.Name}} sh -xe
` + "```" + `
{{- end}}`)

	g.AddTemplate(`{{define "linux/gh_pkg" -}}
` + "```sh" + `
curl -L https://raw.githubusercontent.com/mh-cbon/latest/master/install.sh \
| GH={{.User}}/{{.Name}} sh -xe
# or
wget -q -O - --no-check-certificate \
https://raw.githubusercontent.com/mh-cbon/latest/master/install.sh \
| GH={{.User}}/{{.Name}} sh -xe
` + "```" + `
{{- end}}`)

	return nil
}
