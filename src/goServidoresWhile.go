package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	inicio := time.Now()

	servidores := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}
	canal := make(chan string, len(servidores))

	i := 0
	for {
		if i > 3 {
			break
		}
		for _, servidor := range servidores {
			go revisarServidor(servidor, canal)
		}
		time.Sleep(1 * time.Second)
		fmt.Println(<-canal)
		i++
		fmt.Println(i)
	}

	tiempoPaso := time.Since(inicio)
	fmt.Printf("Tiempo de ejecución %s\n", tiempoPaso)
}

func revisarServidor(servidor string, canal chan string) {
	_, err := http.Get(servidor)
	if err != nil {
		fmt.Println(servidor, "No esta disponible =(")
		canal <- servidor + " No se encuentra disponible"
	} else {
		fmt.Println(servidor, "Esta funcionando normalmente =)")
		canal <- servidor + " Esta funcionando"
	}
}
