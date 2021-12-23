// Enhanced Markdown template processor.
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"time"

	"github.com/mh-cbon/emd/cli"
	"github.com/mh-cbon/emd/deprecated"
	"github.com/mh-cbon/emd/emd"
	gostd "github.com/mh-cbon/emd/go"
	gononstd "github.com/mh-cbon/emd/go-nonstd"
	"github.com/mh-cbon/emd/provider"
	"github.com/mh-cbon/emd/std"
)

// VERSION defines the running build id.
var VERSION = "1.0.2"

var program = cli.NewProgram("emd", VERSION)

var verbose bool

func logMsg(f string, args ...interface{}) {
	if verbose {
		log.Printf(f+"\n", args...)
	}
}

func main() {
	program.Bind()
	if err := program.Run(os.Args); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

// gen sub command
type gencommand struct {
	*cli.Command
	in        mFlags
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
	gen.Set.Var(&gen.in, "in", "Input src file")
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

	verbose = os.Getenv("VERBOSE") != ""
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

	projectPath, err := getProjectPath()
	if err != nil {
		return err
	}
	logMsg("projectPath %q", projectPath)

	plugins := getPlugins()
	data, err := getData(projectPath)
	if err != nil {
		return err
	}

	gen := emd.NewGenerator()

	gen.SetDataMap(data)

	if len(cmd.in) == 0 {
		b := tryReadOsStdin()
		if b != nil && b.Len() > 0 {
			gen.AddTemplate(b.String())

		} else {
			if s, err := os.Stat("README.e.md"); !os.IsNotExist(err) && s.IsDir() == false {
				err := gen.AddFileTemplate("README.e.md")
				if err != nil {
					return err
				}
			} else {
				gen.AddTemplate(defTemplate)
			}
		}
	}

	for name, plugin := range plugins {
		if err := plugin(gen); err != nil {
			return fmt.Errorf("Failed to register %v package: %v", name, err)
		}
	}

	if cmd.data != "" {
		jData := map[string]interface{}{}
		if err := json.Unmarshal([]byte(cmd.data), &jData); err != nil {
			return fmt.Errorf("Cannot decode JSON data string: %v", err)
		}
		gen.SetDataMap(jData)
	}

	if len(cmd.in) == 0 {
		if err := gen.Execute(out); err != nil {
			return fmt.Errorf("Generator failed: %v", err)
		}
	} else {
		for _, val := range cmd.in {
			if err := gen.AddFileTemplate(val); err != nil {
				return err
			}
			if err := gen.Execute(out); err != nil {
				return fmt.Errorf("Generator failed: %v", err)
			}
		}
	}

	return nil
}

func tryReadOsStdin() *bytes.Buffer {
	copied := make(chan bool)
	timedout := make(chan bool)
	var ret bytes.Buffer
	go func() {
		io.Copy(&ret, os.Stdin)
		copied <- true
	}()
	go func() {
		<-time.After(time.Millisecond * 10)
		timedout <- ret.Len() == 0
	}()
	select {
	case empty := <-timedout:
		if empty {
			return nil
		}
		<-copied
	case <-copied:
	}
	return &ret
}

func getProjectPath() (string, error) {
	originalCwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	logMsg("cwd %q", originalCwd)

	// regular go package
	{
		projectPath, err := matchProjectPath(originalCwd)
		if err == nil {
			return projectPath, nil
		}
	}

	// symlinked go package
	{
		cwd, err := filepath.EvalSymlinks(originalCwd)
		if err == nil {
			projectPath, err := matchProjectPath(cwd)
			if err == nil {
				return projectPath, nil
			}
		}
	}

	// all other cases
	return originalCwd, nil
}

var re = regexp.MustCompile("(src/[^/]+[.](com|org|net)/.+)")

func matchProjectPath(p string) (string, error) {
	res := re.FindAllString(p, -1)
	if len(res) > 0 {
		return res[0][3:], nil
	}
	return "", fmt.Errorf("Invalid working directory %q", p)
}

func getData(cwd string) (map[string]interface{}, error) {
	p := provider.Default(cwd)
	return map[string]interface{}{
		"Name":         p.GetProjectName(),
		"User":         p.GetUserName(),
		"ProviderURL":  p.GetProviderURL(),
		"ProviderName": p.GetProviderID(),
		"URL":          p.GetURL(),
		"ProjectURL":   p.GetProjectURL(),
		"Branch":       "master",
	}, nil
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

// InitFile creates a basic emd file if none exists.
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
		return fmt.Errorf("File exists at %q", out)
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
