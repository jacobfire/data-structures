package main

import "fmt"

type Node struct {
	key   string
	value string
	left  *Node
	right *Node
}

func main() {
	s := NewEmptyStack()

	n := Node{
		key:   "node",
		value: "hello",
	}

	nn := Node{
		key:   "node",
		value: "hello",
	}
	s.Push(n)
	s.Push(nn)
	n2 := s.Top()
	fmt.Println(n2)

	fmt.Println(s.Size())
}
