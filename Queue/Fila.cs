using ListaEncadeada;
namespace Queue;
class Fila : Lista {
  public void Push(string element) {
    No n = new(element);
    this.inserirNoFinal(n);
  }
  public No Pop() {
    No n = this.Buscar(1);
    this.deletarNaPos(1);
    return n;
  }
  public void Print() {
    this.printar();
  }
}