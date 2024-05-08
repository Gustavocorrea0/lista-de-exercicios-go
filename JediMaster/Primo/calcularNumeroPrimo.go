package main

import (
	"fmt"
)

func main() {
	var numeroUsuario int
	for {
		fmt.Println("+--------------------------------------+")
		fmt.Println("|       Calculo de Número Primo        |")
		fmt.Println("+--------------------------------------+")
		fmt.Print("| > Qual é o número: ")
		_, err := fmt.Scanln(&numeroUsuario)

		if err != nil {
			fmt.Println("+--------------------------------------+")
			fmt.Println("|            Valor Inválido            |")
			fmt.Println("+--------------------------------------+")
		} else {
			break
		}
	}

	validarNumero(numeroUsuario)
}

func validarNumero(numero int) {
	if numero <= 1 {
		fmt.Println("+--------------------------------------+")
		fmt.Println("|          O número não é primo        |")
		fmt.Println("+--------------------------------------+")
		return
	}

	if numero <= 3 {
		fmt.Println("+--------------------------------------+")
		fmt.Println("|          O número é primo            |")
		fmt.Println("+--------------------------------------+")
		return
	}

	if numero%2 == 0 || numero%3 == 0 {
		fmt.Println("+--------------------------------------+")
		fmt.Println("|          O número não é primo        |")
		fmt.Println("+--------------------------------------+")
		return
	}

	for i := 5; i*i <= numero; i += 6 {
		if numero%i == 0 || numero%(i+2) == 0 {
			fmt.Println("+--------------------------------------+")
			fmt.Println("|          O número não é primo        |")
			fmt.Println("+--------------------------------------+")
			return
		}
	}

	fmt.Println("+--------------------------------------+")
	fmt.Println("|          O número é primo            |")
	fmt.Println("+--------------------------------------+")
}
