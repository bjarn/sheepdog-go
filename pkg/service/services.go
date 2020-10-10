package service

import (
	"fmt"
	"github.com/bjarn/sheepdog/utils"
)

// Restart a single service.
func RestartSingle(service Service) {
	fmt.Printf("Restarting %s... ", service.Name)
	err := service.Restart()
	if err != nil {
		fmt.Printf("An error occurred whilst restarting %s: %s\n", service.Name, err)
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

	for _, service := range Services {
		RestartSingle(service)
	}
}

// Stop a single service.
func StopSingle(service Service) {
	fmt.Printf("Stopping %s... ", service.Name)
	err := service.Stop()
	if err != nil {
		fmt.Printf("An error occurred whilst stopping %s: %s\n", service.Name, err)
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

	for _, service := range Services {
		StopSingle(service)
	}
}

// Start a single service.
func StartSingle(service Service) {
	fmt.Printf("Start %s... ", service.Name)
	err := service.Start()
	if err != nil {
		fmt.Printf("An error occurred whilst start %s: %s\n", service.Name, err)
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

	for _, service := range Services {
		StartSingle(service)
	}
}