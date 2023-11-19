package main

import "di/message"

func main() {
	// MessageSender'a hangi servisi kullanacagi bilgisini veriyoruz
	emailSender := &message.EmailSender{}
	messageSender := &message.MessageSender{MessageService: emailSender}
	messageSender.Send("hello")

	smsSender := &message.SMSSender{}
	messageSender = &message.MessageSender{MessageService: smsSender}
	messageSender.Send("hello")
}
