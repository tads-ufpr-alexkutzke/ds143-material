package main

// Preprocessamento guarda o que Preprocessa calcula e Distancia consulta.
type Preprocessamento struct {
	// TODO: declarar os campos necessários
}

// Preprocessa recebe um grafo linha, com pelo menos dois vértices e pesos
// positivos, e prepara as consultas de distância. O custo deve ser O(V).
func Preprocessa(g *GrafoP) *Preprocessamento {
	panic("TODO: implementar Preprocessa")
}

// Distancia devolve o comprimento do caminho mais curto entre os vértices v e
// w. O custo deve ser O(1).
func (p *Preprocessamento) Distancia(v, w int) int {
	panic("TODO: implementar Distancia")
}
