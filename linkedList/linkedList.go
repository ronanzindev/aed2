package linkedList

import (
	"fmt"
	"strings"
)

type Node struct {
	element string
	next    *Node
}

func (n *Node) Next() *Node {
	return n.next
}
func (n *Node) Element() string { return n.element }

type LinkedList struct {
	head *Node
	qtd  int
}

func (l *LinkedList) Head() *Node {
	return l.head
}

func NewLinkedList() *LinkedList {
	return &LinkedList{head: &Node{element: "a", next: nil}}
}

func (l *LinkedList) InsertAtBeggining(element string) {
	no := &Node{element: element}
	no.next = l.head.next
	l.head.next = no
	l.qtd++
}

func (l *LinkedList) InsertAtEnd(element string) {
	aux := l.head
	for aux.next != nil {
		aux = aux.next
	}
	aux.next = &Node{element: element}
	l.qtd++
}

// la ele
func (l *LinkedList) InsertAtPosition(pos int, element string) {
	n := &Node{element: element}
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
func (l *LinkedList) InsertSorting(element string) {
	n := &Node{element: element}
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

func (l *LinkedList) DeleteAtBeggining() {
	if l.head.next == nil {
		return
	}
	l.head = l.head.next
	l.qtd--
}

func (l *LinkedList) DeleteAtEnd() {
	aux := l.head
	for aux.next.next != nil {
		aux = aux.next
	}
	aux.next = nil
	l.qtd--
}

func (l *LinkedList) DeleteAtPosition(pos int) {
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

func (l *LinkedList) Search(pos int) *Node {
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

func (l *LinkedList) Print() {
	i := 0
	aux := l.head
	for aux != nil {
		fmt.Printf("%d -> %s\n", i, aux.element)
		aux = aux.next
		i++
	}
}

func (l *LinkedList) Reverse() *LinkedList {
	newLinkedList := &LinkedList{}
	current := l.head
	if l == nil || current == nil {
		return newLinkedList
	}
	var next *Node
	var previous *Node
	for current != nil {
		next = current.next
		current.next = previous
		previous = current
		current = next
	}
	newLinkedList.head = previous
	return newLinkedList
}
