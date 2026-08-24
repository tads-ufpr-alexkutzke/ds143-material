// Fibonacci recursivo "ingênuo" (sem memoização), contando o número de
// chamadas recursivas para ilustrar o crescimento exponencial de custo.

package main

import "fmt"

var cont int64

func fib(n int64) int64 {
	cont++
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	var x int64
	fmt.Scan(&x)

	fmt.Println(fib(x))
	fmt.Println(cont)
}
