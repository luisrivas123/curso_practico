package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isError(e error) {
	if e != nil {
		log.Fatal("Error-> ", e)
	}
}

type calc struct {
}

// Receiver function, le da la propieda a calc de tener este método
func (calc) operate(entrada string, operador string) int {
	entradaLimpia := strings.Split(entrada, operador)
	operador1 := parsear(entradaLimpia[0])
	operador2 := parsear(entradaLimpia[1])
	switch operador {
	case "+":
		fmt.Println(operador1 + operador2)
		return operador1 + operador2
	case "-":
		fmt.Println(operador1 - operador2)
		return operador1 - operador2
	case "*":
		fmt.Println(operador1 * operador2)
		return operador1 * operador2
	case "/":
		fmt.Println(operador1 / operador2)
		return operador1 / operador2
	default:
		fmt.Println(operador, "Operador no soportado")
		return 0
	}
}

func parsear(entrada string) int {
	operador, err := strconv.Atoi(entrada)
	isError(err)
	return operador
}

func leerEntrada() string {
	// Leer input del usuario
	scanner := bufio.NewScanner(os.Stdin)
	// captura y devuele los valores
	scanner.Scan()
	return scanner.Text()
}

func main() {

	entrada := leerEntrada()
	operador := leerEntrada()
	c := calc{}
	fmt.Println(c.operate(entrada, operador))

}
