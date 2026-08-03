// "Puzzle" da Conjectura de Collatz, imprimindo a árvore de chamadas
// recursivas (indentada por profundidade).

package main

import (
	"fmt"
	"strings"
)

func puzzlePrint(n, tabs int) int {
	fmt.Printf("%spuzzle(%d)\n", strings.Repeat("  ", tabs), n)

	if n == 1 {
		return 1
	}
	if n%2 == 0 {
		return puzzlePrint(n/2, tabs+1)
	}
	return puzzlePrint(3*n+1, tabs+1)
}

func main() {
	var x int
	fmt.Scan(&x)

	fmt.Println(puzzlePrint(x, 0))
}
