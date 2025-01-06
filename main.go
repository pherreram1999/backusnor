package main

import "fmt"

const (
	IfsNum = 10
)

func main() {
	stack := NewSymboStack()

	stack.AddFromString("iCtSA")

	for stack.Head != nil {
		if stack.Head.Symbol == 'S' {
			stack.insertString("iCtSA")
			break
		}
		stack.MoveRigth()
	}

	stack.Debug()
	fmt.Println(stack)

}
