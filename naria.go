package main

import (
	"container/list"
	"fmt"
)

type No struct {
	valor  int
	filhos []*No
}

func NovoNo(valor int) *No {
	return &No{
		valor:  valor,
		filhos: make([]*No, 0),
	}
}

func (n *No) AdicionarFilho(filho *No) {
	n.filhos = append(n.filhos, filho)
}

func VarreduraEmLargura(raiz *No) {
	fila := list.New()
	fila.PushBack(raiz)

	for fila.Len() > 0 {
		no := fila.Remove(fila.Front()).(*No)

		fmt.Println(no.valor)

		for _, filho := range no.filhos {
			fila.PushBack(filho)
		}
	}
}

func main() {
	raiz := NovoNo(1)
	raiz.AdicionarFilho(NovoNo(2))
	raiz.AdicionarFilho(NovoNo(3))
	raiz.filhos[0].AdicionarFilho(NovoNo(4))
	raiz.filhos[0].AdicionarFilho(NovoNo(5))
	raiz.filhos[1].AdicionarFilho(NovoNo(6))
	raiz.filhos[1].AdicionarFilho(NovoNo(7))

	VarreduraEmLargura(raiz)
}
