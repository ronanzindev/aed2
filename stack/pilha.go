package stack

import "aed/linkedList"

type Pilha struct {
	lista *linkedList.LinkedList
}

func NewPilha() *Pilha {
	return &Pilha{
		lista: linkedList.NewLinkedList(),
	}
}

func (p *Pilha) Push(element string) {
	p.lista.InsertAtPosition(1, element)
}

func (p *Pilha) Pop() string {
	No := p.lista.Search(1)
	p.lista.DeleteAtPosition(1)
	return No.Element()
}

func (p *Pilha) Printar() {
	p.lista.Print()
}
