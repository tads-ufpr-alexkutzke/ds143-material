// "Puzzle" da Conjectura de Collatz: aplica repetidamente (via recursão)
// n/2 se n for par, ou 3n+1 se n for ímpar, até alcançar 1.

package main

import "fmt"

func puzzle(n int) int {
	if n == 1 {
		return 1
	}
	if n%2 == 0 {
		return puzzle(n / 2)
	}
	return puzzle(3*n + 1)
}

func main() {
	var x int
	fmt.Scan(&x)

	fmt.Println(puzzle(x))
}
