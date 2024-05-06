package main

import "fmt"

func main() {
	var valor1 int
	var valor2 int
	var valor3 int

	fmt.Println("Qual é o primeiro número: ")
	_, err := fmt.Scanln(&valor1)

	if err != nil || valor1 < 0 {
		fmt.Println("Primeiro valor inválido")
		return
	}

	fmt.Println("Qual é o segundo número: ")
	_, err2 := fmt.Scanln(&valor2)

	if err2 != nil || valor2 < 0 {
		fmt.Println("Segundo valor inválido")
		return
	}

	fmt.Println("Qual é o terceiro número: ")
	_, err3 := fmt.Scanln(&valor3)

	if err3 != nil || valor3 < 0 {
		fmt.Println("Segundo valor inválido")
		return
	}

	calcularQuadrados(valor1, valor2, valor3)

}

func calcularQuadrados(valorUm int, valorDois int, valorTres int) {
	var somaDosQuadrados int

	valorUm = valorUm * valorUm
	valorDois = valorDois * valorDois
	valorTres = valorTres * valorTres

	somaDosQuadrados = valorUm + valorDois + valorTres
	fmt.Println("| A soma dos quadrados é:", somaDosQuadrados)
}
