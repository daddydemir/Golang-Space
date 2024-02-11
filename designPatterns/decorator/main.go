package main

import (
	"decorator/example"
)


func oldMain() {
	n := example.Notification{Content: "Hello World!"}
	n.Send()

	s := example.SMSNotify{example.Notification{Content: "Sms content"}}
	s.Send()

	slack := example.SlackNotify{example.Notification{Content: "Slack content"}}
	slack.Send()
}

func main() {
	
	normalNotify := new(Notify)
	slackNotify := new(SlackNotify)
	
	notifies := make([]Notifier, 0)
	notifies = append(notifies, normalNotify)
	notifies = append(notifies, slackNotify)
	
	decorator := NotifyDecorator{notifies}
	decorator.Send("message!!!")
	
	
}