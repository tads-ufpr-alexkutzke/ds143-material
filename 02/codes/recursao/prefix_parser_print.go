// Parser recursivo de expressões em notação prefixa, imprimindo a
// árvore de chamadas recursivas (indentada por profundidade).

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var expr string
var pos int

func eval(tabs int) int {
	x := 0

	for pos < len(expr) && expr[pos] == ' ' {
		pos++
	}

	fmt.Printf("%seval()\n", strings.Repeat("  ", tabs))

	if pos < len(expr) && expr[pos] == '+' {
		pos++
		return eval(tabs+1) + eval(tabs+1)
	}
	if pos < len(expr) && expr[pos] == '*' {
		pos++
		return eval(tabs+1) * eval(tabs+1)
	}
	for pos < len(expr) && expr[pos] >= '0' && expr[pos] <= '9' {
		x = 10*x + int(expr[pos]-'0')
		pos++
	}
	return x
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	linha, _ := reader.ReadString('\n')
	expr = linha
	pos = 0

	fmt.Println(eval(0))
}
