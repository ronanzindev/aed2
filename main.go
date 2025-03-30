package main

import (
	"aed/linkedList"
)

func main() {
	l := linkedList.NewLinkedList()
	l.InsertAtEnd("b")
	l.InsertAtEnd("c")
	l.InsertAtEnd("d")
	l.Print()
	// l = l.Reverse()
	// l.Print()
}
