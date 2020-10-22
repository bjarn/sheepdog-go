package service

type MySql56Service Service

var MySql56 = &MySql56Service{
	Name:        "mysql@5.6",
	RequireRoot: false,
}

func (service *MySql56Service) Control(action Action) error   {
	return Control(service.Name, service.RequireRoot, action)
}

func (service *MySql56Service) Start() error     {
	return Start(service)
}

func (service *MySql56Service) Stop() error {
	return Stop(service)
}

func (service *MySql56Service) Restart() error   {
	return Restart(service)
}

func (service *MySql56Service) Install() error   {
	return Install(service, service.Name)
}

func (service *MySql56Service) Configure() error {
	return nil
}

type MySql57Service Service

var MySql57 = &MySql57Service{
	Name:        "mysql@5.7",
	RequireRoot: false,
}

func (service *MySql57Service) Control(action Action) error   {
	return Control(service.Name, service.RequireRoot, action)
}

func (service *MySql57Service) Start() error     {
	return Start(service)
}

func (service *MySql57Service) Stop() error {
	return Stop(service)
}

func (service *MySql57Service) Restart() error   {
	return Restart(service)
}

func (service *MySql57Service) Install() error   {
	return Install(service, service.Name)
}

func (service *MySql57Service) Configure() error {
	return nil
}

type MariaDbService Service

var MariaDb = &MariaDbService{
	Name:        "mariadb",
	RequireRoot: false,
}

func (service *MariaDbService) Control(action Action) error   {
	return Control(service.Name, service.RequireRoot, action)
}

func (service *MariaDbService) Start() error     {
	return Start(service)
}

func (service *MariaDbService) Stop() error {
	return Stop(service)
}

func (service *MariaDbService) Restart() error   {
	return Restart(service)
}

func (service *MariaDbService) Install() error   {
	return Install(service, service.Name)
}

func (service *MariaDbService) Configure() error {
	return nil
}
