//Em Go, panic, defer e recover são construções usadas para lidar com erros e exceções.
//
//panic é usado para interromper o fluxo normal do programa e sinalizar que ocorreu um erro crítico que não pode ser tratado. Quando uma chamada para panic é feita, a execução do programa é interrompida imediatamente e a execução retorna para a função que chamou a função que causou o pânico. O fluxo de controle é então passado de volta pela pilha de chamadas de função até que um mecanismo de recuperação seja ativado.
//
//defer é usado para adiar a execução de uma função até que a função que a contém retorne. As funções defer são executadas em ordem inversa, ou seja, a última função defer adicionada é executada primeiro e assim por diante. Isso é útil para garantir que os recursos alocados por uma função sejam sempre liberados, independentemente do resultado da função.
//
//recover é usado em conjunto com panic para recuperar o controle do fluxo de execução quando ocorre um pânico. recover é usado para recuperar o valor passado para panic e retornar o controle do fluxo de execução para a função que chamou a função que causou o pânico. O uso do recover permite que o programa lide com um pânico e continue a execução em vez de interrompê-lo completamente.
package main

import "fmt"

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()

	fmt.Println("Start")
	panic("Something went wrong")
	fmt.Println("End")
}


