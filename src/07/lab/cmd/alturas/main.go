// Programa da Parte 3: compara a altura da árvore binária de busca com a
// da AVL, sobre três sequências de inserção diferentes.
//
// Cada tabela traz duas colunas para a AVL: o valor guardado no campo
// Alt da raiz e a altura obtida percorrendo a árvore. Em uma AVL
// consistente as duas são iguais.
//
// Este arquivo está pronto. Ele usa a RotacaoEsquerda e a VerificaAVL
// que você escreveu, então só roda depois da Parte 2.
//
//	go build -o alturas ./cmd/alturas
//	./alturas
package main

import (
	"fmt"
	"math/rand"

	"ds143/atividade2/abb"
	"ds143/atividade2/avl"
)

var tamanhos = []int{10, 100, 1000, 10000}

func cabecalho(titulo string) {
	fmt.Println()
	fmt.Println(titulo)
	fmt.Println()
	fmt.Println("       n | altura ABB | AVL campo Alt | AVL recalculada")
	fmt.Println("---------+------------+---------------+----------------")
}

func linha(n int, t1 *abb.No, t2 *avl.No) {
	fmt.Printf("%8d | %10d | %13d | %15d\n",
		n, abb.Altura(t1), avl.AlturaArmazenada(t2), avl.AlturaRecalculada(t2))
}

// sequencia devolve os n valores na ordem em que devem ser inseridos.
func sequencia(ordem string, n int) []int {
	valores := make([]int, 0, n)
	switch ordem {
	case "crescente":
		for v := 1; v <= n; v++ {
			valores = append(valores, v)
		}
	case "decrescente":
		for v := n; v >= 1; v-- {
			valores = append(valores, v)
		}
	case "aleatoria":
		aleatorio := rand.New(rand.NewSource(42))
		for _, v := range aleatorio.Perm(n) {
			valores = append(valores, v+1)
		}
	}
	return valores
}

func constroi(ordem string, n int) (*abb.No, *avl.No) {
	var t1 *abb.No
	var t2 *avl.No
	for _, v := range sequencia(ordem, n) {
		t1 = abb.Insere(t1, v)
		t2 = avl.Insere(t2, v)
	}
	return t1, t2
}

func main() {
	cabecalho("Tabela 1: inserindo 1, 2, 3, ..., n (ordem crescente)")
	for _, n := range tamanhos {
		t1, t2 := constroi("crescente", n)
		linha(n, t1, t2)
	}

	cabecalho("Tabela 2: inserindo n, n-1, ..., 2, 1 (ordem decrescente)")
	for _, n := range tamanhos {
		t1, t2 := constroi("decrescente", n)
		linha(n, t1, t2)
	}

	cabecalho("Tabela 3: inserindo os mesmos n valores em ordem aleatória\n(semente fixa, a mesma permutação nas duas estruturas)")
	for _, n := range tamanhos {
		t1, t2 := constroi("aleatoria", n)
		linha(n, t1, t2)
	}

	fmt.Println()
	fmt.Println("Verificação das três AVLs com n = 1000:")
	for _, ordem := range []string{"crescente", "decrescente", "aleatoria"} {
		_, t := constroi(ordem, 1000)
		ok, culpado := avl.VerificaAVL(t)
		if ok {
			fmt.Printf("  %-12s passou\n", ordem)
			continue
		}
		fmt.Printf("  %-12s falhou no nó %d: campo Alt %d, altura recalculada %d, fator %d\n",
			ordem, culpado.Info, culpado.Alt, avl.AlturaRecalculada(culpado), avl.Fator(culpado))
	}
}
