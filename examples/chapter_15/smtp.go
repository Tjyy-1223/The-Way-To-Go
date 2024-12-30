package main

import (
	"bytes"
	"log"
	"net/smtp"
)

func main() {
	// 连接到远程 SMTP 服务器。
	client, err := smtp.Dial("mail.example.com:25")
	if err != nil {
		log.Fatal(err)
	}

	// 设置寄件人和收件人
	client.Mail("sender@example.org")
	client.Rcpt("recipient@example.net")

	// 发送邮件主体。
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
