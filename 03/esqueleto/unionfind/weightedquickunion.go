package unionfind

// WeightedQuickUnion é o Quick-Union com balanceamento por tamanho: a raiz da
// árvore menor passa a apontar para a raiz da maior, o que mantém a altura em
// O(log n). Implementação pronta, e ponto de partida da que falta.
type WeightedQuickUnion struct {
	parent []int
	size   []int // size[i] = número de objetos na subárvore de raiz i
	count  int
}

var _ UF = (*WeightedQuickUnion)(nil)

// NewWeightedQuickUnion cria a estrutura com n objetos.
func NewWeightedQuickUnion(n int) *WeightedQuickUnion {
	uf := &WeightedQuickUnion{
		parent: make([]int, n),
		size:   make([]int, n),
		count:  n,
	}
	for i := range uf.parent {
		uf.parent[i] = i
		uf.size[i] = 1
	}
	return uf
}

// Find sobe até a raiz, sem alterar a árvore no caminho.
func (uf *WeightedQuickUnion) Find(p int) int {
	for p != uf.parent[p] {
		p = uf.parent[p]
	}
	return p
}

func (uf *WeightedQuickUnion) Connected(p, q int) bool {
	return uf.Find(p) == uf.Find(q)
}

// Union pendura a árvore menor sob a maior e atualiza o tamanho da nova raiz.
func (uf *WeightedQuickUnion) Union(p, q int) {
	raizP, raizQ := uf.Find(p), uf.Find(q)
	if raizP == raizQ {
		return
	}
	if uf.size[raizP] < uf.size[raizQ] {
		uf.parent[raizP] = raizQ
		uf.size[raizQ] += uf.size[raizP]
	} else {
		uf.parent[raizQ] = raizP
		uf.size[raizP] += uf.size[raizQ]
	}
	uf.count--
}

func (uf *WeightedQuickUnion) Count() int {
	return uf.count
}
