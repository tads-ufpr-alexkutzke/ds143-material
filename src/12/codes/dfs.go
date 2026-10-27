// Busca em profundidade (DFS) recursiva a partir do vértice 0. Cada chamada
// imprime uma linha, recuada pela profundidade da recursão, e as ordens de
// visita (pré-ordem) e de término (pós-ordem) aparecem no fim.
//
// Execute com:  go run dfs.go

package main

import (
	"fmt"
	"slices"
	"strings"
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

type DFS struct {
	g       *Grafo
	marcado []bool
	pai     []int
	pre     []int // vértices na ordem em que a busca chega a eles
	pos     []int // vértices na ordem em que a busca termina com eles
}

func BuscaEmProfundidade(g *Grafo, s int) *DFS {
	d := &DFS{g: g, marcado: make([]bool, g.V()), pai: make([]int, g.V())}
	d.pai[s] = -1
	d.visita(s, 0)
	return d
}

func (d *DFS) visita(v, nivel int) {
	recuo := strings.Repeat("  ", nivel)
	fmt.Printf("  %svisita(%d)\n", recuo, v)
	d.marcado[v] = true
	d.pre = append(d.pre, v)
	for _, w := range d.g.adj[v] {
		if !d.marcado[w] {
			d.pai[w] = v
			d.visita(w, nivel+1)
		}
	}
	d.pos = append(d.pos, v)
	fmt.Printf("  %stermina %d\n", recuo, v)
}

func (d *DFS) Caminho(v int) []int {
	if !d.marcado[v] {
		return nil
	}
	var cam []int
	for x := v; x != -1; x = d.pai[x] {
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

	fmt.Println("Busca em profundidade a partir de 0")
	d := BuscaEmProfundidade(g, 0)

	fmt.Println("\nPré-ordem:", d.pre)
	fmt.Println("Pós-ordem:", d.pos)
	fmt.Println("\n  v | pai")
	fmt.Println("  --+----")
	for v := 0; v < g.V(); v++ {
		fmt.Printf("  %d | %3d\n", v, d.pai[v])
	}
	fmt.Println("\nCaminho de 0 até 7:", d.Caminho(7))
}
