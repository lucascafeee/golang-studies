package main

import (
	"fmt"
)

type ArvoreBinaria struct {
	arvore []int
}

type Fila []int

func (f *Fila) Enfileirar(valor int) {
	*f = append(*f, valor)
}

func (f *Fila) Desenfileirar() int {
	valor := (*f)[0]
	*f = (*f)[1:]
	return valor
}

func (f *Fila) Vazia() bool {
	return len(*f) == 0
}

type Pilha []int

func (p *Pilha) Empilhar(valor int) {
	*p = append(*p, valor)
}

func (p *Pilha) Desempilhar() int {
	valor := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return valor
}

func (p *Pilha) Vazia() bool {
	return len(*p) == 0
}

func (t *ArvoreBinaria) Inserir(valor int) {
	t.arvore = append(t.arvore, valor)
}

func (t *ArvoreBinaria) indiceFilhoEsquerdo(indice int) int {
	return 2*indice + 1
}

func (t *ArvoreBinaria) indiceFilhoDireito(indice int) int {
	return 2*indice + 2
}

func (t *ArvoreBinaria) buscaEmLargura() {
	fila := Fila{}
	fila.Enfileirar(0)

	for !fila.Vazia() {
		indice := fila.Desenfileirar()
		if indice >= len(t.arvore) {
			continue
		}

		fmt.Print(t.arvore[indice], " ")

		fila.Enfileirar(t.indiceFilhoEsquerdo(indice))
		fila.Enfileirar(t.indiceFilhoDireito(indice))
	}
}

func (t *ArvoreBinaria) buscaEmProfundidade() {
	pilha := Pilha{}
	pilha.Empilhar(0)

	for !pilha.Vazia() {
		indice := pilha.Desempilhar()
		if indice >= len(t.arvore) {
			continue
		}

		fmt.Print(t.arvore[indice], " ")

		pilha.Empilhar(t.indiceFilhoDireito(indice))
		pilha.Empilhar(t.indiceFilhoEsquerdo(indice))
	}
}

func (t *ArvoreBinaria) preOrdem(indice int) {
	if indice >= len(t.arvore) {
		return
	}

	fmt.Print(t.arvore[indice], " ")
	t.preOrdem(t.indiceFilhoEsquerdo(indice))
	t.preOrdem(t.indiceFilhoDireito(indice))
}

func (t *ArvoreBinaria) emOrdem(indice int) {
	if indice >= len(t.arvore) {
		return
	}

	t.emOrdem(t.indiceFilhoEsquerdo(indice))
	fmt.Print(t.arvore[indice], " ")
	t.emOrdem(t.indiceFilhoDireito(indice))
}

func (t *ArvoreBinaria) posOrdem(indice int) {
	if indice >= len(t.arvore) {
		return
	}

	t.posOrdem(t.indiceFilhoEsquerdo(indice))
	t.posOrdem(t.indiceFilhoDireito(indice))
	fmt.Print(t.arvore[indice], " ")
}

func main() {
	arvore := ArvoreBinaria{}
	valores := []int{10, 5, 15, 2, 8, 12, 18}

	for _, valor := range valores {
		arvore.Inserir(valor)
	}

	fmt.Println("Varredura em Largura:")
	arvore.buscaEmLargura()
	fmt.Println()

	fmt.Println("Varredura em Profundidade:")
	arvore.buscaEmProfundidade()
	fmt.Println()

	fmt.Println("Pré-ordem:")
	arvore.preOrdem(0)
	fmt.Println()

	fmt.Println("Em-ordem:")
	arvore.emOrdem(0)
	fmt.Println()

	fmt.Println("Pós-ordem:")
	arvore.posOrdem(0)
	fmt.Println()
}
