package stack

import (
	"fmt"
)

type Stack []string

// check if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// push element to the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

// remove an element from the stack
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func StackDs() {
	var stack Stack

	stack.Push("Data Structures")
	stack.Push("Golang")

	for _, val := range stack {
		fmt.Println(val)
	}
	stack.Pop()
	for _, val := range stack {
		fmt.Println(val)
	}
}
