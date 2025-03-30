package queue

import "aed/linkedList"

type Fila struct {
	lista *linkedList.LinkedList
}

func NewFila() *Fila {
	return &Fila{
		lista: linkedList.NewLinkedList(),
	}
}

func (f *Fila) Push(e string) {
	f.lista.InsertAtEnd(e)
}

func (f *Fila) Pop() string {
	n := f.lista.Search(1)
	f.lista.DeleteAtPosition(1)
	if n == nil {
		return ""
	}
	return n.Element()
}

func (f *Fila) Printar() {
	f.lista.Print()
}
