package main

import "fmt"

func main() {
	var valorEmDolar float32
	var cotacaoDolar float32

	fmt.Println("Digite a cotação do dolar: ")
	_, err := fmt.Scanln(&cotacaoDolar)

	if err != nil || cotacaoDolar <= 0 {
		fmt.Println("Digite um valor valido para cotação")
		return
	}

	fmt.Println("Digite a quantidade de dolares: ")
	_, err1 := fmt.Scanln(&valorEmDolar)

	if err1 != nil || valorEmDolar <= 0 {
		fmt.Println("Digite um valor valido")
		return
	}

	converterParaReal(cotacaoDolar, valorEmDolar)
}

func converterParaReal(cotacao float32, quantidadeDolar float32) {
	var valorConvertido = quantidadeDolar * cotacao
	fmt.Println("O valor convertido é igual a R$", valorConvertido)
}
