package service

import (
	"github.com/bjarn/sheepdog/pkg/command"
	"os/exec"
)

type Service struct {
	Name        string
	RequireRoot bool
}

var Services = []Service{Nginx, DnsMasq, MySql57, Redis, Mailhog}

// Restart the service using Brew service
func (service Service) Restart() error {
	var cmd *exec.Cmd
	if service.RequireRoot {
		cmd = command.Sudo("brew", "services", "restart", service.Name)
	} else {
		cmd = command.Brew("services", "restart", service.Name)
	}

	return cmd.Run()
}

// Start the service using Brew service
func (service Service) Start() error {
	var cmd *exec.Cmd
	if service.RequireRoot {
		cmd = command.Sudo("brew", "services", "start", service.Name)
	} else {
		cmd = command.Brew("services", "start", service.Name)
	}

	return cmd.Run()
}

// Stop the service using Brew service
func (service Service) Stop() error {
	var cmd *exec.Cmd
	if service.RequireRoot {
		cmd = command.Sudo("brew", "services", "stop", service.Name)
	} else {
		cmd = command.Brew("services", "stop", service.Name)
	}

	return cmd.Run()
}
