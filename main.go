package main

import (
	"fmt"
	"github.com/lucascafeee/estudosGo/app"
	"log"
	"os"
)

func main() {
	fmt.Println("is running...")

	aplication := app.Gerar()
	if erro := aplication.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
