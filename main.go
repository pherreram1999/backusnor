package main

import (
	"fmt"
)

const (
	IfsNum = 10
)

func main() {
	stack := NewSymboStack()

	stack.AddFromString("iCtSA")

	var ifCounter uint = 0

	move := "R"

	for stack.Head != nil && ifCounter < IfsNum {

		fmt.Printf("%s , %c, %s  => %s -> %p\n",
			stack.State,
			stack.Head.Symbol,
			move,
			stack.String(),
			stack.Head,
		)

		if stack.State == "q0" && stack.Head.Symbol == 'S' {
			// se pasa por referencia para no mover el cabezal
			fmt.Println("Expand")
			stack.insertString("iCtSA")
			stack.State = "q1"
			move = "R"
			ifCounter++
			stack.MoveRigth()
		} else if stack.State == "q1" && stack.Head.Symbol == 'A' {
			stack.insertString(";eS")
			stack.State = "q0"
			move = "L"
			ifCounter++
			stack.MoveLeft()
		} else if stack.State == "q0" {
			stack.MoveRigth()
		} else if stack.State == "q1" {
			stack.MoveLeft()
		}

		if stack.Head == nil {
			fmt.Println("head is null")
		}

	}

	stack.Debug()

}
