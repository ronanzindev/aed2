namespace Arvores;

class No {
  private int elemento;
  private No left;
  private No rigth;
  public No(int e ) => elemento = e;
  public No getLeft() => left;
  public No getRigth() => rigth;
  public int getElemento() => elemento;
  public void setElemento(int e) => elemento = e;
  public void setLeft(No l) => left = l;
  public void setRigth(No r) => rigth = r;
  public void Add(int elem) {
    if(elem < elemento) {
      if(left == null) {
        left = new(elem);
      } else {
        left.Add(elem);
      }
    } else {
      if(rigth == null) {
        rigth = new(elem);
      } else {
        rigth.Add(elem);
      }
    }
  }
}
class Arvore {
  private No raiz;
  public Arvore(int elemento) => raiz = new(elemento);
  public void inserir(int e) {
    if(raiz == null) {
      raiz = new(e);
      return;
    }
    raiz.Add(e);
  }
}