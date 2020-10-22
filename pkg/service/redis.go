package service

type RedisService Service

var Redis = &RedisService{
	Name:        "redis",
	RequireRoot: false,
}

func (service *RedisService) Control(action Action) error   {
	return Control(service.Name, service.RequireRoot, action)
}

func (service *RedisService) Start() error     {
	return Start(service)
}

func (service *RedisService) Stop() error {
	return Stop(service)
}

func (service *RedisService) Restart() error   {
	return Restart(service)
}

func (service *RedisService) Install() error   {
	return Install(service, service.Name)
}

func (service *RedisService) Configure() error {
	return nil
}
