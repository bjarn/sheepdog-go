package service

import "fmt"

type PhpFpm74Service Service

var PhpFpm74 = &PhpFpm74Service{
	Name:        "php@7.4",
	RequireRoot: false,
}

func (service *PhpFpm74Service) Control(action Action) error   {
	return Control(service.Name, service.RequireRoot, action)
}

func (service *PhpFpm74Service) Start() error     {
	return Start(service)
}

func (service *PhpFpm74Service) Stop() error {
	fmt.Println("Nginx Stop")
	return nil
}

func (service *PhpFpm74Service) Restart() error   {
	return Restart(service)
}

func (service *PhpFpm74Service) Install() error   {
	return Install(service, service.Name)
}

func (service *PhpFpm74Service) Configure() error {
	return nil
}

type PhpFpm73Service Service

var PhpFpm73 = &PhpFpm73Service{
	Name:        "php@7.3",
	RequireRoot: false,
}

func (service *PhpFpm73Service) Control(action Action) error   {
	return Control(service.Name, service.RequireRoot, action)
}

func (service *PhpFpm73Service) Start() error     {
	return Start(service)
}

func (service *PhpFpm73Service) Stop() error {
	return Stop(service)
}

func (service *PhpFpm73Service) Restart() error   {
	return Restart(service)
}

func (service *PhpFpm73Service) Install() error   {
	return Install(service, service.Name)
}

func (service *PhpFpm73Service) Configure() error {
	return nil
}

type PhpFpm72Service Service

var PhpFpm72 = &PhpFpm72Service{
	Name:        "php@7.2",
	RequireRoot: false,
}

func (service *PhpFpm72Service) Control(action Action) error   {
	return Control(service.Name, service.RequireRoot, action)
}

func (service *PhpFpm72Service) Start() error     {
	return Start(service)
}

func (service *PhpFpm72Service) Stop() error {
	return Stop(service)
}

func (service *PhpFpm72Service) Restart() error   {
	return Restart(service)
}

func (service *PhpFpm72Service) Install() error   {
	return Install(service, service.Name)
}

func (service *PhpFpm72Service) Configure() error {
	return nil
}
