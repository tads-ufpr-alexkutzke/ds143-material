// Algoritmo de Kruskal no grafo G com pesos: as arestas são examinadas da mais
// leve para a mais pesada, e cada uma entra na árvore se não fechar um ciclo.
// O Union-Find responde se os dois extremos já estão ligados.
//
// Execute com:  go run kruskal.go

package main

import (
	"fmt"
	"sort"
)

// Aresta liga v e w, com um peso. Num grafo não dirigido, v-w e w-v são a
// mesma aresta.
type Aresta struct {
	V, W int
	Peso float64
}

// UF é o Union-Find com união por tamanho, do início do semestre.
type UF struct {
	pai     []int
	tamanho []int
}

func NovoUF(n int) *UF {
	uf := &UF{pai: make([]int, n), tamanho: make([]int, n)}
	for i := range uf.pai {
		uf.pai[i] = i
		uf.tamanho[i] = 1
	}
	return uf
}

func (uf *UF) Find(p int) int {
	for p != uf.pai[p] {
		p = uf.pai[p]
	}
	return p
}

func (uf *UF) Connected(p, q int) bool {
	return uf.Find(p) == uf.Find(q)
}

func (uf *UF) Union(p, q int) {
	rp, rq := uf.Find(p), uf.Find(q)
	if rp == rq {
		return
	}
	if uf.tamanho[rp] < uf.tamanho[rq] {
		rp, rq = rq, rp
	}
	uf.pai[rq] = rp
	uf.tamanho[rp] += uf.tamanho[rq]
}

// Kruskal devolve as arestas da árvore geradora mínima de um grafo conexo com
// n vértices, na ordem em que entraram.
func Kruskal(n int, arestas []Aresta) []Aresta {
	// Ordena uma cópia, para não alterar a lista de quem chamou.
	ordem := make([]Aresta, len(arestas))
	copy(ordem, arestas)
	sort.Slice(ordem, func(i, j int) bool {
		return ordem[i].Peso < ordem[j].Peso
	})

	uf := NovoUF(n)
	var arvore []Aresta
	for _, a := range ordem {
		if len(arvore) == n-1 {
			break // a árvore já tem V-1 arestas
		}
		if uf.Connected(a.V, a.W) {
			fmt.Printf("  %d-%d  %4g   descarta: %d e %d já estão ligados\n",
				a.V, a.W, a.Peso, a.V, a.W)
			continue
		}
		uf.Union(a.V, a.W)
		arvore = append(arvore, a)
		fmt.Printf("  %d-%d  %4g   entra\n", a.V, a.W, a.Peso)
	}
	return arvore
}

func main() {
	// O grafo G da aula de representação e percursos, agora com pesos.
	arestas := []Aresta{
		{0, 1, 4}, {0, 2, 3}, {0, 5, 7}, {1, 3, 2}, {2, 3, 5},
		{2, 4, 6}, {3, 6, 9}, {4, 5, 1}, {4, 6, 8}, {6, 7, 12},
	}
	n := 8

	fmt.Println("Kruskal: arestas da mais leve para a mais pesada")
	arvore := Kruskal(n, arestas)

	total := 0.0
	fmt.Print("\nÁrvore geradora mínima:")
	for _, a := range arvore {
		fmt.Printf(" %d-%d", a.V, a.W)
		total += a.Peso
	}
	fmt.Printf("\n%d arestas, peso total %g\n", len(arvore), total)
}
