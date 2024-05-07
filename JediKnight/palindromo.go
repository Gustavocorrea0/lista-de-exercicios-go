package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	menuPalidromo()
}

func menuPalidromo() {
	var opcao string
	for {
		fmt.Println("O que deseja inserir")
		fmt.Println("P - Palavra ou Frase")
		fmt.Println("N - Numero")
		fmt.Print("| > ")
		_, err := fmt.Scanln(&opcao)

		if err != nil {
			fmt.Println("Opção inválida")
		}

		switch strings.ToUpper(opcao) {
		case "P":
			palindromoPalavraOuFrase()
		case "N":
			palindromoNumero()
		default:
			fmt.Println("Opção Inválida")
		}
	}
}

func palindromoPalavraOuFrase() {
	var palavraOuFrase string
	var fraseInvertida strings.Builder

	scannerBufio := bufio.NewScanner(os.Stdin)
	fmt.Println("Digite a palavra ou frase desejada:")
	for scannerBufio.Scan() {
		palavraOuFrase = scannerBufio.Text()
		if len(palavraOuFrase) > 0 {
			break
		}
		fmt.Println("Entrada inválida. Digite novamente:")
	}

	palavraOuFraseSemEspacos := strings.ReplaceAll(strings.ToLower(palavraOuFrase), " ", "")

	for i := len(palavraOuFraseSemEspacos) - 1; i >= 0; i-- {
		fraseInvertida.WriteByte(palavraOuFraseSemEspacos[i])
	}

	fmt.Println("Palavra ou Frase:", palavraOuFrase)
	fmt.Println("Palavra ou Frase invertida:", fraseInvertida.String())

	if palavraOuFraseSemEspacos == fraseInvertida.String() {
		fmt.Println("A Palavra ou Frase é um Palíndromo")
	} else {
		fmt.Println("A Palavra ou Frase não é um Palíndromo")
	}

	os.Exit(0)
}

func palindromoNumero() {
	var numero int
	var palindromo bool = true

	for {
		fmt.Println("Digite o numero desejado:")
		_, errNumero := fmt.Scanln(&numero)

		if errNumero != nil {
			fmt.Println("Valor Inválido")
		} else {
			break
		}
	}

	numeroConvertidoParaString := strconv.Itoa(numero)

	for i, x := 0, len(numeroConvertidoParaString)-1; i < x; i, x = i+1, x-1 {
		if numeroConvertidoParaString[i] != numeroConvertidoParaString[x] {
			palindromo = false
			break
		}
	}

	if palindromo {
		fmt.Println("O numero é um Palíndromo")
	} else {
		fmt.Println("O numero não é um Palíndromo")
	}
	os.Exit(0)
}
