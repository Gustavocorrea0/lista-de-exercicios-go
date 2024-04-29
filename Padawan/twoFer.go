package main

import (
	"fmt"
	"strings"
)

func main() {
	var nomeUsuario string
	fmt.Println("Qual é o seu nome: ")
	fmt.Scanln(&nomeUsuario)

	nomeUsuario = strings.TrimSpace(nomeUsuario)
	if nomeUsuario == "" {
		fmt.Println("Um para você, um para mim")
	} else {
		fmt.Println("Um para ", nomeUsuario, " um para mim")
	}
}
