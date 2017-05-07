// Package emd provides support to process .md files.
package emd

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"text/template"

	"github.com/mh-cbon/emd/utils"
)

//Generator generates an emd content.
type Generator struct {
	t     *template.Template
	o     io.Writer
	tpls  []string
	funcs map[string]interface{}
	data  map[string]interface{}
	post  []func(string) string
}

// NewGenerator creates Generator Pointers.
func NewGenerator() *Generator {
	return &Generator{
		tpls:  []string{},
		funcs: map[string]interface{}{},
		data:  map[string]interface{}{},
	}
}

//AddPostProcess registers a post process function.
// Post process are registered by template func call and are removed after next template generation.
func (g *Generator) AddPostProcess(f func(string) string) {
	g.post = append(g.post, f)
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
	// read the prelude
	newT, newData, err := utils.GetPrelude(t)
	if err != nil {
		panic(err)
	}
	g.SetDataMap(newData)

	g.tpls = append(g.tpls, newT)
}

//AddFileTemplate registers a template file.
func (g *Generator) AddFileTemplate(t string) error {
	s, err := ioutil.ReadFile(t)
	if err != nil {
		return err
	}
	// add the template string.
	g.AddTemplate(string(s))
	return nil
}

//GetTemplate returns the compiled templates.
//It is available only during Execute.
func (g *Generator) GetTemplate() *template.Template {
	return g.t
}

//GetOut returns the out writer.
//It is available only during Execute.
func (g *Generator) GetOut() io.Writer {
	return g.o
}

//WriteString writes a string on out.
//It is available only during Execute.
func (g *Generator) WriteString(s string) (int, error) {
	return g.o.Write([]byte(s))
}

//GetData returns a copy of the template's data.
//It is available only during Execute.
func (g *Generator) GetData() map[string]interface{} {
	ret := map[string]interface{}{}
	for k, v := range g.data {
		ret[k] = v
	}
	return ret
}

//GetKey returns value of K.
func (g *Generator) GetKey(K string) interface{} {
	return g.data[K]
}

//GetSKey returns a string value of K.
func (g *Generator) GetSKey(K string) string {
	v := g.GetKey(K)
	if v == nil {
		return ""
	}
	return fmt.Sprintf("%v", g.GetKey(K))
}

//Execute the template to out.
func (g *Generator) Execute(out io.Writer) error {
	var b bytes.Buffer
	g.o = &b // becasue of post process, we need to buffer out.
	var err error
	g.t = template.New("").Funcs(g.funcs)
	for _, tpl := range g.tpls {
		g.t, err = g.t.Parse(tpl)
		if err != nil {
			return err
		}
	}
	err = g.t.Execute(g.o, g.data)
	g.t = nil
	g.o = nil
	s := b.String()
	b.Truncate(0)
	for _, p := range g.post {
		s = p(s)
	}
	g.post = g.post[:0]
	out.Write([]byte(s))
	return err
}
