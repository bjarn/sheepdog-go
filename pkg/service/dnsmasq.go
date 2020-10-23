package service

import "github.com/bjarn/sheepdog/pkg/sheepdog"

const ResolverPath = "/etc/resolver"
const DnsMasqConfigPath = "/usr/local/etc/dnsmasq.conf"

type DnsMasqService Service

var DnsMasq = &DnsMasqService{
	Name:        "dnsmasq",
	RequireRoot: true,
}

func (service *DnsMasqService) Control(action Action) error   {
	return Control(service.Name, service.RequireRoot, action)
}

func (service *DnsMasqService) Start() error     {
	return Start(service)
}

func (service *DnsMasqService) Stop() error {
	return Stop(service)
}

func (service *DnsMasqService) Restart() error   {
	return Restart(service)
}

func (service *DnsMasqService) Install() error   {
	return Install(service, service.Name)
}

func (service *DnsMasqService) Configure() error {
	config, err := sheepdog.GetConfig()

	if err != nil {
		panic(err)
	}

	return nil
}