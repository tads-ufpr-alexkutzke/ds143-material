// Máximo Divisor Comum (MDC) via Algoritmo de Euclides, recursivo,
// imprimindo a árvore de chamadas recursivas (indentada por profundidade).

package main

import (
	"fmt"
	"strings"
)

func mdcPrint(m, n, tabs int) int {
	fmt.Printf("%smdc(%d,%d)\n", strings.Repeat("  ", tabs), m, n)

	if n == 0 {
		return m
	}
	return mdcPrint(n, m%n, tabs+1)
}

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	fmt.Println(mdcPrint(x, y, 0))
}
