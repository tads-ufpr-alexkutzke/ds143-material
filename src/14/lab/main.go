// Programa que lê um grafo linha com pesos e responde à distância entre pares
// de vértices. O grafo e as consultas têm este formato:
//
//	n              número de vértices (pelo menos 2)
//	v w peso       uma linha para cada uma das n-1 arestas
//	v w            uma linha para cada consulta
//
// Os vértices são numerados de 0 a n-1, e os pesos são inteiros positivos. As
// consultas terminam no fim da entrada ou em um par com número negativo. Cada
// resposta é uma linha com a distância.
//
// Grafo em arquivo e consultas digitadas (no fim, Ctrl+D ou um par negativo):
//
//	go run . dados/tinyLinha.txt
//
// Grafo em arquivo e consultas em outro arquivo:
//
//	go run . dados/tinyLinha.txt < dados/tinyConsultas.txt
//
// Sem argumento, o grafo e as consultas vêm da entrada padrão:
//
//	go run . < dados/tinyTudo.txt
package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// leInts lê uma linha com qtd inteiros separados por espaço. Devolve false em
// caso de erro ou de fim de arquivo.
func leInts(scanner *bufio.Scanner, qtd int) ([]int, bool) {
	if !scanner.Scan() {
		return nil, false
	}
	partes := strings.Fields(scanner.Text())
	if len(partes) != qtd {
		return nil, false
	}
	numeros := make([]int, qtd)
	for i, parte := range partes {
		numero, err := strconv.Atoi(parte)
		if err != nil {
			return nil, false
		}
		numeros[i] = numero
	}
	return numeros, true
}

// leGrafo lê o número de vértices e as n-1 arestas. Confere o formato e os
// limites dos números, e supõe que as arestas de fato formam uma linha.
func leGrafo(scanner *bufio.Scanner) (*GrafoP, error) {
	primeira, ok := leInts(scanner, 1)
	if !ok {
		return nil, errors.New("a primeira linha deve ter o número de vértices")
	}
	n := primeira[0]
	if n < 2 {
		return nil, fmt.Errorf("o grafo precisa de pelo menos 2 vértices, e não %d", n)
	}

	g := NovoGrafoP(n)
	for i := 1; i < n; i++ {
		aresta, ok := leInts(scanner, 3)
		if !ok {
			return nil, fmt.Errorf("a aresta %d de %d deveria ter o formato v w peso", i, n-1)
		}
		v, w, peso := aresta[0], aresta[1], aresta[2]
		if v < 0 || v >= n || w < 0 || w >= n {
			return nil, fmt.Errorf("a aresta %d de %d tem vértice fora do intervalo de 0 a %d", i, n-1, n-1)
		}
		if peso <= 0 {
			return nil, fmt.Errorf("a aresta %d de %d tem peso %d, que não é positivo", i, n-1, peso)
		}
		g.AdicionaAresta(Aresta{V: v, W: w, Peso: peso})
	}
	return g, nil
}

// respondeConsultas lê pares de vértices e escreve a distância de cada par,
// até o fim da entrada, uma linha fora do formato ou um par com número
// negativo. Um vértice maior que o último da linha é avisado em erros, e a
// leitura continua.
func respondeConsultas(p *Preprocessamento, n int, scanner *bufio.Scanner, saida, erros io.Writer) {
	for {
		par, ok := leInts(scanner, 2)
		if !ok {
			return
		}
		v, w := par[0], par[1]
		if v < 0 || w < 0 {
			return
		}
		if v >= n || w >= n {
			fmt.Fprintf(erros, "vértice fora do intervalo de 0 a %d\n", n-1)
			continue
		}
		fmt.Fprintln(saida, p.Distancia(v, w))
	}
}

func main() {
	if len(os.Args) > 2 {
		fmt.Fprintln(os.Stderr, "uso: go run . [arquivo do grafo]")
		os.Exit(1)
	}

	consultas := bufio.NewScanner(os.Stdin)
	entradaGrafo := consultas
	if len(os.Args) == 2 {
		arquivo, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro: %v\n", err)
			os.Exit(1)
		}
		defer arquivo.Close()
		entradaGrafo = bufio.NewScanner(arquivo)
	}

	g, err := leGrafo(entradaGrafo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro: %v\n", err)
		os.Exit(1)
	}

	p := Preprocessa(g)
	respondeConsultas(p, g.V(), consultas, os.Stdout, os.Stderr)
}
