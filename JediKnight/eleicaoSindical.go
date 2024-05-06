package main

import (
	"fmt"
	"strings"
)

var candidatoA int = 0
var candidatoB int = 0
var candidatoC int = 0
var votosNulos int = 0
var votosBranco int = 0
var quantidadeDeVotos int = 0

func main() {
	menu()
}

func menu() {
	for {
		var opcao string
		fmt.Println("MENU DE VOTAÇÃO")
		fmt.Println("1 - VOTAR")
		fmt.Println("0 - FINALIZAR VOTACAO")
		_, err := fmt.Scanln(&opcao)

		if err != nil {
			fmt.Println("OPÇÃO INVÁLIDA")
		}

		switch opcao {
		case "1":
			votar()
		case "2":
			finalizarVotacao()
		default:
			fmt.Println("Opção inválida")
		}

	}
}

func votar() {
	var voto string
	for {
		fmt.Println("VOTAR")
		fmt.Println("99 - Candidato A")
		fmt.Println("10 - Candidato B")
		fmt.Println("37 - Candidato C")
		fmt.Println("B - Branco")
		fmt.Println("N - Nulo")
		_, err3 := fmt.Scanln(&voto)

		if err3 != nil {
			fmt.Println("Voto inválido")
			votar()
		}

		switch strings.ToUpper(voto) {
		case "90":
			votarEmCandidatoA()
		case "10":
			votarEmCandidatoB()
		case "37":
			votarEmCandidatoC()
		case "B":
			votarEmBranco()
		case "N":
			votarEmNulo()
		default:
			fmt.Println("Opção inválida")
			votar()
		}
	}
}

func votarEmCandidatoA() {
	var confirmar string
	for {
		fmt.Println("CONFIRMAR VOTO")
		fmt.Println("S para sim")
		fmt.Println("N para não")
		_, err4 := fmt.Scanln(&confirmar)

		if err4 != nil {
			fmt.Println("Opção inválida")
		}

		if strings.ToUpper(confirmar) == "S" {
			candidatoA++
			fmt.Println("Voto Realizado com SUCESSO")
			menu()
		}

		if strings.ToUpper(confirmar) == "N" {
			fmt.Println("Voto cancelado")
			votar()
		} else {
			fmt.Println("Opção inválida")
		}
	}
}

func votarEmCandidatoB() {
	var confirmar string
	for {
		fmt.Println("CONFIRMAR VOTO")
		fmt.Println("S para sim")
		fmt.Println("N para não")
		_, err4 := fmt.Scanln(&confirmar)

		if err4 != nil {
			fmt.Println("Opção inválida")
		}

		if strings.ToUpper(confirmar) == "S" {
			candidatoB++
			fmt.Println("Voto Realizado com SUCESSO")
			menu()
		}

		if strings.ToUpper(confirmar) == "N" {
			fmt.Println("Voto cancelado")
			votar()
		} else {
			fmt.Println("Opção inválida")
		}
	}
}

func votarEmCandidatoC() {
	var confirmar string
	for {
		fmt.Println("CONFIRMAR VOTO")
		fmt.Println("S para sim")
		fmt.Println("N para não")
		_, err4 := fmt.Scanln(&confirmar)

		if err4 != nil {
			fmt.Println("Opção inválida")
		}

		if strings.ToUpper(confirmar) == "S" {
			candidatoC++
			fmt.Println("Voto Realizado com SUCESSO")
			menu()
		}

		if strings.ToUpper(confirmar) == "N" {
			fmt.Println("Voto cancelado")
			votar()
		} else {
			fmt.Println("Opção inválida")
		}
	}
}

func votarEmBranco() {
	var confirmar string
	for {
		fmt.Println("CONFIRMAR VOTO")
		fmt.Println("S para sim")
		fmt.Println("N para não")
		_, err4 := fmt.Scanln(&confirmar)

		if err4 != nil {
			fmt.Println("Opção inválida")
		}

		if strings.ToUpper(confirmar) == "S" {
			votosBranco++
			fmt.Println("Voto Realizado com SUCESSO")
			menu()
		}

		if strings.ToUpper(confirmar) == "N" {
			fmt.Println("Voto cancelado")
			votar()
		} else {
			fmt.Println("Opção inválida")
		}
	}
}

func votarEmNulo() {
	var confirmar string
	for {
		fmt.Println("CONFIRMAR VOTO")
		fmt.Println("S para sim")
		fmt.Println("N para não")
		_, err4 := fmt.Scanln(&confirmar)

		if err4 != nil {
			fmt.Println("Opção inválida")
		}

		if strings.ToUpper(confirmar) == "S" {
			votosNulos++
			fmt.Println("Voto Realizado com SUCESSO")
			menu()
		}

		if strings.ToUpper(confirmar) == "N" {
			fmt.Println("Voto cancelado")
			votar()
		} else {
			fmt.Println("Opção inválida")
		}
	}
}

func finalizarVotacao() {
	var totalDeVotos int
	var percentualCandidatoA float32
	var percentualCandidatoB float32
	var percentualCandidatoC float32
	var percentualBrancos float32
	var percentualNulos float32

	totalDeVotos = candidatoA + candidatoB + candidatoC

	percentualCandidatoA = (float32(candidatoA) / float32(totalDeVotos)) * 100
	percentualCandidatoB = (float32(candidatoB) / float32(totalDeVotos)) * 100
	percentualCandidatoC = (float32(candidatoC) / float32(totalDeVotos)) * 100
	percentualBrancos = (float32(votosBranco) / float32(totalDeVotos)) * 100
	percentualNulos = (float32(votosNulos) / float32(totalDeVotos)) * 100

	fmt.Println("TOTAL DE VOTOS:", totalDeVotos)
	if candidatoA > candidatoB && candidatoA > candidatoC {

	}

	fmt.Println("> CANDIDATO A:", candidatoA, "| > PERCENTUAL: ", percentualCandidatoA, "%")
	fmt.Println("> CANDIDATO B:", candidatoB, "| > PERCENTUAL: ", percentualCandidatoB, "%")
	fmt.Println("> CANDIDATO C:", candidatoC, "| > PERCENTUAL: ", percentualCandidatoC, "%")
	fmt.Println("> BRANCOS:", votosBranco, "| > PERCENTUAL: ", percentualBrancos, "%")
	fmt.Println("> NULOS:", votosNulos, "| > PERCENTUAL: ", percentualNulos, "%")

}
