// Fibonacci recursivo "ingênuo", imprimindo a árvore de chamadas
// recursivas (indentada por profundidade) para visualizar a repetição
// de subproblemas.

package main

import (
	"fmt"
	"strings"
)

func fibPrint(n, tabs int) int {
	fmt.Printf("%sfib(%d)\n", strings.Repeat("  ", tabs), n)

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibPrint(n-1, tabs+1) + fibPrint(n-2, tabs+1)
}

func main() {
	var x int
	fmt.Scan(&x)

	fmt.Println(fibPrint(x, 0))
}
