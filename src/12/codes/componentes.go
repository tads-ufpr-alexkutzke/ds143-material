// Componentes conexas e detecção de ciclo em grafos não dirigidos, as duas
// com busca em profundidade.
//
// Execute com:  go run componentes.go

package main

import "fmt"

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

// Componentes devolve id[v], o número da componente de v, e quantas
// componentes o grafo tem. Cada busca iniciada no laço de fora descobre uma
// componente inteira.
func Componentes(g *Grafo) ([]int, int) {
	id := make([]int, g.V())
	marcado := make([]bool, g.V())
	var visita func(v, c int)
	visita = func(v, c int) {
		marcado[v] = true
		id[v] = c
		for _, w := range g.adj[v] {
			if !marcado[w] {
				visita(w, c)
			}
		}
	}
	quantas := 0
	for v := 0; v < g.V(); v++ {
		if !marcado[v] {
			visita(v, quantas)
			quantas++
		}
	}
	return id, quantas
}

// TemCiclo procura, durante a busca, uma aresta v-w que leve a um vértice w
// já marcado e que não seja a aresta pela qual a busca chegou a v. Devolve
// essa aresta, que fecha o ciclo.
func TemCiclo(g *Grafo) (bool, int, int) {
	marcado := make([]bool, g.V())
	var visita func(v, pai int) (bool, int, int)
	visita = func(v, pai int) (bool, int, int) {
		marcado[v] = true
		for _, w := range g.adj[v] {
			if !marcado[w] {
				if ok, a, b := visita(w, v); ok {
					return true, a, b
				}
			} else if w != pai {
				return true, v, w
			}
		}
		return false, 0, 0
	}
	for v := 0; v < g.V(); v++ {
		if !marcado[v] {
			if ok, a, b := visita(v, -1); ok {
				return true, a, b
			}
		}
	}
	return false, 0, 0
}

func monta(n int, arestas [][2]int) *Grafo {
	g := NovoGrafo(n)
	for _, a := range arestas {
		g.AdicionaAresta(a[0], a[1])
	}
	return g
}

func main() {
	arestas := [][2]int{
		{0, 1}, {0, 2}, {1, 2}, {3, 4},
		{5, 6}, {5, 7}, {6, 8}, {7, 8},
	}
	g := monta(10, arestas)

	id, quantas := Componentes(g)
	fmt.Printf("Grafo com 10 vértices e %d arestas: %d componentes\n", len(arestas), quantas)
	fmt.Println("\n  v  | 0 1 2 3 4 5 6 7 8 9")
	fmt.Print("  id |")
	for v := 0; v < g.V(); v++ {
		fmt.Printf(" %d", id[v])
	}
	fmt.Println()
	fmt.Println("\n0 e 2 na mesma componente?", id[0] == id[2])
	fmt.Println("2 e 5 na mesma componente?", id[2] == id[5])

	if ok, a, b := TemCiclo(g); ok {
		fmt.Printf("\nCiclo: a aresta %d-%d leva a um vértice já marcado\n", a, b)
	}

	// Sem as arestas 1-2 e 7-8, cada componente vira uma árvore.
	floresta := monta(10, [][2]int{{0, 1}, {0, 2}, {3, 4}, {5, 6}, {5, 7}, {6, 8}})
	ok, _, _ := TemCiclo(floresta)
	fmt.Println("Sem as arestas 1-2 e 7-8, tem ciclo?", ok)
}
