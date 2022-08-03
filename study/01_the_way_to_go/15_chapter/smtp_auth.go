package main

import (
	"log"
	"net/smtp"
)

func main() {
	auth := smtp.PlainAuth(
		"",
		"lianliang2014@126.com",
		"",
		"mail.126.com",
	)

	err := smtp.SendMail(
		"mail.126.com",
		auth,
		"lianliang2014@126.com",
		[]string{"3114833418@qq.com"},
		[]byte("This is the email test body"),
	)

	if err != nil {
		log.Fatal(err)
	}
}
