package main

import "fmt"

type Tv struct {
	isRunning bool
}

func (t *Tv) on() {
	t.isRunning = true
	fmt.Printf("TV is running \n")
}

func (t *Tv) off() {
	t.isRunning = false
	fmt.Printf("TV is not running \n")
}
