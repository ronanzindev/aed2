using ListaEncadeada;
namespace BuscarOrdernacao;

class Buscar
{
  Lista lista;
  public Buscar()
  {
    lista = new();
    lista.inserirNoInicio(new("c"));
    lista.inserirNoInicio(new("b"));
    lista.inserirNoInicio(new("a"));
  }
  public No BuscarSequencial(string element)
  {
    No aux = lista.Head();
    for (int i = 0; i <= lista.Qtd(); i++)
    {
      if (aux.getElement() == element) return aux;
      aux = aux.getNext();
    }
    return null;
  }

  public int BuscarBinaria(int[] vetor, int alvo) {
    // supor vetor = [1,2,3,4,5,6,7,8,9,10]
    int tam = vetor.Length;
    if(tam == 0) {
      return -1;
    }
    // 0
    int inicio = 0;
    // 9
    int final = tam -1;
    while(inicio <= final) {
      int meio = (final + inicio) / 2;
      // 5
      int numeroAtual = vetor[meio]; 
      if(numeroAtual ==alvo) return numeroAtual;
      if(numeroAtual > alvo)  final = meio - 1;
      if(numeroAtual < alvo) inicio = meio  + 1;
    }
    return -1;
  }
}

class Ordernar { }
