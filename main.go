package main

import (
	"fmt"

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
}
