// Package std contains standard helpers.
package std

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	yaml "gopkg.in/yaml.v2"

	"github.com/mh-cbon/emd/emd"
	"github.com/mh-cbon/emd/utils"
)

// Cat displays a file header, returns file body.
func Cat(g *emd.Generator) func(string) (string, error) {
	return func(f string) (string, error) {
		s, err := Read(f)
		if err != nil {
			return "", err
		}
		pre := g.GetSKey("emd_cat_pre")
		_, err = g.WriteString(pre + f + "\n")
		return strings.TrimSpace(string(s)), err
	}
}

// Read returns file body.
func Read(f string) (string, error) {
	s, err := ioutil.ReadFile(f)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(s)), err
}

// makeMapOf given arguments.
func makeMapOf(keyValuesMap ...interface{}) (map[string]interface{}, error) {
	ret := map[string]interface{}{}
	if len(keyValuesMap) > 0 {
		if len(keyValuesMap)%2 != 0 {
			return ret, fmt.Errorf("Incorrect arguments number in call to render template function, args are: %#v", keyValuesMap)
		}
		for i := 0; i < len(keyValuesMap); i += 2 {
			key, ok := keyValuesMap[i].(string)
			if ok == false {
				return ret, fmt.Errorf("Incorrect key type %T of arg %#v in call to render template function, expected a string, args are: %#v",
					keyValuesMap[i],
					keyValuesMap[i],
					keyValuesMap)
			}
			ret[key] = keyValuesMap[i+1]
		}
	}
	return ret, nil
}

// Render renders template wih given name, and map arguments to its data argument.
func Render(g *emd.Generator) func(string, map[string]interface{}, ...interface{}) (string, error) {
	return func(name string, data map[string]interface{}, keyValuesMap ...interface{}) (string, error) {
		extraData := map[string]interface{}{}
		for k, v := range data {
			extraData[k] = v
		}
		y, err := makeMapOf(keyValuesMap...)
		if err != nil {
			return "", err
		}
		for k, v := range y {
			extraData[k] = v
		}
		return "", g.GetTemplate().ExecuteTemplate(g.GetOut(), name, extraData)
	}
}

// Exec display a program invokation header, returns its output.
func Exec(g *emd.Generator) func(string, ...string) (string, error) {
	return func(bin string, args ...string) (string, error) {
		out, err := utils.Exec(bin, args)
		if err != nil {
			return "", err
		}

		f := utils.GetCmdStr(bin, args)
		pre := g.GetSKey("emd_exec_pre")
		_, err = g.WriteString(pre + f + "\n")

		return strings.TrimSpace(out), err
	}
}

// Shell display a cli invokation header, returns the cli output.
func Shell(g *emd.Generator) func(string) (string, error) {
	return func(s string) (string, error) {
		out, err := utils.Shell("", s)
		if err != nil {
			return "", err
		}

		pre := g.GetSKey("emd_shell_pre")
		_, err = g.WriteString(pre + s + "\n")

		return strings.TrimSpace(string(out)), err
	}
}

// Color embeds a text block with markdown triple backquotes syntax makrup.
func Color(syntax, content string) string {
	if content == "" && syntax != "" {
		content = syntax
		syntax = "sh" // set the default color
	}
	return fmt.Sprintf("```%v\n%v\n```", syntax, content)
}

// Yaml parses given file as yaml, locate given path, build a new map, yaml encode it, returns its string.
func Yaml(file string, paths ...string) (string, error) {
	s, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}

	m := yaml.MapSlice{}
	err = yaml.Unmarshal(s, &m)
	if err != nil {
		return "", err
	}

	res := yaml.MapSlice{}
	if len(paths) > 0 {
		// make it more complex later
		for _, p := range paths {
			for _, k := range m {
				if k.Key.(string) == p {
					res = append(res, k)

				}
			}
		}
	} else {
		res = m
	}

	var d []byte
	d, err = yaml.Marshal(&res)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(d)), nil
}

// Preline prepends every line of content with pre.
func Preline(pre, content string) string {
	res := ""
	for _, c := range content {
		res += string(c)
		if c == '\n' {
			res += pre
		}
	}
	if res != "" {
		res = pre + res
	}
	return res
}

// Echo prints given strings.
func Echo(content ...string) string {
	return strings.Join(content, " ")
}

var replaceIndex = 0

