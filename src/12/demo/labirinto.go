// Busca em largura e busca em profundidade no mesmo labirinto, de S até E.
// Cada célula livre é um vértice, e duas células livres vizinhas (acima,
// abaixo, à esquerda ou à direita) são ligadas por uma aresta. O programa não
// guarda as arestas: os vizinhos de uma célula são calculados quando a busca
// precisa deles.
//
// Execute com:  go run labirinto.go

package main

import "fmt"

var mapa = []string{
	"#####################",
	"#S       #          #",
	"#        #   ####   #",
	"#####    #      #   #",
	"#        #####  #   #",
	"#   ###             #",
	"#E    #             #",
	"#####################",
}

var linhas, colunas = len(mapa), len(mapa[0])

// Uma célula é identificada por um único inteiro, linha*colunas + coluna, que
// serve de índice para os vetores das buscas.
func celula(l, c int) int { return l*colunas + c }

// vizinhos devolve as células livres vizinhas de v, na ordem direita, abaixo,
// esquerda, acima.
func vizinhos(v int) []int {
	l, c := v/colunas, v%colunas
	var viz []int
	for _, d := range [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
		nl, nc := l+d[0], c+d[1]
		if mapa[nl][nc] != '#' {
			viz = append(viz, celula(nl, nc))
		}
	}
	return viz
}

// bfs devolve o vetor pai da busca em largura a partir de s, e o número de
// células visitadas até chegar a e.
func bfs(s, e int) ([]int, int) {
	pai := make([]int, linhas*colunas)
	marcado := make([]bool, linhas*colunas)
	visitadas := 0
	fila := []int{s}
	marcado[s] = true
	for len(fila) > 0 {
		v := fila[0]
		fila = fila[1:]
		visitadas++
		if v == e {
			break
		}
		for _, w := range vizinhos(v) {
			if !marcado[w] {
				marcado[w] = true
				pai[w] = v
				fila = append(fila, w)
			}
		}
	}
	return pai, visitadas
}

// dfs visita v e, recursivamente, os vizinhos ainda não marcados, até chegar
// a e. Devolve true quando chegou.
func dfs(v, e int, pai []int, marcado []bool, visitadas *int) bool {
	marcado[v] = true
	*visitadas++
	if v == e {
		return true
	}
	for _, w := range vizinhos(v) {
		if !marcado[w] {
			pai[w] = v
			if dfs(w, e, pai, marcado, visitadas) {
				return true
			}
		}
	}
	return false
}

// imprime mostra o labirinto com o caminho de s até e marcado com pontos, e
// devolve o número de passos do caminho.
func imprime(pai []int, s, e int) int {
	tela := make([][]byte, linhas)
	for l := range mapa {
		tela[l] = []byte(mapa[l])
	}
	passos := 0
	for v := pai[e]; v != s; v = pai[v] {
		tela[v/colunas][v%colunas] = '.'
		passos++
	}
	for _, linha := range tela {
		fmt.Println("  " + string(linha))
	}
	return passos + 1
}

func main() {
	var s, e int
	for l := range mapa {
		for c := range mapa[l] {
			switch mapa[l][c] {
			case 'S':
				s = celula(l, c)
			case 'E':
				e = celula(l, c)
			}
		}
	}

	pai, visitadas := bfs(s, e)
	fmt.Println("Busca em largura")
	passos := imprime(pai, s, e)
	fmt.Printf("  caminho de %d passos, %d células visitadas\n\n", passos, visitadas)

	pai = make([]int, linhas*colunas)
	marcado := make([]bool, linhas*colunas)
	visitadas = 0
	dfs(s, e, pai, marcado, &visitadas)
	fmt.Println("Busca em profundidade")
	passos = imprime(pai, s, e)
	fmt.Printf("  caminho de %d passos, %d células visitadas\n", passos, visitadas)
}
