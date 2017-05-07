// Package gostd contains go standard helpers.
package gostd

import (
	"fmt"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"

	"github.com/mh-cbon/emd/emd"
	"github.com/mh-cbon/emd/utils"
)

// PkgDoc Reads the first of the files, or `main.go`, lookup for its package comment and returns it as plain text.
func PkgDoc(files ...string) (string, error) {
	file := "main.go"
	if len(files) > 0 {
		file = files[0]
	}
	if _, err := os.Stat(file); len(files) == 0 && os.IsNotExist(err) {
		return "", nil
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, file, nil, parser.ParseComments)
	if err != nil {
		return "", fmt.Errorf("Failed to parse input file: %v", err)
	}

	if f.Comments == nil || len(f.Comments) == 0 {
		return "Go package documentation not found!", nil
	}

	return f.Comments[0].Text(), nil
}

// GoTest Reads the first of the files, or `main.go`, lookup for its package comment and returns it as plain text.
func GoTest(g *emd.Generator) func(string, string, ...string) (string, error) {
	return func(rpkg, run string, args ...string) (string, error) {
		if rpkg != "" {
			if _, err := os.Stat(rpkg); !os.IsNotExist(err) {
				rpkg, err = filepath.Abs(rpkg)
				if err != nil {
					return "", err
				}
				GOPATH := filepath.Join(os.Getenv("GOPATH"), "src")
				rpkg = strings.Replace(rpkg, GOPATH, "", -1)
				rpkg = strings.Replace(rpkg, "\\", "/", -1) // windows..
			}
		}
		nargs := []string{"test", "-v"}
		if rpkg != "" {
			nargs = append(nargs, rpkg[1:]) // rm front /
		}
		if run != "" {
			nargs = append(nargs, []string{"-run", run}...)
		}
		nargs = append(nargs, args...)
		out, err := utils.Exec("go", nargs)
		if err != nil {
			return "", err
		}

		s := utils.GetCmdStr("go", nargs)
		pre := g.GetSKey("emd_gotest_pre")
		_, err = g.WriteString(pre + s + "\n")

		return strings.TrimSpace(string(out)), err
	}
}

// InstructionGoGetInstall is a template to show instructions to install the package with go get.
var InstructionGoGetInstall = `{{define "go/install" -}}
` + "```sh" + `
go get {{.ProviderURL}}/{{.User}}/{{.Name}}
` + "```" + `
{{- end}}`

// BadgeGoDoc is a template to show a godoc badge.
var BadgeGoDoc = `{{define "badge/godoc" -}}
[!` +
	`[GoDoc]` +
	`(https://godoc.org/{{.ProviderURL}}/{{.User}}/{{.Name}}?status.svg)` +
	`]` +
	`(http://godoc.org/{{.ProviderURL}}/{{.User}}/{{.Name}})
{{- end}}`

// BadgeGoReport is a template to show a goreport badge.
var BadgeGoReport = `{{define "badge/goreport" -}}
[!` +
	`[Go Report Card]` +
	`(https://goreportcard.com/badge/{{.ProviderURL}}/{{.User}}/{{.Name}})` +
	`]` +
	`(https://goreportcard.com/report/{{.ProviderURL}}/{{.User}}/{{.Name}})
{{- end}}`

// Register go standard helpers to the generator.
func Register(g *emd.Generator) error {
	g.AddFunc("pkgdoc", PkgDoc)
	g.AddFunc("gotest", GoTest(g))
	g.AddTemplate(InstructionGoGetInstall)
	g.AddTemplate(BadgeGoDoc)
	g.AddTemplate(BadgeGoReport)
	return nil
}
