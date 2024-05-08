package main

import (
	"fmt"
	"os"
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
		fmt.Println("+--------------------------------------+")
		fmt.Println("|           MENU DE VOTAÇÃO            |")
		fmt.Println("+--------------------------------------+")
		fmt.Println("|              1 - VOTAR               |")
		fmt.Println("|          0 - FINALIZAR VOTACAO       |")
		fmt.Println("+--------------------------------------+")
		fmt.Println("| > Opção escolhida:")

		_, err := fmt.Scanln(&opcao)

		if err != nil {
			fmt.Println("+--------------------------------------+")
			fmt.Println("|             OPÇÃO INVÁLIDA           |")
		}

		switch opcao {
		case "1":
			votar()
		case "0":
			finalizarVotacao()
		default:
			fmt.Println("+--------------------------------------+")
			fmt.Println("|             OPÇÃO INVÁLIDA           |")
		}

	}
}

func votar() {
	var voto string
	for {
		fmt.Println("+--------------------------------------+")
		fmt.Println("|                 VOTAR                |")
		fmt.Println("+--------------------------------------+")
		fmt.Println("|           90 - Candidato A           |")
		fmt.Println("|           10 - Candidato B           |")
		fmt.Println("|           37 - Candidato C           |")
		fmt.Println("|           B - Branco                 |")
		fmt.Println("|           N - Nulo                   |")
		fmt.Print("| > OPÇÃO DESEJADA:")
		_, err3 := fmt.Scanln(&voto)

		if err3 != nil {
			fmt.Println("+--------------------------------------+")
			fmt.Println("|             VOTO INVÁLIDO            |")
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
			fmt.Println("+--------------------------------------+")
			fmt.Println("|             OPÇÃO INVÁLIDA           |")
			votar()
		}
	}
}

func votarEmCandidatoA() {
	var confirmar string
	for {
		fmt.Println("+--------------------------------------+")
		fmt.Println("|            CONFIRMAR VOTO            |")
		fmt.Println("+--------------------------------------+")
		fmt.Println("|              S - sim                 |")
		fmt.Println("|              N - não                 |")
		fmt.Println("| > OPÇÃO DESEJADA")
		_, err4 := fmt.Scanln(&confirmar)

		if err4 != nil {
			fmt.Println("+--------------------------------------+")
			fmt.Println("|             OPÇÃO INVÁLIDA           |")
		}

		if strings.ToUpper(confirmar) == "S" {
			candidatoA++
			fmt.Println("+--------------------------------------+")
			fmt.Println("|      Voto Realizado com SUCESSO      |")
			menu()
		}

		if strings.ToUpper(confirmar) == "N" {
			fmt.Println("+--------------------------------------+")
			fmt.Println("|             Voto CANCELADO           |")
			votar()
		} else {
			fmt.Println("+--------------------------------------+")
			fmt.Println("|             OPÇÃO INVÁLIDA           |")
		}
	}
}

func votarEmCandidatoB() {
	var confirmar string
	for {
		fmt.Println("+--------------------------------------+")
		fmt.Println("|            CONFIRMAR VOTO            |")
		fmt.Println("+--------------------------------------+")
		fmt.Println("|              S - sim                 |")
		fmt.Println("|              N - não                 |")
		fmt.Println("| > OPÇÃO DESEJADA")
		_, err4 := fmt.Scanln(&confirmar)

		if err4 != nil {
			fmt.Println("+--------------------------------------+")
			fmt.Println("|             OPÇÃO INVÁLIDA           |")
		}

		if strings.ToUpper(confirmar) == "S" {
			candidatoB++
			fmt.Println("+--------------------------------------+")
			fmt.Println("|      Voto Realizado com SUCESSO      |")
			menu()
		}

		if strings.ToUpper(confirmar) == "N" {
			fmt.Println("+--------------------------------------+")
			fmt.Println("|             Voto CANCELADO           |")
			votar()
		} else {
			fmt.Println("+--------------------------------------+")
			fmt.Println("|             OPÇÃO INVÁLIDA           |")
		}
	}
}

