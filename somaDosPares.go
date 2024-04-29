package main

import "fmt"

func main() {
	var numero1 int
	var numero2 int

	fmt.Println("Digite o primeiro número: ")
	_, err := fmt.Scanln(&numero1)

	if err != nil {
		fmt.Println("Erro ao ler o primeiro número!")
		return
	}

	fmt.Println("Digite o segundo número: ")
	_, err1 := fmt.Scanln(&numero2)

	if err1 != nil {
		fmt.Println("Erro ao ler o segundo número!")
		return
	}

	fmt.Println("+------------------------------------------+")
	calcularPares(numero1, numero2)

}

func calcularPares(primeiroNumero int, segundoNumero int) {
	var valorDosPares int

	for i := primeiroNumero; i <= segundoNumero; i++ {
		fmt.Println(i)

		if i%2 == 0 {
			valorDosPares = valorDosPares + i
		}

	}

	fmt.Print("O valor da soma dos pares:", valorDosPares)
}
