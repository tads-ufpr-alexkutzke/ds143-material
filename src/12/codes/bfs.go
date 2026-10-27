// Busca em largura (BFS) a partir do vértice 0, com o estado da fila a cada
// passo, e o caminho com menos arestas de 0 até 7.
//
// Execute com:  go run bfs.go

package main

import (
	"fmt"
	"slices"
)

type Grafo struct {
	adj [][]int
}

func NovoGrafo(n int) *Grafo {
	return &Grafo{adj: make([][]int, n)}
}

func (g *Grafo) V() int { return len(g.adj) }

func (g *Grafo) AdicionaAresta(v, w int) {
	g.adj[v] = append(g.adj[v], w)
	g.adj[w] = append(g.adj[w], v)
}

// BFS guarda o resultado da busca: quem foi alcançado, por qual aresta e a
// quantas arestas da origem.
type BFS struct {
	marcado []bool
	pai     []int // pai[w] = v: w foi alcançado pela aresta v-w
	dist    []int // número de arestas no caminho da origem até v
}

func BuscaEmLargura(g *Grafo, s int) *BFS {
	b := &BFS{
		marcado: make([]bool, g.V()),
		pai:     make([]int, g.V()),
		dist:    make([]int, g.V()),
	}
	b.pai[s] = -1
	b.marcado[s] = true
	fila := []int{s}
	for len(fila) > 0 {
		v := fila[0]
		fila = fila[1:]
		for _, w := range g.adj[v] {
			if !b.marcado[w] {
				b.marcado[w] = true
				b.pai[w] = v
				b.dist[w] = b.dist[v] + 1
				fila = append(fila, w)
			}
		}
		fmt.Printf("  retira %d   fila: %v\n", v, fila)
	}
	return b
}

// Caminho sobe pelo vetor pai, de v até a origem, e inverte o resultado.
func (b *BFS) Caminho(v int) []int {
	if !b.marcado[v] {
		return nil
	}
	var cam []int
	for x := v; x != -1; x = b.pai[x] {
		cam = append(cam, x)
	}
	slices.Reverse(cam)
	return cam
}

func main() {
	arestas := [][2]int{
		{0, 1}, {0, 2}, {0, 5}, {1, 3}, {2, 3},
		{2, 4}, {3, 6}, {4, 5}, {4, 6}, {6, 7},
	}
	g := NovoGrafo(8)
	for _, a := range arestas {
		g.AdicionaAresta(a[0], a[1])
	}

	fmt.Println("Busca em largura a partir de 0")
	b := BuscaEmLargura(g, 0)

	fmt.Println("\n  v | dist | pai")
	fmt.Println("  --+------+----")
	for v := 0; v < g.V(); v++ {
		fmt.Printf("  %d | %4d | %3d\n", v, b.dist[v], b.pai[v])
	}

	fmt.Println("\nCaminho de 0 até 7:", b.Caminho(7))
}
