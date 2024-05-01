package main

import (
	"fmt"
	"strings"
)

func main() {
	var nomeUsuario string

	fmt.Println("+-------------------------------------------------+")
	fmt.Println("|                      TwoFer                     |")
	fmt.Println("+-------------------------------------------------+")
	fmt.Println("| > Qual é o seu nome: ")
	fmt.Scanln(&nomeUsuario)

	nomeUsuario = strings.TrimSpace(nomeUsuario)
	if nomeUsuario == "" {
		fmt.Println("+-------------------------------------------------+")
		fmt.Println("|            Um para você, um para mim            |")
		fmt.Println("+-------------------------------------------------+")
	} else {
		fmt.Println("+-------------------------------------------------+")
		fmt.Println("| Um para", nomeUsuario, "um para mim")
		fmt.Println("+-------------------------------------------------+")
	}
}
