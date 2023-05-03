// Json unmarshal

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Pessoa struct {
	Nome   string `json:"nome"`
	Idade  int    `json:"idade"`
	Cidade string `json:"cidade"`
}

func main() {
	file, err := ioutil.ReadFile("pessoa.json")
	if err != nil {
		log.Fatal(err)
	}

	var p Pessoa
	err = json.Unmarshal(file, &p)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Nome:", p.Nome)
	fmt.Println("Idade:", p.Idade)
	fmt.Println("Cidade:", p.Cidade)
}
