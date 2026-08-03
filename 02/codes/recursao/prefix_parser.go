// Parser (e avaliador) recursivo de expressões em notação prefixa.
// Exemplo de entrada: "+ 3 4" avalia para 7; "* 2 + 3 4" avalia para 14.

package main

import (
	"bufio"
	"fmt"
	"os"
)

var expr string
var pos int

func eval() int {
	x := 0

	for pos < len(expr) && expr[pos] == ' ' {
		pos++
	}

	if pos < len(expr) && expr[pos] == '+' {
		pos++
		return eval() + eval()
	}
	if pos < len(expr) && expr[pos] == '*' {
		pos++
		return eval() * eval()
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

	fmt.Println(eval())
}
