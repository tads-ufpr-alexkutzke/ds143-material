// Programa de medição empírica dos algoritmos de ordenação.
//
// Cada linha é um algoritmo, cada par de colunas é um tamanho n com o tempo
// em segundos e a razão em relação ao tamanho anterior. Todos os algoritmos
// recebem exatamente o mesmo vetor de entrada, gerado com semente fixa.
//
// Uso:
//
//	go build -o benchmark ./cmd/benchmark
//	./benchmark              # entrada aleatória
//	./benchmark -ordenada    # entrada já ordenada
package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sort"
	"time"

	"ds143/aula04/algoritmos"
)

// pequenos são os tamanhos usados na tabela que inclui os algoritmos O(n²).
var pequenos = []int{5000, 10000, 20000, 40000}

// grandes são os tamanhos usados na tabela só com os algoritmos O(n log n).
var grandes = []int{125000, 250000, 500000, 1000000}

type algoritmo struct {
	nome    string
	ordenar func([]int)
	rapido  bool // entra também na tabela dos tamanhos grandes
}

var catalogo = []algoritmo{
	{"SelectionSort", algoritmos.SelectionSort, false},
	{"InsertionSort", algoritmos.InsertionSort, false},
	{"MergeSort", algoritmos.MergeSort, true},
	{"QuickSort", algoritmos.QuickSort, true},
	{"QuickSort emb.", algoritmos.QuickSortEmbaralhado, true},
	{"sort.Ints", sort.Ints, true},
}

// gerar produz o vetor de entrada com n elementos. Com ordenada verdadeiro, o
// vetor já vem em ordem crescente.
func gerar(n int, ordenada bool) []int {
	v := make([]int, n)
	if ordenada {
		for i := range v {
			v[i] = i
		}
		return v
	}
	r := rand.New(rand.NewSource(42))
	for i := range v {
		v[i] = r.Intn(10 * n)
	}
	return v
}

// medir cronometra uma ordenação sobre uma cópia de v e confere o resultado.
// Devolve tempo negativo quando o vetor não termina ordenado.
//
// A ordenação é repetida até acumular pelo menos 200 ms, e o tempo devolvido é
// o menor dos tempos observados. Sem isso, os tamanhos pequenos dariam medidas
// dominadas por ruído do sistema operacional e do coletor de lixo.
func medir(ordenar func([]int), v []int) time.Duration {
	melhor := time.Duration(-1)
	acumulado := time.Duration(0)

	for repeticao := 0; repeticao < 50; repeticao++ {
		c := append([]int(nil), v...)

		inicio := time.Now()
		ordenar(c)
		gasto := time.Since(inicio)

		if !sort.IntsAreSorted(c) {
			return -1
		}
		if melhor < 0 || gasto < melhor {
			melhor = gasto
		}

		acumulado += gasto
		if acumulado >= 200*time.Millisecond {
			break
		}
	}

	return melhor
}

func tabela(titulo string, tamanhos []int, escolhidos []algoritmo, ordenada bool) {
	fmt.Printf("\n%s\n\n", titulo)

	fmt.Printf("%-16s", "algoritmo")
	for _, n := range tamanhos {
		fmt.Printf("%9d%6s", n, "razao")
	}
	fmt.Println()

	entradas := make([][]int, len(tamanhos))
	for i, n := range tamanhos {
		entradas[i] = gerar(n, ordenada)
	}

	for _, a := range escolhidos {
		fmt.Printf("%-16s", a.nome)
		var anterior time.Duration
		for i := range tamanhos {
			gasto := medir(a.ordenar, entradas[i])
			if gasto < 0 {
				// Uma vez que o algoritmo erra, os tamanhos maiores são
				// pulados: um QuickSort com particionamento errado degenera
				// para tempo quadrático e levaria minutos nos maiores.
				for range tamanhos[i:] {
					fmt.Printf("%9s%6s", "ERRO", "")
				}
				break
			}
			fmt.Printf("%9.4f", gasto.Seconds())
			if anterior > 0 {
				fmt.Printf("%6.1f", gasto.Seconds()/anterior.Seconds())
			} else {
				fmt.Printf("%6s", "")
			}
			anterior = gasto
		}
		fmt.Println()
	}
}

func main() {
	ordenada := flag.Bool("ordenada", false, "usa vetor de entrada já ordenado")
	flag.Parse()

	entrada := "aleatoria"
	if *ordenada {
		entrada = "ja ordenada"
	}
	fmt.Printf("Tempos em segundos. Entrada %s. ERRO significa vetor nao ordenado ao final.\n", entrada)

	tabela("Tabela 1: todos os algoritmos", pequenos, catalogo, *ordenada)

	if *ordenada {
		fmt.Println("\nA tabela 2 nao roda com -ordenada: o QuickSort sem embaralhamento")
		fmt.Println("levaria tempo demais nesses tamanhos. Explique por que.")
		return
	}

	var rapidos []algoritmo
	for _, a := range catalogo {
		if a.rapido {
			rapidos = append(rapidos, a)
		}
	}
	tabela("Tabela 2: apenas os algoritmos O(n log n)", grandes, rapidos, *ordenada)
}
