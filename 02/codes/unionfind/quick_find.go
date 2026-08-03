// Quick Find para Union-Find
// Prof. Alex Kutzke
//
// Baseado no exemplo dado no livro Algorithms, 4th Edition
// de Robert Sedgewick e Kevin Wayne.
// Versão original em Java: https://algs4.cs.princeton.edu/15uf/QuickFindUF.java

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// UF representa uma estrutura Union-Find (conectividade dinâmica)
// usando a estratégia Quick-Find.
// id[i] armazena a componente à qual o elemento i pertence.
type UF struct {
	id    []int
	n     int
	count int
}

// NewUF inicializa N itens (0 até N-1), cada um em sua própria componente.
func NewUF(n int) *UF {
	uf := &UF{
		id:    make([]int, n),
		n:     n,
		count: n,
	}
	for i := 0; i < n; i++ {
		uf.id[i] = i
	}
	return uf
}

// Count retorna o número atual de componentes.
func (uf *UF) Count() int {
	return uf.count
}

// Connected retorna true se p e q pertencem à mesma componente.
// Custo: O(1).
func (uf *UF) Connected(p, q int) bool {
	return uf.id[p] == uf.id[q]
}

// Find retorna o identificador da componente do elemento p.
// Custo: O(1).
func (uf *UF) Find(p int) int {
	return uf.id[p]
}

// Union conecta os elementos p e q, unindo suas componentes.
// Custo: O(n), pois é preciso percorrer todo o slice id para
// renomear a componente de p para a componente de q.
func (uf *UF) Union(p, q int) {
	pID := uf.id[p]
	qID := uf.id[q]

	if pID == qID {
		return
	}

	for i := 0; i < uf.n; i++ {
		if uf.id[i] == pID {
			uf.id[i] = qID
		}
	}
	uf.count--
}

// readInts lê dois inteiros separados por espaço de uma linha do scanner.
// Retorna os inteiros e true se a leitura foi bem-sucedida, ou 0, 0, false
// em caso de erro ou fim de arquivo.
func readInts(scanner *bufio.Scanner) (int, int, bool) {
	if !scanner.Scan() {
		return 0, 0, false
	}
	line := strings.TrimSpace(scanner.Text())
	parts := strings.Fields(line)
	if len(parts) < 2 {
		return 0, 0, false
	}
	p, err1 := strconv.Atoi(parts[0])
	q, err2 := strconv.Atoi(parts[1])
	if err1 != nil || err2 != nil {
		return 0, 0, false
	}
	return p, q, true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Lê o número de elementos (primeira linha: um único inteiro)
	if !scanner.Scan() {
		fmt.Fprintln(os.Stderr, "Erro: arquivo vazio")
		os.Exit(1)
	}
	n, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro: número inválido na primeira linha: %v\n", err)
		os.Exit(1)
	}

	uf := NewUF(n)

	// Lê pares p q até encontrar números negativos ou EOF
	for {
		p, q, ok := readInts(scanner)
		if !ok {
			break
		}
		if p < 0 || q < 0 {
			break
		}

		if !uf.Connected(p, q) {
			fmt.Printf("%d %d\n", p, q)
			uf.Union(p, q)
		}
	}
}
