package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var frase string
	scannerBufio := bufio.NewScanner(os.Stdin)
	fmt.Println("Qual é a frase: ")
	for scannerBufio.Scan() {
		frase = scannerBufio.Text()
		if len(frase) > 0 {
			break
		}
		fmt.Println("+--------------------------------------+")
		fmt.Println("| Entrada inválida. Digite novamente:")
	}

	if letraRepetida, ok := validarLetrasRepetidas2(frase); ok {
		fmt.Println("A letra:", strings.ToUpper(string(letraRepetida)), ",se repete")
	} else {
		fmt.Println("Nenhuma Letra se repete")
	}

}

func buscarLetrasRepetidas(fraseRecebida string) {
	fraseSemEspacos := strings.ReplaceAll(strings.ToLower(fraseRecebida), " ", "")

	for _, letra := range fraseSemEspacos {
		fmt.Println("Letra: ", string(letra))
	}
}

func validarLetrasRepetidas2(frase string) (rune, bool) {
	palavraOuFraseSemEspacos := strings.ReplaceAll(strings.ToLower(frase), " ", "")
	sliceLetras := make(map[rune]bool)

	for _, letra := range palavraOuFraseSemEspacos {
		if sliceLetras[letra] {
			return letra, true
		}
		sliceLetras[letra] = true
	}

	return 0, false
}
