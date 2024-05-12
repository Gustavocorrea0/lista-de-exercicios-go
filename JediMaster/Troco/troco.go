package main

import (
	"fmt"
	"strconv"
)

func main() {
	var valorTotal, valorPago float64

	for {
		fmt.Println("+-----------------------------------------+")
		fmt.Println("|                   Troco                 |")
		fmt.Println("+-----------------------------------------+")
		fmt.Print("| > Digite o valor total a ser pago:     ")

		if inputIsValid, value := getInput(); inputIsValid {
			valorTotal = value
			break
		}
	}

	for {
		fmt.Println("+-----------------------------------------+")
		fmt.Print("| > Digite o valor efetivamente pago:    ")
		if inputIsValid, value := getInput(); inputIsValid {
			valorPago = value
			if valorPago >= valorTotal {
				break
			}
			fmt.Println("+-------------------------------------------------------------------------------------------+")
			fmt.Println("|    O valor pago deve ser igual ou maior que o valor total a ser pago. Tente novamente.    |")
			fmt.Println("+-------------------------------------------------------------------------------------------+")
		}
	}

	troco := valorPago - valorTotal

	fmt.Println("+-----------------------------------------+")
	fmt.Println("|                 Troco:                  |")
	fmt.Println("+-----------------------------------------+")
	for _, nota := range []float64{100, 50, 10, 5, 1} {
		qtdNotas := int(troco / nota)
		if qtdNotas > 0 {
			fmt.Printf("| > %d nota(s) de R$%.2f\n", qtdNotas, nota)
			troco -= float64(qtdNotas) * nota
		}
	}

	for _, moeda := range []float64{0.5, 0.1, 0.05, 0.01} {
		qtdMoedas := int(troco / moeda)
		if qtdMoedas > 0 {
			fmt.Printf("| > %d moeda(s) de R$%.2f\n", qtdMoedas, moeda)
			troco -= float64(qtdMoedas) * moeda
		}
	}
}

func getInput() (bool, float64) {
	var input string
	fmt.Scanln(&input)
	value, err := strconv.ParseFloat(input, 64)
	if err != nil || value <= 0 {
		fmt.Println("+-----------------------------------------------------------+")
		fmt.Println("|    Por favor, digite um valor válido e maior que zero.    |")
		fmt.Println("+-----------------------------------------------------------+")
		return false, 0
	}
	return true, value
}
