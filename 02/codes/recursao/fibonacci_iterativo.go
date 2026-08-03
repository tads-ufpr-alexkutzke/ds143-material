// Fibonacci iterativo: mesmo resultado do recursivo, mas em tempo O(n)
// e sem o custo de pilha de chamadas.

package main

import "fmt"

func fib(n int64) int64 {
	var first, second, next int64 = 1, 1, 0

	for c := int64(0); c < n; c++ {
		if c <= 1 {
			next = c
		} else {
			next = first + second
			first = second
			second = next
		}
	}

	return next
}

func main() {
	var x int64
	fmt.Scan(&x)

	fmt.Println(fib(x))
}
