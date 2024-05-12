package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var jogadaPlayer1 string
var jogadaPlayer2 string

func main() {
	capturarJogadaPlayer1()
}

func capturarJogadaPlayer1() {
	for {
		fmt.Println("+-----------------------------+")
		fmt.Println("|           JoKenPô           |")
		fmt.Println("+-----------------------------+")
		fmt.Println("|        Jogada Player 1      |")
		fmt.Println("+-----------------------------+")
		fmt.Println("|         P  - Pedra          |")
		fmt.Println("|         A  - Papel          |")
		fmt.Println("|         T  - Tesoura        |")
		fmt.Println("+-----------------------------+")
		fmt.Print("| > ")
		_, err := fmt.Scanln(&jogadaPlayer1)

		if err != nil {
			fmt.Println("+-----------------------------+")
			fmt.Println("|        Valor Invalido       |")
		}

		switch strings.ToUpper(jogadaPlayer1) {
		case "P", "A", "T":
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
		fmt.Println("|           JoKenPô           |")
		fmt.Println("+-----------------------------+")
		fmt.Println("|        Jogada Player 2      |")
		fmt.Println("+-----------------------------+")
		fmt.Println("|         P  - Pedra          |")
		fmt.Println("|         A  - Papel          |")
		fmt.Println("|         T  - Tesoura        |")
		fmt.Println("+-----------------------------+")
		fmt.Print("| > ")
		_, err := fmt.Scanln(&jogadaPlayer2)

		if err != nil {
			fmt.Println("+-----------------------------+")
			fmt.Println("|        Valor Invalido       |")
		}

		switch strings.ToUpper(jogadaPlayer2) {
		case "P", "A", "T":
			jogar(jogadaPlayer1, jogadaPlayer2)
			break
		default:
			fmt.Println("+-----------------------------+")
			fmt.Println("|        Valor Invalido       |")
		}
	}
}

func jogar(jogadaP1 string, jogadaP2 string) {
	limparTerminal()
	var j1 = strings.ToUpper(jogadaP1)
	var j2 = strings.ToUpper(jogadaP2)

	fmt.Println("+------------------------------+")
	fmt.Println("| > Player 1:", jogadaP1)
	fmt.Println("| > Player 2:", jogadaP2)
	fmt.Println("+------------------------------+")

	if j1 == j2 {
		fmt.Println("| > RESULTADO: EMPATE         |")
		os.Exit(0)
	}

	if (j1 == "P" && j2 == "T") ||
		(j1 == "T" && j2 == "A") ||
		(j1 == "A" && j2 == "P") {

		fmt.Println("| > RESULTADO: Player 1 venceu |")
		fmt.Println("+------------------------------+")
		os.Exit(0)
	}

	if (j2 == "P" && j1 == "T") ||
		(j2 == "T" && j1 == "A") ||
		(j2 == "A" && j1 == "P") {

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
