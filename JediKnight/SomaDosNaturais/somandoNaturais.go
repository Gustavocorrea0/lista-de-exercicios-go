package main

import (
	"fmt"
)

func main() {
	var resultado int
	var primeiroNumero int = 0
	var segundoNumero int = 99

	for i := primeiroNumero; i < segundoNumero; i++ {
		resultado = resultado + i
	}

	fmt.Println("+------------------------------------------+")
	fmt.Println("| > Primeiro número:", primeiroNumero)
	fmt.Println("| > Segundo número:", segundoNumero)
	fmt.Println("| > O resultado da soma é:", resultado)
	fmt.Println("+------------------------------------------+")
}
