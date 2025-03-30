package tree

import "fmt"

type Node struct {
	element int
	left    *Node
	right   *Node
}

func (n *Node) Add(element int) {
	if n == nil {
		return
	}
	if element <= n.element {
		if n.left == nil {
			n.left = &Node{element: element}
		} else {
			n.left.Add(element)
		}
	} else {
		if n.right == nil {
			n.right = &Node{element: element}
		} else {
			n.right.Add(element)
		}
	}

}
func (n *Node) preOrder() {
	if n == nil {
		return
	}
	fmt.Printf("%d ", n.element)
	n.left.preOrder()
	n.right.preOrder()
}
func (n *Node) inOrder() {
	if n == nil {
		return
	}
	n.left.inOrder()
	fmt.Printf("%d ", n.element)
	n.right.inOrder()
}

func (n *Node) postOrder() {
	if n == nil {
		return
	}
	n.left.postOrder()
	n.right.postOrder()
	fmt.Printf("%d ", n.element)
}

type Tree struct {
	root *Node
}

func NewTree(rootElement int) *Tree {
	return &Tree{
		root: &Node{element: rootElement},
	}
}

func (a *Tree) Insert(elemento int) {
	no := &Node{element: elemento}
	if a.root == nil {
		a.root = no
	} else {
		a.root.Add(elemento)
	}
}

func (a *Tree) PreOrder() {
	a.root.preOrder()
}

func (a *Tree) InOrder() {
	a.root.inOrder()
}
func (a *Tree) PostOrder() {
	a.root.postOrder()
}

func (n *Node) findMin() *Node {
	aux := n
	for aux != nil && aux.left != nil {
		aux = aux.left
	}
	return aux
}

func (n *Node) delete(elemento int) *Node {
	if n == nil {
		return nil
	}
	// caso seja menor
	if elemento < n.element {
		n.left = n.left.delete(elemento)
		// caso seja maior
	} else if elemento > n.element {
		n.right = n.right.delete(elemento)
		// caso seja igual
	} else {
		// caso não tenha nem direita e nem esquerda, só deletar a filha
		if n.left == nil && n.right == nil {
			return nil
		}
		// caso tenha um filho na direita e nenhum na esquerda
		if n.left == nil {
			return n.right
			// caso tenha um filha na esquerda e nenhuma na direita
		} else if n.right == nil {
			return n.left
		}
		// caso tenha 2 filhos(esquerda e direita)
		minRight := n.right.findMin()
		n.element = minRight.element
		n.right = n.right.delete(minRight.element)
	}
	return n
}

func (a *Tree) Delete(elemento int) {
	a.root = a.root.delete(elemento)
}
