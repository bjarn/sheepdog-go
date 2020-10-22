package service

import (
	"github.com/bjarn/sheepdog/pkg/command"
	"os/exec"
	"strings"
)

type IService interface {
	Control(action Action) error
	Restart() error
	Start() error
	Stop() error
	Install() error
	Configure() error
}

type Service struct {
	IService
	Name        string
	RequireRoot bool
}

type Action string

const (
	StartAction   Action = "start"
	StopAction    Action = "stop"
	RestartAction Action = "restart"
)

func Control(service string, requireRoot bool, action Action) error {
	var cmd *exec.Cmd

	if requireRoot {
		cmd = command.Sudo("brew", "services", string(action), service)
	} else {
		cmd = command.Brew("services", string(action), service)
	}

	return cmd.Run()
}

// Restart the service using Brew service
func Restart(s IService) error {
	return s.Control(RestartAction)
}

// Start the service using Brew service
func Start(s IService) error {
	return s.Control(StartAction)
}

// Stop the service using Brew service
func Stop(s IService) error {
	return s.Control(StopAction)
}

// Install and start the service
func Install(service IService, serviceName string) error {
	err := command.Brew("install", serviceName).Run()
	if err != nil {
		// Brew throws exit status 1 as warning, just go on...
		if !strings.Contains(err.Error(), "exit status 1") {
			panic(err)
		}
	}

	err = service.Start()

	return err
}
