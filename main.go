package main

import (
	"fmt"
	"time"
)

const (
	IfsNum = 10
)

func main() {
	stack := NewSymboStack()

	stack.AddFromString("iCtSA")

	for stack.Head != nil {
		if stack.Head.Symbol == 'S' {
			// se pasa por referencia para no mover el cabezal
			stack.insertString("iCtSA")
			fmt.Println(stack.Head)
			fmt.Println(stack)

			time.Sleep(time.Second * 3)
		}
		stack.MoveRigth()
	}

}
