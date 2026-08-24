// Programa que compara duas estratégias para responder a k consultas de busca
// em um vetor de n elementos:
//
//	A: k buscas sequenciais no vetor como ele veio, sem ordenar;
//	B: ordenar o vetor uma vez e fazer k buscas binárias.
//
// O tempo da estratégia B inclui o tempo de ordenação. A ordenação usa
// sort.Ints, para que a comparação não dependa do MergeSort da pasta
// algoritmos. A busca binária é a sua, e precisa estar correta.
//
// Uso:
//
//	go build -o consultas ./cmd/consultas
//	./consultas
package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"

	"ds143/aula04/algoritmos"
)

const n = 200000

var quantidades = []int{1, 10, 100, 500, 1000, 5000, 10000}

func main() {
	r := rand.New(rand.NewSource(42))

	original := make([]int, n)
	for i := range original {
		original[i] = r.Intn(10 * n)
	}

	ordenado := append([]int(nil), original...)
	inicio := time.Now()
	sort.Ints(ordenado)
	tempoOrdenacao := time.Since(inicio)

	fmt.Printf("n = %d elementos\n", n)
	fmt.Printf("ordenar o vetor uma vez: %.4f s\n\n", tempoOrdenacao.Seconds())

	fmt.Printf("%9s%14s%14s%14s\n", "k", "A: seq.", "B: ord+bin", "B - ordenar")
	for _, k := range quantidades {
		consultas := make([]int, k)
		for i := range consultas {
			consultas[i] = original[r.Intn(n)]
		}

		inicio = time.Now()
		for _, x := range consultas {
			if algoritmos.BuscaSequencial(original, x) < 0 {
				fmt.Println("erro: busca sequencial nao encontrou um valor presente")
				return
			}
		}
		tempoA := time.Since(inicio)

		inicio = time.Now()
		for _, x := range consultas {
			if algoritmos.BuscaBinaria(ordenado, x) < 0 {
				fmt.Println("erro: busca binaria nao encontrou um valor presente. Corrija-a antes de medir.")
				return
			}
		}
		tempoBuscas := time.Since(inicio)
		tempoB := tempoOrdenacao + tempoBuscas

		fmt.Printf("%9d%14.4f%14.4f%14.4f\n", k, tempoA.Seconds(), tempoB.Seconds(), tempoBuscas.Seconds())
	}
}
