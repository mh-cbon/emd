// Package cli handles command line arguments.
package cli

import (
	"flag"
	"fmt"
	"os"
)

// NewProgram makes a new Program instance.
func NewProgram(name string, version string) *Program {
	return &Program{map[string]Commander{}, name, version, false, false, false}
}

// NewCommand makes a new Program instance.
func NewCommand(name string, desc string, fn func(s Commander) error) *Command {
	return &Command{name, desc, flag.NewFlagSet(name, flag.ExitOnError), fn}
}

// Program is a struct to define a program and its command.
type Program struct {
	commands       map[string]Commander
	ProgramName    string
	ProgramVersion string
	Help           bool
	ShortHelp      bool
	Version        bool
}

// Bind help and version flag.
func (p *Program) Bind() {
	flag.CommandLine.Init(p.ProgramName, flag.ExitOnError)
	flag.BoolVar(&p.ShortHelp, "h", false, "Show help")
	flag.BoolVar(&p.Help, "help", false, "Show help")
	flag.BoolVar(&p.Version, "version", false, "Show version")
}

// Add a new sub command.
func (p *Program) Add(c Commander) bool {
	p.commands[c.getName()] = c
	return true
}

// ShowVersion prints program name and version on stderr.
func (p *Program) ShowVersion() error {
	fmt.Fprintf(os.Stderr, "%s - %v\n", p.ProgramName, p.ProgramVersion)
	return nil
}

// ShowUsage prints program usage of given subCmd on stderr. If subCmd is empty, prints general usage.
func (p *Program) ShowUsage(subCmd string) error {
	if subCmd == "" {
		p.ShowVersion()
		fmt.Fprintln(os.Stderr, "\nUsage")
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr, "\nCommands")
		for name, c := range p.commands {
			fmt.Fprintf(os.Stderr, "\t%v\t%v\n", name, c.getDesc())
		}
		return nil
	}
	if cmd := p.commands[subCmd]; cmd != nil {
		return p.ShowCmdUsage(cmd)
	}
	return fmt.Errorf("No such command %q", subCmd)
}

// ShowCmdUsage prints command usage of given subCmd on stderr.
func (p *Program) ShowCmdUsage(cmd Commander) error {
	p.ShowVersion()
	fmt.Fprintf(os.Stderr, "\nCommand %q: %v\n", cmd.getName(), cmd.getDesc())
	cmd.getSet().PrintDefaults()
	return nil
}

// Run the program against given set of arguments.
func (p *Program) Run(args []string) error {
	if len(args) > 1 {

		if cmd, ok := p.commands[args[1]]; ok {
			if err := cmd.getSet().Parse(args[2:]); err != nil {
				return nil
			}
		} else {
			flag.Parse()
		}

		for name, cmd := range p.commands {
			if cmd.getSet().Parsed() {
				err := cmd.getFn()(cmd)
				if err != nil {
					err = fmt.Errorf("command %q failed: %v", name, err)
				}
				return err
			}
		}
	}

	if p.Version {
		return p.ShowVersion()
	}

	return p.ShowUsage("")
}

// Commander is an generalizer of Command.
type Commander interface {
	getDesc() string
	getName() string
	getSet() *flag.FlagSet
	getFn() func(s Commander) error
}

// Command describe a program sub command.
type Command struct {
	name string
	desc string
	Set  *flag.FlagSet
	fn   func(s Commander) error
}

func (c *Command) getDesc() string {
	return c.desc
}
func (c *Command) getName() string {
	return c.name
}
func (c *Command) getSet() *flag.FlagSet {
	return c.Set
}
func (c *Command) getFn() func(s Commander) error {
	return c.fn
}
