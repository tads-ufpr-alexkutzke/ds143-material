// Pacote abb: árvore binária de busca.
//
// Este arquivo está pronto e é o mesmo código da Aula 06. Leia antes de
// escrever qualquer coisa em exercicios.go: as funções que faltam seguem
// o mesmo estilo.
package abb

import "fmt"

// No é um nó da árvore binária de busca.
type No struct {
	Info int
	Esq  *No
	Dir  *No
}

// Busca procura um valor descendo por um único caminho: em cada nó, a
// comparação descarta uma das duas subárvores inteiras.
func Busca(a *No, procurado int) bool {
	if a == nil {
		return false
	}
	if procurado < a.Info {
		return Busca(a.Esq, procurado)
	}
	if procurado > a.Info {
		return Busca(a.Dir, procurado)
	}
	return true
}

// Insere acrescenta um valor e devolve a raiz da árvore resultante. O nó
// novo entra sempre como folha. Valores repetidos são ignorados.
func Insere(a *No, v int) *No {
	if a == nil {
		return &No{Info: v}
	}
	if v < a.Info {
		a.Esq = Insere(a.Esq, v)
	} else if v > a.Info {
		a.Dir = Insere(a.Dir, v)
	}
	return a
}

// Minimo devolve o menor valor da árvore, que está no nó mais à
// esquerda. Devolve 0 e false para a árvore vazia.
func Minimo(a *No) (int, bool) {
	if a == nil {
		return 0, false
	}
	for a.Esq != nil {
		a = a.Esq
	}
	return a.Info, true
}

// Maximo devolve o maior valor da árvore, que está no nó mais à direita.
func Maximo(a *No) (int, bool) {
	if a == nil {
		return 0, false
	}
	for a.Dir != nil {
		a = a.Dir
	}
	return a.Info, true
}

// Remove tira um valor da árvore e devolve a raiz resultante. São os
// três casos da Aula 06, com o antecessor subindo no caso de duas
// subárvores.
func Remove(a *No, v int) *No {
	if a == nil {
		return nil
	}

	if v < a.Info {
		a.Esq = Remove(a.Esq, v)
		return a
	}
	if v > a.Info {
		a.Dir = Remove(a.Dir, v)
		return a
	}

	if a.Esq == nil {
		return a.Dir
	}
	if a.Dir == nil {
		return a.Esq
	}

	antecessor := a.Esq
	for antecessor.Dir != nil {
		antecessor = antecessor.Dir
	}
	a.Info = antecessor.Info
	a.Esq = Remove(a.Esq, antecessor.Info)
	return a
}

// EmOrdem devolve as chaves em ordem crescente. Em uma ABB correta, o
// resultado sai sempre ordenado.
func EmOrdem(a *No) []int {
	if a == nil {
		return nil
	}
	valores := EmOrdem(a.Esq)
	valores = append(valores, a.Info)
	return append(valores, EmOrdem(a.Dir)...)
}

// Altura devolve o comprimento do caminho mais longo até uma folha,
// contado em arestas. A árvore vazia tem altura -1.
func Altura(a *No) int {
	if a == nil {
		return -1
	}
	esq := Altura(a.Esq)
	dir := Altura(a.Dir)
	if esq > dir {
		return 1 + esq
	}
	return 1 + dir
}

// Imprime desenha a árvore deitada, com a raiz à esquerda.
func Imprime(a *No, profundidade int) {
	if a == nil {
		return
	}
	Imprime(a.Dir, profundidade+1)
	for i := 0; i < profundidade; i++ {
		fmt.Print("    ")
	}
	fmt.Println(a.Info)
	Imprime(a.Esq, profundidade+1)
}
