package main

import (
	"fmt"
)

func main() {
	var numero int = 10
	for {
		fmt.Println("Qual é o numero:")
		_, err := fmt.Scanln(&numero)

		if err != nil || numero <= 0 {
			fmt.Printf("Número Inválido")
		} else {
			break
		}
	}
	calcularFatorial(numero)
}

func calcularFatorial(numero int) {
	resultado := 1
	for i := 1; i <= numero; i++ {
		resultado *= i
	}

	fmt.Println("O fatorial de", numero, "é", resultado)
}
