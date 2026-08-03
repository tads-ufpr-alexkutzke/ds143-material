// Máximo Divisor Comum (MDC) via Algoritmo de Euclides, recursivo.

package main

import "fmt"

func mdc(m, n int) int {
	if n == 0 {
		return m
	}
	return mdc(n, m%n)
}

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	fmt.Println(mdc(x, y))
}
