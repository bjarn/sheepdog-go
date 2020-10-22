package service

const NginxPath = "/usr/local/etc/nginx"

type NginxService Service

var Nginx = &NginxService{
	Name:        "nginx",
	RequireRoot: false,
}

func (service *NginxService) Control(action Action) error   {
	return Control(service.Name, service.RequireRoot, action)
}

func (service *NginxService) Start() error     {
	return service.Control(StartAction)
}

func (service *NginxService) Stop() error {
	return service.Control(StopAction)
}

func (service NginxService) Restart() error   {
	return service.Control(RestartAction)
}

func (service *NginxService) Install() error   {
	return Install(service, service.Name)
}

func (service *NginxService) Configure() error {
	return nil
}
