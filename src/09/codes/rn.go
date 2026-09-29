// Árvore rubro-negra na variante left-leaning (LLRB), como no
// RedBlackBST.java de Sedgewick e Wayne (Algorithms, 4a ed., seção 3.3):
// inserção, busca e verificação das propriedades.
//
// Execute com:  go run rn.go

package main

import (
	"fmt"
	"sort"
)

// Cor de um nó. A cor pertence ao nó e, por convenção, é a cor da ligação
// que vem do pai até ele.
type Cor bool

const (
	Vermelho Cor = true
	Preto    Cor = false
)

// No é um nó da árvore rubro-negra. Em relação ao nó da ABB, o único campo
// novo é a cor.
type No struct {
	Info int
	Cor  Cor
	Esq  *No
	Dir  *No
}

// vermelho responde se a ligação que chega em a é vermelha. Nó nulo conta
// como preto, o que dispensa testar nil em todo lugar.
func vermelho(a *No) bool {
	return a != nil && a.Cor == Vermelho
}

// Busca é a mesma da árvore binária de busca: a cor não participa da
// comparação.
func Busca(a *No, procurado int) bool {
	for a != nil {
		if procurado < a.Info {
			a = a.Esq
		} else if procurado > a.Info {
			a = a.Dir
		} else {
			return true
		}
	}
	return false
}

// rotacionaEsquerda joga para a esquerda uma ligação vermelha que está à
// direita. A cor acompanha o nó que sobe.
func rotacionaEsquerda(a *No) *No {
	b := a.Dir
	a.Dir = b.Esq
	b.Esq = a
	b.Cor = a.Cor
	a.Cor = Vermelho
	return b
}

// rotacionaDireita é a imagem espelhada de rotacionaEsquerda.
func rotacionaDireita(a *No) *No {
	b := a.Esq
	a.Esq = b.Dir
	b.Dir = a
	b.Cor = a.Cor
	a.Cor = Vermelho
	return b
}

// inverteCores repassa para o pai o vermelho dos dois filhos. É a
// recoloração: nenhum ponteiro muda. Em Sedgewick, a função troca cada cor
// pela oposta, para servir também à remoção; na inserção, dá no mesmo.
func inverteCores(a *No) {
	a.Cor = Vermelho
	a.Esq.Cor = Preto
	a.Dir.Cor = Preto
}

// Insere acrescenta v à árvore e devolve a nova raiz, sempre preta.
func Insere(a *No, v int) *No {
	a = insere(a, v)
	a.Cor = Preto
	return a
}

// insere é a inserção da ABB com três consertos na volta da recursão.
func insere(a *No, v int) *No {
	if a == nil {
		// Todo nó novo entra vermelho: assim ele não altera a quantidade
		// de nós pretos de nenhum caminho.
		return &No{Info: v, Cor: Vermelho}
	}

	if v < a.Info {
		a.Esq = insere(a.Esq, v)
	} else if v > a.Info {
		a.Dir = insere(a.Dir, v)
	}

	if vermelho(a.Dir) && !vermelho(a.Esq) {
		a = rotacionaEsquerda(a) // vermelho à direita: passa para a esquerda
	}
	if vermelho(a.Esq) && vermelho(a.Esq.Esq) {
		a = rotacionaDireita(a) // dois vermelhos seguidos: equilibra
	}
	if vermelho(a.Esq) && vermelho(a.Dir) {
		inverteCores(a) // dois filhos vermelhos: sobe o vermelho
	}

	return a
}

// Altura devolve a altura da árvore contando todos os nós. A árvore vazia
// tem altura -1.
func Altura(a *No) int {
	if a == nil {
		return -1
	}
	return 1 + maior(Altura(a.Esq), Altura(a.Dir))
}

// AlturaPreta conta apenas os nós pretos de um caminho qualquer da raiz até
// uma folha. Se a árvore for rubro-negra, o valor é o mesmo para todos os
// caminhos, e é isso que EhRubroNegra confere.
func AlturaPreta(a *No) int {
	if a == nil {
		return 0
	}
	if a.Cor == Preto {
		return 1 + AlturaPreta(a.Esq)
	}
	return AlturaPreta(a.Esq)
}

// EhRubroNegra confere as três propriedades que interessam: raiz preta,
// nenhum nó vermelho com filho vermelho, e mesma altura preta em todos os
// caminhos.
func EhRubroNegra(a *No) bool {
	if vermelho(a) {
		return false
	}
	return propriedadesLocais(a) && alturaPretaUniforme(a, AlturaPreta(a), 0)
}

func propriedadesLocais(a *No) bool {
	if a == nil {
		return true
	}
	if vermelho(a) && (vermelho(a.Esq) || vermelho(a.Dir)) {
		return false
	}
	return propriedadesLocais(a.Esq) && propriedadesLocais(a.Dir)
}

func alturaPretaUniforme(a *No, esperada, acumulada int) bool {
	if a == nil {
		return acumulada == esperada
	}
	if a.Cor == Preto {
		acumulada++
	}
	return alturaPretaUniforme(a.Esq, esperada, acumulada) &&
		alturaPretaUniforme(a.Dir, esperada, acumulada)
}

// Imprime desenha a árvore deitada, com a raiz à esquerda e a subárvore
// direita em cima. (v) marca o nó cuja ligação com o pai é vermelha, e (p),
// o nó cuja ligação é preta.
func Imprime(a *No, nivel int) {
	if a == nil {
		return
	}
	Imprime(a.Dir, nivel+1)
	for i := 0; i < nivel; i++ {
		fmt.Print("    ")
	}
	cor := "p"
	if a.Cor == Vermelho {
		cor = "v"
	}
	fmt.Printf("%d (%s)\n", a.Info, cor)
	Imprime(a.Esq, nivel+1)
}

// EmOrdem imprime as chaves em ordem crescente.
func EmOrdem(a *No) {
	if a == nil {
		return
	}
	EmOrdem(a.Esq)
	fmt.Print(a.Info, " ")
	EmOrdem(a.Dir)
}

func maior(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	valores := []int{50, 30, 90, 20, 40, 95, 10, 35, 45}

	var a *No
	for _, v := range valores {
		a = Insere(a, v)
	}

	fmt.Println("Inserindo", valores)
	fmt.Println()
	Imprime(a, 0)
	fmt.Println()
	fmt.Print("in-ordem: ")
	EmOrdem(a)
	fmt.Println()
	fmt.Printf("altura: %d   altura preta: %d   é rubro-negra: %v\n",
		Altura(a), AlturaPreta(a), EhRubroNegra(a))

	fmt.Println()
	fmt.Println("Os mesmos valores em ordem crescente:")
	fmt.Println()

	crescentes := append([]int{}, valores...)
	sort.Ints(crescentes)

	var b *No
	for _, v := range crescentes {
		b = Insere(b, v)
	}
	Imprime(b, 0)
	fmt.Printf("\naltura: %d   altura preta: %d   é rubro-negra: %v\n",
		Altura(b), AlturaPreta(b), EhRubroNegra(b))
}
