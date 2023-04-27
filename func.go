package main

import "fmt"

func main() {
	// Exemplo de função básica
	soma := somar(2, 3)
	fmt.Println("A soma é:", soma)

	// Exemplo de função que retorna mais de um valor
	res, resto := dividir(10, 3)
	fmt.Println("O resultado da divisão é:", res)
	fmt.Println("O resto da divisão é:", resto)

	// Exemplo de função com parâmetros variádicos
	nomes := []string{"João", "Maria", "Pedro", "Ana"}
	imprimirNomes(nomes...)

	// Exemplo de função anônima
	func() {
		fmt.Println("Esta é uma função anônima.")
	}()

	// Exemplo de função com closure
	funcaoClosure := retornaFuncao()
	fmt.Println(funcaoClosure())
}

// Função básica que soma dois números
func somar(a, b int) int {
	return a + b
}

// Função que divide dois números e retorna o resultado e o resto da divisão
func dividir(a, b int) (int, int) {
	return a / b, a % b
}

// Função com parâmetros variádicos que imprime uma lista de nomes
func imprimirNomes(nomes ...string) {
	fmt.Println("Lista de nomes:")
	for _, nome := range nomes {
		fmt.Println(nome)
	}
}

// Função que retorna outra função (closure)
func retornaFuncao() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
