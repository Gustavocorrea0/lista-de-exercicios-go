package main

import "fmt"

func main() {
	var primeiroNumeroDoIntervalo int
	var segundoNumeroDoIntervalo int

	fmt.Println("Qual é o primeiro numero do intervalor: ")
	_, err1 := fmt.Scanln(&primeiroNumeroDoIntervalo)

	if err1 != nil || primeiroNumeroDoIntervalo < 0 {
		fmt.Printf("Erro ao ler primeiro numero")
		return
	}

	fmt.Println("Qual é o segundo numero do intervalo: ")
	_, err2 := fmt.Scanln(&segundoNumeroDoIntervalo)

	if err2 != nil || segundoNumeroDoIntervalo < 0 {
		fmt.Printf("Erro ao ler segundo numero")
		return
	}

	fmt.Println("+--------------------------------------------+")

	if primeiroNumeroDoIntervalo > segundoNumeroDoIntervalo {
		fmt.Println("O primeiro valor deve ser menor que o segundo")
		return
	}

	for i := primeiroNumeroDoIntervalo; i < segundoNumeroDoIntervalo; i++ {
		validarFizzBuzz(i)
	}

}

func validarFizzBuzz(numero int) {
	if numero%3 == 0 && numero%5 == 0 {
		fmt.Println("FizzBuzz")
		return
	} else if numero%3 == 0 {
		fmt.Println("Fizz")
		return
	} else if numero%5 == 0 {
		fmt.Println("Buzz")
		return
	} else {
		fmt.Println(numero)
	}
}
