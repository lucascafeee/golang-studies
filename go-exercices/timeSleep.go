package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "Olá, Mundo!"
	}()

	select {
	case msg := <-ch:
		// Recebeu um valor do canal
		fmt.Println(msg)
	case <-time.After(1 * time.Second):
		// Timeout - nenhum valor recebido dentro do prazo
		fmt.Println("Timeout! Nenhum valor recebido.")
	}
}