// Toc generates and prints a TOC.
func Toc(g *emd.Generator) func(int, ...string) string {
	return func(depth int, toctitles ...string) string {
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
				x := titles[:0]
				for _, t := range titles {
					if t.Line > lineIndex {
						x = append(x, t)
					}
				}
				root := utils.MakeTitleTree(x)
				toc := ""
				e := -1
				ee := -1
				root.Traverse(utils.PowerLess(5, func(n *utils.MdTitleTree) {
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
					if n.Duplicate > -1 {
						link += fmt.Sprintf("-%v", n.Duplicate)
					}
					x := strings.Repeat("  ", e)
					toc += fmt.Sprintf("%v- [%v](#%v)\n", x, n.Title, link)
				}))

				lines[lineIndex] = strings.Replace(lines[lineIndex], replaceToken, toctitle, -1)
				lines = append(lines[:lineIndex+1], lines[lineIndex:]...)
				lines[lineIndex+1] = strings.TrimRight(toc, "\n")
				return strings.Join(lines, "\n")
			}
			log.Println("weird, a toc was generated, but it was not added to the final content.")
			return s
		})
		return replaceToken
	}
}

// GHReleasePages is a template to show a notice about the gh releases page.
var GHReleasePages = `{{define "gh/releases" -}}
Check the [release page](https://github.com/{{.User}}/{{.Name}}/releases)!
{{- end}}`

// BadgeTravis is a template to show a travis badge.
var BadgeTravis = `{{define "badge/travis" -}}
[!` +
	`[travis Status]` +
	`(https://travis-ci.org/{{.User}}/{{.Name}}.svg?branch={{.Branch}})` +
	`]` +
	`(https://travis-ci.org/{{.User}}/{{.Name}})
{{- end}}`

// BadgeAppveyor is a template to show an appveyor badge.
var BadgeAppveyor = `{{define "badge/appveyor"}}
[!` +
	`[appveyor Status]` +
	`(https://ci.appveyor.com/api/projects/status/{{.ProviderName}}/{{.User}}/{{.Name}}?branch={{.Branch}}&svg=true)` +
	`]` +
	`(https://ci.appveyor.com/project/{{.User}}/{{.Name}})
{{- end}}`

// BadgeCodeship is a template to show a codehsip badge.
var BadgeCodeship = `{{define "badge/codeship" -}}
[!` +
	`[codeship Status]` +
	`(https://codeship.com/projects/{{.CsUUID}}/status?branch={{.Branch}})` +
	`]` +
	`(https://codeship.com/projects/{{.CsProjectID}})` + `
{{- end}}`

// InstructionChocoInstall is a template to show instructions to install the package with chocolatey.
var InstructionChocoInstall = `{{define "choco/install" -}}
` + "```sh" + `
choco install {{.Name}}
` + "```" + `
{{- end}}`

// InstructionGhRepo is a template to show instructions to install the rpm/deb repositories with gh-pages.
var InstructionGhRepo = `{{define "linux/gh_src_repo" -}}
` + "```sh" + `
wget -O - https://raw.githubusercontent.com/mh-cbon/latest/master/source.sh \
| GH={{.User}}/{{.Name}} sh -xe
# or
curl -L https://raw.githubusercontent.com/mh-cbon/latest/master/source.sh \
| GH={{.User}}/{{.Name}} sh -xe
` + "```" + `
{{- end}}`

// InstructionGhPkg is a template to show instructions to install the rpm/deb package with gh-pages.
var InstructionGhPkg = `{{define "linux/gh_pkg" -}}
` + "```sh" + `
curl -L https://raw.githubusercontent.com/mh-cbon/latest/master/install.sh \
| GH={{.User}}/{{.Name}} sh -xe
# or
wget -q -O - --no-check-certificate \
https://raw.githubusercontent.com/mh-cbon/latest/master/install.sh \
| GH={{.User}}/{{.Name}} sh -xe
` + "```" + `
{{- end}}`

// BadgeLicense shows a badge for a license.
var BadgeLicense = `{{define "license/shields" -}}
[!` +
	`[{{.License}} License]` +
	`(http://img.shields.io/badge/License-{{.License}}-{{or .LicenseColor "blue"}}.svg)` +
	`]` +
	`({{or .LicenseFile "LICENSE"}})
{{- end}}`

// Register standard helpers to the generator.
func Register(g *emd.Generator) error {

	g.AddFunc("cat", Cat(g))
	g.AddFunc("read", Read)
	g.AddFunc("render", Render(g))
	g.AddFunc("exec", Exec(g))
	g.AddFunc("shell", Shell(g))
	g.AddFunc("color", Color)
	g.AddFunc("toc", Toc(g))
	g.AddFunc("yaml", Yaml)
	g.AddFunc("preline", Preline)
	g.AddFunc("echo", Echo)

	g.AddTemplate(GHReleasePages)
	g.AddTemplate(BadgeTravis)
	g.AddTemplate(BadgeAppveyor)
	g.AddTemplate(BadgeCodeship)
	g.AddTemplate(InstructionChocoInstall)
	g.AddTemplate(InstructionGhRepo)
	g.AddTemplate(InstructionGhPkg)
	g.AddTemplate(BadgeLicense)

	return nil
}
