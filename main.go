package main

import (
	listaencadeada "aed/ListaEncadeada"
	"fmt"
)

func main() {
	l := listaencadeada.NewListaEncadeada()
	l.InserirNoInicio("c")
	l.InserirOrdenado("b")
	l.Printar()
	b()
	l.DeletarNaPos(1)
	l.Printar()
}

func b() {
	fmt.Println("------------------------------------")
}
