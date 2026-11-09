// Estradas de um município fictício, o mesmo da aula de árvores geradoras
// mínimas. A sede (S) e as escolas A a K são vértices, e cada estrada é uma
// aresta, com o comprimento como peso. Compara duas árvores sobre as estradas:
// a árvore geradora mínima, que usa o menor total de estrada, e a árvore de
// caminhos mínimos a partir da sede, que dá a rota mais curta até cada escola.
//
// No mapa, cada coluna vale 1 km e cada linha, 2 km, porque o caractere do
// terminal é mais alto que largo.
//
// Execute com:  go run rotas.go

package main

import (
	"container/heap"
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

// As estradas, em pares de locais.
var estradas = []string{
	"S-B", "S-E", "S-F", "S-H", "S-K", "S-J",
	"A-B", "A-C", "B-C", "B-K", "C-D", "C-E", "D-E",
	"E-J", "F-G", "F-H", "F-K", "G-H", "H-I", "H-J", "I-J",
}

const largura, altura = 31, 11

type Aresta struct {
	V, W int
	Peso float64
}

func (a Aresta) Outro(v int) int {
	if v == a.V {
		return a.W
	}
	return a.V
}

// distancia em km, em linha reta.
func distancia(a, b Local) float64 {
	dx := float64(a.col - b.col)
	dy := 2 * float64(a.lin-b.lin)
	return math.Hypot(dx, dy)
}

func indice(nome string) int {
	for i, x := range locais {
		if x.nome == nome {
			return i
		}
	}
	panic("local desconhecido: " + nome)
}

// Kruskal, com Union-Find, como na aula de árvores geradoras mínimas.
func Kruskal(n int, arestas []Aresta) []Aresta {
	ordem := append([]Aresta(nil), arestas...)
	sort.Slice(ordem, func(i, j int) bool { return ordem[i].Peso < ordem[j].Peso })
	pai := make([]int, n)
	for i := range pai {
		pai[i] = i
	}
	var find func(int) int
	find = func(p int) int {
		for p != pai[p] {
			p = pai[p]
		}
		return p
	}
	var arvore []Aresta
	for _, a := range ordem {
		if rv, rw := find(a.V), find(a.W); rv != rw {
			pai[rv] = rw
			arvore = append(arvore, a)
		}
	}
	return arvore
}

type Candidato struct {
	V    int
	Dist float64
}

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

// Dijkstra devolve a distância de s até cada vértice e a aresta do caminho
// mínimo que chega a cada um.
func Dijkstra(n int, arestas []Aresta, s int) ([]float64, []Aresta) {
	adj := make([][]Aresta, n)
	for _, a := range arestas {
		adj[a.V] = append(adj[a.V], a)
		adj[a.W] = append(adj[a.W], a)
	}
	dist := make([]float64, n)
	chegada := make([]Aresta, n)
	for v := range dist {
		dist[v] = math.Inf(1)
	}
	naArvore := make([]bool, n)
	dist[s] = 0
	fila := &FilaCandidatos{{s, 0}}
	for fila.Len() > 0 {
		v := heap.Pop(fila).(Candidato).V
		if naArvore[v] {
			continue
		}
		naArvore[v] = true
		for _, a := range adj[v] {
			w := a.Outro(v)
			if !naArvore[w] && dist[v]+a.Peso < dist[w] {
				dist[w] = dist[v] + a.Peso
				chegada[w] = a
				heap.Push(fila, Candidato{w, dist[w]})
			}
		}
	}
	return dist, chegada
}

// desenha devolve as linhas do mapa, com as estradas traçadas por pontos.
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

	var todas []Aresta
	for _, e := range estradas {
		partes := strings.Split(e, "-")
		v, w := indice(partes[0]), indice(partes[1])
		todas = append(todas, Aresta{v, w, distancia(locais[v], locais[w])})
	}

	agm := Kruskal(n, todas)
	distCaminho, chegada := Dijkstra(n, todas, 0)
	arvoreCaminhos := chegada[1:] // a sede não tem aresta de chegada

	// A distância pela árvore mínima é a de Dijkstra só sobre as arestas dela,
	// porque numa árvore há um único caminho entre dois vértices.
	distAGM, _ := Dijkstra(n, agm, 0)

	m1, m2 := desenha(agm), desenha(arvoreCaminhos)
	fmt.Printf("%-35s %s\n", "Árvore geradora mínima", "Árvore de caminhos mínimos, de S")
	for l := range m1 {
		fmt.Printf("%-35s %s\n", m1[l], m2[l])
	}
	fmt.Printf("%-35s %s\n", fmt.Sprintf("%.1f km de estrada", total(agm)),
		fmt.Sprintf("%.1f km de estrada", total(arvoreCaminhos)))

	fmt.Printf("\nDistância de S, entre %d estradas:\n", len(todas))
	fmt.Println("  escola   pela árvore mínima   caminho mínimo")
	for v := 1; v < n; v++ {
		fmt.Printf("  %-8s %15.1f km %13.1f km\n", locais[v].nome, distAGM[v], distCaminho[v])
	}
}
