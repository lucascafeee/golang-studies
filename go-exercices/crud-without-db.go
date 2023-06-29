package main

// testar o crud por meio das rotas no postman!

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Funcionario struct {
	ID        int    `json:"id"`
	Nome      string `json:"nome"`
	Sobrenome string `json:"sobrenome"`
	Email     string `json:"email"`
}

var funcionarios = make(map[int]Funcionario)
var contadorID = 0

func criarFuncionario(w http.ResponseWriter, r *http.Request) {
	var funcionario Funcionario
	json.NewDecoder(r.Body).Decode(&funcionario)

	contadorID++
	funcionario.ID = contadorID

	funcionarios[funcionario.ID] = funcionario

	json.NewEncoder(w).Encode(funcionario)
}

func lerFuncionarioEspecifico(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := parseInt(params["id"])

	funcionario, ok := funcionarios[id]
	if !ok {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(funcionario)
}

func atualizarFuncionario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := parseInt(params["id"])

	_, ok := funcionarios[id]
	if !ok {
		http.NotFound(w, r)
		return
	}

	var funcionarioAtualizado Funcionario
	json.NewDecoder(r.Body).Decode(&funcionarioAtualizado)

	funcionarioAtualizado.ID = id
	funcionarios[id] = funcionarioAtualizado

	json.NewEncoder(w).Encode(funcionarioAtualizado)
}

func deletarFuncionario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := parseInt(params["id"])

	_, ok := funcionarios[id]
	if !ok {
		http.NotFound(w, r)
		return
	}

	delete(funcionarios, id)

	fmt.Fprint(w, "Funcionário deletado")
}

func listarFuncionarios(w http.ResponseWriter, r *http.Request) {
	listaFuncionarios := []Funcionario{}
	for _, funcionario := range funcionarios {
		listaFuncionarios = append(listaFuncionarios, funcionario)
	}
	json.NewEncoder(w).Encode(listaFuncionarios)
}

func parseInt(s string) int {
	var n int
	fmt.Sscanf(s, "%d", &n)
	return n
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/funcionarios", criarFuncionario).Methods("POST")
	r.HandleFunc("/funcionarios/{id}", lerFuncionarioEspecifico).Methods("GET")
	r.HandleFunc("/funcionarios/{id}", atualizarFuncionario).Methods("PUT")
	r.HandleFunc("/funcionarios/{id}", deletarFuncionario).Methods("DELETE")
	r.HandleFunc("/funcionarios", listarFuncionarios).Methods("GET")
	fmt.Println("Servidor rodando na porta 8080")
	http.ListenAndServe(":8080", r)
}
