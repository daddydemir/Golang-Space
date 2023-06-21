package main

import "fmt"

type Publisher struct {
	subscribers []Subscriber
}

func (p Publisher) addSubscriber(subscriber Subscriber) {
	p.subscribers = append(p.subscribers, subscriber)
}

func (p Publisher) removeSubscriber(subscriber Subscriber) {
	for i := 0; i < len(p.subscribers); i++ {
		if p.subscribers[i] == subscriber {
			p.subscribers = append(p.subscribers[:i], p.subscribers[i+1:]...)
			break
		}
	}
}

func (p Publisher) notifySubscribers() {
	for _, sub := range p.subscribers {
		sub.update()
	}
}

type Subscriber struct {
	id       int
	username string
}

func (subscriber Subscriber) update() {
	fmt.Println("state is updated.")
}
