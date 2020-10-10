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

type Action string

const (
	StartAction Action = "start"
	StopAction Action = "stop"
	RestartAction Action = "restart"
)

func (service Service) Control(action Action) error {
	var cmd *exec.Cmd

	if service.RequireRoot {
		cmd = command.Sudo("brew", "services", string(action), service.Name)
	} else {
		cmd = command.Brew("services", string(action), service.Name)
	}

	return cmd.Run()
}

// Restart the service using Brew service
func (service Service) Restart() error {
	return service.Control(RestartAction)
}

// Start the service using Brew service
func (service Service) Start() error {
	return service.Control(StartAction)
}

// Stop the service using Brew service
func (service Service) Stop() error {
	return service.Control(StopAction)
}
