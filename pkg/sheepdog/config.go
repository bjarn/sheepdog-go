package sheepdog

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Service struct {
	Enabled bool   `json:"enabled"`
	Name    string `json:"name"`
}

type Database struct {
	Password string `json:"password"`
}

type Config struct {
	Domain   string    `json:"domain"`
	Database Database  `json:"database"`
	Services []Service `json:"service"`
}

func GetConfigPath() string {
	return HomeDir() + "/config.json"
}

// Create the initial config for Sheepdog
func CreateConfig(domain string) error {
	if _, err := os.Stat(HomeDir()); os.IsNotExist(err) {
		err := CreateHomeDir()
		if err != nil {
			panic(err)
		}
	}

	_, err := os.Create(GetConfigPath())
	if err != nil {
		panic(err)
	}

	config := Config{
		Domain:   domain,
		Database: Database{Password: "root"},
		Services: nil,
	}

	data, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(GetConfigPath(), data, 0600)

	if err != nil {
		panic(err)
	}

	return nil
}

// Get the current Sheepdog config
func GetConfig() (Config, error) {
	vhosts, err := ioutil.ReadFile(GetConfigPath())

	var config Config

	if err != nil {
		err = CreateConfig("test")
		if err != nil {
			panic(err)
		}
	}

	err = json.Unmarshal(vhosts, &config)

	return config, err
}
