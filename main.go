package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	MaxSteps = 10
)

func main() {
	stack := NewSymboStack()
	stack.AddFromString("iCtSA")

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	move := "R"
	ifCounter := 0

	for stack.Head != nil && ifCounter < MaxSteps {

		//fmt.Println("Movement:", move, "- Stack:", stack, "Symbol Input", stack.Head)

		if stack.State == "q0" && stack.Head.Symbol == 'A' {
			if r.Float64() > .50 || stack.Head.Right == nil {
				stack.insertString(";eS")
			} else {
				stack.Remove(stack.Head)
			}
			if stack.Head != nil && stack.Head.Right == nil {
				move = "L" // porque es el ultimo nodo
				stack.State = "q1"
			} else {
				move = "R"
			}
		} else if stack.State == "q0" && stack.Head.Right == nil {
			stack.State = "q1"
			move = "L"
		} else if stack.State == "q1" && stack.Head.Symbol == 'S' {
			stack.insertString("iCtSA")
			ifCounter++
			stack.State = "q1"
			move = "L"
		} else if stack.State == "q1" && stack.Head.Left == nil {
			stack.State = "q0"
			move = "R"
		}

		if move == "R" {
			stack.MoveRigth()
		} else {
			stack.MoveLeft()
		}

	}

	fmt.Println(stack)

}
