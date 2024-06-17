package arvores

import "fmt"

type No struct {
	elemento int
	left     *No
	right    *No
}

func (n *No) Add(elemento int) {
	if n == nil {
		return
	}
	if elemento <= n.elemento {
		if n.left == nil {
			n.left = &No{elemento: elemento}
		} else {
			n.left.Add(elemento)
		}
	} else {
		if n.right == nil {
			n.right = &No{elemento: elemento}
		} else {
			n.right.Add(elemento)
		}
	}

}
func (n *No) preOrder() {
	if n == nil {
		return
	}
	fmt.Printf("%d ", n.elemento)
	n.left.preOrder()
	n.right.preOrder()
}
func (n *No) inOrder() {
	if n == nil {
		return
	}
	n.left.inOrder()
	fmt.Printf("%d ", n.elemento)
	n.right.inOrder()
}

func (n *No) postOrder() {
	if n == nil {
		return
	}
	n.left.postOrder()
	n.right.postOrder()
	fmt.Printf("%d ", n.elemento)
}

type Arvore struct {
	raiz *No
}

func NovaArvore(elementoRaiz int) *Arvore {
	return &Arvore{
		raiz: &No{elemento: elementoRaiz},
	}
}

func (a *Arvore) Inserir(elemento int) {
	no := &No{elemento: elemento}
	if a.raiz == nil {
		a.raiz = no
	} else {
		a.raiz.Add(elemento)
	}
}

func (a *Arvore) PreOrder() {
	a.raiz.preOrder()
}

func (a *Arvore) InOrder() {
	a.raiz.inOrder()
}
func (a *Arvore) PostOrder() {
	a.raiz.postOrder()
}

func (n *No) findMin() *No {
	aux := n
	for aux != nil && aux.left != nil {
		aux = aux.left
	}
	return aux
}

func (n *No) deletar(elemento int) *No {
	if n == nil {
		return nil
	}
	// caso seja menor
	if elemento < n.elemento {
		n.left = n.left.deletar(elemento)
		// caso seja maior
	} else if elemento > n.elemento {
		n.right = n.right.deletar(elemento)
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
		n.elemento = minRight.elemento
		n.right = n.right.deletar(minRight.elemento)
	}
	return n
}

func (a *Arvore) Deletar(elemento int) {
	a.raiz = a.raiz.deletar(elemento)
}
