package main

import "fmt"

func main() {
	var num int = 10
	var ptr *int

	// Atribui o endereço de num ao ponteiro ptr
	ptr = &num

	// Exibe o valor de num através do ponteiro
	fmt.Println("Valor de num:", *ptr)

	// Altera o valor de num através do ponteiro
	*ptr = 20
	fmt.Println("Novo valor de num:", num)

	// Passa um ponteiro como argumento para uma função
	changeValue(&num)
	fmt.Println("Novo valor de num após a chamada da função:", num)
}

// Altera o valor de uma variável através de um ponteiro
func changeValue(ptr *int) {
	*ptr = 30
}
