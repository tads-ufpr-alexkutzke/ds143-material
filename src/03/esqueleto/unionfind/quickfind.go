package unionfind

// QuickFind mantém, para cada objeto, o identificador da componente a que ele
// pertence. Esta implementação já está pronta: use-a como modelo do idioma
// esperado nas demais.
type QuickFind struct {
	id    []int // id[i] = componente do objeto i
	count int   // número de componentes
}

// Garante em tempo de compilação que *QuickFind cumpre a interface UF.
// Enquanto algum método estiver faltando, o pacote não compila.
var _ UF = (*QuickFind)(nil)

// NewQuickFind cria a estrutura com n objetos, cada um em sua própria
// componente.
func NewQuickFind(n int) *QuickFind {
	uf := &QuickFind{id: make([]int, n), count: n}
	for i := range uf.id {
		uf.id[i] = i
	}
	return uf
}

// Find custa O(1): o identificador da componente está armazenado.
func (uf *QuickFind) Find(p int) int {
	return uf.id[p]
}

func (uf *QuickFind) Connected(p, q int) bool {
	return uf.id[p] == uf.id[q]
}

// Union custa O(n): renomear a componente de p exige percorrer todo o slice.
func (uf *QuickFind) Union(p, q int) {
	pID, qID := uf.id[p], uf.id[q]
	if pID == qID {
		return
	}
	for i := range uf.id {
		if uf.id[i] == pID {
			uf.id[i] = qID
		}
	}
	uf.count--
}

func (uf *QuickFind) Count() int {
	return uf.count
}
