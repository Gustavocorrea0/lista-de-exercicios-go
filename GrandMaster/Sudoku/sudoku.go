package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

const (
	nR0 = 0
	nR1 = 1
	nR2 = 2
	nR3 = 3
	nR4 = 4
	nR5 = 5
	nR6 = 6
	nR7 = 7
	nR8 = 8
)

var tabuleiro = [][]int{
	{nR3, nR2, nR4, nR7, nR1, nR3, nR2, nR2, nR8},
	{nR0, nR1, nR2, nR3, nR4, nR1, nR3, nR4, nR2},
	{nR2, nR3, nR4, nR3, nR4, nR7, nR2, nR8, nR0},
	{nR0, nR3, nR2, nR8, nR5, nR3, nR6, nR3, nR4},
	{nR8, nR7, nR5, nR3, nR0, nR6, nR1, nR2, nR7},
	{nR4, nR1, nR4, nR8, nR2, nR8, nR8, nR5, nR3},
	{nR2, nR3, nR2, nR4, nR8, nR2, nR5, nR6, nR5},
	{nR8, nR4, nR2, nR0, nR7, nR1, nR4, nR3, nR2},
	{nR1, nR7, nR2, nR0, nR4, nR5, nR2, nR2, nR7},
}

var linha int
var coluna int
var numeroDoJogador int

func main() {
	fmt.Println("+----------------------------------+")
	fmt.Println("|              SUDOKU              |")
	apresentaTabuleiro()
}

func apresentaTabuleiro() {
	var possuiZero bool
	for _, linha := range tabuleiro {
		for _, valor := range linha {
			if valor == nR0 {
				possuiZero = true
				break
			}
		}
		if possuiZero {
			break
		}
	}

	if possuiZero {
		fmt.Println("+----------------------------------+")
		fmt.Println("|            Tabuleiro             |")
		fmt.Println("+----------------------------------+")
		for _, linha := range tabuleiro {
			fmt.Println(linha, " |")
		}
		lerLinha()
	} else {
		fmt.Println("+----------------------------------+")
		fmt.Println("|            Tabuleiro             |")
		fmt.Println("+----------------------------------+")
		for _, linha := range tabuleiro {
			fmt.Println(linha, " |")
		}
		fmt.Println("+----------------------------------+")
		fmt.Println("|           FIM DE JOGO            |")
		fmt.Println("+----------------------------------+")
	}
}

func lerLinha() {
	for {
		fmt.Println("+----------------------------------+")
		fmt.Print("| > Insira a linha (entre 1-9): ")
		_, err := fmt.Scanln(&linha)

		if (err != nil) || (linha > 9) || (linha < 1) {
			fmt.Println("+----------------------------------+")
			fmt.Println("|      VALOR DE LINHA INVÁLIDO     |")
		} else {
			lerColuna()
			break
		}
	}
}

func lerColuna() {
	for {
		fmt.Println("+----------------------------------+")
		fmt.Print("| > Insira a coluna (entre 1-9): ")
		_, err := fmt.Scanln(&coluna)

		if (err != nil) || (coluna > 9) || (coluna < 1) {
			fmt.Println("+----------------------------------+")
			fmt.Println("|      VALOR DE COLUNA INVÁLIDO    |")
		} else {
			numeroJogador()
			break
		}
	}
}

func numeroJogador() {
	for {
		fmt.Println("+----------------------------------+")
		fmt.Print("| > Insira um numero (entre 1-9): ")
		_, err := fmt.Scanln(&numeroDoJogador)

		if (err != nil) || (numeroDoJogador > 9) || (numeroDoJogador < 1) {
			fmt.Println("+----------------------------------+")
			fmt.Println("|          NUMERO INVÁLIDO         |")
		} else {
			jogarSudoku(linha, coluna, numeroDoJogador)
			break
		}
	}
}

func jogarSudoku(valorLinha int, valorColuna int, numeroJogador int) {
	valorColuna = valorColuna - 1
	valorLinha = valorLinha - 1

	if tabuleiro[valorLinha][valorColuna] != 0 {
		fmt.Println("+----------------------------------+")
		fmt.Println("|    POSICAO JÁ ESTÁ PREENCHIDA    |")
		apresentaTabuleiro()
	}

	tabuleiro[valorLinha][valorColuna] = numeroJogador
	limparTerminal()
	apresentaTabuleiro()
}

func limparTerminal() {
	switch runtime.GOOS {
	case "linux", "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Println("Sistema operacional não suportado.")
	}
}
