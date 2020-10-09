package command

import "os/exec"

var Sudo = func(args ...string) *exec.Cmd {
	return exec.Command("sudo", args...)
}

var Brew = func(args ...string) *exec.Cmd {
	return exec.Command("brew", args...)
}