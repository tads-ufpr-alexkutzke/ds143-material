// Quick Union para Union-Find
// Prof. Alex Kutzke
//
// Baseado no exemplo dado no livro Algorithms, 4th Edition
// de Robert Sedgewick e Kevin Wayne.
// Versão original em Java: https://algs4.cs.princeton.edu/15uf/QuickUnionUF.java

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// UF representa uma estrutura Union-Find (conectividade dinâmica)
// usando a estratégia Quick-Union.
// parent[i] armazena o pai do elemento i na árvore da sua componente;
// a raiz de uma árvore é o elemento que é pai de si mesmo.
type UF struct {
	parent []int
	n      int
	count  int
}

// NewUF inicializa N itens (0 até N-1), cada um em sua própria componente.
func NewUF(n int) *UF {
	uf := &UF{
		parent: make([]int, n),
		n:      n,
		count:  n,
	}
	for i := 0; i < n; i++ {
		uf.parent[i] = i
	}
	return uf
}

// Count retorna o número atual de componentes.
func (uf *UF) Count() int {
	return uf.count
}

// Connected retorna true se p e q pertencem à mesma componente.
// Custo: proporcional à altura da árvore.
func (uf *UF) Connected(p, q int) bool {
	return uf.Find(p) == uf.Find(q)
}

// Find retorna a raiz da árvore que contém o elemento p, subindo
// pela cadeia de pais até encontrar um elemento que é pai de si mesmo.
// Custo: proporcional à altura da árvore (pode degenerar para O(n)).
func (uf *UF) Find(p int) int {
	for p != uf.parent[p] {
		p = uf.parent[p]
	}
	return p
}

// Union conecta os elementos p e q, unindo suas componentes: a raiz
// da árvore de p passa a apontar para a raiz da árvore de q.
func (uf *UF) Union(p, q int) {
	rootP := uf.Find(p)
	rootQ := uf.Find(q)

	if rootP == rootQ {
		return
	}

	uf.parent[rootP] = rootQ
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
