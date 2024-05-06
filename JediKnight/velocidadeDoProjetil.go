package main

import "fmt"

func main() {
	var distaciaUsuario float32
	var tempoUsuario float32

	fmt.Println("Qual a distacia percorrida pela projetil:")
	_, err := fmt.Scanln(&distaciaUsuario)

	if err != nil || distaciaUsuario < 0 {
		fmt.Println("Distância Inválida")
		return
	}

	fmt.Println("Qual o tempo o projetil levou para chegar ao objetivo:")
	_, err1 := fmt.Scanln(&tempoUsuario)

	if err1 != nil || tempoUsuario < 0 {
		fmt.Println("Distância Inválida")
		return
	}

	calcularVelocidade(distaciaUsuario, tempoUsuario)
}

func calcularVelocidade(distacia float32, tempo float32) {
	var velocidade float32
	velocidade = (distacia * 1000) / (tempo * 60)
	fmt.Println("| A velocidade é:", velocidade, "M/S")
}
