// Algoritmo de Dijkstra no grafo G com pesos, a partir do vértice 0. A árvore
// cresce um vértice por vez, sempre pelo vértice de fora mais próximo da
// origem. Uma fila de prioridade (heap de mínimo) guarda os candidatos.
//
// Execute com:  go run dijkstra.go

package main

import (
	"container/heap"
	"fmt"
	"math"
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

// Candidato é um vértice que pode entrar na árvore, com a distância da origem
// até ele pelo caminho encontrado até aqui.
type Candidato struct {
	V    int
	Dist float64
}

// FilaCandidatos implementa heap.Interface: o menor Dist sai primeiro.
type FilaCandidatos []Candidato

func (f FilaCandidatos) Len() int           { return len(f) }
func (f FilaCandidatos) Less(i, j int) bool { return f[i].Dist < f[j].Dist }
func (f FilaCandidatos) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }
func (f *FilaCandidatos) Push(x any)        { *f = append(*f, x.(Candidato)) }
func (f *FilaCandidatos) Pop() any {
	antigo := *f
	n := len(antigo)
	c := antigo[n-1]
	*f = antigo[:n-1]
	return c
}

// Dijkstra devolve, para cada vértice, a distância a partir de s e o vértice
// anterior no caminho mínimo. Os pesos não podem ser negativos.
func Dijkstra(g *GrafoP, s int) (dist []float64, anterior []int) {
	dist = make([]float64, g.V())
	anterior = make([]int, g.V())
	for v := range dist {
		dist[v] = math.Inf(1) // ainda sem caminho
		anterior[v] = -1
	}
	naArvore := make([]bool, g.V())
	fila := &FilaCandidatos{}

	dist[s] = 0
	heap.Push(fila, Candidato{s, 0})
	mostra("início", fila)

	for fila.Len() > 0 {
		c := heap.Pop(fila).(Candidato)
		v := c.V
		if naArvore[v] {
			mostra(fmt.Sprintf("%d(%g) descarta", v, c.Dist), fila)
			continue // v já entrou por um caminho menor
		}
		naArvore[v] = true

		// Relaxa as arestas de v: passar por v melhora o caminho até w?
		for _, a := range g.adj[v] {
			w := a.Outro(v)
			if !naArvore[w] && dist[v]+a.Peso < dist[w] {
				dist[w] = dist[v] + a.Peso
				anterior[w] = v
				heap.Push(fila, Candidato{w, dist[w]})
			}
		}
		mostra(fmt.Sprintf("%d(%g) entra", v, c.Dist), fila)
	}
	return dist, anterior
}

// Caminho devolve os vértices do caminho mínimo de s até v, seguindo o vetor
// anterior de trás para a frente.
func Caminho(anterior []int, v int) []int {
	var caminho []int
	for x := v; x != -1; x = anterior[x] {
		caminho = append(caminho, x)
	}
	slices.Reverse(caminho)
	return caminho
}

// mostra imprime a fila em ordem de distância. O slice do heap não está
// ordenado: a cópia ordenada é só para leitura.
func mostra(passo string, f *FilaCandidatos) {
	copia := slices.Clone(*f)
	slices.SortFunc(copia, func(a, b Candidato) int {
		if a.Dist < b.Dist {
			return -1
		}
		if a.Dist > b.Dist {
			return 1
		}
		return 0
	})
	fmt.Printf("  %-14s fila:", passo)
	for _, c := range copia {
		fmt.Printf(" %d(%g)", c.V, c.Dist)
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

	fmt.Println("Dijkstra a partir de 0")
	dist, anterior := Dijkstra(g, 0)

	fmt.Println("\nv  dist  anterior  caminho")
	for v := range dist {
		fmt.Printf("%d  %4g  %8d  %v\n", v, dist[v], anterior[v], Caminho(anterior, v))
	}
}
