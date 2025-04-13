package notificator

import (
	"log"
	"net/smtp"
)

const GMAIL_EMAIL = "codesrelease@gmail.com"
const GMAIL_PASSWORD = "kyjdEk-xiwjif-7gitku"

type Email struct {
	To      string
	Subject string
	Body    string
}

func SendEmail(email Email) error {
	msg := "From: " + GMAIL_EMAIL + "\n" +
		"To: " + email.To + "\n" +
		"Subject: " + email.Subject + "\n\n" +
		email.Body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", GMAIL_EMAIL, GMAIL_PASSWORD, "smtp.gmail.com"),
		GMAIL_EMAIL, []string{email.To}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return err
	}

	return nil
}
