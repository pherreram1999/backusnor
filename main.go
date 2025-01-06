package main

import (
	"fmt"
	"time"
)

const (
	IfsNum   = 10
	MaxSteps = 1000
)

func main() {
	stack := NewSymboStack()
	stack.AddFromString("iCtSA")

	var ifCounter uint = 0
	move := "R"
	stepCounter := 0

	for stack.Head != nil && ifCounter < IfsNum && stepCounter < MaxSteps {
		stepCounter++

		fmt.Printf("Estado: %s, Cabezal: %c, Movimiento: %s, Cinta: [%s]\n",
			stack.State,
			stack.Head.Symbol,
			move,
			stack.String(),
		)

		if stack.State == "q0" && stack.Head.Symbol == 'S' {
			stack.insertString("iCtSA")
			ifCounter++

			stack.State = "q1"
			move = "R"
		} else if stack.State == "q1" && stack.Head.Symbol == 'A' {
			stack.insertString(";eS")
			stack.State = "q0"
			move = "L"
		}

		// Movimiento del cabezal
		if move == "R" {
			stack.MoveRigth()
		} else {
			stack.MoveLeft()
		}

		time.Sleep(1 * time.Second)
	}

	if stepCounter >= MaxSteps {
		fmt.Println("Máquina estancada. Terminando ejecución.")
	} else {
		fmt.Println("Ejecución terminada correctamente.")
	}
}
