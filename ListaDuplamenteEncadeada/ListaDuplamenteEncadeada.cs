namespace ListaDuplamenteEncadeada;
class No {
  private string element;
  private No anterior,proximo;
  public No(string e) {
    element = e;
    anterior = null;
    proximo = null;
  }
  public No(string e, No prox) {
    element = e;
    proximo = prox;
  }
  public No(No ant, string e ) {
    element = e;
    anterior = ant;
  }
  public void setElement(string e) => element = e;
  public void setAnterior(No n) => anterior = n;
  public  void setProximo(No n) => proximo = n;
  public string getElement() => element;
  public No getProximo() => proximo;
  public No getAnterior() =>anterior;
} 
class Lista {
  private No head, tail;
  private int qtd;
  public Lista() {
    head = new("head");
    head.setAnterior(null);
    tail = new("tail");
    head.setProximo(tail);
    tail.setAnterior(head);
    tail.setProximo(null);
    qtd = 0;
  }
  public void inserirNoInicio(No n) {
    n.setAnterior(head);
    n.setProximo(head.getProximo());
    head.getProximo().setAnterior(n);
    head.setProximo(n);
    qtd++;
  }
  public void inserirNoFinal(No n) {
    n.setAnterior(tail.getAnterior());
    n.setProximo(tail);
    tail.getAnterior().setProximo(n);
    tail.setAnterior(n);
    qtd++;
  }
  public void inserirNaPos(int pos, No n) {
    if(pos < 0)  return;
    if(pos > qtd) pos = qtd -1;
    if(pos == 0) {
      n.setProximo(head.getProximo());
      head.setProximo(n);
      qtd++;
      return;
    }
    No aux = head;
    for(int i =0 ; i < pos-1; i++) {
      aux = aux.getProximo();
    } 
    n.setAnterior(aux);
    n.setProximo(aux.getProximo());
    aux.getProximo().setAnterior(n);
    aux.setProximo(n);
    qtd++;
  }
  public void retirar(int pos) {
    if(pos < 0) return;
    if(pos > qtd) pos = qtd -1;
    No aux = head.getProximo();
    for(int i = 0; i < pos -1; i++) {
      aux = aux.getProximo();
    }
    aux.getAnterior().setProximo(aux.getProximo());
    qtd--;
  }
  public string buscar(int pos) {
    if(pos <= 0) pos = 1;
    if(pos > qtd) pos = qtd -1;
    No aux = head.getProximo();
    for(int i = 0; i < pos -1; i++) aux = aux.getProximo();
    return aux.getElement();
  }
  public No getFirst() => head.getProximo();
  public No getLast() => tail.getAnterior();
  public void printar() {
    int i = 0;
    No aux = head;
    while (aux != null)
    {
      Console.WriteLine($"{i} -> {aux.getElement()}");
      aux = aux.getProximo();
      i++;
    }
  }

}