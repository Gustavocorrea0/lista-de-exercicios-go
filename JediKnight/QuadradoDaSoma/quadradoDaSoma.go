package main

import "fmt"

func main() {
	var valor1 int
	var valor2 int
	var valor3 int
	fmt.Println("+--------------------------------------+")
	fmt.Println("|           Quadrado da Soma           |")
	fmt.Println("+--------------------------------------+")
	fmt.Println("| > Qual é o primeiro número: ")
	_, err := fmt.Scanln(&valor1)

	if err != nil || valor1 < 0 {
		fmt.Println("+--------------------------------------+")
		fmt.Println("|         Quarto valor inválido        |")
		fmt.Println("+--------------------------------------+")
		return
	}

	fmt.Println("+--------------------------------------+")
	fmt.Println("| > Qual é o segundo número: ")
	_, err2 := fmt.Scanln(&valor2)

	if err2 != nil || valor2 < 0 {
		fmt.Println("+--------------------------------------+")
		fmt.Println("|        Segundo valor inválido        |")
		fmt.Println("+--------------------------------------+")
		return
	}

	fmt.Println("+--------------------------------------+")
	fmt.Println("| > Qual é o terceiro número: ")
	_, err3 := fmt.Scanln(&valor3)

	if err3 != nil || valor3 < 0 {
		fmt.Println("+--------------------------------------+")
		fmt.Println("|       Terceiro valor inválido        |")
		fmt.Println("+--------------------------------------+")
		return
	}

	quadradoDaSoma(valor1, valor2, valor3)

}

func quadradoDaSoma(valorUm int, valorDois int, valorTres int) {
	var soma int
	soma = valorUm + valorDois + valorTres
	soma = soma * soma
	fmt.Println("+--------------------------------------+")
	fmt.Println("| > O quadrado da soma é:", soma)
	fmt.Println("+--------------------------------------+")
}
