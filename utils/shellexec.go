package utils

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

// Command Return a new exec.Cmd object for the given command string
func Command(cwd string, cmd string) (*TempCmd, error) {
	return NewTempCmd(cwd, cmd)
}

// TempCmd ...
type TempCmd struct {
	*exec.Cmd
	f string
}

var isWindows = runtime.GOOS == "windows"

// NewTempCmd is a cmd wrapped into a tmp file
func NewTempCmd(cwd string, cmd string) (*TempCmd, error) {
	f, err := ioutil.TempDir("", "stringexec")
	if err != nil {
		return nil, err
	}
	fp := filepath.Join(f, "s")
	if isWindows {
		fp += ".bat"
	}
	err = ioutil.WriteFile(fp, []byte(cmd), 0766)
	if err != nil {
		return nil, err
	}
	ret := &TempCmd{Cmd: exec.Command("sh", "-c", fp), f: fp}
	if isWindows {
		ret.Cmd = exec.Command("cmd", "/C", fp)
	}
	ret.Cmd.Dir = cwd
	return ret, nil
}

// Run the cmd
func (t *TempCmd) Run() error {
	if err := t.Cmd.Start(); err != nil {
		return err
	}
	return t.Wait()
}

// Wait wait for command then delete the tmp file.
func (t *TempCmd) Wait() error {
	err := t.Cmd.Wait()
	os.Remove(t.f)
	return err
}

// Shell exec a string.
func Shell(wd, s string) (string, error) {
	if wd == "" {
		var err error
		wd, err = os.Getwd()
		if err != nil {
			return "", err
		}
	}
	cmd, err := NewTempCmd(wd, s)
	if err != nil {
		return "", &CliError{Err: err, Cmd: s}
	}
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", &CliError{Err: err, Cmd: s}
	}
	return string(out), nil
}
