package algoritmos

import (
	"math/rand"
	"sort"
	"testing"
)

// entradas reúne os casos usados nos testes de ordenação e de busca.
// A entrada aleatória usa semente fixa: falhas são sempre reproduzíveis.
func entradas() map[string][]int {
	aleatorio := make([]int, 500)
	r := rand.New(rand.NewSource(7))
	for i := range aleatorio {
		aleatorio[i] = r.Intn(1000)
	}

	return map[string][]int{
		"vazio":              {},
		"um elemento":        {42},
		"dois fora de ordem": {2, 1},
		"ja ordenado":        {1, 2, 3, 4, 5, 6, 7, 8},
		"invertido":          {8, 7, 6, 5, 4, 3, 2, 1},
		"com repetidos":      {5, 1, 5, 3, 3, 1, 9, 5},
		"aleatorio":          aleatorio,
	}
}

func copia(v []int) []int {
	return append([]int(nil), v...)
}

var ordenadores = map[string]func([]int){
	"SelectionSort":        SelectionSort,
	"InsertionSort":        InsertionSort,
	"MergeSort":            MergeSort,
	"QuickSort":            QuickSort,
	"QuickSortEmbaralhado": QuickSortEmbaralhado,
}

func TestOrdenacao(t *testing.T) {
	for nome, ordenar := range ordenadores {
		t.Run(nome, func(t *testing.T) {
			for caso, entrada := range entradas() {
				obtido := copia(entrada)
				esperado := copia(entrada)
				sort.Ints(esperado)

				ordenar(obtido)

				if len(obtido) != len(esperado) {
					t.Fatalf("%s: devolveu %d elementos, esperado %d", caso, len(obtido), len(esperado))
				}
				for i := range esperado {
					if obtido[i] != esperado[i] {
						t.Fatalf("%s: posicao %d vale %d, esperado %d (obtido %v)",
							caso, i, obtido[i], esperado[i], resumo(obtido))
					}
				}
			}
		})
	}
}

// TestParticiona verifica a pós-condição de Particiona, sem olhar o resultado
// final da ordenação: o pivô precisa terminar na posição devolvida, com os
// menores à esquerda e os demais à direita.
func TestParticiona(t *testing.T) {
	casos := [][]int{
		{3, 8, 1, 9, 4},
		{1, 2, 3, 4, 5},
		{5, 4, 3, 2, 1},
		{7, 7, 7, 7},
		{2, 1},
	}

	for _, entrada := range casos {
		v := copia(entrada)
		hi := len(v) - 1
		pivo := v[hi]

		p := Particiona(v, 0, hi)

		if p < 0 || p > hi {
			t.Fatalf("%v: posicao devolvida %d fora do intervalo [0, %d]", entrada, p, hi)
		}
		if v[p] != pivo {
			t.Fatalf("%v: o pivo era %d, mas v[%d] vale %d apos particionar (%v)",
				entrada, pivo, p, v[p], v)
		}
		for i := 0; i < p; i++ {
			if v[i] >= pivo {
				t.Fatalf("%v: v[%d]=%d esta a esquerda do pivo %d (%v)", entrada, i, v[i], pivo, v)
			}
		}
		for i := p + 1; i <= hi; i++ {
			if v[i] < pivo {
				t.Fatalf("%v: v[%d]=%d esta a direita do pivo %d (%v)", entrada, i, v[i], pivo, v)
			}
		}
	}
}

func TestMerge(t *testing.T) {
	casos := []struct{ a, b []int }{
		{[]int{1, 3, 5}, []int{2, 4, 6}},
		{[]int{1, 2, 3}, []int{4, 5, 6}},
		{[]int{4, 5, 6}, []int{1, 2, 3}},
		{[]int{1, 1, 1}, []int{1, 1}},
		{[]int{}, []int{1, 2}},
		{[]int{1, 2}, []int{}},
	}

	for _, caso := range casos {
		esperado := append(copia(caso.a), caso.b...)
		sort.Ints(esperado)

		obtido := Merge(caso.a, caso.b)

		if len(obtido) != len(esperado) {
			t.Fatalf("Merge(%v, %v) devolveu %d elementos (%v), esperado %d",
				caso.a, caso.b, len(obtido), obtido, len(esperado))
		}
		for i := range esperado {
			if obtido[i] != esperado[i] {
				t.Fatalf("Merge(%v, %v) = %v, esperado %v", caso.a, caso.b, obtido, esperado)
			}
		}
	}
}

func TestBuscaSequencial(t *testing.T) {
	v := []int{4, 8, 15, 16, 23, 42}

	for i, x := range v {
		if obtido := BuscaSequencial(v, x); obtido != i {
			t.Errorf("BuscaSequencial(v, %d) = %d, esperado %d", x, obtido, i)
		}
	}
	if obtido := BuscaSequencial(v, 99); obtido != -1 {
		t.Errorf("BuscaSequencial(v, 99) = %d, esperado -1", obtido)
	}
}

func TestBuscaBinaria(t *testing.T) {
	vetores := [][]int{
		{},
		{42},
		{1, 2},
		{4, 8, 15, 16, 23, 42},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
	}

	for _, v := range vetores {
		for i, x := range v {
			obtido := BuscaBinaria(v, x)
			if obtido != i {
				t.Errorf("BuscaBinaria(%v, %d) = %d, esperado %d", v, x, obtido, i)
			}
		}
		for _, ausente := range []int{-1, 0, 100} {
			if obtido := BuscaBinaria(v, ausente); obtido != -1 {
				t.Errorf("BuscaBinaria(%v, %d) = %d, esperado -1", v, ausente, obtido)
			}
		}
	}
}

// TestBuscaBinariaGrande confere a busca binária contra a busca sequencial em
// um vetor grande, com semente fixa.
func TestBuscaBinariaGrande(t *testing.T) {
	const n = 10000
	r := rand.New(rand.NewSource(11))

	v := make([]int, n)
	for i := range v {
		v[i] = r.Intn(5 * n)
	}
	sort.Ints(v)

	for i := 0; i < 200; i++ {
		x := r.Intn(5 * n)
		obtido := BuscaBinaria(v, x)
		esperado := BuscaSequencial(v, x)

		if esperado == -1 {
			if obtido != -1 {
				t.Fatalf("busca por %d: binaria devolveu %d, mas o valor nao esta no vetor", x, obtido)
			}
			continue
		}
		if obtido < 0 || obtido >= n || v[obtido] != x {
			t.Fatalf("busca por %d: binaria devolveu %d, sequencial devolveu %d", x, obtido, esperado)
		}
	}
}

// resumo encurta vetores longos nas mensagens de erro.
func resumo(v []int) []int {
	if len(v) <= 12 {
		return v
	}
	return v[:12]
}
