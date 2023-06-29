// json marchal

package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Nome  string `json:"nome"`
	Idade int    `json:"idade"`
}

func main() {
	p := Person{"Café Lucas", 18}

	// Convert struct to JSON
	jsonData, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(jsonData))
}
