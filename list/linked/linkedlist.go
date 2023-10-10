package main

import "fmt"

type LinkedList struct {
	Value    int
	Last     bool
	Previous *LinkedList
	Next     *LinkedList
}

func main() {
	root := New(5)
	root.Add(6)
	root.Add(7)

	root, _ = root.Delete(6)

	fmt.Printf("qty: %d\n", root.CountNodes())
	fmt.Printf("count: %d\n", root.CountValues())
}

func New(v int) *LinkedList {
	return &LinkedList{
		Value: v,
	}
}

func (r *LinkedList) Add(v int) (*LinkedList, bool) {
	if r == nil {
		return nil, false
	}
	newNode := &LinkedList{
		Value: v,
	}
	if r.Next == nil {
		r.Next = newNode
	} else {
		r.Next.Add(newNode.Value)
	}

	return newNode, true
}

func (r *LinkedList) CountNodes() int {
	var counter int
	if r == nil {
		return 0
	}
	counter++
	for {
		if r.Next != nil {
			counter++
			r = r.Next
		} else {
			break
		}
	}
	return counter
}

func (r *LinkedList) CountValues() int {
	var sum int
	if r == nil {
		return 0
	}
	sum += r.Value
	for {
		if r.Next != nil {
			r = r.Next
			sum += r.Value
		} else {
			break
		}
	}
	return sum
}

func (r *LinkedList) Delete(v int) (*LinkedList, bool) {
	if r == nil {
		return nil, false
	}
	//var previousNode *LinkedList
	if r.Value == v {
		next := r.Next
		if next.Value == v {
			previous := r
			previous.Next = next
		}
		//if previousNode != nil {
		//	previousNode.Next = r
		//	return previousNode, true
		//}
		return r, true
	} else {
		r.Next.Delete(v)
	}

	return r, false
}
