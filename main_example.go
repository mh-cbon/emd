// Demonstrates the generation
// of the given README.e.md source file
// to os.Stdout.

package main

import (
	"os"

	"github.com/mh-cbon/emd/emd"
	"github.com/mh-cbon/emd/std"
)

func genExample() {

	// make a new instance of emd.Generator.
	gen := emd.NewGenerator()

	// set the main template.
	if err := gen.AddFileTemplate("README.e.md"); err != nil {
		panic(err)
	}

	// set the data available in templates.
	gen.SetDataMap(map[string]interface{}{"Name": "projectname"})

	// register a plugin
	if err := std.Register(gen); err != nil {
		panic(err)
	}

	// process the template.
	if err := gen.Execute(os.Stdout); err != nil {
		panic(err)
	}
}
