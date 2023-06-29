package main

import "fmt"

// definindo a struct Pessoa
type Pessoa struct { // Aqui coloca os atributos, é como se fosse a "Class" do golang
	nome  string
	idade int
}

func main() {
	// criando uma nova instância da struct Pessoa
	pessoa1 := Pessoa{
		nome:  "João",
		idade: 30,
	}

	// imprimindo os valores dos campos da struct
	fmt.Println("Nome:", pessoa1.nome)
	fmt.Println("Idade:", pessoa1.idade)

	// modificando o valor do campo "idade" da struct
	pessoa1.idade = 31

	// imprimindo o valor modificado
	fmt.Println("Nova idade:", pessoa1.idade)
}
