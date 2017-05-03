package main

import (
	"log"
	"net/smtp"
)

func main() {
	// Set up authentication information.
	auth := smtp.PlainAuth(
		"",
		"huan.com",
		"",
		"smtp.163.com",
	)
	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(
		"smtp.163.com:25",
		auth,
		"huan.com",
		[]string{""},
		[]byte("Hahahahahah."),
	)
	if err != nil {
		log.Fatal(err)
	}
}