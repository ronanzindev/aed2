namespace TabelaHash;

class Tabela
{
  private int chave;
  private string valor;
  public Tabela(int c, string v)
  {
    chave = c;
    valor = v;
  }
  public void setChave(int c) => chave = c;
  public void setValor(string v) => valor = v;
  public int getChave() => chave;
  public string getValor() => valor;
}
class TabelaHashEnderecamentoAberto
{
  private int tamanho;
  private Tabela[] items;
  public TabelaHashEnderecamentoAberto(int tam)
  {
    tamanho = tam;
    items = new Tabela[tam];
  }
  public int Hash(int chave) => chave % tamanho;

  public int sondagemLinear(int pos, int i) => (pos + i) % tamanho;

  public void inserir(int chave, string valor)
  {
    int pos = Hash(chave);
    for (int i = 0; i < tamanho; i++)
    {
      int novaPos = sondagemLinear(pos, i);
      if (items[novaPos] == null)
      {
        items[novaPos] = new Tabela(chave, valor);
        return;
      }
    }
  }
  public string buscar(int chave)
  {
    int pos = Hash(chave);
    for (int i = 0; i < tamanho; i++)
    {
      int novaPos = sondagemLinear(pos, i);
      if (items[novaPos] == null) return "";
      if (items[novaPos].getChave() == chave) return items[novaPos].getValor();
    }
    return "";
  }
}

class TabelaEncadeamento {
   private int chave;
  private string valor;
  private TabelaEncadeamento proximo;
  public TabelaEncadeamento(int c, string v, TabelaEncadeamento prox)
  {
    chave = c;
    valor = v;
    proximo = prox;
  }
  public void setChave(int c) => chave = c;
  public void setValor(string v) => valor = v;
  public void setProximo(TabelaEncadeamento t) => proximo = t;
  public TabelaEncadeamento getProximo() => proximo;
  public int getChave() => chave;
  public string getValor() => valor;
}

class TabelaHashEncadeamento {
  private int tamanho;
  private TabelaEncadeamento[] items;
  public TabelaHashEncadeamento(int tam) {
    tamanho = tam;
    items = new TabelaEncadeamento[tam]; 
  } 
  public int Hash(int chave) => chave % tamanho;
  public void inserir(int chave, string valor) {
    int pos = Hash(chave);
    TabelaEncadeamento t = new TabelaEncadeamento(chave, valor, items[pos]);
    items[pos] = t; 
  }
  public string buscar(int chave) {
    int pos = Hash(chave);
    TabelaEncadeamento aux = items[pos];
    while(aux != null) {
      if(aux.getChave() == chave) return aux.getValor();
      aux = aux.getProximo();
    }
     return "";
  }
}