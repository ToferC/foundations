package services

import (
	"fmt"
	"log"
	"net/smtp"
	"net/textproto"

	"github.com/jordan-wright/email"
)

// SendMail sends an email notification
func SendMail(to, subject, body, mailPassword string) {

	fmt.Println(mailPassword)

	e := &email.Email{
		To:      []string{to},
		From:    "Foundations App <foundationsappmail@gmail.com>",
		Subject: subject,
		HTML:    []byte(body),
		Headers: textproto.MIMEHeader{},
	}

	err := e.Send("smtp.gmail.com:587", smtp.PlainAuth("", "foundationsappmail@gmail.com", mailPassword, "smtp.gmail.com"))
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Email Sent")
	fmt.Printf("%s", e)
}
