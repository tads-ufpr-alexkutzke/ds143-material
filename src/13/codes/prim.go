// Algoritmo de Prim no grafo G com pesos, a partir do vértice 0. A árvore
// cresce um vértice por vez, sempre pela aresta mais leve que sai dela. Uma
// fila de prioridade (heap de mínimo) guarda as arestas candidatas.
//
// Esta é a versão "preguiçosa" (lazy): arestas que deixaram de servir ficam na
// fila e são descartadas quando saem dela.
//
// Execute com:  go run prim.go

package main

import (
	"container/heap"
	"fmt"
	"slices"
)

type Aresta struct {
	V, W int
	Peso float64
}

// Outro devolve o extremo da aresta que não é v.
func (a Aresta) Outro(v int) int {
	if v == a.V {
		return a.W
	}
	return a.V
}

// GrafoP é um grafo não dirigido com pesos, em listas de adjacência. A aresta
// v-w aparece na lista de v e na lista de w.
type GrafoP struct {
	adj [][]Aresta
}

func NovoGrafoP(n int) *GrafoP {
	return &GrafoP{adj: make([][]Aresta, n)}
}

func (g *GrafoP) V() int { return len(g.adj) }

func (g *GrafoP) AdicionaAresta(a Aresta) {
	g.adj[a.V] = append(g.adj[a.V], a)
	g.adj[a.W] = append(g.adj[a.W], a)
}

// FilaArestas implementa heap.Interface: o pacote container/heap mantém o
// slice organizado como heap de mínimo pelo peso, com Push e Pop em
// O(log n).
type FilaArestas []Aresta

func (f FilaArestas) Len() int           { return len(f) }
func (f FilaArestas) Less(i, j int) bool { return f[i].Peso < f[j].Peso }
func (f FilaArestas) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }
func (f *FilaArestas) Push(x any)        { *f = append(*f, x.(Aresta)) }
func (f *FilaArestas) Pop() any {
	antigo := *f
	n := len(antigo)
	a := antigo[n-1]
	*f = antigo[:n-1]
	return a
}

// Prim devolve as arestas da árvore geradora mínima de um grafo conexo, na
// ordem em que entraram, começando por s.
func Prim(g *GrafoP, s int) []Aresta {
	naArvore := make([]bool, g.V())
	fila := &FilaArestas{}
	var arvore []Aresta

	// visita põe v na árvore e as arestas de v para fora dela na fila.
	visita := func(v int) {
		naArvore[v] = true
		for _, a := range g.adj[v] {
			if !naArvore[a.Outro(v)] {
				heap.Push(fila, a)
			}
		}
	}

	visita(s)
	mostra("início", fila)
	for fila.Len() > 0 && len(arvore) < g.V()-1 {
		a := heap.Pop(fila).(Aresta)
		if naArvore[a.V] && naArvore[a.W] {
			mostra(fmt.Sprintf("%d-%d %2g descarta", a.V, a.W, a.Peso), fila)
			continue // os dois extremos já estão na árvore
		}
		arvore = append(arvore, a)
		novo := a.V
		if naArvore[a.V] {
			novo = a.W
		}
		visita(novo)
		mostra(fmt.Sprintf("%d-%d %2g entra %d", a.V, a.W, a.Peso, novo), fila)
	}
	return arvore
}

// mostra imprime a fila em ordem de peso. O slice do heap não está ordenado:
// a cópia ordenada é só para leitura.
func mostra(passo string, f *FilaArestas) {
	copia := slices.Clone(*f)
	slices.SortFunc(copia, func(a, b Aresta) int {
		if a.Peso < b.Peso {
			return -1
		}
		if a.Peso > b.Peso {
			return 1
		}
		return 0
	})
	fmt.Printf("  %-17s fila:", passo)
	for _, a := range copia {
		fmt.Printf(" %d-%d(%g)", a.V, a.W, a.Peso)
	}
	fmt.Println()
}

func main() {
	arestas := []Aresta{
		{0, 1, 4}, {0, 2, 3}, {0, 5, 7}, {1, 3, 2}, {2, 3, 5},
		{2, 4, 6}, {3, 6, 9}, {4, 5, 1}, {4, 6, 8}, {6, 7, 12},
	}
	g := NovoGrafoP(8)
	for _, a := range arestas {
		g.AdicionaAresta(a)
	}

	fmt.Println("Prim a partir de 0")
	arvore := Prim(g, 0)

	total := 0.0
	fmt.Print("\nÁrvore geradora mínima:")
	for _, a := range arvore {
		fmt.Printf(" %d-%d", a.V, a.W)
		total += a.Peso
	}
	fmt.Printf("\n%d arestas, peso total %g\n", len(arvore), total)
}
