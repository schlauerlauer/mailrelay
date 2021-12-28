package controllers

import (
	"fmt"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	log "github.com/sirupsen/logrus"
	"mailrelay/models"
)

func login(cfg models.Config) *client.Client {
	log.Debug("Connecting to server.")
	// Connect to server
	c, err := client.DialTLS(fmt.Sprintf("%s:%d", cfg.Mail.Server, cfg.Mail.Port), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Debug("Connected")
	// Login
	if err := c.Login(cfg.Mail.Address, cfg.Mail.Password); err != nil {
		log.Fatal(err)
	}
	log.Debug("Logged in")
	return c
}

func selectDirectory(dir string) {

}

func GetMail(cfg models.Config) {

	
	c := login(cfg)
	defer c.Logout()

	// List mailboxes
	mailboxes := make(chan *imap.MailboxInfo, 10)
	done := make(chan error, 1)
	go func () {
		done <- c.List("", "*", mailboxes)
	}()

	log.Info("Mailboxes:")
	for m := range mailboxes {
		log.Info("* " + m.Name)
	}

	if err := <-done; err != nil {
		log.Fatal(err)
	}

	// Select INBOX
	mbox, err := c.Select("INBOX", false)
	if err != nil {
		log.Fatal(err)
	}
	//log.Info("Flags for INBOX:", mbox.Flags)

	// Get the last 4 messages
	from := uint32(1)
	to := mbox.Messages
	if mbox.Messages > 3 {
		// We're using unsigned integers here, only subtract if the result is > 0
		from = mbox.Messages - 3
	}
	seqset := new(imap.SeqSet)
	seqset.AddRange(from, to)

	messages := make(chan *imap.Message, 10)
	done = make(chan error, 1)
	go func() {
		done <- c.Fetch(seqset, []imap.FetchItem{imap.FetchEnvelope}, messages)
	}()

	log.Info("Last 4 messages:")
	for msg := range messages {
		fmt.Print(msg)
		log.WithFields(log.Fields{
			"domain": msg.Envelope.From[0].HostName,
			"sender": msg.Envelope.From[0].Address(),
			"name": msg.Envelope.From[0].PersonalName,
			"recv": msg.Envelope.To[0].Address(),
		}).Info("Mail")
	}

	if err := <-done; err != nil {
		log.Fatal(err)
	}

	log.Info("Done!")
}
