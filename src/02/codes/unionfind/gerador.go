// gerador.go: gera uma entrada aleatória para os programas Union-Find.
// Uso: go run gerador.go <n> <pares> > entrada.txt

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "uso: go run gerador.go <n> <pares>")
		os.Exit(1)
	}
	n, _ := strconv.Atoi(os.Args[1])
	pares, _ := strconv.Atoi(os.Args[2])

	saida := bufio.NewWriter(os.Stdout)
	defer saida.Flush()

	fmt.Fprintln(saida, n)
	for i := 0; i < pares; i++ {
		fmt.Fprintln(saida, rand.Intn(n), rand.Intn(n))
	}
	fmt.Fprintln(saida, "-1 -1")
}
