package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	for {
		fmt.Println("+--------------------------------------------------------+")
		fmt.Println("|                  Sequência de Fibonacci                |")
		fmt.Println("+--------------------------------------------------------+")
		fmt.Print("| > Informe um número:")
		fmt.Scanln(&input)
		if isValidNumber(input) {
			break
		}
		fmt.Println("+-----------------------------------------------------------+")
		fmt.Println("|    Por favor, digite apenas Números!. Tente novamente    |")
		fmt.Println("+-----------------------------------------------------------+")
	}

	n, _ := strconv.Atoi(input)
	fibSequence := fibonacciSequence(n)
	fmt.Println("+------------------------------------------+")
	fmt.Println("+    Sequência de Fibonacci até", n, ":    ")
	fmt.Println("+------------------------------------------+")
	for _, num := range fibSequence {
		fmt.Print(num, " ")
	}
	fmt.Println(n)
}

func isValidNumber(input string) bool {
	_, err := strconv.Atoi(input)
	return err == nil
}

func fibonacciSequence(n int) []int {
	sequence := make([]int, 0)
	a, b := 0, 1
	for b <= n {
		sequence = append(sequence, a)
		a, b = b, a+b
	}
	return sequence
}
