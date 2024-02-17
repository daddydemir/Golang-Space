package main

func main() {

	tv := &Tv{}

	onCommand := OnCommand{tv}
	onButton := Button{onCommand}
	onButton.press()

	offCommand := OffCommand{tv}
	offButton := Button{offCommand}
	offButton.press()

}
