package unionfind

// QuickUnion representa as componentes como uma floresta: cada objeto aponta
// para um pai, e a raiz identifica a componente. Union custa uma subida até
// cada raiz, e as árvores podem degenerar em listas encadeadas.
type QuickUnion struct {
	// TODO: declare os campos necessários.
}

var _ UF = (*QuickUnion)(nil)

// NewQuickUnion cria a estrutura com n objetos, cada um raiz de sua própria
// árvore.
func NewQuickUnion(n int) *QuickUnion {
	panic("TODO: implementar NewQuickUnion")
}

func (uf *QuickUnion) Union(p, q int) {
	panic("TODO: implementar QuickUnion.Union")
}

func (uf *QuickUnion) Find(p int) int {
	panic("TODO: implementar QuickUnion.Find")
}

func (uf *QuickUnion) Connected(p, q int) bool {
	panic("TODO: implementar QuickUnion.Connected")
}

func (uf *QuickUnion) Count() int {
	panic("TODO: implementar QuickUnion.Count")
}
