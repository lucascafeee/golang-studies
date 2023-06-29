package main

import "fmt"

func main() {
	var x int = 42
	var p *int = &x

	fmt.Println("Valor de x:", x)
	fmt.Println("Endereço de x:", &x)
	fmt.Println("Valor de p:", p)
	fmt.Println("Valor apontado por p:", *p)

	*p = 99
	fmt.Println("Novo valor de x:", x) // Output: Novo valor de x: 99
}

//Os ponteiros em Go são usados para uma variedade de tarefas, incluindo passagem de parâmetros por referência em funções, alocação dinâmica de memória e acesso a recursos de baixo nível, como dispositivos de hardware.
//É importante lembrar que, como ponteiros permitem acesso direto à memória, eles também podem causar problemas de segurança se usados incorretamente.

//Os ponteiros em Go permitem que você passe referências para valores em vez de cópias de valores para funções e métodos, economizando tempo e memória.
//
//Algumas vantagens do uso de ponteiros em Go incluem:
//
//Passar ponteiros para valores é mais eficiente em termos de memória do que passar cópias de valores, especialmente para valores grandes.
//Com o uso de ponteiros, você pode modificar diretamente um valor original sem fazer uma cópia, economizando tempo e memória.
//Ponteiros permitem que você manipule valores em um nível mais baixo, acessando diretamente a memória em vez de copiá-la ou movê-la.
//Algumas desvantagens do uso de ponteiros em Go incluem:
//
//O uso excessivo de ponteiros pode tornar o código menos legível e mais difícil de entender.
//Os ponteiros podem introduzir riscos de segurança, como acessar endereços inválidos de memória ou permitir que um invasor modifique dados arbitrários na memória.
//Os ponteiros podem ser mais difíceis de depurar em comparação com valores normais, pois você precisa estar ciente de onde os ponteiros estão apontando e se eles estão sendo modificados corretamente.