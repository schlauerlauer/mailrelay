package main

import (
	"fmt"
	"mailrelay/models"
	log "github.com/sirupsen/logrus"
)

var (
	cfg = NewConfig("./config.yaml")
)

func init() {
	log.SetLevel(log.DebugLevel)
	log.Print("Log level ", log.GetLevel(), ".")
	CheckEmptyString(cfg.Server.Port, "api port")
	CheckEmptyString(cfg.Mail.Server, "mail server")
	CheckEmptyString(cfg.Mail.Address, "mail address")
	CheckEmptyString(cfg.Mail.Password, "mail password")
	CheckEmptyString(fmt.Sprint(cfg.Mail.Port), "mail port")
}

func main() {
	models.ConnectDatabase(cfg.Database)
	GetMail()
}
