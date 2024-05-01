package main

import "fmt"

func main() {
	var salarioSemCorrecao float32
	var percentualAumento float32

	fmt.Println("+----------------------------------------+")
	fmt.Println("|            Ajuste Salarial             |")
	fmt.Println("+----------------------------------------+")
	fmt.Println("| > Qual é o salario atual")
	_, err := fmt.Scanln(&salarioSemCorrecao)

	if err != nil || salarioSemCorrecao <= 0 {
		fmt.Println("+----------------------------------------+")
		fmt.Println("|      Erro ao ler o salario digitado    |")
		fmt.Println("+----------------------------------------+")
		return
	}

	fmt.Println("+----------------------------------------+")
	fmt.Println("| > Qual é o percentual de aumento: ")
	_, err = fmt.Scanln(&percentualAumento)

	if err != nil || percentualAumento <= 0 {
		fmt.Println("+----------------------------------------------+")
		fmt.Println("| Erro ao ler o percentual de aumento digitado |")
		fmt.Println("+----------------------------------------------+")
		return
	}

	var aumento = (salarioSemCorrecao * percentualAumento) / 100
	var novoSalario = salarioSemCorrecao + aumento

	fmt.Println("+----------------------------------------------+")
	fmt.Println("| > O novo salario é", novoSalario)
	fmt.Println("+----------------------------------------------+")
}
