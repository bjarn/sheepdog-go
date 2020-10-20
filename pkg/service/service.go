package service

import (
	"github.com/bjarn/sheepdog/pkg/command"
	"os/exec"
	"strings"
)

type Service struct {
	Name        string
	RequireRoot bool
}

var Services = []Service{Nginx, DnsMasq, MySql57, Redis, MailHog}

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

// Install and start the service
func (service Service) Install() error {
	err := command.Brew("install", service.Name).Run()
	if err != nil {
		// Brew throws exit status 1 as warning, just go on...
		if !strings.Contains(err.Error(), "exit status 1") {
			panic(err)
		}
	}

	err = service.Start()

	return err
}