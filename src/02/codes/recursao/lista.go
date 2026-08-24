// Lista encadeada simples, com Tamanho e Maximo implementados
// recursivamente. Diferente do C, não há função para liberar a lista:
// o coletor de lixo (garbage collector) do Go cuida disso automaticamente.

package main

import "fmt"

type Lista struct {
	Info int
	Next *Lista
}

func criar() *Lista {
	return nil
}

func inserir(lista *Lista, elem int) *Lista {
	return &Lista{Info: elem, Next: lista}
}

func imprimir(lista *Lista) {
	for v := lista; v != nil; v = v.Next {
		fmt.Printf("Valor: %d\n", v.Info)
	}
}

// tamanho conta os elementos da lista recursivamente.
func tamanho(lista *Lista) int {
	if lista == nil {
		return 0
	}
	return 1 + tamanho(lista.Next)
}

// maximo encontra o maior valor da lista recursivamente.
func maximo(lista *Lista) int {
	fmt.Printf("--- %d\n", lista.Info)
	if lista.Next == nil {
		return lista.Info
	}
	tmp := maximo(lista.Next)
	if tmp < lista.Info {
		return lista.Info
	}
	return tmp
}

func main() {
	lista := criar()
	lista = inserir(lista, 1)
	lista = inserir(lista, 2)
	lista = inserir(lista, 4)
	lista = inserir(lista, 400)
	lista = inserir(lista, 4)
	lista = inserir(lista, 4)
	lista = inserir(lista, 4)

	imprimir(lista)
	fmt.Printf("Tamanho da lista: %d\n", tamanho(lista))
	fmt.Printf("Maior da lista: %d\n", maximo(lista))
}
