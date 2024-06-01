package queue

import listaencadeada "aed/ListaEncadeada"

type Fila struct {
	lista *listaencadeada.ListaEncadeada
}

func NewFila() *Fila {
	return &Fila{
		lista: listaencadeada.NewListaEncadeada(),
	}
}

func (f *Fila) Push(e string) {
	f.lista.InserirNoFinal(e)
}

func (f *Fila) Pop() string {
	n := f.lista.Buscar(1)
	f.lista.DeletarNaPos(1)
	if n == nil {
		return ""
	}
	return n.Element()
}

func (f *Fila) Printar() {
	f.lista.Printar()
}
