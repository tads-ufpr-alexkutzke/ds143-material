// A forma da árvore binária de busca depende da ordem de inserção.
//
// Execute com:  go run forma.go

package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
)

type No struct {
	Info int
	Esq  *No
	Dir  *No
}

func Insere(a *No, v int) *No {
	if a == nil {
		return &No{Info: v}
	}
	if v < a.Info {
		a.Esq = Insere(a.Esq, v)
	} else if v > a.Info {
		a.Dir = Insere(a.Dir, v)
	}
	return a
}

func Altura(a *No) int {
	if a == nil {
		return -1
	}
	esq := Altura(a.Esq)
	dir := Altura(a.Dir)
	if esq > dir {
		return 1 + esq
	}
	return 1 + dir
}

// SomaProfundidades soma a profundidade de todos os nós. Dividida pelo
// número de nós e somada a 1, dá o número médio de comparações de uma
// busca bem sucedida: um nó na profundidade d é achado com d+1
// comparações.
func SomaProfundidades(a *No, profundidade int) int {
	if a == nil {
		return 0
	}
	return profundidade +
		SomaProfundidades(a.Esq, profundidade+1) +
		SomaProfundidades(a.Dir, profundidade+1)
}

func Constroi(valores []int) *No {
	var a *No
	for _, v := range valores {
		a = Insere(a, v)
	}
	return a
}

func Imprime(a *No, profundidade int) {
	if a == nil {
		return
	}
	Imprime(a.Dir, profundidade+1)
	for i := 0; i < profundidade; i++ {
		fmt.Print("    ")
	}
	fmt.Println(a.Info)
	Imprime(a.Esq, profundidade+1)
}

func main() {
	valores := []int{50, 30, 90, 20, 40, 95, 10, 35, 45}

	crescente := append([]int(nil), valores...)
	sort.Ints(crescente)

	fmt.Println("Inserindo 50 30 90 20 40 95 10 35 45:")
	Imprime(Constroi(valores), 0)
	fmt.Println("altura:", Altura(Constroi(valores)))

	fmt.Println("\nInserindo os mesmos valores em ordem crescente:")
	Imprime(Constroi(crescente), 0)
	fmt.Println("altura:", Altura(Constroi(crescente)))

	// Altura de árvores construídas a partir de permutações aleatórias,
	// comparada com a altura mínima possível e com a da degenerada.
	fmt.Println("\nInserções em ordem aleatória (média de 100 árvores, 5 para n >= 100 mil):")
	fmt.Println("       n | mínima | aleatória | degenerada | comparações médias")
	fmt.Println("---------+--------+-----------+------------+-------------------")

	r := rand.New(rand.NewSource(1))
	for _, n := range []int{1000, 10000, 100000, 1000000} {
		somaAlturas, somaMedias := 0, 0.0
		repeticoes := 100
		if n >= 100000 {
			repeticoes = 5
		}
		for i := 0; i < repeticoes; i++ {
			permutacao := r.Perm(n)
			a := Constroi(permutacao)
			somaAlturas += Altura(a)
			somaMedias += float64(SomaProfundidades(a, 0))/float64(n) + 1
		}
		minima := int(math.Ceil(math.Log2(float64(n)+1))) - 1
		fmt.Printf("%8d | %6d | %9.1f | %10d | %18.1f\n",
			n, minima,
			float64(somaAlturas)/float64(repeticoes),
			n-1,
			somaMedias/float64(repeticoes))
	}
}
