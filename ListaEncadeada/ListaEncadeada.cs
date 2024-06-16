namespace ListaEncadeada;
public class Lista
{
  private No head;
  private int qtd;

  public Lista()
  {
    head = new No("a");
    qtd = 0;
  }
  public No Head() => head;
  public int Qtd() => qtd;
   public void inserirNoInicio(No n)
  {
    n.setNext(head.getNext());
    head.setNext(n);
    qtd++;
  }

  public void inserirNoFinal(No n)
  {
    No aux = head;
    while (aux.getNext() != null) aux = aux.getNext();
    aux.setNext(n);
    qtd++;
  }
  // la ele
  public void inserirNaPos(int pos, No n)
  {
    if (pos < 0)
    {
      return;
    }
    if (pos > qtd) pos = qtd - 1;
    if (pos == 0)
    {
      n.setNext(head);
      head = n;
      qtd++;
      return;
    }
    No aux = head;
    for (int i = 0; i < pos - 1; pos++)
    {
      aux = aux.getNext();
    }
    n.setNext(aux.getNext());
    aux.setNext(n);
    qtd++;
  }

  public void inserirOrdenado(No n)
  {
    if (string.Compare(n.getElement(), head.getElement()) == -1)
    {
      n.setNext(head);
      head = n;
      qtd++;
      return;
    }
    No anterior = head;
    No aux = head.getNext();
    while (string.Compare(n.getElement(), aux.getElement()) == 1)
    {
      anterior = aux;
      aux = aux.getNext();
    }
    n.setNext(aux);
    anterior.setNext(n);
    qtd++;
  }

  public void deletarNoInicio() {
    if(head.getNext() == null) return;
    head = head.getNext();
    qtd--;
  }

  public void deletarNoFinal() {
    No aux = head;
    while(aux.getNext().getNext() != null) {
      aux = aux.getNext();
    }
    aux.setNext(null);
    qtd--;
  }

  public void deletarNaPos(int pos) {
    if(pos  < 0) return;
    if(pos > qtd) pos = qtd -1;
    if(pos == 0) head = head.getNext();
    No anterior = head;
    No aux = head.getNext();
    for(int i = 0; i  < pos -1; i++ ) {
      anterior = aux;
      aux = aux.getNext();
    } 
    anterior.setNext(aux.getNext());
    qtd--;
  }
  public No Buscar(int pos) {
    if(pos < 0) return null;
    if(pos > qtd) pos = qtd -1;
    No aux = head;
    for(int i = 0; i < pos; i++) {
      aux = aux.getNext();
    }
    return aux;
  }
  public void printar()
  {
    No aux = head;
    int i = 0;
    while (aux != null)
    {
      Console.WriteLine($"{i} -> {aux.getElement()}");
      aux = aux.getNext();
      i++;
    }
  }
}