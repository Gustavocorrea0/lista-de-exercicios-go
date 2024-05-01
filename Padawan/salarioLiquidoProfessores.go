package main

import "fmt"

func main() {
	var valorDaHoraTrabalhada float32
	var horasTrabalhadas float32
	var descontoINSS float32

	fmt.Println("+-------------------------------------------------+")
	fmt.Println("|          Salario Liquido Professores            |")
	fmt.Println("+-------------------------------------------------+")
	fmt.Println("| > Qual a quantidade horas trabalhadas: ")
	_, err := fmt.Scanln(&horasTrabalhadas)

	if err != nil || horasTrabalhadas < 0 {
		fmt.Println("+-------------------------------------------------+")
		fmt.Println("|          Quantidade de Horas inválidas          |")
		fmt.Println("+-------------------------------------------------+")
	}

	fmt.Println("+-------------------------------------------------+")
	fmt.Println("| > Qual é o valor da hora trabalhada: ")
	_, err2 := fmt.Scanln(&valorDaHoraTrabalhada)

	if err2 != nil || valorDaHoraTrabalhada < 0 {
		fmt.Println("+-------------------------------------------------+")
		fmt.Println("|    Quantidade de horas trabalhadas inválidas    |")
		fmt.Println("+-------------------------------------------------+")
	}

	fmt.Println("+-------------------------------------------------+")
	fmt.Println("| > Qual é o percentual de desconto INSS: ")
	_, err3 := fmt.Scanln(&descontoINSS)

	if err3 != nil || descontoINSS <= 0 {
		fmt.Println("+-------------------------------------------------+")
		fmt.Println("|    Erro ao ler percentual de desconto INSS      |")
		fmt.Println("+-------------------------------------------------+")
	}

	calcularSalario(valorDaHoraTrabalhada, horasTrabalhadas, descontoINSS)
}

func calcularSalario(valorDaHora float32, quantidaddeHorasTrabalhadas float32, percentualDescontoINSS float32) {
	var salarioBruto = valorDaHora * quantidaddeHorasTrabalhadas
	var totalDescontos = (salarioBruto * percentualDescontoINSS) / 100
	var salárioLiquido = salarioBruto - totalDescontos

	fmt.Println("+-------------------------------------------------+")
	fmt.Println("| > O salario bruto é:", salarioBruto)
	fmt.Println("+-------------------------------------------------+")
	fmt.Println("| > O total de descontos é:", totalDescontos)
	fmt.Println("+-------------------------------------------------+")
	fmt.Println("| > O salario liquido é:", salárioLiquido)
	fmt.Println("+-------------------------------------------------+")
}
