package main

import (
	"fmt"
	"strconv"
)

func main() {
	var values []int
	var saida string
	var sum int
	count := 0

	fmt.Println("----------)&&&(---------")
	fmt.Println("-----SOMA DOS PARES-----")
	fmt.Println("----------)&&&(---------")

	for {
		fmt.Print("Digite um número (ou 'S' para sair): ")
		fmt.Scanln(&saida)

		if saida == "S" {
			break
		}

		// strconv.Atoi é uma converte uma string que representa um número em um valor inteiro
		num, err := strconv.Atoi(saida)
		if err != nil {
			fmt.Println("INSIRA UM VALOR VALIDO!.")
			continue
		}

		//adiciona um valor inteiro a fatias de um values
		values = append(values, num)
		sum += num
		count++
	}

	if count == 0 {
		fmt.Println("NENHUM NUMERO FOI INSERIDO.")
		return
	}

	average := float64(sum) / float64(count)
	fmt.Printf("A média dos números digitados é de: %.2f\n", average)
}