func votarEmCandidatoC() {
	var confirmar string
	for {
		fmt.Println("+--------------------------------------+")
		fmt.Println("|            CONFIRMAR VOTO            |")
		fmt.Println("+--------------------------------------+")
		fmt.Println("|              S - sim                 |")
		fmt.Println("|              N - não                 |")
		fmt.Println("| > OPÇÃO DESEJADA")
		_, err4 := fmt.Scanln(&confirmar)

		if err4 != nil {
			fmt.Println("+--------------------------------------+")
			fmt.Println("|             OPÇÃO INVÁLIDA           |")
		}

		if strings.ToUpper(confirmar) == "S" {
			candidatoC++
			fmt.Println("+--------------------------------------+")
			fmt.Println("|      Voto Realizado com SUCESSO      |")
			menu()
		}

		if strings.ToUpper(confirmar) == "N" {
			fmt.Println("+--------------------------------------+")
			fmt.Println("|             Voto CANCELADO           |")
			votar()
		} else {
			fmt.Println("+--------------------------------------+")
			fmt.Println("|             OPÇÃO INVÁLIDA           |")
		}
	}
}

func votarEmBranco() {
	var confirmar string
	for {
		fmt.Println("+--------------------------------------+")
		fmt.Println("|            CONFIRMAR VOTO            |")
		fmt.Println("+--------------------------------------+")
		fmt.Println("|              S - sim                 |")
		fmt.Println("|              N - não                 |")
		fmt.Println("| > OPÇÃO DESEJADA")
		_, err4 := fmt.Scanln(&confirmar)

		if err4 != nil {
			fmt.Println("+--------------------------------------+")
			fmt.Println("|             OPÇÃO INVÁLIDA           |")
		}

		if strings.ToUpper(confirmar) == "S" {
			votosBranco++
			fmt.Println("+--------------------------------------+")
			fmt.Println("|      Voto Realizado com SUCESSO      |")
			menu()
		}

		if strings.ToUpper(confirmar) == "N" {
			fmt.Println("+--------------------------------------+")
			fmt.Println("|             Voto CANCELADO           |")
			votar()
		} else {
			fmt.Println("+--------------------------------------+")
			fmt.Println("|             OPÇÃO INVÁLIDA           |")
		}
	}
}

func votarEmNulo() {
	var confirmar string
	for {
		fmt.Println("+--------------------------------------+")
		fmt.Println("|            CONFIRMAR VOTO            |")
		fmt.Println("+--------------------------------------+")
		fmt.Println("|              S - sim                 |")
		fmt.Println("|              N - não                 |")
		fmt.Println("| > OPÇÃO DESEJADA")
		_, err4 := fmt.Scanln(&confirmar)

		if err4 != nil {
			fmt.Println("+--------------------------------------+")
			fmt.Println("|             OPÇÃO INVÁLIDA           |")
		}

		if strings.ToUpper(confirmar) == "S" {
			votosNulos++
			fmt.Println("+--------------------------------------+")
			fmt.Println("|      Voto Realizado com SUCESSO      |")
			menu()
		}

		if strings.ToUpper(confirmar) == "N" {
			fmt.Println("+--------------------------------------+")
			fmt.Println("|             Voto CANCELADO           |")
			votar()
		} else {
			fmt.Println("+--------------------------------------+")
			fmt.Println("|             OPÇÃO INVÁLIDA           |")
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

	if totalDeVotos == 0 {
		fmt.Println("NÃO HOUVERAM VOTOS")
	} else {
		fmt.Println("+------------------------------------------------+")
		if candidatoA > candidatoB && candidatoA > candidatoC {
			fmt.Println("| > VENCEDOR: CANDIDATO A")
		} else if candidatoB > candidatoA && candidatoB > candidatoC {
			fmt.Println("| > VENCEDOR: CANDIDATO B")
		} else if candidatoC > candidatoA && candidatoC > candidatoB {
			fmt.Println("| > VENCEDOR: CANDIDATO C")
		}
		fmt.Println("+------------------------------------------------+")
		fmt.Println("| > TOTAL DE VOTOS:", totalDeVotos)
		fmt.Println("+------------------------------------------------+")
		fmt.Println("| > CANDIDATO A:", candidatoA, "| > PERCENTUAL: ", percentualCandidatoA, "%")
		fmt.Println("| > CANDIDATO B:", candidatoB, "| > PERCENTUAL: ", percentualCandidatoB, "%")
		fmt.Println("| > CANDIDATO C:", candidatoC, "| > PERCENTUAL: ", percentualCandidatoC, "%")
		fmt.Println("| > BRANCOS:", votosBranco, "| > PERCENTUAL: ", percentualBrancos, "%")
		fmt.Println("| > NULOS:", votosNulos, "| > PERCENTUAL: ", percentualNulos, "%")
		fmt.Println("+------------------------------------------------+")
	}

	os.Exit(0)
}
