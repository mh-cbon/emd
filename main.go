// Enhanced Markdown template processor.

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mh-cbon/emd/emd"
	gostd "github.com/mh-cbon/emd/go"
	gononstd "github.com/mh-cbon/emd/go-nonstd"
	"github.com/mh-cbon/emd/std"
)

// VERSION defines the running build id.
var VERSION = "0.0.0"

type commander interface {
	getDesc() string
	getName() string
	getSet() *flag.FlagSet
	getFn() func(s interface{}) error
}
type command struct {
	name string
	desc string
	set  *flag.FlagSet
	fn   func(s interface{}) error
}

func (c *command) getDesc() string {
	return c.desc
}
func (c *command) getName() string {
	return c.name
}
func (c *command) getSet() *flag.FlagSet {
	return c.set
}
func (c *command) getFn() func(s interface{}) error {
	return c.fn
}

func newCommand(name string, desc string, fn func(s interface{}) error) *command {
	return &command{name, desc, flag.NewFlagSet(name, flag.ExitOnError), fn}
}

type gencommand struct {
	*command
	in   string
	out  string
	data string
	help bool
}

func newGenCommand(desc string) *gencommand {
	return &gencommand{command: newCommand("gen", desc, generate)}
}

func main() {

	flag.Bool("help", false, "Show help")
	versionFlag := flag.Bool("version", false, "Show version")

	cmds := map[string]commander{}

	g := newGenCommand("Process an emd file.")
	g.set.StringVar(&g.in, "in", "README.e.md", "Input src file")
	g.set.StringVar(&g.out, "out", "-", "Output destination, defaults to stdout")
	g.set.StringVar(&g.data, "data", "", "JSON map of data")
	g.set.BoolVar(&g.help, "help", false, "Show help")
	cmds["gen"] = g

	if len(os.Args) > 1 {
		if cmd, ok := cmds[os.Args[1]]; ok {
			cmd.getSet().Parse(os.Args[2:])
		} else {
			flag.Parse()
		}

		for name, cmd := range cmds {
			if cmd.getSet().Parsed() {
				mustNotPanic(
					cmd.getFn()(cmd),
					name+" failed: %v",
				)
			}
		}
	}

	if *versionFlag {
		vers("emd")
	} else {
		usage("emd", "", cmds)
	}
}

func vers(name string) {
	fmt.Fprintf(os.Stderr, "%s - %v\n", name, VERSION)
}
func usage(name string, cmd string, cmds map[string]commander) {
	if cmd == "" {
		vers(name)
		fmt.Fprintln(os.Stderr, "\nUsage")
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr, "\nCommands")
		for name, c := range cmds {
			fmt.Fprintf(os.Stderr, "\t%v\t%v\n", name, c.getDesc())
		}
	} else {
		usagecmd(name, cmds[cmd])
	}
}
func usagecmd(name string, cmd commander) {
	vers(name)
	fmt.Fprintf(os.Stderr, "\nCommand %q: %v\n", cmd.getName(), cmd.getDesc())
	cmd.getSet().PrintDefaults()
}

func generate(s interface{}) error {

	cmd, ok := s.(*gencommand)
	if ok == false {
		return fmt.Errorf("Invalid command type %T", s)
	}

	if cmd.help {
		usagecmd("emd", cmd)
		return nil
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

func mustNotPanic(err error, formats ...string) {
	if err != nil {
		format := "err is not nil: %v"
		if len(formats) > 0 {
			format = formats[0]
		}
		panic(fmt.Errorf(format, err))
	}
}

var defTemplate = `# {{.Name}}

{{template "badge/travis" .}}{{template "badge/godoc" .}}

{{pkgdoc}}

# Install

{{goinstall}}
`
