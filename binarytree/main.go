package main

import (
	"fmt"
)

func main() {

	t := &TreeNode{val: 8}

	t.Insert(1)
	t.Insert(2)
	t.Insert(3)
	t.Insert(4)
	t.Insert(5)
	t.Insert(6)
	t.Insert(7)

	notExistValue := 12
	finding, ok := t.Find(notExistValue)
	if !ok {
		fmt.Println("node with value %d not fount in tree", notExistValue)
	}

	existValue := 6
	finding, ok = t.Find(existValue)
	if !ok {
		fmt.Println("node with value %d not fount in tree", existValue)
	}

	fmt.Printf("%+v", finding)

	t.Delete(5)
	t.Delete(7)
	t.Delete(8)

	t.PrintInorder()
	fmt.Println("")

	fmt.Println("min is %d", t.FindMin())
	fmt.Println("max is %d", t.FindMax())

	t.PrintInorder()
}
