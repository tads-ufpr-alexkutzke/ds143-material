// Fatorial: versão recursiva e versão iterativa, para comparação.

package main

import "fmt"

// fat calcula n! recursivamente.
func fat(n int) int {
	if n == 1 {
		return 1
	}
	return n * fat(n-1)
}

// fat2 calcula n! iterativamente.
func fat2(n int) int {
	t := 1
	for i := 1; i <= n; i++ {
		t *= i
	}
	return t
}

func main() {
	var x int
	fmt.Scan(&x)

	fmt.Println(fat(x))
	fmt.Println(fat2(x))
}
