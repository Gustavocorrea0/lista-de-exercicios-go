package main

import (
	"fmt"
	"strings"
	"unicode"
)

var listaAlgarismosRomanos = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

var valorDoNumero int = 0

func main() {
	receberAlgarismos()
}

func receberAlgarismos() {
	var algarismo string
	for {
		fmt.Println("+--------------------------------------+")
		fmt.Println("|           Algarismos Romanos         |")
		fmt.Println("+--------------------------------------+")
		fmt.Println("|                Valores               |")
		fmt.Println("+--------------------------------------+")
		fmt.Println("|      Símbolo      |       Valor      |")
		fmt.Println("|         I         |        1         |")
		fmt.Println("|         V         |        5         |")
		fmt.Println("|         X         |        10        |")
		fmt.Println("|         L         |        50        |")
		fmt.Println("|         C         |        100       |")
		fmt.Println("|         D         |        500       |")
		fmt.Println("|         M         |        1000      |")
		fmt.Println("+--------------------------------------+")
		fmt.Print("| > Qual é o algarismo romano: ")
		_, err := fmt.Scanln(&algarismo)

		if err != nil || !validarAlgarismo(algarismo) {
			fmt.Println("+--------------------------------------+")
			fmt.Println("|          Algarismos Inválidos        |")
		} else {
			break
		}
	}
	converterNumeroRomanoParaComum(strings.ToUpper(algarismo))
}

func converterNumeroRomanoParaComum(algarismos string) {
	for i := 0; i < len(algarismos); i++ {
		algarismoAtual := listaAlgarismosRomanos[string(algarismos[i])]
		proximoAlgarismo := 0
		if i+1 < len(algarismos) {
			proximoAlgarismo = listaAlgarismosRomanos[string(algarismos[i+1])]
		}
		if algarismoAtual < proximoAlgarismo {
			valorDoNumero += proximoAlgarismo - algarismoAtual
			i++
		} else {
			valorDoNumero += algarismoAtual
		}
	}
	fmt.Println("+--------------------------------------+")
	fmt.Println("| > O valor final é:", valorDoNumero)
	fmt.Println("+--------------------------------------+")
}

func validarAlgarismo(conjuntoLetra string) bool {
	for _, r := range conjuntoLetra {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
