// Enhanced Markdown template processor.
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mh-cbon/emd/cli"
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
	gen.Set.StringVar(&gen.in, "in", "README.e.md", "Input src file")
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

	out := os.Stdout

	if cmd.out != "-" {
		f, err := os.Create(cmd.out)
		if err != nil {
			f, err = os.Open(cmd.out)
			if err != nil {
				return fmt.Errorf("Cannot open out destination: %v", err)
			}
		}
		defer f.Close()
		out = f
	}

	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("Failed to determmine cwd: %v", err)
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
	} else {
		gen.AddTemplate(defTemplate)
	}

	gen.SetDataMap(data)

	if err := std.Register(gen); err != nil {
		return fmt.Errorf("Failed to register std package: %v", err)
	}

	if err := gostd.Register(gen); err != nil {
		return fmt.Errorf("Failed to register gostd package: %v", err)
	}

	if err := gononstd.Register(gen); err != nil {
		return fmt.Errorf("Failed to register gononstd package: %v", err)
	}

	if err := gen.Execute(out); err != nil {
		return fmt.Errorf("Generator failed: %v", err)
	}

	if cmd.out != "-" {
		fmt.Println("")
		fmt.Println("Success")
	}

	return nil
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

{{template "badge/travis" .}}{{template "badge/godoc" .}}

{{pkgdoc}}

# Install

{{goinstall}}
`
