package main

import "fmt"

func main() {
	var primeiroNumeroDoIntervalor int
	var segundoNumeroDoIntervalor int

	fmt.Println("Qual é o primeiro numero do intervalor: ")
	_, err1 := fmt.Scanln(&primeiroNumeroDoIntervalor)

	if err1 != nil {
		fmt.Printf("Erro ao ler primeiro numero")
		return
	}

	fmt.Println("Qual é o segundo numero do intervalo: ")
	_, err2 := fmt.Scanln(&segundoNumeroDoIntervalor)

	if err2 != nil {
		fmt.Printf("Erro ao ler segundo numero")
		return
	}

	fmt.Printf("+--------------------------------------------+")

	for i := primeiroNumeroDoIntervalor; i < segundoNumeroDoIntervalor; i++ {
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
