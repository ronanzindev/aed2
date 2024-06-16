package listaduplamenteencadeada

import (
	"fmt"
)

type No struct {
	element           string
	anterior, proximo *No
}

func (n *No) Element() string { return n.element }

type Lista struct {
	head, tail *No
	qtd        int
}

func NewLista() *Lista {
	l := &Lista{head: &No{element: "head", anterior: nil, proximo: nil}, qtd: 0, tail: &No{element: "tail"}}
	l.tail.anterior = l.head
	l.head.proximo = l.tail
	return l
}
func (l *Lista) Head() *No { return l.head }

func (l *Lista) InserirNoInicio(element string) {
	n := &No{element: element}
	n.anterior = l.head
	n.proximo = l.head.proximo
	l.head.proximo.anterior = n
	l.head.proximo = n
	// l.tail.anterior = n.proximo
	l.qtd++
}

func (l *Lista) InserirNoFinal(element string) {
	n := &No{element: element}
	n.anterior = l.tail.anterior
	n.proximo = l.tail
	l.tail.anterior.proximo = n
	l.tail.anterior = n
	l.qtd++
}
func (l *Lista) InserirApos(pos int, element string) {
	if pos < 0 {
		return
	}
	if pos > l.qtd {
		pos = l.qtd - 1
	}
	n := &No{element: element}
	if pos == 0 {
		n.proximo = l.head.proximo
		l.head.proximo = n
		l.qtd++
		return
	}
	count := 0
	aux := l.head.proximo
	for count < pos-1 {
		aux = aux.proximo
		count++
	}
	n.anterior = aux
	n.proximo = aux.proximo
	aux.proximo.anterior = n
	aux.proximo = n
	l.qtd++
}

func (l *Lista) Retirar(pos int) {
	if pos <= 0 {
		return
	}
	if pos > l.qtd {
		pos = l.qtd
	}
	aux := l.head.proximo
	count := 0
	for count < pos-1 {
		aux = aux.proximo
		count++
	}
	aux.anterior.proximo = aux.proximo
	l.qtd--
}
func (l *Lista) Buscar(pos int) string {
	if pos <= 0 {
		pos = 1
	}
	if pos > l.qtd {
		pos = l.qtd - 1
	}
	count := 0
	aux := l.head.proximo
	for count < pos-1 {
		aux = aux.proximo
		count++
	}
	return aux.element
}
func (l *Lista) GetFirst() *No {
	return l.head.proximo
}
func (l *Lista) GetLast() *No {
	return l.tail.anterior
}
func (l *Lista) Printar() {
	i := 1
	aux := l.head.proximo
	for aux != l.tail {
		fmt.Printf("%d -> %s\n", i, aux.element)
		aux = aux.proximo
		i++
	}
}
