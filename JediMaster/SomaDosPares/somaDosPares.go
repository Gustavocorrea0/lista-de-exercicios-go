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

	fmt.Println("+--------------------------------------+")
	fmt.Println("|            SOMA DOS PARES            |")
	fmt.Println("+--------------------------------------+")

	for {
		fmt.Println("+--------------------------------------+")
		fmt.Print("| > Digite um número (ou 'S' para sair): ")
		fmt.Scanln(&saida)

		if saida == "S" {
			break
		}

		num, err := strconv.Atoi(saida)
		if err != nil {
			fmt.Println("+--------------------------------------+")
			fmt.Println("|      INSIRA UM VALOR VALIDO!.        |")
			continue
		}

		values = append(values, num)
		sum += num
		count++
	}

	if count == 0 {
		fmt.Println("+--------------------------------------+")
		fmt.Println("|      NENHUM NUMERO FOI INSERIDO      |")
		return
	}

	media := float64(sum) / float64(count)
	fmt.Println("+--------------------------------------+")
	fmt.Println("| > A média dos números digitados é de: ", media)
}
