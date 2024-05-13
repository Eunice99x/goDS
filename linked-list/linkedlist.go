package linkedlist

import (
	"fmt"
	"strings"
)

type node struct {
	data int
	next *node
}

func (n node) String() string {
	return fmt.Sprintf("%d", n.data)
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) Add(data int) {
	newNode := new(node)
	newNode.data = data
	if l.head == nil {
		l.head = newNode
	} else {
		iterator := l.head
		for ; iterator.next != nil; iterator = iterator.next {
		}
		iterator.next = newNode
	}
}

func (l *linkedList) Remove(date int) {
	var prev *node
	for current := l.head; current.next != nil; current = current.next {
		if current.data == date {
			if l.head == current {
				l.head = current.next
			} else {
				prev.next = current.next
				return
			}
		}
		prev = current
	}

}

func (l linkedList) Get(data int) *node {
	for iterator := l.head; iterator != nil; iterator = iterator.next {
		if iterator.data == data {
			return iterator
		}
	}

	return nil
}

func (l linkedList) String() string {
	sb := strings.Builder{}

	for iterator := l.head; iterator != nil; iterator = iterator.next {
		sb.WriteString(fmt.Sprintf("%d", iterator.data))
	}

	return sb.String()
}

var L = linkedList{}
