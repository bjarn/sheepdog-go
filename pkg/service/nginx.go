package service

const NginxPath = "/usr/local/etc/nginx"

var Nginx = Service{
	Name: "nginx",
	RequireRoot: false,
}
