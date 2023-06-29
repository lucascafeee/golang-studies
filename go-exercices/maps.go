package main

import "fmt"

func main() {
	// Criando um map vazio
	m := make(map[string]int)

	// Adicionando elementos ao map
	m["foo"] = 1
	m["bar"] = 2
	m["baz"] = 3

	// Imprimindo o map
	fmt.Println("Mapa original:", m)

	// Obtendo o valor de uma chave
	fooValue := m["foo"]
	fmt.Println("Valor de 'foo':", fooValue)

	// Verificando se uma chave existe no map
	_, barExists := m["bar"]
	_, quxExists := m["qux"]
	fmt.Println("'bar' existe?", barExists)
	fmt.Println("'qux' existe?", quxExists)

	// Removendo um elemento do map
	//A função delete é uma função nativa da linguagem Go que permite remover um elemento de um mapa.
	delete(m, "baz")
	fmt.Println("Mapa após remoção de 'baz':", m)
}
