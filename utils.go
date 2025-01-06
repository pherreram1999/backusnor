package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const MaxDimensionsLen = 1_000

func input(label string) string {
	var s string
	// creamos una lectura sobre la entrada el de stdin para captuar lo escbrime sobre la terminal
	r := bufio.NewReader(os.Stdin)
	for {
		_, _ = fmt.Fprintf(os.Stderr, "%s: ", label+" ")
		s, _ = r.ReadString('\n') // rompemos con salto de linea
		if s != "" {              // si leemos un string vacio, sin \n detenemos
			break
		}
	}
	// regresamos la cadena sin espacios extras
	return strings.TrimSpace(s)
}

func AskNumSteps(r *rand.Rand) {
	validNumberString := regexp.MustCompile("^[0-9]*$")
	for {
		MaxSteps = r.Intn(MaxDimensionsLen)
		if MaxSteps == 0 { // en el caso de que rand retorne 0
			MaxSteps = 1
		}
		nString := input(fmt.Sprintf("Valor de n (por defecto %d) entre 1 y %d: ", MaxSteps, MaxDimensionsLen))
		if nString == "" {
			break // usamos nuestra n por defecto
		}
		if !validNumberString.MatchString(nString) {
			fmt.Printf("** Favor de ingresar un numero valido entre 0 y %d\n", MaxDimensionsLen)
			continue
		}
		nNumber, err := strconv.Atoi(nString)

		if err != nil {
			panic(err)
		}

		if nNumber > MaxDimensionsLen || nNumber < 1 {
			fmt.Printf("!!!!! n debe ser un valor entre 1 y %d\n", MaxDimensionsLen)
			continue
		}

		MaxSteps = nNumber
		break
	}
}
