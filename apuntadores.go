package main

import "fmt"

func main() {
	x := 25
	fmt.Println(x)
	fmt.Println(&x)
	cambiarValor(x)
	// y tiene la referencia de x
	y := &x
	fmt.Println(y)
	// trae el valor almacenado en x
	fmt.Println(*y)
}

func cambiarValor(a int) {
	fmt.Println(&a)
	a = 36
}
