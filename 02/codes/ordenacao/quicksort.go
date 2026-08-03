// QuickSort com particionamento de Hoare/Sedgewick, embaralhando o
// vetor de entrada antes de ordenar (para evitar o pior caso O(n²) em
// entradas já ordenadas ou adversariais).

package main

import (
	"fmt"
	"math/rand"
)

func shuffle(a []int, r *rand.Rand) {
	n := len(a)
	for i := 0; i < n-1; i++ {
		j := i + r.Intn(n-i)
		a[i], a[j] = a[j], a[i]
	}
}

func swap(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}

func partition(a []int, lo, hi int) int {
	i, j := lo, hi+1
	for {
		for {
			i++
			if a[i] >= a[lo] || i == hi {
				break
			}
		}
		for {
			j--
			if a[lo] >= a[j] || j == lo {
				break
			}
		}
		if i >= j {
			break
		}
		swap(a, i, j)
	}

	swap(a, lo, j)
	return j
}

func quicksortRange(a []int, lo, hi int) {
	if hi <= lo {
		return
	}
	j := partition(a, lo, hi)
	quicksortRange(a, lo, j-1)
	quicksortRange(a, j+1, hi)
}

func quicksort(a []int, r *rand.Rand) {
	shuffle(a, r)
	quicksortRange(a, 0, len(a)-1)
}

func main() {
	a := []int{5, 2, 8, 4, 10, 9, 7, 6, 3, 1}
	r := rand.New(rand.NewSource(1))

	quicksort(a, r)
	for _, x := range a {
		fmt.Printf("%d, ", x)
	}
	fmt.Println()
}
