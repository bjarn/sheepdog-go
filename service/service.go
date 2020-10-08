package service

import (
	"github.com/bjarn/sheepdog/command"
)

type Service struct {
	Name string
}

var Services = []Service{Nginx, DnsMasq, MySql57, Redis, Mailhog}

// Restart the service using Brew service
func (service Service) Restart() error {
	cmd := command.Brew("services", "restart", service.Name)

	return cmd.Run()
}

// Start the service using Brew service
func (service Service) Start() error {
	cmd := command.Brew("services", "start", service.Name)

	return cmd.Run()
}

// Stop the service using Brew service
func (service Service) Stop() error {
	cmd := command.Brew("services", "stop", service.Name)

	return cmd.Run()
}