package main

import (
	"fmt"
)

func main() {
	var totalNumero int = 12
	var numeroDigitado int
	var numeros []int

	for i := 0; i < totalNumero; i++ {
		fmt.Println("+--------------------------------------+")
		fmt.Println("|       Ordenar números inteiro        |")
		for {
			fmt.Println("+------------------------------------------------+")
			fmt.Println("| > Insira um número( faltam", totalNumero-i, "):")
			_, err := fmt.Scanln(&numeroDigitado)

			if err != nil || numeroDigitado < 0 {
				fmt.Println("+--------------------------------------+")
				fmt.Println("|             Valor Inválido           |")
				fmt.Println("+--------------------------------------+")
			} else {
				break
			}
		}

		//       INSERE | ARRAY | VALOR INSERIDO
		numeros = append(numeros, numeroDigitado)
	}

	fmt.Println("+-------------------------------------------------------------------------------------+")
	fmt.Println("| > Numeros Digitados:", numeros)
	ordenarComSort(numeros)
	fmt.Println("| > Numeros Ordenados:", numeros)
	fmt.Println("+-------------------------------------------------------------------------------------+")
}

func ordenarComSort(arrayDeNumeros []int) {
	comprimentoArray := len(arrayDeNumeros)
	for i := 0; i < comprimentoArray; i++ {
		for x := 0; x < comprimentoArray-i-1; x++ {
			if arrayDeNumeros[x] > arrayDeNumeros[x+1] {
				arrayDeNumeros[x], arrayDeNumeros[x+1] = arrayDeNumeros[x+1], arrayDeNumeros[x]
			}
		}
	}
}
