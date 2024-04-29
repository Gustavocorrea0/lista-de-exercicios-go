package main

import "fmt"

func main() {
	var numeroUm int
	var numeroDois int

	fmt.Println("Digite o primeiro Valor: ")
	_, err := fmt.Scan(&numeroUm)

	if err != nil {
		fmt.Println("Erro ao ler primeiro valor")
		return
	}

	fmt.Println("Digite o segundo Valor: ")
	_, err1 := fmt.Scan(&numeroDois)

	if err1 != nil {
		fmt.Println("Erro ao ler segundo valor")
		return
	}

	fmt.Println("+----------------------------+")
	fmt.Println("|      Valores Originais     |")
	fmt.Println("| > Valor 1: ", numeroUm)
	fmt.Println("| > Valor 2: ", numeroDois)

	var guardaValor = numeroUm
	numeroUm = numeroDois
	numeroDois = guardaValor

	fmt.Println("+----------------------------+")
	fmt.Println("|      Valores Trocados      |")
	fmt.Println("| > Valor 1: ", numeroUm)
	fmt.Println("| > Valor 2: ", numeroDois)
	fmt.Println("+----------------------------+")

}