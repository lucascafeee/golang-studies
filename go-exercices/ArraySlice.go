package main

import "fmt"

func main() {
	// Exemplo de array
	numeros := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array de números:", numeros)

	// Exemplo de slice
	cores := []string{"vermelho", "verde", "azul"}
	fmt.Println("Slice de cores:", cores)
	fmt.Println("Comprimento do slice de cores:", len(cores))
	fmt.Println("Capacidade do slice de cores:", cap(cores))

	// Adicionando um elemento ao slice
	cores = append(cores, "amarelo")
	//o append é útil quando você precisa adicionar elementos a uma fatia existente sem precisar criar uma nova fatia e copiar todos os elementos existentes para ela. Isso pode economizar tempo e memória, especialmente quando trabalhando com fatias grandes.
	fmt.Println("Slice de cores após adicionar um elemento:", cores)
	fmt.Println("Comprimento do slice de cores após adicionar um elemento:", len(cores))
	fmt.Println("Capacidade do slice de cores após adicionar um elemento:", cap(cores))
}

//Outra diferença importante entre arrays e slices é que um slice é uma referência a um array subjacente, o que significa que se você modificar um elemento do slice, você também estará modificando o array original. Por outro lado, modificar um elemento de um array não afetará outros arrays ou slices que fazem referência a ele.
//Em resumo, enquanto os arrays são usados para armazenar um conjunto fixo de elementos do mesmo tipo, os slices fornecem uma estrutura mais flexível e dinâmica para armazenar e manipular coleções de elementos em Go.
