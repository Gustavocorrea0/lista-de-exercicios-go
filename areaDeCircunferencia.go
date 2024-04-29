package main

import "fmt"

func main() {

	var raio float64

	fmt.Println("Calculo de area da circunferencia")
	fmt.Println("Qual é o raio: ")
	_, err := fmt.Scanln(&raio)

	if err != nil {
		fmt.Println("Erro ao ler o valor")
		return
	}

	calcular(raio)

}

func calcular(valorRaio float64) {
	var pi = 3.14159265
	var area = pi * (valorRaio * valorRaio)
	fmt.Println("O resultado é:", area)
	return
}
