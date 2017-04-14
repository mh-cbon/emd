// Package std contains standard helpers.
package std

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"

	"github.com/mh-cbon/emd/emd"
	"github.com/mh-cbon/emd/utils"
)

var mdTitle = regexp.MustCompile(`^([#]{1,6})\s*(.+)`)

type tocTitle struct {
	w int
	t string
}

// Register standard helpers to the generator.
func Register(g *emd.Generator) error {

	// cat a file and returns its body.
	g.AddFunc("cat", func(f string) (string, error) {
		s, err := ioutil.ReadFile(f)
		if err != nil {
			return "", err
		}
		pre := g.GetSKey("emd_cat_pre")
		_, err = g.WriteString(pre + f + "\n")
		return strings.TrimSpace(string(s)), err
	})

	// render a template with args
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

	// exec a program with args.
	g.AddFunc("exec", func(bin string, args ...string) (string, error) {
		out, err := utils.Exec(bin, args)
		if err != nil {
			return "", err
		}

		f := utils.GetCmdStr(bin, args)
		pre := g.GetSKey("emd_exec_pre")
		_, err = g.WriteString(pre + f + "\n")

		return strings.TrimSpace(out), err
	})

	// exec a program with args.
	g.AddFunc("shell", func(s string) (string, error) {
		out, err := utils.Shell("", s)
		if err != nil {
			return "", err
		}

		pre := g.GetSKey("emd_shell_pre")
		_, err = g.WriteString(pre + s + "\n")

		return strings.TrimSpace(string(out)), err
	})

	// surround a text block with markdown triple backquotes syntax makrup.
	g.AddFunc("color", func(syntax, content string) string {
		if content == "" && syntax != "" {
			content = syntax
			syntax = "sh" // set the default color
		}
		return fmt.Sprintf("```%v\n%v\n```", syntax, content)
	})

	replaceIndex := 0
	// generate a toc
	g.AddFunc("toc", func(depth int, toctitles ...string) string {
		toctitle := "TOC"
		if len(toctitles) > 0 {
			toctitle = toctitles[0]
		}
		replaceToken := fmt.Sprintf("%v%v", "REPLACETOKENGOESHERE", replaceIndex)
		replaceIndex++

		g.AddPostProcess(func(s string) string {
			// a quick and dirty md parser of titles (###) and block (```)
			lineIndex := utils.LineIndex(s, replaceToken)
			if lineIndex > -1 {

				lines := strings.Split(s, "\n")

				titles := utils.GetAllMdTitles(s)
				root := utils.MakeTitleTree(titles)
				toc := ""
				e := -1
				ee := -1
				root.Traverse(utils.LineGreater(lineIndex, utils.PowerLess(5, func(n *utils.MdTitleTree) {
					if n.Power > ee {
						e++
					} else if n.Power < ee {
						e--
						if e < 0 {
							e = 0
						}
					}
					ee = n.Power
					link := utils.GetMdLinkHash(n.Title)
					x := strings.Repeat("  ", e)
					toc += fmt.Sprintf("%v- [%v](#%v)\n", x, n.Title, link)
				})))

				lines[lineIndex] = strings.Replace(lines[lineIndex], replaceToken, toctitle, -1)
				lines = append(lines[:lineIndex+1], lines[lineIndex:]...)
				lines[lineIndex+1] = strings.TrimRight(toc, "\n")
				return strings.Join(lines, "\n")
			}
			log.Println("weird, a toc was generated, but it was not added to the final content.")
			return s
		})
		return replaceToken
	})

	g.AddTemplate(`{{define "gh/releases" -}}
Check the [release page](https://github.com/{{.User}}/{{.Name}}/releases)!
{{- end}}`)

	g.AddTemplate(`{{define "badge/travis" -}}
[!` +
		`[travis Status]` +
		`(https://travis-ci.org/{{.User}}/{{.Name}}.svg?branch={{.Branch}})` +
		`]` +
		`(https://travis-ci.org/{{.User}}/{{.Name}})
{{- end}}`)

	g.AddTemplate(`{{define "badge/appveyor"}}
[!` +
		`[appveyor Status]` +
		`(https://ci.appveyor.com/api/projects/status/{{.ProviderName}}/{{.User}}/{{.Name}}?branch={{.Branch}}&svg=true)` +
		`]` +
		`(https://ci.appveyor.com/project/{{.User}}/{{.Name}})
	{{- end}}`)

	g.AddTemplate(`{{define "badge/codeship" -}}
[!` +
		`[codeship Status]` +
		`(https://codeship.com/projects/{{.CsUUID}}/status?branch={{.Branch}})` +
		`]` +
		`(https://codeship.com/projects/{{.CsProjectID}})` + `
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

	g.AddTemplate(`{{define "license/shields" -}}
[!` +
		`[{{.License}} License]` +
		`(http://img.shields.io/badge/License-{{.License}}-{{or .LicenseColor "blue"}}.svg)` +
		`]` +
		`({{or .LicenseFile "LICENSE"}})
{{- end}}`)

	return nil
}
