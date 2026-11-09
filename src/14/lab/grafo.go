package main

// Aresta liga v e w, com um peso. Num grafo não dirigido, v-w e w-v são a
// mesma aresta.
type Aresta struct {
	V, W, Peso int
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

// NovoGrafoP cria um grafo com n vértices, numerados de 0 a n-1, sem arestas.
func NovoGrafoP(n int) *GrafoP {
	return &GrafoP{adj: make([][]Aresta, n)}
}

// V devolve o número de vértices.
func (g *GrafoP) V() int { return len(g.adj) }

// AdicionaAresta coloca a aresta na lista de cada um dos seus extremos.
func (g *GrafoP) AdicionaAresta(a Aresta) {
	g.adj[a.V] = append(g.adj[a.V], a)
	g.adj[a.W] = append(g.adj[a.W], a)
}

// Adj devolve as arestas que saem de v.
func (g *GrafoP) Adj(v int) []Aresta { return g.adj[v] }

// Grau devolve o número de arestas que saem de v.
func (g *GrafoP) Grau(v int) int { return len(g.adj[v]) }
