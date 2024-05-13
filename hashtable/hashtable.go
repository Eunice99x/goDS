package hashtable

import (
	"fmt"
)

const ArraySize = 7

type Hashtable struct {
	array [ArraySize]*bucket
}

type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	key  string
	next *bucketNode
}

func (h *Hashtable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

func (h *Hashtable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

func (h *Hashtable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Print(k, "Already Exists")
	}
}

func (b *bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

func (b *bucket) delete(k string) {
	if b.head.key == k {
		b.head = b.head.next
		return
	}

	prevNode := b.head
	for prevNode.next != nil {
		if prevNode.next.key == k {
			prevNode.next = prevNode.next.next
		}
		prevNode = prevNode.next
	}
}

func hash(key string) int {
	s := 0
	for _, i := range key {
		s += int(i)
	}
	return s % ArraySize
}

func Init() *Hashtable {
	result := &Hashtable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func Hash() {
	test := Init()
	fmt.Println(test)
}
