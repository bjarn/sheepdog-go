package command

import "os/exec"

var Brew = func(args ...string) *exec.Cmd {
	return exec.Command("brew", args...)
}
