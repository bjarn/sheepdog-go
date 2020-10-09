package utils

import "github.com/bjarn/sheepdog/command"

func RequireSudo() error {
	cmd := command.Sudo("-v")
	return cmd.Run()
}