package main

import (
	log "github.com/sirupsen/logrus"
	ctl "mailrelay/controllers"
)

var (
	cfg = ctl.NewConfig("./config.yaml")
)

func init() {
	log.SetLevel(log.DebugLevel)
	log.Print("Log level ", log.GetLevel(), ".")
	ctl.CheckEmptyString(cfg.Server.Port, "api port")
	ctl.CheckEmptyString(cfg.Mail.Server, "mail server")
	ctl.CheckEmptyString(cfg.Mail.Address, "mail address")
	ctl.CheckEmptyString(cfg.Mail.Password, "mail password")
	ctl.CheckEmptyInt(cfg.Mail.Port, "mail port")
}

func main() {
	ctl.ConnectDatabase(cfg.Database)
	ctl.GetMail(*cfg)
}
