package shell

import (
	"os"
	"os/exec"
)

// LookPath is alias exec.LookPath
func LookPath(fn string) (path string, err error) {
	return exec.LookPath(fn)
}

// Run runs external commands and print realtime stdout and stderr
func Run(name string, args ...string) (err error) {
	c := exec.Command(name, args...)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	if err := c.Run(); err != nil {
		return err
	}
	return nil
}

// Sudo creates exec.Cmd as root user
func Sudo(name string, args ...string) (c *exec.Cmd) {
	c = exec.Command("sudo", append([]string{name}, args...)...)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return
}
