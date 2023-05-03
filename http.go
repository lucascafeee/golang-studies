package main

import (
	"log"
	"net/http"
)

//O protocolo HTTP (Hypertext Transfer Protocol) é o protocolo de comunicação mais utilizado na internet.
//Ele é responsável por estabelecer as regras de comunicação entre clientes e servidores, permitindo a transferência de informações, como arquivos, páginas web, imagens, vídeos e áudios.
//O HTTP funciona a partir de uma arquitetura cliente-servidor, onde o cliente envia uma requisição ao servidor e este envia uma resposta ao cliente.
//As requisições e respostas HTTP contêm informações importantes, como os cabeçalhos (headers) e o corpo (body).

func main() {
	//Em Go, é possível implementar um servidor HTTP utilizando a biblioteca padrão "net/http".
	//Para criar um servidor HTTP, é necessário definir um manipulador (handler) para cada rota que o servidor irá responder.
	//O manipulador é uma função que recebe uma requisição HTTP e retorna uma resposta HTTP.

	// HTTP É UM PROTOCOLO DE COMUNICAÇÃO - BASE DE COMUNICAÇÃO DA WEB

	// Cliente (Faz requisição) - Servidor (Processa a requisição e envia a resposta)

	// Request - response

	//Rotas
	//URI - Identificador do recurso / Caminho
	//Método - GET, POST, PUT, DELETE

	http.HandleFunc("/home", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Olá mundo"))
	})

	http.HandleFunc("/users", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("rota para usuarios"))
	})

	log.Fatal(http.ListenAndServe(":5000", nil))
}
