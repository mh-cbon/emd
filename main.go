// Enhanced Markdown template processor.
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/mh-cbon/emd/cli"
	"github.com/mh-cbon/emd/deprecated"
	"github.com/mh-cbon/emd/emd"
	gostd "github.com/mh-cbon/emd/go"
	gononstd "github.com/mh-cbon/emd/go-nonstd"
	"github.com/mh-cbon/emd/provider"
	"github.com/mh-cbon/emd/std"
)

// VERSION defines the running build id.
var VERSION = "0.0.0"

var program = cli.NewProgram("emd", VERSION)

func main() {
	program.Bind()
	if err := program.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
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

// init sub command
type initcommand struct {
	*cli.Command
	help      bool
	shortHelp bool
	out       string
	force     bool
}

func init() {
	gen := &gencommand{Command: cli.NewCommand("gen", "Process an emd file.", Generate)}
	gen.Set.StringVar(&gen.in, "in", "", "Input src file")
	gen.Set.StringVar(&gen.out, "out", "-", "Output destination, defaults to stdout")
	gen.Set.StringVar(&gen.data, "data", "", "JSON map of data")
	gen.Set.BoolVar(&gen.help, "help", false, "Show help")
	gen.Set.BoolVar(&gen.shortHelp, "h", false, "Show help")

	program.Add(gen)

	ini := &initcommand{Command: cli.NewCommand("init", "Init a basic emd file.", InitFile)}
	ini.Set.BoolVar(&ini.help, "help", false, "Show help")
	ini.Set.BoolVar(&ini.shortHelp, "h", false, "Show help")
	ini.Set.BoolVar(&ini.force, "force", false, "Force write")
	ini.Set.StringVar(&ini.out, "out", "README.e.md", "Out file")
	program.Add(ini)
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

	cwd, err := getCwd()
	if err != nil {
		return err
	}
	gopath := filepath.Join(os.Getenv("GOPATH"), "src")
	gopath = strings.Replace(gopath, "\\", "/", -1)
	projectPath := cwd[len(gopath):]

	plugins := getPlugins()
	data := getData(projectPath)

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

func getData(cwd string) map[string]interface{} {
	p := provider.Default(cwd)
	if p.Match() == false {
		log.Printf("Failed to identify this project url %v\n", cwd)
	}
	return map[string]interface{}{
		"Name":         p.GetProjectName(),
		"User":         p.GetUserName(),
		"ProviderURL":  p.GetProviderURL(),
		"ProviderName": p.GetProviderID(),
		"URL":          p.GetURL(),
		"ProjectURL":   p.GetProjectURL(),
		"Branch":       "master",
	}
}

func getPlugins() map[string]func(*emd.Generator) error {
	return map[string]func(*emd.Generator) error{
		"std":        std.Register,
		"gostd":      gostd.Register,
		"gononstd":   gononstd.Register,
		"deprecated": deprecated.Register,
	}
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

func getCwd() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("Failed to determmine cwd: %v", err)
	}
	// for those who uses symlinks to relocate their code,
	// the path must be evaluated.
	cwd, err = filepath.EvalSymlinks(cwd)
	if err != nil {
		return "", fmt.Errorf("Failed to determmine eval path: %v", err)
	}
	return cwd, nil
}

// InitFile creates a basic emd file if none exits.
func InitFile(s cli.Commander) error {

	cmd, ok := s.(*initcommand)
	if ok == false {
		return fmt.Errorf("Invalid command type %T", s)
	}

	if cmd.help || cmd.shortHelp {
		return program.ShowCmdUsage(cmd)
	}

	out := cmd.out
	if cmd.out == "" {
		out = "README.e.md"
	}
	if _, err := os.Stat(out); !cmd.force && !os.IsNotExist(err) {
		return fmt.Errorf("File exits at %q", out)
	}
	return ioutil.WriteFile(out, []byte(defTemplate), os.ModePerm)
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
