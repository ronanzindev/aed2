package stack

import listaencadeada "aed/ListaEncadeada"

type Pilha struct {
	lista *listaencadeada.ListaEncadeada
}

func NewPilha() *Pilha {
	return &Pilha{
		lista: listaencadeada.NewListaEncadeada(),
	}
}

func (p *Pilha) Push(element string) {
	p.lista.InserirNaPos(1, element)
}

func (p *Pilha) Pop() string {
	No := p.lista.Buscar(1)
	p.lista.DeletarNaPos(1)
	return No.Element()
}

func (p *Pilha) Printar() {
	p.lista.Printar()
}
