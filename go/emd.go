// Package gostd contains go standard helpers.
package gostd

import (
	"fmt"
	"go/parser"
	"go/token"
	"strings"

	"github.com/mh-cbon/emd/emd"
	"github.com/mh-cbon/emd/utils"
)

// Register go standard helpers to the generator.
func Register(g *emd.Generator) error {

	g.AddFunc("pkgdoc", func(files ...string) (string, error) {
		file := "main.go"
		if len(files) > 0 {
			file = files[0]
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
	})

	g.AddFunc("gotest", func(run string, args ...string) (string, error) {
		args = append([]string{"test", "-v", "-run", run}, args...)
		out, err := utils.Exec("go", args)
		if err != nil {
			return "", err
		}
		title := "\n###### $ " + utils.GetCmdStr("go", args) + "\n"
		_, err = g.GetOut().Write([]byte(title))
		return strings.TrimSpace(string(out)), err
	})

	g.AddTemplate(`{{define "go/install"}}
` + "```" + `sh
go get {{.ProviderURL}}/{{.User}}/{{.Name}}
` + "```" + `
{{end}}`)

	g.AddTemplate(`{{define "badge/godoc"}}
[![GoDoc](https://godoc.org/{{.ProviderURL}}/{{.User}}/{{.Name}}?status.svg)](http://godoc.org/{{.ProviderURL}}/{{.User}}/{{.Name}})
{{end}}`)

	g.AddTemplate(`{{define "badge/goreport"}}
[![Go Report Card](https://goreportcard.com/badge/{{.ProviderURL}}/{{.User}}/{{.Name}})](https://goreportcard.com/report/{{.ProviderURL}}/{{.User}}/{{.Name}})
{{end}}`)

	return nil
}
