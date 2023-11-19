package message

type MessageSender struct {
	MessageService Message
	// Dependeny Inversion: yuksek seviyeli modul dusuk seviyeli modulun detaylarina bagimli degil
}

func (ms *MessageSender) Send(msg string) error {
	return ms.MessageService.SendMessage(msg)
}
