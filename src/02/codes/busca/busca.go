// Busca sequencial, busca binária (iterativa) e busca binária (recursiva)
// em um vetor ordenado de 1.000.000 de elementos, para comparar o número
// de comparações realizado por cada estratégia.

package main

import "fmt"

const n = 1_000_000

func preencheVetor(n int) []int {
	v := make([]int, n)
	for i, j := 0, 0; j < n; i += 2 {
		v[j] = i
		j++
	}
	return v
}

// buscaSequencial percorre o vetor do início ao fim.
// Custo: O(n) no pior caso.
func buscaSequencial(v []int, elem int) int {
	i := 0
	for ; i < len(v); i++ {
		if v[i] == elem {
			fmt.Printf("%d comparações (sequencial)\n", i+1)
			return i
		}
	}
	fmt.Printf("%d comparações (sequencial)\n", i)
	return -1
}

// buscaBinaria divide o intervalo de busca pela metade a cada passo.
// Custo: O(log n) no pior caso.
func buscaBinaria(v []int, elem int) int {
	l, r := 0, len(v)-1
	comp := 0

	for l <= r {
		m := (l + r) / 2
		comp++
		if v[m] == elem {
			fmt.Printf("%d comparações (binária iterativa)\n", comp)
			return m
		} else if v[m] < elem {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	fmt.Printf("%d comparações (binária iterativa)\n", comp)
	return -1
}

// buscaBinariaRecursiva é a mesma estratégia da busca binária, escrita
// recursivamente em vez de com um laço.
func buscaBinariaRecursiva(v []int, elem, l, r int) int {
	if l > r {
		return -1
	}

	m := (l + r) / 2
	if v[m] == elem {
		return m
	} else if v[m] < elem {
		return buscaBinariaRecursiva(v, elem, m+1, r)
	}
	return buscaBinariaRecursiva(v, elem, l, m-1)
}

func imprimeResultado(nome string, elem, r int) {
	if r != -1 {
		fmt.Printf("[%s] O elemento %d está na posição %d do vetor.\n", nome, elem, r)
	} else {
		fmt.Printf("[%s] O elemento %d não pertence ao vetor.\n", nome, elem)
	}
}

func main() {
	v := preencheVetor(n)

	var x int
	fmt.Scan(&x)

	imprimeResultado("sequencial", x, buscaSequencial(v, x))
	imprimeResultado("binária iterativa", x, buscaBinaria(v, x))
	imprimeResultado("binária recursiva", x, buscaBinariaRecursiva(v, x, 0, n-1))
}
