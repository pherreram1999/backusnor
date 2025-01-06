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
			// se pasa por referencia para no mover el cabezal
			stack.insertString(&stack.Head, "iCtSA")
			break
		}
		stack.MoveRigth()
	}

	stack.Debug()
	fmt.Println(stack)

}
