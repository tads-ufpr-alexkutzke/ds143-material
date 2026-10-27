// O mesmo grafo não dirigido em duas representações: matriz de adjacência e
// listas de adjacência.
//
// Execute com:  go run representacao.go

package main

import "fmt"

// Matriz guarda, para cada par de vértices (v, w), se existe a aresta v-w.
type Matriz struct {
	adj [][]bool
}

func NovaMatriz(n int) *Matriz {
	adj := make([][]bool, n)
	for v := range adj {
		adj[v] = make([]bool, n)
	}
	return &Matriz{adj: adj}
}

// Num grafo não dirigido, a aresta v-w aparece duas vezes: em (v, w) e em
// (w, v). A matriz é simétrica.
func (g *Matriz) AdicionaAresta(v, w int) {
	g.adj[v][w] = true
	g.adj[w][v] = true
}

// Grafo guarda, para cada vértice, a lista dos seus vizinhos.
type Grafo struct {
	adj [][]int
}

func NovoGrafo(n int) *Grafo {
	return &Grafo{adj: make([][]int, n)}
}

func (g *Grafo) V() int { return len(g.adj) }

// A aresta v-w entra na lista de v e na lista de w.
func (g *Grafo) AdicionaAresta(v, w int) {
	g.adj[v] = append(g.adj[v], w)
	g.adj[w] = append(g.adj[w], v)
}

// Existe percorre a lista de v: custa o grau de v, e não uma posição só, como
// na matriz.
func (g *Grafo) Existe(v, w int) bool {
	for _, x := range g.adj[v] {
		if x == w {
			return true
		}
	}
	return false
}

func main() {
	arestas := [][2]int{
		{0, 1}, {0, 2}, {0, 5}, {1, 3}, {2, 3},
		{2, 4}, {3, 6}, {4, 5}, {4, 6}, {6, 7},
	}
	n := 8

	m := NovaMatriz(n)
	g := NovoGrafo(n)
	for _, a := range arestas {
		m.AdicionaAresta(a[0], a[1])
		g.AdicionaAresta(a[0], a[1])
	}

	fmt.Printf("%d vértices, %d arestas\n\n", n, len(arestas))

	fmt.Println("Matriz de adjacência")
	fmt.Print("    ")
	for w := 0; w < n; w++ {
		fmt.Printf(" %d", w)
	}
	fmt.Println()
	for v := 0; v < n; v++ {
		fmt.Printf("  %d ", v)
		for w := 0; w < n; w++ {
			if m.adj[v][w] {
				fmt.Print(" 1")
			} else {
				fmt.Print(" .")
			}
		}
		fmt.Println()
	}

	fmt.Println("\nListas de adjacência")
	for v := 0; v < g.V(); v++ {
		fmt.Printf("  %d: %v\n", v, g.adj[v])
	}

	fmt.Println("\nExiste a aresta 2-5?")
	fmt.Println("  matriz:", m.adj[2][5], " listas:", g.Existe(2, 5))
	fmt.Println("Existe a aresta 4-6?")
	fmt.Println("  matriz:", m.adj[4][6], " listas:", g.Existe(4, 6))
}
