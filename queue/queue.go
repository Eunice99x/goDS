package queue

import "fmt"

type Queue []string

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) Enqueue(str string) {
	*q = append(*q, str)
}

func (q *Queue) Dequeue() (string, bool) {
	if q.IsEmpty() {
		return "", false
	} else {
		element := (*q)[0]
		*q = (*q)[1:]
		return element, true
	}
}

func QueueDs() {
	var queue Queue

	queue.Enqueue("Data Structures")
	queue.Enqueue("Golang")
	queue.Enqueue("Queue")

	for i, val := range queue {
		fmt.Println(i, val)
	}
	queue.Dequeue()
	for i, val := range queue {
		fmt.Println(i, val)
	}
}
