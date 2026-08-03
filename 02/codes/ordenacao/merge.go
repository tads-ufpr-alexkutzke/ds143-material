// Intercalação (merge) de dois vetores já ordenados em um terceiro vetor.
// É a operação básica por trás do MergeSort.
//
// Entrada: N, os N inteiros de a, M, os M inteiros de b (a e b já ordenados).

package main

import "fmt"

func merge(a []int, b []int) []int {
	n, m := len(a), len(b)
	c := make([]int, n+m)
	i, j := 0, 0

	for k := 0; k < n+m; k++ {
		switch {
		case i == n:
			c[k] = b[j]
			j++
		case j == m:
			c[k] = a[i]
			i++
		case a[i] < b[j]:
			c[k] = a[i]
			i++
		default:
			c[k] = b[j]
			j++
		}
	}
	return c
}

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	var m int
	fmt.Scan(&m)
	b := make([]int, m)
	for i := range b {
		fmt.Scan(&b[i])
	}

	c := merge(a, b)
	for _, x := range c {
		fmt.Printf("%d ", x)
	}
	fmt.Println()
}
