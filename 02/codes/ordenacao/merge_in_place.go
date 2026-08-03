// Intercalação (merge) "in place": recebe um único vetor a, com as duas
// metades a[l..m] e a[m+1..r] já ordenadas, e as intercala usando um
// vetor auxiliar em que a metade direita é copiada de trás para frente
// — truque clássico (Sedgewick) que evita checagem explícita de limites
// ao percorrer as duas metades a partir das pontas.
//
// Entrada: N (tamanho da primeira metade), M (tamanho da segunda metade),
// seguidos pelos N+M inteiros já concatenados (metade esquerda ordenada,
// seguida da metade direita ordenada).

package main

import "fmt"

func mergeInPlace(a []int, l, m, r int) {
	aux := make([]int, r+1)

	for i := m + 1; i > l; i-- {
		aux[i-1] = a[i-1]
	}
	for j := m; j < r; j++ {
		aux[r+m-j] = a[j+1]
	}

	i, j := l, r
	for k := l; k <= r; k++ {
		if aux[j] < aux[i] {
			a[k] = aux[j]
			j--
		} else {
			a[k] = aux[i]
			i++
		}
	}
}

func main() {
	var n, m int
	fmt.Scan(&n)
	fmt.Scan(&m)

	a := make([]int, n+m)
	for i := range a {
		fmt.Scan(&a[i])
	}

	mergeInPlace(a, 0, n-1, n+m-1)
	for _, x := range a {
		fmt.Printf("%d ", x)
	}
	fmt.Println()
}
