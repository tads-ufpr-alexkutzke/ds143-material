package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// UF representa a estrutura Union-Find para conectividade dinâmica.
// id[i] armazena o identificador da componente à qual o elemento i pertence.
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
func (uf *UF) Connected(p, q int) bool {
	return uf.Find(p) == uf.Find(q)
}

// Find retorna o identificador da componente do elemento p.
// ATENÇÃO: esta implementação está incompleta — é o desafio de vocês!
func (uf *UF) Find(p int) int {
	// TODO: implementar uma estratégia de busca da componente
	return 0
}

// Union conecta os elementos p e q, unindo suas componentes.
// ATENÇÃO: esta implementação está incompleta — é o desafio de vocês!
func (uf *UF) Union(p, q int) {
	// TODO: implementar uma estratégia de união entre componentes
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
