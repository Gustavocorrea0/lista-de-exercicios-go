package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var jogadaPlayer1s string
var jogadaPlayer2s string

func main() {
	fmt.Println("+-----------------------------------+")
	fmt.Println("| Pedra Papel Tesoura Lagarto Spock |")
	fmt.Println("+-----------------------------------+")
	capturarJogadaPlayer1()
}

func capturarJogadaPlayer1() {
	for {
		fmt.Println("|        Jogada Player 1      |")
		fmt.Println("+-----------------------------+")
		fmt.Println("|         P  - Pedra          |")
		fmt.Println("|         A  - Papel          |")
		fmt.Println("|         T  - Tesoura        |")
		fmt.Println("|         L  - Lagarto        |")
		fmt.Println("|         S  - Spock          |")
		fmt.Println("+-----------------------------+")
		fmt.Print("| >  Qual a escolha: ")
		_, err := fmt.Scanln(&jogadaPlayer1s)

		if err != nil {
			fmt.Println("+-----------------------------+")
			fmt.Println("|        Valor Invalido       |")
		}

		switch strings.ToUpper(jogadaPlayer1s) {
		case "P", "A", "T", "L", "S":
			capturarJogadaPlayer2()
			break
		default:
			fmt.Println("+-----------------------------+")
			fmt.Println("|        Valor Invalido       |")
		}
	}
}

func capturarJogadaPlayer2() {
	limparTerminal()
	for {
		fmt.Println("+-----------------------------+")
		fmt.Println("|        Jogada Player 2      |")
		fmt.Println("+-----------------------------+")
		fmt.Println("|         P  - Pedra          |")
		fmt.Println("|         A  - Papel          |")
		fmt.Println("|         T  - Tesoura        |")
		fmt.Println("|         L  - Lagarto        |")
		fmt.Println("|         S  - Spock          |")
		fmt.Println("+-----------------------------+")
		fmt.Print("| >  Qual a escolha: ")
		_, err1 := fmt.Scanln(&jogadaPlayer2s)

		if err1 != nil {
			fmt.Println("+-----------------------------+")
			fmt.Println("|        Valor Invalido       |")
		}

		switch strings.ToUpper(jogadaPlayer2s) {
		case "P", "A", "T", "L", "S":
			jogar1(jogadaPlayer1s, jogadaPlayer2s)
		default:
			fmt.Println("+-----------------------------+")
			fmt.Println("|        Valor Invalido       |")
		}
	}
}

func jogar1(jogadaP1 string, jogadaP2 string) {
	limparTerminal()
	var escolhaJogador1 = strings.ToUpper(jogadaP1)
	var escolhaJogador2 = strings.ToUpper(jogadaP2)

	fmt.Println("+------------------------------+")
	fmt.Println("| > Player 1:", escolhaJogador1)
	fmt.Println("| > Player 2:", escolhaJogador2)
	fmt.Println("+------------------------------+")

	if (escolhaJogador1 == "T" && escolhaJogador2 == "A") ||
		(escolhaJogador1 == "A" && escolhaJogador2 == "P") ||
		(escolhaJogador1 == "P" && escolhaJogador2 == "L") ||
		(escolhaJogador1 == "L" && escolhaJogador2 == "S") ||
		(escolhaJogador1 == "S" && escolhaJogador2 == "T") ||
		(escolhaJogador1 == "T" && escolhaJogador2 == "L") ||
		(escolhaJogador1 == "L" && escolhaJogador2 == "A") ||
		(escolhaJogador1 == "A" && escolhaJogador2 == "S") ||
		(escolhaJogador1 == "S" && escolhaJogador2 == "P") ||
		(escolhaJogador1 == "P" && escolhaJogador2 == "T") {
		fmt.Println("| > RESULTADO: Player 1 venceu |")
		fmt.Println("+------------------------------+")
		os.Exit(0)
	}

	if (escolhaJogador2 == "T" && escolhaJogador1 == "A") ||
		(escolhaJogador2 == "A" && escolhaJogador1 == "P") ||
		(escolhaJogador2 == "P" && escolhaJogador1 == "L") ||
		(escolhaJogador2 == "L" && escolhaJogador1 == "S") ||
		(escolhaJogador2 == "S" && escolhaJogador1 == "T") ||
		(escolhaJogador2 == "T" && escolhaJogador1 == "L") ||
		(escolhaJogador2 == "L" && escolhaJogador1 == "A") ||
		(escolhaJogador2 == "A" && escolhaJogador1 == "S") ||
		(escolhaJogador2 == "S" && escolhaJogador1 == "P") ||
		(escolhaJogador2 == "P" && escolhaJogador1 == "T") {
		fmt.Println("| > RESULTADO: Player 2 venceu |")
		fmt.Println("+------------------------------+")
		os.Exit(0)
	}
}

func limparTerminal() {
	// Verifica qual sistema operacional está sendo utilizado
	switch runtime.GOOS {
	case "linux", "darwin": // Linux ou macOS
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows": // Windows
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Println("Sistema operacional não suportado.")
	}
}
