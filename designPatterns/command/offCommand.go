package main

type OffCommand struct {
	Device Device
}

func (o OffCommand) execute() {
	o.Device.off()
}
