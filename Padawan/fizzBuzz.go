package main

import "fmt"

func main() {
	var primeiroNumeroDoIntervalo int
	var segundoNumeroDoIntervalo int
	fmt.Println("+------------------------------------------------+")
	fmt.Println("|                    FizzBuzz                    |")
	fmt.Println("+------------------------------------------------+")
	fmt.Println("| > Qual é o primeiro numero do intervalo: ")
	_, err1 := fmt.Scanln(&primeiroNumeroDoIntervalo)

	if err1 != nil || primeiroNumeroDoIntervalo < 0 {
		fmt.Println("+------------------------------------------------+")
		fmt.Println("|            Erro ao ler primeiro numero         |")
		fmt.Println("+------------------------------------------------+")
		return
	}

	fmt.Println("+------------------------------------------------+")
	fmt.Println("| > Qual é o segundo numero do intervalo: ")
	_, err2 := fmt.Scanln(&segundoNumeroDoIntervalo)

	if err2 != nil || segundoNumeroDoIntervalo < 0 {
		fmt.Println("+------------------------------------------------+")
		fmt.Println("|            Erro ao ler segundo numero          |")
		fmt.Println("+------------------------------------------------+")
		return
	}

	if primeiroNumeroDoIntervalo > segundoNumeroDoIntervalo {
		fmt.Println("+-------------------------------------------------+")
		fmt.Println("|  O primeiro valor deve ser menor que o segundo  |")
		fmt.Println("+-------------------------------------------------+")
		return
	}

	fmt.Println("+------------------------------------------------+")
	fmt.Println("|           	  Resultado 			 |")
	fmt.Println("")
	for i := primeiroNumeroDoIntervalo; i < segundoNumeroDoIntervalo; i++ {
		validarFizzBuzz(i)
	}

}

func validarFizzBuzz(numero int) {
	if numero%3 == 0 && numero%5 == 0 {
		fmt.Println("| > FizzBuzz")
		return
	} else if numero%3 == 0 {
		fmt.Println("| > Fizz")
		return
	} else if numero%5 == 0 {
		fmt.Println("| > Buzz")
		return
	} else {
		fmt.Println("| >", numero)
	}
}
