1)d

2)c

3)b 

4)c

5)a) Não pois são privados. Só podem ser acessados diretamento dentro de sua classe

5)b) Não pois o valor do atributo posX é um double.  Poderia adicionamos um .toString() ai final do return

5)c) Não ira retornar error, ja que é permitido multiplos construtore C# tendo apenas suas assinaturas diferentes.

5)d) Irá retornar o nome do corpo em minusculo.

6)
```c#
class No {
   private string element;
  private No next;
  public No(string e) {
    element = e;
    next = null;
  }
  public No() {}
  public string getElement() => element;
  public No getNext() => next;
  public void setElement(string e) => element = e;
  public void setNext(No n) => next = n;
}
```

7)
```c#
class ListaEncadeada {
  private No head;
  private int qtd;
  public ListaEncadeada() {
    head = new No();
  }
  public void inserirNaPos(int pos, No n) {
    if(pos < 0) return;
    if(pos > qtd) pos = qtd -1;
    if(pos == 0) {
      n.setNext(head.getNext());
      head.setNext(n);
      qtd++;
      return;
    }
    No aux = head;
    for(int i = 0; i < pos - 1; i++) {
      aux = aux.getNext();
    } 
    n.setNext(aux.getNext());
    aux.setNext(n);
    qtd++;
  }
  public void excluirNaPos(int pos) {
    if(pos < 0) return;
    if(pos > 0) pos = qtd;
    if(pos == 0) {
      head.setNext(heat.getNext().getNext());
      qtd--;
      return;
    }
    No aux = head.getNext();
    No temp = aux;
    for(int i = 0; i < pos -1; i++){
      if(aux.getNext() != null) {
        temp = aux;   
        aux = aux.getNext();
      }
    } 
    temp.setNext(aux.getNext());
    qtd--;
  }
}
```
10) c

11) a

12) b  

13) d

14) b

15) a

16) a

17) c

18) b

19) a

20) b

21)a)v

21)b)iii

21)c)iv

22)c

23)b

24)d
