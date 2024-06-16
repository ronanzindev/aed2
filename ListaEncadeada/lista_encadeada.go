package listaencadeada

import (
	"fmt"
	"strings"
)

type No struct {
	element string
	next    *No
}

func (n *No) Next() *No {
	return n.next
}
func (n *No) Element() string { return n.element }

type ListaEncadeada struct {
	head *No
	qtd  int
}

func (l *ListaEncadeada) Head() *No {
	return l.head
}

func NewListaEncadeada() *ListaEncadeada {
	return &ListaEncadeada{head: &No{element: "a", next: nil}}
}

func (l *ListaEncadeada) InserirNoInicio(element string) {
	no := &No{element: element}
	no.next = l.head.next
	l.head.next = no
	l.qtd++
}

func (l *ListaEncadeada) InserirNoFinal(element string) {
	aux := l.head
	for aux.next != nil {
		aux = aux.next
	}
	aux.next = &No{element: element}
	l.qtd++
}

// la ele
func (l *ListaEncadeada) InserirNaPos(pos int, element string) {
	n := &No{element: element}
	if pos < 0 {
		return
	}
	if pos > l.qtd {
		pos = l.qtd - 1
	}
	if pos == 0 {
		n.next = l.head
		l.head = n
		l.qtd++
		return
	}
	count := 1
	aux := l.head
	for count < pos {
		aux = aux.next
		count++
	}
	n.next = aux.next
	aux.next = n
	l.qtd++
}
func (l *ListaEncadeada) InserirOrdenado(element string) {
	n := &No{element: element}
	if strings.Compare(n.element, l.head.element) == -1 {
		n.next = l.head
		l.head = n
		l.qtd++
		return
	}
	anterior := l.head
	aux := l.head.next
	for strings.Compare(n.element, aux.element) == 1 {
		anterior = aux
		aux = aux.next
	}
	n.next = aux
	anterior.next = n
	l.qtd++
}

func (l *ListaEncadeada) DeletarNoInicio() {
	if l.head.next == nil {
		return
	}
	l.head = l.head.next
	l.qtd--
}

func (l *ListaEncadeada) DeletarNoFinal() {
	aux := l.head
	for aux.next.next != nil {
		aux = aux.next
	}
	aux.next = nil
	l.qtd--
}

func (l *ListaEncadeada) DeletarNaPos(pos int) {
	if pos < 0 {
		return
	}
	if pos > l.qtd {
		pos = l.qtd - 1
	}
	if pos == 0 {
		l.head = l.head.next
		l.qtd--
		return
	}
	count := 1
	anterior := l.head
	aux := l.head.next
	for count < pos {
		anterior = aux
		aux = aux.next
		count++
	}
	anterior.next = aux.next
	l.qtd--
}

func (l *ListaEncadeada) Buscar(pos int) *No {
	if pos < 0 {
		return nil
	}
	if pos > l.qtd {
		pos = l.qtd - 1
	}
	count := 0
	aux := l.head
	for count < pos {
		aux = aux.next
		count++
	}
	return aux
}

func (l *ListaEncadeada) Printar() {
	i := 0
	aux := l.head
	for aux != nil {
		fmt.Printf("%d -> %s\n", i, aux.element)
		aux = aux.next
		i++
	}
}
