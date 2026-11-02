// Rede de fibra óptica para as escolas rurais de um município fictício. A sede
// (S) já tem conexão, e as escolas A a K precisam ser ligadas a ela. Compara
// duas redes: cada escola ligada diretamente à sede, e a árvore geradora
// mínima, calculada pelo algoritmo de Kruskal sobre todas as ligações
// possíveis.
//
// No mapa, cada coluna vale 1 km e cada linha, 2 km, porque o caractere do
// terminal é mais alto que largo.
//
// Execute com:  go run escolas.go

package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

type Local struct {
	nome     string
	col, lin int
}

var locais = []Local{
	{"S", 15, 5},
	{"A", 3, 1}, {"B", 8, 2}, {"C", 5, 7}, {"D", 1, 9}, {"E", 11, 9},
	{"F", 20, 1}, {"G", 27, 0}, {"H", 23, 6}, {"I", 29, 9}, {"J", 18, 10},
	{"K", 12, 0},
}

const largura, altura = 31, 11

type Aresta struct {
	V, W int
	Peso float64
}

// distancia em km, em linha reta.
func distancia(a, b Local) float64 {
	dx := float64(a.col - b.col)
	dy := 2 * float64(a.lin-b.lin)
	return math.Hypot(dx, dy)
}

type UF struct{ pai, tamanho []int }

func NovoUF(n int) *UF {
	uf := &UF{make([]int, n), make([]int, n)}
	for i := range uf.pai {
		uf.pai[i], uf.tamanho[i] = i, 1
	}
	return uf
}

func (uf *UF) Find(p int) int {
	for p != uf.pai[p] {
		p = uf.pai[p]
	}
	return p
}

func (uf *UF) Union(p, q int) bool {
	rp, rq := uf.Find(p), uf.Find(q)
	if rp == rq {
		return false
	}
	if uf.tamanho[rp] < uf.tamanho[rq] {
		rp, rq = rq, rp
	}
	uf.pai[rq] = rp
	uf.tamanho[rp] += uf.tamanho[rq]
	return true
}

func Kruskal(n int, arestas []Aresta) []Aresta {
	sort.Slice(arestas, func(i, j int) bool { return arestas[i].Peso < arestas[j].Peso })
	uf := NovoUF(n)
	var arvore []Aresta
	for _, a := range arestas {
		if uf.Union(a.V, a.W) {
			arvore = append(arvore, a)
		}
	}
	return arvore
}

// desenha devolve as linhas do mapa, com as ligações traçadas por pontos.
func desenha(rede []Aresta) []string {
	grade := make([][]byte, altura)
	for l := range grade {
		grade[l] = []byte(strings.Repeat(" ", largura))
	}
	for _, a := range rede {
		p, q := locais[a.V], locais[a.W]
		passos := max(abs(p.col-q.col), abs(p.lin-q.lin))
		for i := 1; i < passos; i++ {
			c := p.col + int(math.Round(float64((q.col-p.col)*i)/float64(passos)))
			l := p.lin + int(math.Round(float64((q.lin-p.lin)*i)/float64(passos)))
			grade[l][c] = '.'
		}
	}
	for _, x := range locais {
		grade[x.lin][x.col] = x.nome[0]
	}
	linhas := make([]string, altura)
	for l := range grade {
		linhas[l] = "|" + string(grade[l]) + "|"
	}
	return linhas
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func total(rede []Aresta) float64 {
	t := 0.0
	for _, a := range rede {
		t += a.Peso
	}
	return t
}

func main() {
	n := len(locais)

	// Rede 1: cada escola ligada direto à sede.
	var estrela []Aresta
	for v := 1; v < n; v++ {
		estrela = append(estrela, Aresta{0, v, distancia(locais[0], locais[v])})
	}

	// Rede 2: árvore geradora mínima do grafo completo, com todas as ligações
	// possíveis entre dois locais.
	var todas []Aresta
	for v := 0; v < n; v++ {
		for w := v + 1; w < n; w++ {
			todas = append(todas, Aresta{v, w, distancia(locais[v], locais[w])})
		}
	}
	agm := Kruskal(n, todas)

	m1, m2 := desenha(estrela), desenha(agm)
	fmt.Printf("%-35s %s\n", "Tudo ligado à sede", "Árvore geradora mínima")
	for l := range m1 {
		fmt.Printf("%-35s %s\n", m1[l], m2[l])
	}
	fmt.Printf("%-35s %s\n", fmt.Sprintf("%d ligações, %.1f km", len(estrela), total(estrela)),
		fmt.Sprintf("%d ligações, %.1f km", len(agm), total(agm)))
	fmt.Printf("\n%d ligações possíveis entre %d locais.\n", len(todas), n)
	fmt.Printf("A árvore usa %.0f%% do cabo da rede ligada à sede.\n",
		100*total(agm)/total(estrela))

	fmt.Println("\nLigações da árvore e escolas isoladas da sede se a ligação for cortada:")
	for _, a := range agm {
		fmt.Printf("  %s-%s %5.1f km   %s\n", locais[a.V].nome, locais[a.W].nome,
			a.Peso, isoladas(agm, a))
	}
}

// isoladas lista os locais que perdem o caminho até a sede quando a ligação
// cortada sai da árvore.
func isoladas(arvore []Aresta, cortada Aresta) string {
	uf := NovoUF(len(locais))
	for _, a := range arvore {
		if a != cortada {
			uf.Union(a.V, a.W)
		}
	}
	var nomes []string
	for v := 1; v < len(locais); v++ {
		if uf.Find(v) != uf.Find(0) {
			nomes = append(nomes, locais[v].nome)
		}
	}
	return strings.Join(nomes, " ")
}
