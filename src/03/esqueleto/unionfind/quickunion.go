package unionfind

// QuickUnion representa as componentes como uma floresta: cada objeto aponta
// para um pai, e a raiz identifica a componente. Implementação pronta.
type QuickUnion struct {
	parent []int // parent[i] = pai de i; parent[i] == i indica raiz
	count  int
}

var _ UF = (*QuickUnion)(nil)

// NewQuickUnion cria a estrutura com n objetos, cada um raiz de sua própria
// árvore.
func NewQuickUnion(n int) *QuickUnion {
	uf := &QuickUnion{parent: make([]int, n), count: n}
	for i := range uf.parent {
		uf.parent[i] = i
	}
	return uf
}

// Find sobe até a raiz. O custo é a altura da árvore, que sem balanceamento
// pode chegar a n-1.
func (uf *QuickUnion) Find(p int) int {
	for p != uf.parent[p] {
		p = uf.parent[p]
	}
	return p
}

func (uf *QuickUnion) Connected(p, q int) bool {
	return uf.Find(p) == uf.Find(q)
}

// Union pendura a raiz de p sob a raiz de q, sem critério nenhum. É essa
// escolha arbitrária que permite a degeneração em lista encadeada.
func (uf *QuickUnion) Union(p, q int) {
	raizP, raizQ := uf.Find(p), uf.Find(q)
	if raizP == raizQ {
		return
	}
	uf.parent[raizP] = raizQ
	uf.count--
}

func (uf *QuickUnion) Count() int {
	return uf.count
}
