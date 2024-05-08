package main

import "fmt"

func main() {
	var distaciaUsuario float32
	var tempoUsuario float32
	fmt.Println("+------------------------------------------------------------+")
	fmt.Print("| > Qual a distâcia percorrida pela projetil:")
	_, err := fmt.Scanln(&distaciaUsuario)

	if err != nil || distaciaUsuario < 0 {
		fmt.Println("+------------------------------------------------------------+")
		fmt.Println("|                   Distância Inválida                       |")
		fmt.Println("+------------------------------------------------------------+")
		return
	}

	fmt.Println("+------------------------------------------------------------+")
	fmt.Print("| > Qual o tempo o projetil levou para chegar ao objetivo:")
	_, err1 := fmt.Scanln(&tempoUsuario)

	if err1 != nil || tempoUsuario < 0 {
		fmt.Println("+------------------------------------------------------------+")
		fmt.Println("|                     Tempo Inválido                         |")
		fmt.Println("+------------------------------------------------------------+")
		return
	}

	calcularVelocidade(distaciaUsuario, tempoUsuario)
}

func calcularVelocidade(distacia float32, tempo float32) {
	var velocidade float32
	velocidade = (distacia * 1000) / (tempo * 60)

	fmt.Println("+------------------------------------------------------------+")
	fmt.Println("| > A velocidade é:", velocidade, "M/S")
	fmt.Println("+------------------------------------------------------------+")
}
