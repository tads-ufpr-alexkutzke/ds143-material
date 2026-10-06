// Demonstração do início da aula de tabelas de dispersão: o map de Go não
// guarda as chaves em ordem, e a busca nele não fica mais lenta quando o
// número de chaves cresce.
//
// Execute com:  go run mapa.go

package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// buscaSequencial percorre o vetor inteiro no pior caso: Θ(n).
func buscaSequencial(v []int, chave int) bool {
	for _, x := range v {
		if x == chave {
			return true
		}
	}
	return false
}

// buscaBinaria exige o vetor ordenado e custa Θ(log n), como a busca em uma
// árvore balanceada.
func buscaBinaria(v []int, chave int) bool {
	ini, fim := 0, len(v)-1
	for ini <= fim {
		meio := (ini + fim) / 2
		if v[meio] == chave {
			return true
		} else if v[meio] < chave {
			ini = meio + 1
		} else {
			fim = meio - 1
		}
	}
	return false
}

// mede devolve o tempo médio, em nanossegundos, de uma busca feita pela
// função f para cada chave da lista de consultas.
func mede(consultas []int, f func(int) bool) float64 {
	achadas := 0
	inicio := time.Now()
	for _, c := range consultas {
		if f(c) {
			achadas++
		}
	}
	decorrido := time.Since(inicio)
	if achadas != len(consultas) {
		panic("alguma chave não foi encontrada")
	}
	return float64(decorrido.Nanoseconds()) / float64(len(consultas))
}

func main() {
	// Parte 1: a ordem de percurso do map.
	idade := map[string]int{
		"Ana": 19, "Bruno": 22, "Carla": 20, "Davi": 21,
		"Elisa": 23, "Fábio": 19, "Gabi": 24, "Hugo": 20,
	}
	fmt.Println("Três percursos do mesmo map, sem alterar nada entre eles:")
	for rodada := 1; rodada <= 3; rodada++ {
		fmt.Printf("  %d:", rodada)
		for nome := range idade {
			fmt.Printf(" %s", nome)
		}
		fmt.Println()
	}

	// Parte 2: o tempo de uma busca conforme n cresce.
	fmt.Println()
	fmt.Println("Tempo médio de uma busca bem-sucedida, em nanossegundos:")
	fmt.Println()
	fmt.Println("         n | sequencial |  binária |     map")
	fmt.Println("-----------+------------+----------+--------")

	gerador := rand.New(rand.NewSource(1))
	for _, n := range []int{1000, 10000, 100000, 1000000} {
		chaves := gerador.Perm(n)
		ordenadas := append([]int{}, chaves...)
		sort.Ints(ordenadas)
		mapa := make(map[int]bool, n)
		for _, c := range chaves {
			mapa[c] = true
		}

		// A busca sequencial é medida com menos consultas, senão a linha de
		// um milhão de chaves levaria vários segundos.
		consultas := make([]int, 200000)
		for i := range consultas {
			consultas[i] = gerador.Intn(n)
		}

		seq := mede(consultas[:1000], func(c int) bool { return buscaSequencial(chaves, c) })
		bin := mede(consultas, func(c int) bool { return buscaBinaria(ordenadas, c) })
		mp := mede(consultas, func(c int) bool { return mapa[c] })

		fmt.Printf("%10d | %10.0f | %8.0f | %7.0f\n", n, seq, bin, mp)
	}
}
