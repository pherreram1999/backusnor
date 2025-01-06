package main

import (
	"fmt"
	"os"
)

const tabSize = 1

func SaveCode(production string) {
	_ = os.Remove("code.txt")
	fileCode, err := os.Create("code.txt")
	if err != nil {
		panic(err)
	}
	defer fileCode.Close()

	tabCount := 0

	printab := func() {
		for i := 0; i < tabCount; i++ {
			_, _ = fmt.Fprintf(fileCode, "\t") // el espaciado
		}
	}

	for _, sys := range production {
		if sys == 'i' {
			printab()
			_, _ = fmt.Fprintf(fileCode, "if ")
		} else if sys == 'C' {
			_, _ = fmt.Fprintf(fileCode, "condition ")
		} else if sys == 't' {
			_, _ = fmt.Fprintf(fileCode, "then\n")
			tabCount += tabSize
		} else if sys == 'S' {
			printab()
			_, _ = fmt.Fprintf(fileCode, "statment\n")
		} else if sys == ';' {
			tabCount -= tabSize
		} else if sys == 'e' {
			printab()
			_, _ = fmt.Fprintf(fileCode, "else\n")
			tabCount += tabSize
		}
	}
}
