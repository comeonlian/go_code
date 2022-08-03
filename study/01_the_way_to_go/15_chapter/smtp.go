package main

import (
	"bytes"
	"log"
	"net/smtp"
)

func main() {
	client, err := smtp.Dial("mail.126.com:25")
	if err != nil {
		log.Fatal(err)
	}

	client.Mail("leolian2015@126.com")
	client.Rcpt("3114833418@qq.com")

	wc, err := client.Data()
	if err != nil {
		log.Fatal(err)
	}

	defer wc.Close()

	buf := bytes.NewBufferString("This is the email body.")
	if _, err = buf.WriteTo(wc); err != nil {
		log.Fatal(err)
	}
}
