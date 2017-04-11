package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

// CliError is an error of cli command
type CliError struct {
	Err error
	Cmd string
}

func (c *CliError) Error() string {
	return fmt.Sprintf("%v\n\nThe command was:\n%v", c.Err, c.Cmd)
}

// GetCmdStr returns the string of a command
func GetCmdStr(bin string, args []string) string {
	s := bin
	for _, a := range args {
		if strings.Index(a, "\"") > -1 {
			a = strings.Replace(a, "\"", "\\\"", -1)
		}
		s += fmt.Sprintf(" %v", a)
	}
	return s
}

// Exec a command
func Exec(bin string, args []string) (string, error) {
	cmd := exec.Command(bin, args...)
	out, err := cmd.CombinedOutput()
	cmdStr := GetCmdStr(bin, args)
	if err != nil {
		return "", &CliError{Err: err, Cmd: cmdStr}
	}
	return string(out), nil
}
