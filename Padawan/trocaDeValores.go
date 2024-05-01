package main

import "fmt"

func main() {
	var numeroUm int
	var numeroDois int
	fmt.Println("+-------------------------------------------------+")
	fmt.Println("|                 Troca de Valores                |")
	fmt.Println("+-------------------------------------------------+")
	fmt.Println("| > Digite o primeiro Valor: ")
	_, err := fmt.Scan(&numeroUm)

	if err != nil || numeroUm < 0 {
		fmt.Println("+-------------------------------------------------+")
		fmt.Println("|             Primeiro valor inválido             |")
		fmt.Println("+-------------------------------------------------+")
		return
	}

	fmt.Println("+-------------------------------------------------+")
	fmt.Println("| > Digite o segundo Valor: ")
	_, err1 := fmt.Scan(&numeroDois)

	if err1 != nil || numeroDois < 0 {
		fmt.Println("+-------------------------------------------------+")
		fmt.Println("|             Segundo valor inválido              |")
		fmt.Println("+-------------------------------------------------+")
		return
	}

	fmt.Println("+-------------------------------------------------+")
	fmt.Println("|                Valores Originais                |")
	fmt.Println("| > Valor 1: ", numeroUm)
	fmt.Println("| > Valor 2: ", numeroDois)
	fmt.Println("+-------------------------------------------------+")

	var guardaValor = numeroUm
	numeroUm = numeroDois
	numeroDois = guardaValor

	fmt.Println("|                Valores Trocados                 |")
	fmt.Println("| > Valor 1: ", numeroUm)
	fmt.Println("| > Valor 2: ", numeroDois)
	fmt.Println("+-------------------------------------------------+")

}
