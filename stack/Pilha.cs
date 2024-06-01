using ListaEncadeada;
namespace Stack;
class Pilha : Lista {
  public Pilha() {}
  public void Push(string e) {
    No n = new(e);
     this.inserirNoInicio(n);
  }
  public No Pop() {
    No n = this.Buscar(1);
    this.deletarNaPos(1);
    return n;
  }
  public void Print()  => this.printar(); 
}