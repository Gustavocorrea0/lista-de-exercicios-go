package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var frase string
	var err error
	for {
		fmt.Println("+-----------------------------------------+")
		fmt.Print("| > Digite a frase:")
		frase, err = obterFrase(reader)
		if err == nil {
			break
		}
		fmt.Println("| Erro:", err)
	}

	var colunas int
	for {
		fmt.Println("+-----------------------------------------+")
		fmt.Print("| > Digite o número de colunas:")
		colunas, err = obterNumeroColunas(reader)
		if err == nil {
			break
		}
		fmt.Println("| Erro:", err)
	}

	fraseQuebrada := quebrarFrase(frase, colunas)

	fmt.Println("+-----------------------------------------+")
	fmt.Println("| Frase quebrada:")
	fmt.Println(fraseQuebrada)
}

func obterFrase(reader *bufio.Reader) (string, error) {
	frase, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	frase = strings.TrimSpace(frase)

	if !contemApenasPalavras(frase) {
		return "", errors.New("| A frase deve conter apenas palavras")
	}

	return frase, nil
}

func contemApenasPalavras(s string) bool {
	palavras := strings.Fields(s)
	for _, palavra := range palavras {
		for _, char := range palavra {
			if !unicode.IsLetter(char) {
				return false
			}
		}
	}
	return true
}

func obterNumeroColunas(reader *bufio.Reader) (int, error) {
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	input = strings.TrimSpace(input)

	colunas, err := strconv.Atoi(input)
	if err != nil {
		return 0, errors.New("| > Digite apenas o número de colunas")
	}

	return colunas, nil
}

func quebrarFrase(frase string, colunas int) string {
	palavras := strings.Fields(frase)
	var resultado []string

	linhaAtual := ""
	for _, palavra := range palavras {
		if len(linhaAtual)+len(palavra)+1 <= colunas {
			linhaAtual += palavra + " "
		} else {
			resultado = append(resultado, linhaAtual)
			linhaAtual = palavra + " "
		}
	}

	resultado = append(resultado, linhaAtual)

	return strings.Join(resultado, "\n")
}
