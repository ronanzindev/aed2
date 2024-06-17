using static System.Console;
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
  public void PreOrder() {
    if(this == null) return;
    WriteLine($"{elemento}");
    left.PreOrder(); 
    rigth.PreOrder(); 
  }
  public void InOrder() {
    if(this == null)  return;
    left.InOrder();
    WriteLine($"{elemento}");
    rigth.InOrder();
  }
  public void PosOrder(){
    if(this == null) return;
    left.PosOrder();
    rigth.PosOrder();
    WriteLine($"{elemento}");
  }
  public No econtrarMenorElemento() {
    No aux = this;
    while(aux != null && aux.left != null) {
      aux = aux.left;
    }
    return aux;
  }
  public No deletar(int e) {
    if(this == null) return null;
    if(e < elemento) { 
      left = left.deletar(elemento);
    }
    else if(e > elemento) { 
     rigth = rigth.deletar(elemento);
    } else {
      if(left == null && rigth == null) return null;
      if(left == null) return rigth;
      if(rigth == null) return left;
      No minimo = rigth.econtrarMenorElemento();
      elemento = minimo.elemento;
      rigth = rigth.deletar(minimo.elemento);
    } 
    return this;
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
  public void PreOrder() => raiz.PreOrder();
  public void EmOrder() =>  raiz.InOrder();
  public void PosOrder() => raiz.PosOrder();

  public void deletar(int elemento) => raiz = raiz.deletar(elemento);
  

}