package main

import (
	"bufio"
	"math/rand"
	"strings"
	"testing"
)

// linhaDeTeste é um grafo linha cuja sequência de vértices é conhecida pelo
// teste, que calcula por ela a distância esperada de cada par.
type linhaDeTeste struct {
	g         *GrafoP
	posicao   []int
	acumulado []int
}

// novaLinha monta a linha que passa pelos vértices de ordem, com pesos[i] na
// aresta entre ordem[i] e ordem[i+1]. As arestas entram no grafo na sequência
// dada por entrada, que guarda índices de pesos.
func novaLinha(ordem, pesos, entrada []int) *linhaDeTeste {
	l := &linhaDeTeste{
		g:         NovoGrafoP(len(ordem)),
		posicao:   make([]int, len(ordem)),
		acumulado: make([]int, len(ordem)),
	}
	for i, v := range ordem {
		l.posicao[v] = i
		if i > 0 {
			l.acumulado[i] = l.acumulado[i-1] + pesos[i-1]
		}
	}
	for _, i := range entrada {
		l.g.AdicionaAresta(Aresta{V: ordem[i], W: ordem[i+1], Peso: pesos[i]})
	}
	return l
}

// esperada devolve a distância entre v e w.
func (l *linhaDeTeste) esperada(v, w int) int {
	d := l.acumulado[l.posicao[v]] - l.acumulado[l.posicao[w]]
	if d < 0 {
		return -d
	}
	return d
}

// linhaSorteada monta uma linha de n vértices com a numeração, os pesos e a
// sequência das arestas sorteados.
func linhaSorteada(n int, semente int64) *linhaDeTeste {
	r := rand.New(rand.NewSource(semente))
	pesos := make([]int, n-1)
	for i := range pesos {
		pesos[i] = 1 + r.Intn(1000)
	}
	return novaLinha(r.Perm(n), pesos, r.Perm(n-1))
}

// sequencia devolve os inteiros de 0 a n-1, em ordem.
func sequencia(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = i
	}
	return s
}

// invertida devolve os inteiros de n-1 a 0.
func invertida(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = n - 1 - i
	}
	return s
}

// confereTodosOsPares compara a distância de todos os pares de vértices.
func confereTodosOsPares(t *testing.T, l *linhaDeTeste) {
	t.Helper()
	p := Preprocessa(l.g)
	for v := 0; v < l.g.V(); v++ {
		for w := 0; w < l.g.V(); w++ {
			if d := p.Distancia(v, w); d != l.esperada(v, w) {
				t.Fatalf("Distancia(%d, %d) = %d, esperado %d", v, w, d, l.esperada(v, w))
			}
		}
	}
}

// O grafo de dados/tinyLinha.txt: 3 -5- 0 -2- 6 -7- 1 -1- 4 -9- 7 -3- 2 -4- 5.
const tinyLinha = `8
4 7 9
6 0 2
1 6 7
5 2 4
3 0 5
7 2 3
1 4 1
`

func TestExemploPequeno(t *testing.T) {
	g, err := leGrafo(bufio.NewScanner(strings.NewReader(tinyLinha)))
	if err != nil {
		t.Fatal(err)
	}
	p := Preprocessa(g)

	casos := []struct{ v, w, esperada int }{
		{3, 5, 31}, {0, 1, 9}, {7, 0, 19}, {5, 4, 16}, {6, 6, 0},
		{3, 0, 5}, {1, 4, 1}, {2, 5, 4}, {4, 3, 15}, {6, 2, 20},
	}
	for _, c := range casos {
		if d := p.Distancia(c.v, c.w); d != c.esperada {
			t.Errorf("Distancia(%d, %d) = %d, esperado %d", c.v, c.w, d, c.esperada)
		}
	}
}

func TestDoisVertices(t *testing.T) {
	confereTodosOsPares(t, novaLinha([]int{0, 1}, []int{7}, []int{0}))
	confereTodosOsPares(t, novaLinha([]int{1, 0}, []int{7}, []int{0}))
}

func TestPontasNosVerticesExtremos(t *testing.T) {
	pesos := []int{4, 1, 8, 2, 6, 3, 5}
	confereTodosOsPares(t, novaLinha(sequencia(8), pesos, sequencia(7)))
	confereTodosOsPares(t, novaLinha(invertida(8), pesos, sequencia(7)))
}

func TestVerticeZeroNoMeio(t *testing.T) {
	ordem := []int{4, 2, 6, 0, 5, 1, 3}
	confereTodosOsPares(t, novaLinha(ordem, []int{3, 9, 1, 7, 2, 4}, sequencia(6)))
}

func TestSequenciaDasArestasNaoMudaResposta(t *testing.T) {
	ordem := []int{5, 3, 0, 6, 2, 7, 1, 4}
	pesos := []int{2, 7, 1, 8, 2, 8, 1}
	confereTodosOsPares(t, novaLinha(ordem, pesos, sequencia(7)))
	confereTodosOsPares(t, novaLinha(ordem, pesos, invertida(7)))
	confereTodosOsPares(t, novaLinha(ordem, pesos, []int{3, 0, 6, 2, 5, 1, 4}))
}

func TestLinhasSorteadas(t *testing.T) {
	for n := 2; n <= 60; n++ {
		for semente := int64(1); semente <= 5; semente++ {
			confereTodosOsPares(t, linhaSorteada(n, 100*int64(n)+semente))
		}
	}
}

func TestDoisGrafosNoMesmoPrograma(t *testing.T) {
	a := linhaSorteada(50, 21)
	b := linhaSorteada(30, 22)
	pa := Preprocessa(a.g)
	pb := Preprocessa(b.g)
	for v := 0; v < 30; v++ {
		for w := 0; w < 30; w++ {
			if d := pa.Distancia(v, w); d != a.esperada(v, w) {
				t.Fatalf("depois de preprocessar outro grafo, Distancia(%d, %d) = %d no primeiro, esperado %d",
					v, w, d, a.esperada(v, w))
			}
			if d := pb.Distancia(v, w); d != b.esperada(v, w) {
				t.Fatalf("Distancia(%d, %d) = %d no segundo grafo, esperado %d", v, w, d, b.esperada(v, w))
			}
		}
	}
}

func TestProgramaCompleto(t *testing.T) {
	entrada := tinyLinha + "3 5\n0 1\n6 6\n9 1\n7 0\n-1 -1\n2 5\n"
	scanner := bufio.NewScanner(strings.NewReader(entrada))
	g, err := leGrafo(scanner)
	if err != nil {
		t.Fatal(err)
	}
	var saida, erros strings.Builder
	respondeConsultas(Preprocessa(g), g.V(), scanner, &saida, &erros)

	if esperada := "31\n9\n0\n19\n"; saida.String() != esperada {
		t.Errorf("saída %q, esperada %q", saida.String(), esperada)
	}
	if erros.String() == "" {
		t.Error("a consulta com o vértice 9, fora do grafo, deveria gerar um aviso")
	}
}

func TestEntradasInvalidas(t *testing.T) {
	casos := map[string]string{
		"sem número de vértices": "",
		"um vértice só":          "1\n",
		"aresta faltando":        "3\n0 1 5\n",
		"vértice fora do grafo":  "2\n0 2 5\n",
		"vértice negativo":       "2\n-1 1 5\n",
		"peso zero":              "2\n0 1 0\n",
		"peso ausente":           "2\n0 1\n",
	}
	for nome, entrada := range casos {
		if _, err := leGrafo(bufio.NewScanner(strings.NewReader(entrada))); err == nil {
			t.Errorf("%s: leGrafo deveria devolver erro", nome)
		}
	}
}
