package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

func main() {
	var str1, str2 string

	fmt.Println("+------------------------------------+")
	fmt.Println("| > Informe a primeira string:")
	str1 = readString()
	fmt.Println("+------------------------------------+")
	fmt.Println("| > Informe a segunda string:")
	str2 = readString()

	str1 = strings.ToLower(strings.ReplaceAll(str1, " ", ""))
	str2 = strings.ToLower(strings.ReplaceAll(str2, " ", ""))

	if len(str1) != len(str2) {
		fmt.Println("+-----------------------------------------------------------------------+")
		fmt.Println("|      As strings não são anagramas por possuir tamanhos diferentes!    |")
		fmt.Println("+-----------------------------------------------------------------------+")
		return
	}

	if str1 == str2 {
		fmt.Println("+------------------------------------------------------------+")
		fmt.Println("|      As strings não são anagramas porque são idênticas     |")
		fmt.Println("+------------------------------------------------------------+")
		return
	}

	runes1 := []rune(str1)
	runes2 := []rune(str2)

	sort.Slice(runes1, func(i, j int) bool { return runes1[i] < runes1[j] })
	sort.Slice(runes2, func(i, j int) bool { return runes2[i] < runes2[j] })

	if string(runes1) == string(runes2) {
		fmt.Println("+-----------------------------------+")
		fmt.Println("|     As strings são anagramas!     |")
		fmt.Println("+-----------------------------------+")
	} else {
		fmt.Println("+---------------------------------------+")
		fmt.Println("|     As strings não são anagramas      |")
		fmt.Println("+---------------------------------------+")
	}
}

func readString() string {
	var input string
	for {
		fmt.Scanln(&input)
		if isValidString(input) {
			break
		}
		fmt.Println("+------------------------------------------------------------+")
		fmt.Println("|    Por favor, digite apenas Letras!. Tente novamente:      |")
		fmt.Println("+------------------------------------------------------------+")
	}
	return input
}

func isValidString(input string) bool {
	for _, char := range input {
		if !unicode.IsLetter(char) {
			return false
		}
	}
	return true
}
