package example

type Notification struct {
	Content string
}

type SlackNotify struct {
	Notification
}

type SMSNotify struct {
	Notification
}

func (n *Notification) Send()  {
	println("Notification sending ... " , n.Content)
}

func (s *SlackNotify) Send()  {
	println("Slack Notification sending ... " , s.Content)
}


func (s *SMSNotify) Send()  {
	println("SMS Notification sending ... ", s.Content)
}
