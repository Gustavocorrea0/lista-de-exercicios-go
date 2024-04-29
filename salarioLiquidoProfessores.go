package main

import "fmt"

func main() {
	var valorDaHoraTrabalhada float64
	var horasTrabalhadas float64
	var descontoINSS float64

	fmt.Println("Qual a quantidade horas trabalhadas: ")
	_, err := fmt.Scanln(&horasTrabalhadas)

	if err != nil {
		fmt.Println("Erro ao ler quantidades de horas")
	}

	fmt.Println("Qual é o valor da hora trabalhada: ")
	_, err2 := fmt.Scanln(&valorDaHoraTrabalhada)

	if err2 != nil {
		fmt.Println("Erro ao ler valor da hora trabalhada")
	}

	fmt.Println("Qual é o percentual de desconto INSS: ")
	_, err3 := fmt.Scanln(&descontoINSS)

	if err3 != nil {
		fmt.Println("Erro ao ler percentual de desconto INSS:")
	}

	calcularSalario(valorDaHoraTrabalhada, horasTrabalhadas, descontoINSS)
}

func calcularSalario(valorDaHora float64, quantidaddeHorasTrabalhadas float64, percentualDescontoINSS float64) {
	var salarioBruto = valorDaHora * quantidaddeHorasTrabalhadas
	var totalDescontos = (salarioBruto * percentualDescontoINSS) / 100
	var salárioLiquido = salarioBruto - totalDescontos

	fmt.Println("+----------------------------------------------------+")
	fmt.Println("| > O salario bruto é:", salarioBruto)
	fmt.Println("+----------------------------------------------------+")
	fmt.Println("| > O total de descontos é:", totalDescontos)
	fmt.Println("+----------------------------------------------------+")
	fmt.Println("| > O salario liquido é:", salárioLiquido)
}
