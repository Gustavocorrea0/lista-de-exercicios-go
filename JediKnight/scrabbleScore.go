package main

import (
	"fmt"
	"os"
	"strings"
)

var AEIOULNRST int = 1
var DG int = 2
var BCMP int = 3
var FHVWY int = 4
var K int = 5
var JX int = 8
var QZ int = 10
var pontos int = 0

func main() {
	fmt.Println("+--------------------------------------+")
	fmt.Println("|           Scrabble Score             |")
	receberPalavra()
}

func receberPalavra() {
	var palavra string
	for {
		fmt.Println("+--------------------------------------+")
		fmt.Print("| > Qual é a palavra: ")
		_, err := fmt.Scanln(&palavra)

		if err != nil {
			fmt.Println("+--------------------------------------+")
			fmt.Println("|          Palavra Inválida            |")
		} else {
			break
		}
	}

	jogarScrablleScore(strings.ToUpper(palavra))
}

func jogarScrablleScore(palavra string) {
	fmt.Println("+--------------------------------------+")
	var letra string
	var pontosDaRodada int = 0
	for _, char := range palavra {
		letra = string(char)
		switch letra {
		case "A", "E", "I", "O", "U", "L", "N", "R", "S", "T":
			fmt.Println("| > Letra:", letra, "| > Score:", AEIOULNRST)
			pontosDaRodada = pontosDaRodada + AEIOULNRST

		case "D", "G":
			fmt.Println("| > Letra:", letra, "| > Score:", DG)
			pontosDaRodada = pontosDaRodada + DG

		case "B", "C", "M", "P":
			fmt.Println("| > Letra:", letra, "| > Score:", BCMP)
			pontosDaRodada = pontosDaRodada + BCMP

		case "F", "H", "V", "W", "Y":
			fmt.Println("| > Letra:", letra, "| > Score:", FHVWY)
			pontosDaRodada = pontosDaRodada + FHVWY

		case "K":
			fmt.Println("| > Letra:", letra, "| > Score:", K)
			pontosDaRodada = pontosDaRodada + K

		case "J", "X":
			fmt.Println("| > Letra:", letra, "| > Score:", JX)
			pontosDaRodada = pontosDaRodada + JX

		case "Q", "Z":
			fmt.Println("| > Letra:", letra, "| > Score:", QZ)
			pontosDaRodada = pontosDaRodada + QZ

		}
	}
	fmt.Println("+--------------------------------------+")
	fmt.Println("| > Score da jogada =", pontosDaRodada)
	fmt.Println("+--------------------------------------+")
	pontos = pontos + pontosDaRodada
	fmt.Println("| > Score =", pontos)

	jogarNovamente()
}

func jogarNovamente() {
	var opcao string
	for {
		fmt.Println("+--------------------------------------+")
		fmt.Println("|           Jogar Novamente?           |")
		fmt.Println("+--------------------------------------+")
		fmt.Println("|               S - sim                |")
		fmt.Println("|               N - não                |")
		fmt.Println("+--------------------------------------+")
		fmt.Print("| Opção desejada:")
		fmt.Scanln(&opcao)

		switch strings.ToUpper(opcao) {
		case "S":
			receberPalavra()
		case "N":
			fmt.Println("+--------------------------------------+")
			fmt.Println("| > Score Final =", pontos)
			fmt.Println("+--------------------------------------+")
			fmt.Println("|                 FIM                  |")
			fmt.Println("+--------------------------------------+")
			os.Exit(0)
		default:
			fmt.Println("+--------------------------------------+")
			fmt.Println("|            Opção Inválida            |")
		}
	}
}
