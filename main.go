// Enhanced Markdown template processor.
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/mh-cbon/emd/cli"
	"github.com/mh-cbon/emd/deprecated"
	"github.com/mh-cbon/emd/emd"
	gostd "github.com/mh-cbon/emd/go"
	gononstd "github.com/mh-cbon/emd/go-nonstd"
	"github.com/mh-cbon/emd/std"
)

// VERSION defines the running build id.
var VERSION = "0.0.0"

var program = cli.NewProgram("emd", VERSION)

func main() {
	program.Bind()
	if err := program.Run(os.Args); err != nil {
		panic(err)
	}
}

// gen sub command
type gencommand struct {
	*cli.Command
	in        string
	out       string
	data      string
	help      bool
	shortHelp bool
}

func init() {
	gen := &gencommand{Command: cli.NewCommand("gen", "Process an emd file.", Generate)}
	gen.Set.StringVar(&gen.in, "in", "", "Input src file")
	gen.Set.StringVar(&gen.out, "out", "-", "Output destination, defaults to stdout")
	gen.Set.StringVar(&gen.data, "data", "", "JSON map of data")
	gen.Set.BoolVar(&gen.help, "help", false, "Show help")
	gen.Set.BoolVar(&gen.shortHelp, "h", false, "Show help")

	program.Add(gen)
}

// Generate is the cli command implementation of gen.
func Generate(s cli.Commander) error {

	cmd, ok := s.(*gencommand)
	if ok == false {
		return fmt.Errorf("Invalid command type %T", s)
	}

	if cmd.help || cmd.shortHelp {
		return program.ShowCmdUsage(cmd)
	}

	out, err := getStdout(cmd.out)
	if err != nil {
		return err
	}
	if x, ok := out.(io.Closer); ok {
		defer x.Close()
	}

	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("Failed to determmine cwd: %v", err)
	}
	// for those who uses symlinks to relocate their code,
	// the path must be evaluated.
	cwd, err = filepath.EvalSymlinks(cwd)
	if err != nil {
		return fmt.Errorf("Failed to determmine eval path: %v", err)
	}

	plugins := map[string]func(*emd.Generator) error{
		"std":        std.Register,
		"gostd":      gostd.Register,
		"gononstd":   gononstd.Register,
		"deprecated": deprecated.Register,
	}

	data := map[string]interface{}{
		"Name":         filepath.Base(cwd),
		"User":         getProjectUser(cwd),
		"ProviderURL":  getProviderURL(cwd),
		"ProviderName": getProviderName(cwd),
		"URL":          getProviderURL(cwd) + "/" + getProjectUser(cwd) + "/" + filepath.Base(cwd),
		"Branch":       "master",
	}

	if cmd.data != "" {
		if err := json.Unmarshal([]byte(cmd.data), &data); err != nil {
			return fmt.Errorf("Cannot decode JSON data string: %v", err)
		}
	}

	gen := emd.NewGenerator()

	if cmd.in != "" {
		if err := gen.AddFileTemplate(cmd.in); err != nil {
			return err
		}
	} else if s, err := os.Stat("README.e.md"); !os.IsNotExist(err) && s.IsDir() == false {
		if err := gen.AddFileTemplate("README.e.md"); err != nil {
			return err
		}
	} else {
		gen.AddTemplate(defTemplate)
	}

	gen.SetDataMap(data)

	for name, plugin := range plugins {
		if err := plugin(gen); err != nil {
			return fmt.Errorf("Failed to register %v package: %v", name, err)
		}
	}

	if err := gen.Execute(out); err != nil {
		return fmt.Errorf("Generator failed: %v", err)
	}

	return nil
}

func getStdout(out string) (io.Writer, error) {

	ret := os.Stdout

	if out != "-" {
		f, err := os.Create(out)
		if err != nil {
			f, err = os.Open(out)
			if err != nil {
				return nil, fmt.Errorf("Cannot open out destination: %v", err)
			}
		}
		ret = f
	}
	return ret, nil
}

func getProjectUser(s string) string {
	ss := strings.Split(s, "/")
	if len(ss) > 2 {
		return ss[len(ss)-2]
	}
	return ""
}

func getProviderName(s string) string {
	if strings.Index(s, "github.com/") > -1 {
		return "github"
	}
	if strings.Index(s, "gitlab.com/") > -1 {
		return "gitlab"
	}
	// etc
	return ""
}

func getProviderURL(s string) string {
	if strings.Index(s, "github.com/") > -1 {
		return "github.com"
	}
	if strings.Index(s, "gitlab.com/") > -1 {
		return "gitlab.com"
	}
	// etc
	return ""
}

var defTemplate = `# {{.Name}}

{{template "badge/goreport" .}} {{template "badge/godoc" .}}

{{pkgdoc}}

# {{toc 5}}

# Install

{{template "gh/releases" .}}

#### go
{{template "go/install" .}}
`
