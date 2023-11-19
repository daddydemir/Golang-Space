package message

import "fmt"

type Message interface {
	SendMessage(string) error
}

type EmailSender struct{}

func (e *EmailSender) SendMessage(msg string) error {
	fmt.Println("Email send :", msg)
	return nil
}

type SMSSender struct{}

func (s *SMSSender) SendMessage(msg string) error {
	fmt.Println("SMS send: ", msg)
	return nil
}
