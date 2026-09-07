// Comparação empírica entre três formas de guardar um conjunto de
// chaves com busca e inserção: slice sem ordem, slice ordenado e árvore
// binária de busca.
//
// Execute com:  go run motivacao.go

package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const (
	n       = 200000 // chaves inseridas
	buscas  = 10000 // buscas medidas
	semente = 1
)

// ---------- slice sem ordem ----------

func insereSlice(s []int, v int) []int {
	return append(s, v)
}

func buscaSlice(s []int, v int) bool {
	for _, x := range s {
		if x == v {
			return true
		}
	}
	return false
}

// ---------- slice ordenado ----------

// insereOrdenado acha a posição por busca binária e abre espaço para a
// chave nova, empurrando para a direita tudo o que vem depois dela.
func insereOrdenado(s []int, v int) []int {
	i := sort.SearchInts(s, v)
	s = append(s, 0)
	copy(s[i+1:], s[i:])
	s[i] = v
	return s
}

func buscaOrdenada(s []int, v int) bool {
	i := sort.SearchInts(s, v)
	return i < len(s) && s[i] == v
}

// ---------- árvore binária de busca ----------

type No struct {
	Info int
	Esq  *No
	Dir  *No
}

func insereArvore(a *No, v int) *No {
	if a == nil {
		return &No{Info: v}
	}
	if v < a.Info {
		a.Esq = insereArvore(a.Esq, v)
	} else if v > a.Info {
		a.Dir = insereArvore(a.Dir, v)
	}
	return a
}

func buscaArvore(a *No, v int) bool {
	for a != nil {
		if v < a.Info {
			a = a.Esq
		} else if v > a.Info {
			a = a.Dir
		} else {
			return true
		}
	}
	return false
}

func main() {
	r := rand.New(rand.NewSource(semente))
	chaves := r.Perm(10 * n)[:n]
	procuradas := make([]int, buscas)
	for i := range procuradas {
		procuradas[i] = chaves[r.Intn(n)]
	}

	fmt.Printf("%d chaves inseridas, %d buscas\n\n", n, buscas)
	fmt.Println("estrutura        | inserção | busca")
	fmt.Println("-----------------+----------+---------")

	inicio := time.Now()
	var s []int
	for _, v := range chaves {
		s = insereSlice(s, v)
	}
	tempoInsercao := time.Since(inicio)

	inicio = time.Now()
	for _, v := range procuradas {
		buscaSlice(s, v)
	}
	fmt.Printf("%-16s | %8v | %v\n", "slice sem ordem", tempoInsercao.Round(time.Millisecond), time.Since(inicio).Round(time.Millisecond))

	inicio = time.Now()
	var o []int
	for _, v := range chaves {
		o = insereOrdenado(o, v)
	}
	tempoInsercao = time.Since(inicio)

	inicio = time.Now()
	for _, v := range procuradas {
		buscaOrdenada(o, v)
	}
	fmt.Printf("%-16s | %8v | %v\n", "slice ordenado", tempoInsercao.Round(time.Millisecond), time.Since(inicio).Round(time.Millisecond))

	inicio = time.Now()
	var a *No
	for _, v := range chaves {
		a = insereArvore(a, v)
	}
	tempoInsercao = time.Since(inicio)

	inicio = time.Now()
	for _, v := range procuradas {
		buscaArvore(a, v)
	}
	fmt.Printf("%-16s | %8v | %v\n", "árvore de busca", tempoInsercao.Round(time.Millisecond), time.Since(inicio).Round(time.Millisecond))
}
