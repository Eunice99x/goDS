package main

import (
	"fmt"

	"github.com/YounesOuterbah/goDS.git/hashtable"
	linkedlist "github.com/YounesOuterbah/goDS.git/linked-list"
	"github.com/YounesOuterbah/goDS.git/queue"
	"github.com/YounesOuterbah/goDS.git/stack"
)

type Stack []string

func main() {
	fmt.Println("Data Structures Using Go")
	fmt.Println("#################################### ")

	stack.StackDs()
	fmt.Println("#################################### ")
	queue.QueueDs()
	fmt.Println("#################################### ")
	linkedlist.L.Add(1)
	linkedlist.L.Add(2)
	linkedlist.L.Add(3)
	fmt.Println(linkedlist.L)
	fmt.Println(linkedlist.L.Get(1))
	linkedlist.L.Remove(2)
	fmt.Println(linkedlist.L)
	fmt.Println("#################################### ")
	hashtable.Hash()
}
