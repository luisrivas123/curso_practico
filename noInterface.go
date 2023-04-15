package main

import "fmt"

type perro struct{}

type pez struct{}

type pajaro struct{}

func (perro) caminar() string {
	return "Soy un perro que camina"
}

func (pez) nada() string {
	return "Soy un pez que nada"
}

func (pajaro) vuela() string {
	return "Soy un pajaro que vuela"
}

func moverPerro(p perro) {
	fmt.Println(p.caminar())
}

func moverPez(p pez) {
	fmt.Println(p.nada())
}

func moverPajaro(p pajaro) {
	fmt.Println(p.vuela())
}

func main() {
	p := perro{}
	moverPerro(p)
	pe := pez{}
	moverPez(pe)
	pa := pajaro{}
	moverPajaro(pa)
}
