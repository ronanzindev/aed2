namespace ListaEncadeada;
public class No {
  private string element;
  private No next;
  public No(string e) {
    element = e;
    next = null;
  }
  public string getElement() => element;
  public No getNext() => next;
  public void setElement(string e) => element = e;
  public void setNext(No n) => next = n;
}