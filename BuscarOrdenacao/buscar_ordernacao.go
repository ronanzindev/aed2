package buscarordenacao

import "aed/linkedList"

type Buscar struct {
	lista *linkedList.LinkedList
}

func NewBuscar() *Buscar {
	lista := linkedList.NewLinkedList()
	lista.InsertAtBeggining("c")
	lista.InsertAtBeggining("b")
	lista.InsertAtBeggining("a")
	return &Buscar{lista: lista}
}
func (b *Buscar) BuscaSequencial(element string) *linkedList.Node {
	aux := b.lista.Head()
	for aux != nil {
		if aux.Element() == element {
			return aux
		}
		aux = aux.Next()
	}
	return nil
}

func (b *Buscar) BuscaBinaria(vetor []int, alvo int) int {
	if len(vetor) == 0 {
		return -1
	}
	final := len(vetor) - 1
	inicio := 0
	for inicio <= final {
		meio := (final - inicio) / 2
		numeroAtual := vetor[meio]
		if numeroAtual == alvo {
			return numeroAtual
		}
		if numeroAtual > alvo {
			final = meio - 1
		}
		if numeroAtual < alvo {
			inicio = meio + 1
		}
	}
	return -1
}
