package main

import "fmt"

func main() {
	var valorEmDolar float32
	var cotacaoDolar float32
	fmt.Println("+----------------------------------------+")
	fmt.Println("|            Conversão Dolar             |")
	fmt.Println("+----------------------------------------+")
	fmt.Println("| > Digite a cotação do dolar: ")
	_, err := fmt.Scanln(&cotacaoDolar)

	if err != nil || cotacaoDolar <= 0 {
		fmt.Println("+----------------------------------------+")
		fmt.Println("|       Digite o valor da cotação        |")
		fmt.Println("+----------------------------------------+")
		return
	}

	fmt.Println("+----------------------------------------+")
	fmt.Println("| > Digite a quantidade de dolares: ")
	_, err1 := fmt.Scanln(&valorEmDolar)

	if err1 != nil || valorEmDolar <= 0 {
		fmt.Println("+----------------------------------------+")
		fmt.Println("|         Digite um valor valido         |")
		fmt.Println("+----------------------------------------+")
		return
	}

	converterParaDolar(cotacaoDolar, valorEmDolar)
}

func converterParaDolar(cotacao float32, quantidadeDolar float32) {
	var valorConvertido = quantidadeDolar / cotacao
	fmt.Println("+----------------------------------------+")
	fmt.Println("| > O valor convertido é igual a R$", valorConvertido)
	fmt.Println("+----------------------------------------+")
}
