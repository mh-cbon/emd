// Package std contains standard helpers.
package std

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/mh-cbon/emd/emd"
	"github.com/mh-cbon/emd/utils"
)

// Register standard helpers to the generator.
func Register(g *emd.Generator) error {

	g.AddFunc("cat", func(f string) (string, error) {
		s, err := ioutil.ReadFile(f)
		if err != nil {
			return "", err
		}
		title := "\n###### > " + f + "\n"
		_, err = g.GetOut().Write([]byte(title))
		return strings.TrimSpace(string(s)), err
	})

	g.AddFunc("render", func(name string, data map[string]interface{}, keyValuesMap ...interface{}) (string, error) {

		extraData := map[string]interface{}{}
		for k, v := range data {
			extraData[k] = v
		}
		if len(keyValuesMap) > 0 {
			if len(keyValuesMap)%2 != 0 {
				return "", fmt.Errorf("Incorrect arguments number in call to render template function, args are: %#v", keyValuesMap)
			}
			for i := 0; i < len(keyValuesMap); i += 2 {
				key, ok := keyValuesMap[i].(string)
				if ok == false {
					return "", fmt.Errorf("Incorrect key type %T of arg %#v in call to render template function, expected a string, args are: %#v",
						keyValuesMap[i],
						keyValuesMap[i],
						keyValuesMap)
				}
				extraData[key] = keyValuesMap[i+1]
			}
		}

		err := g.GetTemplate().ExecuteTemplate(g.GetOut(), name, extraData)

		return "", err
	})

	g.AddFunc("exec", func(bin string, args ...string) (string, error) {
		out, err := utils.Exec(bin, args)
		if err != nil {
			return "", err
		}
		title := "\n###### $ " + utils.GetCmdStr(bin, args) + "\n"
		_, err = g.GetOut().Write([]byte(title))
		return strings.TrimSpace(out), err
	})

	g.AddFunc("color", func(syntax, content string) string {
		if content == "" && syntax != "" {
			content = syntax
			syntax = "sh" // set the default color
		}
		return fmt.Sprintf("```%v\n%v\n```", syntax, content)
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

	g.AddTemplate(`{{define "badge/codeship" -}}
[![codeship Status](https://codeship.com/projects/{{.CsUUID}}/status?branch={{.Branch}})](https://codeship.com/{{.CsUUID}})
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
