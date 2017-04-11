// Package deprecated contains deprecated helpers.
package deprecated

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/mh-cbon/emd/emd"
)

// CliError is an error of cli command
type CliError struct {
	Err error
	Cmd string
}

func (c *CliError) Error() string {
	return fmt.Sprintf("%v\n\nThe command was:\n%v", c.Err, c.Cmd)
}

// Register standard helpers to the generator.
func Register(g *emd.Generator) error {

	// deprecated for cat
	g.AddFunc("file", func(f string, exts ...string) (string, error) {
		log.Println("")
		log.Println("							file function is deprecated , please update")
		log.Println("")
		s, err := ioutil.ReadFile(f)
		if err != nil {
			return "", err
		}
		ext := filepath.Ext(f)
		ext = strings.TrimPrefix(ext, ".")
		if len(exts) > 0 {
			ext = exts[0]
		}
		res := `
###### > ` + f + `
` + "```" + ext + `
` + strings.TrimSpace(string(s)) + `
` + "```"
		return res, err
	})

	// deprecated for exec
	g.AddFunc("cli", func(bin string, args ...string) (string, error) {
		log.Println("")
		log.Println("							cli function is deprecated , please update")
		log.Println("")
		cmd := exec.Command(bin, args...)
		out, err := cmd.CombinedOutput()
		if err != nil {
			s := bin
			for _, a := range args {
				if strings.Index(a, "\"") > -1 {
					a = strings.Replace(a, "\"", "\\\"", -1)
				}
				s += fmt.Sprintf(" %v", a)
			}
			return "", &CliError{Err: err, Cmd: s}
		}
		fbin := filepath.Base(bin)
		res := `
###### $ ` + fbin + ` ` + strings.Join(args, " ") + `
` + "```sh" + `
` + strings.TrimSpace(string(out)) + `
` + "```"
		return res, err
	})

	return nil
}
