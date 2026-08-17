package unionfind

// WeightedQuickUnion é o Quick-Union com balanceamento por tamanho: a raiz da
// árvore menor passa a apontar para a raiz da maior, o que mantém a altura em
// O(log n).
type WeightedQuickUnion struct {
	// TODO: declare os campos necessários.
}

var _ UF = (*WeightedQuickUnion)(nil)

// NewWeightedQuickUnion cria a estrutura com n objetos.
func NewWeightedQuickUnion(n int) *WeightedQuickUnion {
	panic("TODO: implementar NewWeightedQuickUnion")
}

func (uf *WeightedQuickUnion) Union(p, q int) {
	panic("TODO: implementar WeightedQuickUnion.Union")
}

func (uf *WeightedQuickUnion) Find(p int) int {
	panic("TODO: implementar WeightedQuickUnion.Find")
}

func (uf *WeightedQuickUnion) Connected(p, q int) bool {
	panic("TODO: implementar WeightedQuickUnion.Connected")
}

func (uf *WeightedQuickUnion) Count() int {
	panic("TODO: implementar WeightedQuickUnion.Count")
}
