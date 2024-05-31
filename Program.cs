// See https://aka.ms/new-console-template for more information
using static System.Console;
using ListaEncadeada;
Lista l = new();
l.inserirNoInicio(new("c"));
l.inserirOrdenado(new("b"));
l.printar();
WriteLine("-------------------------------------------------");
l.deletarNaPos(1);
l.printar();