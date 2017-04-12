package emd

import (
	"bytes"
	"strings"
	"testing"
)

func TestVVV(t *testing.T) {
	gen := NewGenerator()
	gen.AddTemplate(`The result of the func: {{f "s"}}
The result of the template: {{template "s" .}}
The data: {{.}}
`)
	gen.AddTemplate(`{{define "s"}}s template{{end}}`)
	gen.AddFunc("f", func(s string) string { return strings.ToUpper(s) })
	gen.SetData("the", "data")
	gen.AddPostProcess(func(s string) string {
		return s + "\npostprocess"
	})
	var buf bytes.Buffer
	gen.Execute(&buf)

	expectedOut := `The result of the func: S
The result of the template: s template
The data: map[the:data]

postprocess`
	gotOut := buf.String()
	if expectedOut != gotOut {
		t.Errorf(
			"Unexpected result\n==== expected\n%q\n==== got\n%v",
			expectedOut,
			gotOut,
		)
	}
}
