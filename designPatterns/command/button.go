package main

type Button struct {
	Command Command
}

func (b *Button) press() {
	b.Command.execute()
}
