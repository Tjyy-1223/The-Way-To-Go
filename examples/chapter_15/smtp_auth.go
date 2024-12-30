package main

import (
	"log"
	"net/smtp"
)

func main() {
	// 设置认证信息。
	auth := smtp.PlainAuth(
		"",
		"user@example.com",
		"password",
		"mail.example.com",
	)

	// 连接到服务器, 认证, 设置发件人、收件人、发送的内容,
	// 然后发送邮件。
	err := smtp.SendMail(
		"mail.example.com:25",
		auth,
		"sender@example.org",
		[]string{"recipient@example.net"},
		[]byte("This is the email body."),
	)
	if err != nil {
		log.Fatal(err)
	}
}
