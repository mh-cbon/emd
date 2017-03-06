package main_test

import (
	"os"

	"github.com/mh-cbon/emd/emd"
	"github.com/mh-cbon/emd/std"
)

// ExampleMain demonstrates the generation
// of the given README.e.md source file
// to os.Stdout.
func ExampleMain() {

	// make a new instance of emd.Generator.
	gen := emd.NewGenerator()

	// set the main template.
	gen.AddTemplate("{{.Name}}")

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
	// Output: projectname
}
