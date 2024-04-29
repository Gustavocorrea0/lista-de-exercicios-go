package main

import "fmt"

func main() {
	var comprimentoCaixa float64
	var laguraCaixa float64
	var alturaCaixa float64

	fmt.Println("Qual é o comprimento da caixa:")
	_, err := fmt.Scanln(&comprimentoCaixa)

	if err != nil || comprimentoCaixa <= 0 {
		fmt.Println("Erro ao ler comprimento da caixa")
		return
	}

	fmt.Println("Qual é a lagura de caixa:")
	_, err1 := fmt.Scanln(&laguraCaixa)

	if err1 != nil || laguraCaixa <= 0 {
		fmt.Println("Erro ao ler largura da caixa")
		return
	}

	fmt.Println("Qual é a altura da caixa:")
	_, err2 := fmt.Scanln(&alturaCaixa)

	if err2 != nil || alturaCaixa <= 0 {
		fmt.Println("Erro ao ler altura da caixa")
		return
	}

	calcularVolume(comprimentoCaixa, laguraCaixa, alturaCaixa)
}

func calcularVolume(comprimento float64, largura float64, altura float64) {
	var volume = comprimento * largura * altura
	fmt.Println("O volume da caixa é:", volume)
}
