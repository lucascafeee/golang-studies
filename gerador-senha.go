package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const (
	tamanhoSenha = 15
	caracteres   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789@"
)

func gerarSenha(tamanho int) string {
	senha := make([]byte, tamanho)
	max := big.NewInt(int64(len(caracteres)))

	for i := range senha {
		indice, _ := rand.Int(rand.Reader, max)
		senha[i] = caracteres[indice.Int64()]
	}

	return string(senha)
}

func main() {
	senha := gerarSenha(tamanhoSenha)
	fmt.Println("Senha gerada:", senha)
}
