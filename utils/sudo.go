package utils

import "github.com/bjarn/sheepdog/pkg/command"

func RequireSudo() error {
	cmd := command.Sudo("-v")
	return cmd.Run()
}