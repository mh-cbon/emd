// Package emd provides support to process .md files.
package emd

import (
	"io"
	"io/ioutil"
	"text/template"
)

//Generator generates an emd content.
type Generator struct {
	t     *template.Template
	tpls  []string
	funcs map[string]interface{}
	data  map[string]interface{}
}

// NewGenerator creates Generator Pointers.
func NewGenerator() *Generator {
	return &Generator{
		tpls:  []string{},
		funcs: map[string]interface{}{},
		data:  map[string]interface{}{},
	}
}

//AddFunc registers a template function.
func (g *Generator) AddFunc(name string, f interface{}) {
	g.funcs[name] = f
}

//AddFuncs registers a map of template functions.
func (g *Generator) AddFuncs(fm map[string]interface{}) {
	for name, f := range fm {
		g.funcs[name] = f
	}
}

//SetData registers a template data.
func (g *Generator) SetData(name string, d interface{}) {
	g.data[name] = d
}

//SetDataMap registers a map of template data.
func (g *Generator) SetDataMap(dm map[string]interface{}) {
	for name, d := range dm {
		g.data[name] = d
	}
}

//AddTemplate registers a template string.
func (g *Generator) AddTemplate(t string) {
	g.tpls = append(g.tpls, t)
}

//AddFileTemplate registers a template file.
func (g *Generator) AddFileTemplate(t string) error {
	s, err := ioutil.ReadFile(t)
	if err != nil {
		return err
	}
	g.AddTemplate(string(s))
	return nil
}

//Execute the template to out.
func (g *Generator) Execute(out io.Writer) error {
	var err error
	t := template.New("").Funcs(g.funcs)
	for _, tpl := range g.tpls {
		t, err = t.Parse(tpl)
		if err != nil {
			return err
		}
	}
	return t.Execute(out, g.data)
}
