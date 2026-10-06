// Filtro de Bloom: um vetor de bits e k funções de dispersão. Responde se uma
// chave "talvez esteja" ou "certamente não está" no conjunto, sem guardar as
// chaves.
//
// Execute com:  go run bloom.go

package main

import (
	"fmt"
	"hash/fnv"
	"math"
)

type Bloom struct {
	bits []bool
	k    int
}

func NovoBloom(m, k int) *Bloom {
	return &Bloom{bits: make([]bool, m), k: k}
}

// posicoes devolve as k posições da chave. As k funções são obtidas de duas,
// h1 e h2, pela fórmula h1 + i*h2, a mesma ideia da dispersão dupla.
func (b *Bloom) posicoes(chave string) []int {
	f := fnv.New64a()
	f.Write([]byte(chave))
	h := f.Sum64()
	h1 := h & 0xffffffff
	h2 := h >> 32

	m := uint64(len(b.bits))
	ps := make([]int, b.k)
	for i := range ps {
		ps[i] = int((h1 + uint64(i)*h2) % m)
	}
	return ps
}

// Insere liga os k bits da chave.
func (b *Bloom) Insere(chave string) {
	for _, p := range b.posicoes(chave) {
		b.bits[p] = true
	}
}

// TalvezContenha responde false só quando algum dos k bits está desligado:
// nesse caso, a chave certamente nunca foi inserida. Com os k bits ligados,
// a chave pode ter sido inserida ou os bits podem ter sido ligados por outras.
func (b *Bloom) TalvezContenha(chave string) bool {
	for _, p := range b.posicoes(chave) {
		if !b.bits[p] {
			return false
		}
	}
	return true
}

func main() {
	const n = 10000  // chaves inseridas
	const m = 10 * n // bits do filtro

	fmt.Printf("%d chaves inseridas em %d bits (%d bits por chave)\n", n, m, m/n)
	fmt.Println("Taxa de falsos positivos em 100000 chaves que não foram inseridas")
	fmt.Println()
	fmt.Println(" k | medida | teoria")
	fmt.Println("---+--------+-------")

	for k := 1; k <= 10; k++ {
		b := NovoBloom(m, k)
		for i := 0; i < n; i++ {
			b.Insere(fmt.Sprintf("inserida-%d", i))
		}

		// Nenhuma chave inserida pode ser negada.
		for i := 0; i < n; i++ {
			if !b.TalvezContenha(fmt.Sprintf("inserida-%d", i)) {
				panic("falso negativo")
			}
		}

		falsos := 0
		const consultas = 100000
		for i := 0; i < consultas; i++ {
			if b.TalvezContenha(fmt.Sprintf("ausente-%d", i)) {
				falsos++
			}
		}

		teoria := math.Pow(1-math.Exp(-float64(k)*n/m), float64(k))
		fmt.Printf("%2d | %5.2f%% | %5.2f%%\n", k,
			100*float64(falsos)/consultas, 100*teoria)
	}
}
