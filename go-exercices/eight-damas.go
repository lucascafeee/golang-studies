package main

import "fmt"

const N = 8

var tabuleiro [N][N]int

// Verifica se é seguro colocar uma rainha na posição tabuleiro[linha][coluna]
func isSeguro(linha, coluna int) bool {
	for i := 0; i < coluna; i++ {
		if tabuleiro[linha][i] == 1 {
			return false
		}
	}

	for i, j := linha, coluna; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if tabuleiro[i][j] == 1 {
			return false
		}
	}

	for i, j := linha, coluna; i < N && j >= 0; i, j = i+1, j-1 {
		if tabuleiro[i][j] == 1 {
			return false
		}
	}

	return true
}

// Resolve o problema das oito rainhas usando backtracking
func resolverNQUtil(coluna int) bool {
	if coluna >= N {
		return true
	}

	for i := 0; i < N; i++ {
		if isSeguro(i, coluna) {
			tabuleiro[i][coluna] = 1
			imprimirTabuleiro()
			if resolverNQUtil(coluna + 1) {
				return true
			}
			tabuleiro[i][coluna] = 0
			imprimirTabuleiro()
		}
	}
	return false
}

// Imprime o tabuleiro
func imprimirTabuleiro() {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Printf("%d ", tabuleiro[i][j])
		}
		fmt.Println()
	}
	fmt.Println("-------")
}

// Resolve o problema das oito rainhas e imprime a solução
func resolverNQ() {
	if !resolverNQUtil(0) {
		fmt.Println("Solução não encontrada")
		return
	}
	fmt.Println("Solução encontrada:")
	imprimirTabuleiro()
}

func main() {
	resolverNQ()
}
