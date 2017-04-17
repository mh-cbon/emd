// Package gononstd contains go non-standard helpers.
package gononstd

import "github.com/mh-cbon/emd/emd"

// InstructionGlideInstall is a template to show instructions to install the package with glide.
var InstructionGlideInstall = `{{define "glide/install" -}}
` + "```sh" + `
mkdir -p $GOPATH/src/{{.URL}}
cd $GOPATH/src/{{.URL}}
git clone https://{{.URL}}.git .
glide install
go install
` + "```" + `
{{- end}}`

// Register go non-standard helpers to the generator.
func Register(g *emd.Generator) error {
	g.AddTemplate(InstructionGlideInstall)
	return nil
}
