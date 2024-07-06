package main

import (
	"fmt"
	"strings"
)

func main() {
	println("Linked Lists")

	l := LinkedList{}
	l.add(1)
	l.add(2)
	l.add(3)
	l.add(4)
	l.add(5)

	fmt.Println(l)
	l.remove(3)
	fmt.Println(l)

}

type LinkedList struct {
	Head *node
}

type node struct {
	val  int
	next *node
}

func (n node) String() string {
	return fmt.Sprintf("%d", n.val)
}

func (l *LinkedList) add(value int) {
	newNode := new(node)
	newNode.val = value

	if l.Head == nil {
		l.Head = newNode
	} else {
		iterator := l.Head
		for ; iterator.next != nil; iterator = iterator.next {
		}
		iterator.next = newNode
	}
}

func (l *LinkedList) get(value int) *node {
	for iterator := l.Head; iterator != nil; iterator = iterator.next {
		if iterator.val == value {
			return iterator
		}
	}
	return nil
}

func (l *LinkedList) remove(value int) {

	var prev *node
	for current := l.Head; current != nil; current = current.next {
		if current.val == value {
			prev.next = current.next
			return
		}
		prev = current
	}
}

func (l LinkedList) String() string {
	sb := strings.Builder{}

	if l.Head != nil {
		for iterator := l.Head; iterator != nil; iterator = iterator.next {
			sb.WriteString(fmt.Sprintf("%s->", iterator))
		}
		sb.WriteString("nil")
	}

	return sb.String()
}
