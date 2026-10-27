// Ordenação topológica e detecção de ciclo em um grafo dirigido: os pacotes de
// um projeto Go e as importações entre eles. O arco B -> A significa que A
// importa B, e por isso B precisa ser compilado antes de A.
//
// Execute com:  go run topologica.go

package main

import (
	"fmt"
	"slices"
)

// Digrafo guarda, para cada vértice, a lista dos vértices para os quais ele
// aponta. O arco v -> w entra só na lista de v.
type Digrafo struct {
	adj [][]int
}

func NovoDigrafo(n int) *Digrafo {
	return &Digrafo{adj: make([][]int, n)}
}

func (g *Digrafo) V() int { return len(g.adj) }

func (g *Digrafo) AdicionaArco(v, w int) {
	g.adj[v] = append(g.adj[v], w)
}

// Estados de um vértice durante a busca.
const (
	branco = iota // ainda não visitado
	cinza         // visitado, e a chamada dele ainda não terminou
	preto         // a chamada dele já terminou
)

// Ordena faz uma busca em profundidade a partir de cada vértice ainda branco e
// devolve a pós-ordem invertida. Se encontrar um arco para um vértice cinza,
// o grafo tem ciclo, e Ordena devolve o ciclo no lugar da ordem.
func Ordena(g *Digrafo) (ordem []int, ciclo []int) {
	cor := make([]int, g.V())
	pai := make([]int, g.V())
	var pos []int

	var visita func(v int) bool
	visita = func(v int) bool {
		cor[v] = cinza
		for _, w := range g.adj[v] {
			switch cor[w] {
			case branco:
				pai[w] = v
				if visita(w) {
					return true
				}
			case cinza:
				// w ainda está na pilha de chamadas: o caminho de w até v,
				// mais o arco v -> w, é um ciclo.
				for x := v; x != w; x = pai[x] {
					ciclo = append(ciclo, x)
				}
				ciclo = append(ciclo, w)
				slices.Reverse(ciclo)
				return true
			}
		}
		cor[v] = preto
		pos = append(pos, v)
		return false
	}

	for v := 0; v < g.V(); v++ {
		if cor[v] == branco && visita(v) {
			return nil, ciclo
		}
	}
	slices.Reverse(pos)
	return pos, nil
}

var nomes = []string{"modelo", "config", "log", "banco", "servico", "http", "main"}

func nome(vs []int) []string {
	var s []string
	for _, v := range vs {
		s = append(s, nomes[v])
	}
	return s
}

func main() {
	const (
		modelo = iota
		config
		log
		banco
		servico
		http
		main
	)
	// importa[a] lista os pacotes que a importa.
	importa := map[int][]int{
		banco:   {modelo, config, log},
		servico: {banco, modelo, log},
		http:    {servico, log},
		main:    {http, config, servico},
	}

	g := NovoDigrafo(len(nomes))
	for a := 0; a < len(nomes); a++ {
		for _, b := range importa[a] {
			g.AdicionaArco(b, a)
		}
	}

	fmt.Println("Arcos (B -> A: A importa B)")
	for v := 0; v < g.V(); v++ {
		fmt.Printf("  %-8s -> %v\n", nomes[v], nome(g.adj[v]))
	}

	ordem, _ := Ordena(g)
	fmt.Println("\nOrdem de compilação:", nome(ordem))

	// O pacote modelo passa a importar o servico.
	g.AdicionaArco(servico, modelo)
	_, ciclo := Ordena(g)
	fmt.Println("\nCom modelo importando servico:")
	fmt.Println("  ciclo:", nome(ciclo))
}
