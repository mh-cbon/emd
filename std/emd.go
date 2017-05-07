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

// Concat multiple strings.
func Concat(s ...interface{}) string {
	ss := []string{}
	for _, v := range s {
		ss = append(ss, fmt.Sprintf("%v", v))
	}
	return strings.Join(ss, "")
}

// PathJoin multiple path parts.
func PathJoin(s ...interface{}) string {
	ss := []string{}
	for _, v := range s {
		ss = append(ss, fmt.Sprintf("%v", v))
	}
	return strings.Join(ss, "/")
}

// Link creates a link according to markdown syntax.
func Link(url interface{}, texts ...interface{}) string {
	var text interface{}
	text = ""
	if len(texts) > 0 {
		text = texts[0]
	}
	if text == "" {
		return fmt.Sprintf("%v", url)
	}
	return fmt.Sprintf("[%v](%v)", text, url)
}

// Img creates a img according to markdown syntax.
func Img(url interface{}, alts ...interface{}) string {
	var alt interface{}
	alt = ""
	if len(alts) > 0 {
		alt = alts[0]
	}
	return fmt.Sprintf("![%v](%v)", alt, url)
}

// SetValue saves a value.
func SetValue(g *emd.Generator) func(string, interface{}) string {
	return func(name string, v interface{}) string {
		g.SetData(name, v)
		return ""
	}
}

// GetValue returns a value.
func GetValue(g *emd.Generator) func(string) interface{} {
	return func(name string) interface{} {
		return g.GetKey(name)
	}
}

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

				lastPower := -1
				root.Traverse(utils.PowerLess(5, func(n *utils.MdTitleTree) {
					if n.Power < 2 {
						e = 0
					} else if lastPower < n.Power {
						e++
					} else if lastPower > n.Power {
						e--
					}
					if e < 0 {
						e = 0
					}
					lastPower = n.Power

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
Check the {{link (concat "https://" .URL "/releases") "release page"}}!
{{- end}}`

// BadgeTravis is a template to show a travis badge.
var BadgeTravis = `{{define "badge/travis" -}}
{{- set "travisUrl" (pathjoin "https://travis-ci.org" .User .Name) }}
{{- set "travisImg" (img (concat .travisUrl ".svg?branch=" .Branch) "travis Status") }}
{{- link .travisUrl .travisImg}}
{{- end}}`

// BadgeAppveyor is a template to show an appveyor badge.
var BadgeAppveyor = `{{define "badge/appveyor" -}}
{{- set "appveyorStatusUrl" (pathjoin "https://ci.appveyor.com/api/projects/status" .ProviderName .User .Name) }}
{{- set "appveyorProjectUrl" (pathjoin "https://ci.appveyor.com/projects" .User .Name) }}
{{- set "appveyorImg" (img (concat .appveyorStatusUrl "?branch=" .Branch "&svg=true") "Appveyor Status") }}
{{- link .appveyorProjectUrl .appveyorImg }}
{{- end}}`

// BadgeCodeship is a template to show a codehsip badge.
var BadgeCodeship = `{{define "badge/codeship" -}}
{{- set "csTitle" (or .CsTitle "Codeship Status") }}
{{- set "csStatusUrl" (pathjoin "https://codeship.com/projects" .CsUUID "status") }}
{{- set "csProjectUrl" (pathjoin "https://codeship.com/projects" .CsProjectID) }}
{{- set "csImg" (img (concat (get "csStatusUrl") "?branch=" .Branch) (get "csTitle") ) }}
{{- link (get "csProjectUrl") (get "csImg") }}
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
{{- set "licenseFile" (or .LicenseFile "LICENSE") }}
{{- set "licenseTitle" (concat .License " License") }}
{{- set "licenseImg" (or .LicenseColor "blue") }}
{{- set "licenseImg" (concat "License-" .License "-" (get "licenseImg") ".svg") }}
{{- set "licenseImg" (concat "http://img.shields.io/badge/" (get "licenseImg")) }}
{{- set "licenseImg" (img (get "licenseImg") (get "licenseTitle")) }}
{{- link (get "licenseFile") (get "licenseImg") }}
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
	g.AddFunc("link", Link)
	g.AddFunc("img", Img)
	g.AddFunc("concat", Concat)
	g.AddFunc("pathjoin", PathJoin)
	g.AddFunc("set", SetValue(g))
	g.AddFunc("get", GetValue(g))

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
