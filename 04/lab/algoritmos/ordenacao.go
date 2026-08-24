package algoritmos

import "math/rand"

// SelectionSort ordena v no lugar, escolhendo a cada passo o menor elemento do
// trecho ainda não ordenado. Faz sempre ~n²/2 comparações, qualquer que seja a
// entrada. Implementação correta, serve de referência de estilo.
func SelectionSort(v []int) {
	n := len(v)
	for i := 0; i < n-1; i++ {
		menor := i
		for j := i + 1; j < n; j++ {
			if v[j] < v[menor] {
				menor = j
			}
		}
		v[i], v[menor] = v[menor], v[i]
	}
}

// InsertionSort ordena v no lugar, deslocando cada elemento para a esquerda
// até encontrar sua posição no trecho já ordenado.
//
// Esta implementação está incorreta.
func InsertionSort(v []int) {
	for i := 1; i < len(v); i++ {
		for j := i; j > 1 && v[j] < v[j-1]; j-- {
			v[j], v[j-1] = v[j-1], v[j]
		}
	}
}

// Merge intercala dois vetores já ordenados em um terceiro, também ordenado.
// É a operação básica do MergeSort.
//
// Esta implementação está incorreta.
func Merge(a, b []int) []int {
	c := make([]int, 0, len(a)+len(b))
	i, j := 0, 0

	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			c = append(c, a[i])
			i++
		} else {
			c = append(c, b[j])
			j++
		}
	}

	for i < len(a) {
		c = append(c, a[i])
		i++
	}

	return c
}

// MergeSort ordena v no lugar dividindo o vetor ao meio, ordenando cada
// metade e intercalando as duas com Merge.
//
// As metades são copiadas para vetores novos a cada nível da recursão. Uma
// implementação de biblioteca evita boa parte dessas cópias; aqui a clareza
// vem antes. O custo assintótico é o mesmo, O(n log n).
func MergeSort(v []int) {
	if len(v) < 2 {
		return
	}
	meio := len(v) / 2
	esq := append([]int(nil), v[:meio]...)
	dir := append([]int(nil), v[meio:]...)
	MergeSort(esq)
	MergeSort(dir)
	copy(v, Merge(esq, dir))
}

// Particiona rearranja v[lo..hi] em torno do pivô, que é o último elemento do
// trecho, e devolve a posição p em que o pivô ficou. Ao final da chamada,
// todo elemento de v[lo..p-1] é menor que v[p], e todo elemento de
// v[p+1..hi] é maior ou igual a v[p].
//
// Esta implementação está incorreta: a pós-condição descrita acima não vale.
func Particiona(v []int, lo, hi int) int {
	pivo := v[hi]
	i := lo

	for j := lo; j < hi; j++ {
		if v[j] < pivo {
			v[i], v[j] = v[j], v[i]
			i++
		}
	}

	return i
}

// QuickSort ordena v no lugar, particionando o vetor em torno de um pivô e
// ordenando os dois lados recursivamente.
func QuickSort(v []int) {
	quickSort(v, 0, len(v)-1)
}

func quickSort(v []int, lo, hi int) {
	if lo >= hi {
		return
	}
	p := Particiona(v, lo, hi)
	quickSort(v, lo, p-1)
	quickSort(v, p+1, hi)
}

// QuickSortEmbaralhado embaralha o vetor antes de chamar o mesmo quickSort.
// A semente é fixa para que duas execuções produzam a mesma sequência.
func QuickSortEmbaralhado(v []int) {
	r := rand.New(rand.NewSource(42))
	r.Shuffle(len(v), func(i, j int) { v[i], v[j] = v[j], v[i] })
	quickSort(v, 0, len(v)-1)
}
