package config

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func Parse(path string) ConfigModel {
	if path == "" {
		userHome, _ := os.UserHomeDir()
		path = userHome + "/.restless/config"
	}

	_, err := os.Stat(path)
	if err != nil {
		log.Fatal("Couldn't open config file")
	}

	fd, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error %v", err)
	}

	configBytes, err := ioutil.ReadAll(fd)
	if err != nil {
		log.Fatalf("Error %v", err)
	}

	configModel := NewConfigModel()
	err = yaml.Unmarshal(configBytes, &configModel)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	return *configModel
}
