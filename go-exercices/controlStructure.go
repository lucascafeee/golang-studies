package main

import "fmt"

func main() {
	// Estrutura de controle if/else
	idade := 20
	if idade < 18 {
		fmt.Println("Você é menor de idade.")
	} else if idade < 60 {
		fmt.Println("Você é adulto.")
	} else {
		fmt.Println("Você é idoso.")
	}

	// Estrutura de controle switch
	diaDaSemana := "sábado"
	switch diaDaSemana {
	case "segunda-feira":
		fmt.Println("Hoje é segunda-feira.")
	case "terça-feira":
		fmt.Println("Hoje é terça-feira.")
	case "quarta-feira":
		fmt.Println("Hoje é quarta-feira.")
	case "quinta-feira":
		fmt.Println("Hoje é quinta-feira.")
	case "sexta-feira":
		fmt.Println("Hoje é sexta-feira.")
	case "sábado":
		fmt.Println("Hoje é sábado.")
	case "domingo":
		fmt.Println("Hoje é domingo.")
	default:
		fmt.Println("Dia da semana inválido.")
	}

	// Estrutura de controle for
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// Estrutura de controle for com range
	numeros := []int{1, 2, 3, 4, 5}
	for indice, numero := range numeros {
		fmt.Printf("Índice: %d, Valor: %d\n", indice, numero)
	}

	// Estrutura de controle while (simulado com for)
	x := 1
	for x < 5 {
		fmt.Println(x)
		x++
	}

	// Estrutura de controle do-while (simulado com for)
	y := 1
	for {
		fmt.Println(y)
		y++
		if y > 5 {
			break
		}
	}
}
