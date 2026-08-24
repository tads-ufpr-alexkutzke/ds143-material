// Comparação dos três algoritmos de ordenação O(n²) clássicos: bubble
// sort, insertion sort e selection sort.
//
// Uso: go run sort_basico.go <selection|insertion|bubble> <n> [print]

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func bubbleSort(v []int) {
	n := len(v)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if v[j] > v[j+1] {
				v[j], v[j+1] = v[j+1], v[j]
			}
		}
	}
}

func insertionSort(v []int) {
	for i := 1; i < len(v); i++ {
		x := v[i]
		j := i
		for j > 0 && x < v[j-1] {
			v[j] = v[j-1]
			j--
		}
		v[j] = x
	}
}

func selectionSort(v []int) {
	n := len(v)
	for i := 0; i < n-1; i++ {
		m := i
		for j := i + 1; j < n; j++ {
			if v[j] < v[m] {
				m = j
			}
		}
		v[m], v[i] = v[i], v[m]
	}
}

func printV(v []int) {
	for _, x := range v {
		fmt.Printf("%d ", x)
	}
	fmt.Println()
}

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "uso: sort_basico <selection|insertion|bubble> <n> [print]")
		os.Exit(1)
	}

	alg := os.Args[1]
	n, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "n inválido: %v\n", err)
		os.Exit(1)
	}
	print := len(os.Args) > 3

	r := rand.New(rand.NewSource(1))
	v := make([]int, n)
	for i := range v {
		v[i] = r.Intn(n)
	}

	if print {
		printV(v)
	}

	switch alg {
	case "selection":
		selectionSort(v)
	case "insertion":
		insertionSort(v)
	case "bubble":
		bubbleSort(v)
	default:
		fmt.Fprintf(os.Stderr, "algoritmo desconhecido: %s\n", alg)
		os.Exit(1)
	}

	if print {
		printV(v)
	}

	for i := n - 1; i > 0; i-- {
		if v[i] < v[i-1] {
			fmt.Println("Error!")
			os.Exit(1)
		}
	}
}
