package config

import (
	"log"
	"os"
)

type ConfigData struct {
	Host     string
	Port     string
	Username string
	token    string
}

func Parse(path string) (string, error) {
	_, err := os.Stat(path)
	if err != nil {
		log.Fatal("Couldn't open config file")
	}

	fd, _ := os.Open(path)

}
