package main

import "fmt"

func main() {
	var nota1 float32
	var nota2 float32
	var nota3 float32
	var nota4 float32

	for {
		fmt.Println("Qual é a nota do primeiro Bimestre: ")
		_, err := fmt.Scanln(&nota1)

		if err != nil || nota1 < 0 || nota1 > 10 {
			fmt.Println("Nota primeiro bimestre inválida")
		} else {
			break
		}
	}

	for {
		fmt.Println("Qual é a nota do segundo Bimestre: ")
		_, err1 := fmt.Scanln(&nota2)

		if err1 != nil || nota2 < 0 || nota2 > 10 {
			fmt.Println("Nota segundo bimestre inválida")
		} else {
			break
		}
	}

	for {
		fmt.Println("Qual é a nota do terceiro Bimestre")
		_, err2 := fmt.Scanln(&nota3)

		if err2 != nil || nota3 < 0 || nota3 > 10 {
			fmt.Println("Nota terceiro bimestre inválida")
		} else {
			break
		}
	}

	for {
		fmt.Println("Qual é a nota do quarto Bimestre: ")
		_, err3 := fmt.Scanln(&nota4)

		if err3 != nil || nota4 < 0 || nota4 > 10 {
			fmt.Println("Nota quarto bimestre inválida")
		} else {
			break
		}
	}

	calcularMedia(nota1, nota2, nota3, nota4)
}

func calcularMedia(notaUm float32, notaDois float32, notaTres float32, notaQuatro float32) {
	var somaDeNotas float32
	var media float32

	somaDeNotas = notaUm + notaDois + notaTres + notaQuatro
	media = somaDeNotas / 4

	if media >= 5 {
		fmt.Println("| > Situação: Aprovado")
		fmt.Println("| > Média:", media)
	} else {
		fmt.Println("| > Situação: Reprovado")
		fmt.Println("| > Média:", media)
	}
}
