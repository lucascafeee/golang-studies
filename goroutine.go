package main

import (
	"fmt"
	"time"
)

func main() {
	// Criando uma goroutine para executar a função count até 5
	go count("goroutine", 5)

	// Executando a função count na main thread até 3
	count("main thread", 3)
}

func count(name string, n int) {
	for i := 1; i <= n; i++ {
		fmt.Println(name, ":", i)
		time.Sleep(time.Millisecond * 500) // Atrasando a execução da função em 500ms
	}
}

//Em Go, a concorrência é a capacidade de executar várias tarefas independentes ao mesmo tempo, sem bloquear ou interromper a execução de outras tarefas. Isso é alcançado com as goroutines, que são unidades leves de threads executadas em uma única thread do sistema operacional. As goroutines são gerenciadas pelo runtime do Go, que é responsável por escaloná-las em diferentes threads da CPU para maximizar o uso dos recursos disponíveis.
//A concorrência em Go é implementada usando canais (channels) para coordenar a comunicação entre as goroutines. Canais são objetos que permitem que uma goroutine envie e receba valores de outra goroutine de forma segura e sincronizada. Isso evita condições de corrida e outros problemas que podem surgir quando várias goroutines acessam e modificam o mesmo estado compartilhado.
