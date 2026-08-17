package unionfind

// QuickFind mantém, para cada objeto, o identificador da componente a que ele
// pertence. Find custa O(1) e Union custa O(n).
type QuickFind struct {
	// TODO: declare os campos necessários.
}

// Garante em tempo de compilação que *QuickFind cumpre a interface UF.
// Enquanto algum método estiver faltando, o pacote não compila.
var _ UF = (*QuickFind)(nil)

// NewQuickFind cria a estrutura com n objetos, cada um em sua própria
// componente.
func NewQuickFind(n int) *QuickFind {
	panic("TODO: implementar NewQuickFind")
}

func (uf *QuickFind) Union(p, q int) {
	panic("TODO: implementar QuickFind.Union")
}

func (uf *QuickFind) Find(p int) int {
	panic("TODO: implementar QuickFind.Find")
}

func (uf *QuickFind) Connected(p, q int) bool {
	panic("TODO: implementar QuickFind.Connected")
}

func (uf *QuickFind) Count() int {
	panic("TODO: implementar QuickFind.Count")
}
