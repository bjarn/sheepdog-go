package service

import (
	"fmt"
)

func Restart(serviceName string)  {
	if serviceName == "" {
		RestartAll()
		return
	}

	// restart specific service
}

func RestartAll() {
	for _, service := range Services {
		fmt.Printf("Restarting %s... ", service.Name)
		err := service.Restart()
		if err != nil {
			fmt.Printf("An error occurred whilst restarting %s: %s\n", service.Name, err)
		} else {
			fmt.Print("Done\n")
		}
	}
}

func StopAll() {
	for _, service := range Services {
		fmt.Printf("Stopping %s... ", service.Name)
		err := service.Stop()
		if err != nil {
			fmt.Printf("An error occurred whilst stopping %s: %s\n", service.Name, err)
		} else {
			fmt.Print("Done\n")
		}
	}
}

func StartAll() {
	for _, service := range Services {
		fmt.Printf("Starting %s... ", service.Name)
		err := service.Start()
		if err != nil {
			fmt.Printf("An error occurred whilst starting %s: %s\n", service.Name, err)
		} else {
			fmt.Print("Done\n")
		}
	}
}