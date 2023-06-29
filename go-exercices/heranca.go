package main

import "fmt"

// Neste exemplo, a estrutura Employee tem um campo Person embutido, que contém os campos Name e Age. A estrutura Employee também tem um campo Salary. Podemos acessar os campos da estrutura pai diretamente da estrutura filho usando a notação de ponto.
// Embora Go não tenha o conceito de herança de classe, a composição é uma maneira de utilizar dos "atributos" das outras structs
type Person struct {
	Name string
	Age  int
}

type Funcionario struct {
	Person
	Salary float64
}

func main() {
	emp := Funcionario{
		Person: Person{
			Name: "João",
			Age:  30,
		},
		Salary: 5000.0,
	}

	fmt.Printf("Nome: %s\n", emp.Name)
	fmt.Printf("Idade: %d\n", emp.Age)
	fmt.Printf("Salário: %f\n", emp.Salary)
}
