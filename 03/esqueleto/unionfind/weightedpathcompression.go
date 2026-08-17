package unionfind

// WeightedPathCompression é o Weighted Quick-Union com compressão de caminho:
// a cada subida até a raiz, os nós visitados são aproximados dela.
//
// Esta é a única variante que não tem código pronto no material da disciplina.
// A seção 15 da aula de Análise de Algoritmos descreve a ideia e mostra a
// versão que aponta cada nó para o avô.
type WeightedPathCompression struct {
	// TODO: declare os campos necessários.
}

var _ UF = (*WeightedPathCompression)(nil)

// NewWeightedPathCompression cria a estrutura com n objetos.
func NewWeightedPathCompression(n int) *WeightedPathCompression {
	panic("TODO: implementar NewWeightedPathCompression")
}

func (uf *WeightedPathCompression) Union(p, q int) {
	panic("TODO: implementar WeightedPathCompression.Union")
}

func (uf *WeightedPathCompression) Find(p int) int {
	panic("TODO: implementar WeightedPathCompression.Find")
}

func (uf *WeightedPathCompression) Connected(p, q int) bool {
	panic("TODO: implementar WeightedPathCompression.Connected")
}

func (uf *WeightedPathCompression) Count() int {
	panic("TODO: implementar WeightedPathCompression.Count")
}
