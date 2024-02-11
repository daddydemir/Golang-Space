package main

type Notifier interface {
	Send(content string) 
}

type NotifyDecorator struct {
	wrappe []Notifier
}

func (n *NotifyDecorator) Send(content string)  {
	for _ , o := range n.wrappe {
		o.Send(content)
	}
}

type Notify struct {}
func (n *Notify) Send(content string)  {
	println("Notification sending ... " , content)
}

type SlackNotify struct {}
func (s *SlackNotify) Send(content string)  {
	println("Slack Notification sending ... " , content)
}