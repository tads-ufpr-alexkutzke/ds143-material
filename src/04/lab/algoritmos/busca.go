package algoritmos

// BuscaSequencial devolve o índice de x em v, ou -1 se x não estiver presente.
// Não exige vetor ordenado. Custo O(n) no pior caso.
//
// Implementação correta, serve de referência de estilo.
func BuscaSequencial(v []int, x int) int {
	for i := 0; i < len(v); i++ {
		if v[i] == x {
			return i
		}
	}
	return -1
}

// BuscaBinaria devolve o índice de x em v, que precisa estar ordenado, ou -1
// se x não estiver presente. Custo O(log n) no pior caso.
//
// Esta implementação está incorreta. Rode os testes, observe em que entradas
// ela erra e corrija.
func BuscaBinaria(v []int, x int) int {
	lo, hi := 0, len(v)-1
	for lo < hi {
		meio := lo + (hi-lo)/2
		switch {
		case x < v[meio]:
			hi = meio - 1
		case x > v[meio]:
			lo = meio + 1
		default:
			return meio
		}
	}
	return -1
}
