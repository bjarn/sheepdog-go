package service

import (
	"fmt"
	"github.com/bjarn/sheepdog/utils"
)

// TODO: Make this dynamic by getting the enabled services from the user's config
var Services = []string{DnsMasq.Name, Nginx.Name, PhpFpm72.Name, PhpFpm73.Name, PhpFpm74.Name,
	MySql56.Name, MySql57.Name, MariaDb.Name, Redis.Name}
var IServices = []IService{DnsMasq, Nginx, PhpFpm72, PhpFpm73, PhpFpm74,
	MySql56, MySql57, MariaDb, Redis}

// Restart a single service.
func RestartSingle(service IService, serviceName string) {
	fmt.Printf("Restarting %s... ", serviceName)

	err := service.Restart()

	if err != nil {
		fmt.Printf("An error occurred whilst restarting %s: %s\n", serviceName, err)
	} else {
		fmt.Print("Done\n")
	}
}

// Restart all services by looping through the services and calling RestartSingle.
func RestartAll() {
	err := utils.RequireSudo()
	if err != nil {
		return
	}

	for index, service := range Services {
		RestartSingle(IServices[index], service)
	}
}

// Stop a single service.
func StopSingle(service IService, serviceName string) {
	fmt.Printf("Stopping %s... ", serviceName)
	err := service.Stop()
	if err != nil {
		fmt.Printf("An error occurred whilst stopping %s: %s\n", serviceName, err)
	} else {
		fmt.Print("Done\n")
	}
}

// Stops all services by looping through the services and calling StopSingle.
func StopAll() {
	err := utils.RequireSudo()
	if err != nil {
		return
	}

	for index, service := range Services {
		StopSingle(IServices[index], service)
	}
}

// Start a single service.
func StartSingle(service IService, serviceName string) {
	fmt.Printf("Start %s... ", serviceName)

	err := service.Start()

	if err != nil {
		fmt.Printf("An error occurred whilst start %s: %s\n", serviceName, err)
	} else {
		fmt.Print("Done\n")
	}
}

// Start all services by looping through the services and calling StartSingle.
func StartAll() {
	err := utils.RequireSudo()
	if err != nil {
		return
	}

	for index, service := range Services {
		StartSingle(IServices[index], service)
	}
}