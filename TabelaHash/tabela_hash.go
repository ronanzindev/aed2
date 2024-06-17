package tabelahash

import "fmt"

type TabelaEnderecamentoAberto struct {
	chave int
	valor string
}
type TabelaHashEnderecamentoAberto struct {
	size  int
	items []*TabelaEnderecamentoAberto
}

func NewTabelaHash(size int) *TabelaHashEnderecamentoAberto {
	return &TabelaHashEnderecamentoAberto{
		size:  size,
		items: make([]*TabelaEnderecamentoAberto, size),
	}
}

func (t *TabelaHashEnderecamentoAberto) HahsChaveDivisao(chave int) int {
	// hash := 5381
	// for _, c := range chave {
	// 	hash = ((hash << 5) + hash) + int(c)
	// }
	// return (hash) % t.size
	return chave % t.size
}

func (t *TabelaHashEnderecamentoAberto) sondagemLinear(pos, i int) int {
	return (pos + i) % t.size
}

func (t *TabelaHashEnderecamentoAberto) Inserir(chave int, valor string) {
	// pos é a posicao no vetor
	pos := t.HahsChaveDivisao(chave)
	for i := 0; i < t.size; i++ {
		newPos := t.sondagemLinear(pos, i)
		if t.items[newPos] == nil {
			t.items[newPos] = &TabelaEnderecamentoAberto{chave: chave, valor: valor}
			return
		}
	}
}

func (t *TabelaHashEnderecamentoAberto) Buscar(chave int) string {
	pos := t.HahsChaveDivisao(chave)
	for i := 0; i < t.size; i++ {
		novaPos := t.sondagemLinear(pos, i)
		if t.items[novaPos] == nil {
			return ""
		}
		if t.items[novaPos].chave == chave {
			return t.items[novaPos].valor
		}
	}
	return ""
}
func (t *TabelaHashEnderecamentoAberto) Printar() {
	for i, v := range t.items {
		if v == nil {
			fmt.Printf("%d->nil\n", i)
		} else {
			fmt.Printf("%d->%s\n", i, v.valor)
		}
	}
}

// Encadeamento
type TabelaEncadeamento struct {
	chave   int
	valor   string
	proximo *TabelaEncadeamento
}

type TabelaHashEncadeamento struct {
	size  int
	items []*TabelaEncadeamento
}

func NewTabelaHashEncadeamento(size int) *TabelaHashEncadeamento {
	return &TabelaHashEncadeamento{
		size:  size,
		items: make([]*TabelaEncadeamento, size),
	}
}

func (t *TabelaHashEncadeamento) Hash(chave int) int {
	return chave % t.size
}

func (t *TabelaHashEncadeamento) Inserir(chave int, valor string) {
	pos := t.Hash(chave)
	novaEntrada := &TabelaEncadeamento{chave: chave, valor: valor, proximo: t.items[pos]}
	t.items[pos] = novaEntrada
}

func (t *TabelaHashEncadeamento) Buscar(chave int) string {
	pos := t.Hash(chave)
	current := t.items[pos]
	for current != nil {
		if current.chave == chave {
			return current.valor
		}
		current = current.proximo
	}
	return ""
}

func (t *TabelaHashEncadeamento) Printar() {
	for i, v := range t.items {
		fmt.Printf("%d -> ", i)
		current := v
		for current != nil {
			fmt.Printf("[Chave: %d, Valor: %s] -> ", current.chave, current.valor)
			current = current.proximo
		}
		fmt.Println("nil")
	}
}
