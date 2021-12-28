package controllers

import (
	log "github.com/sirupsen/logrus"
	"mailrelay/models"
	"gopkg.in/yaml.v2"
	"os"
)

func CheckEmptyString(checkThis string, description string) {
	if checkThis == "" {
		log.WithField("field", description).Fatal("Needed config field is empty.")
	}
}

func CheckEmptyInt(checkThis uint, description string) {
	if checkThis == 0 {
		log.WithField("field", description).Fatal("Needed config field is empty.")
	}
}

func NewConfig(configPath string) *models.Config {
	config := &models.Config{}
	file, err := os.Open(configPath)
	if err != nil {
		log.Error(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Error(err)
		}
	}(file)
	d := yaml.NewDecoder(file)
	if err := d.Decode(&config); err != nil {
		log.Error(err)
	}
	return config
}
