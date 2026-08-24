// Fibonacci com memoização (programação dinâmica): recursivo como o
// "ingênuo", mas guardando resultados já calculados para evitar
// recomputar os mesmos subproblemas — custo O(n).

package main

import "fmt"

var knownF [100]int64

func fib(n int64) int64 {
	if knownF[n] != 0 {
		return knownF[n]
	}
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	t := fib(n-1) + fib(n-2)
	knownF[n] = t
	return t
}

func main() {
	var x int64
	fmt.Scan(&x)

	fmt.Println(fib(x))
}
