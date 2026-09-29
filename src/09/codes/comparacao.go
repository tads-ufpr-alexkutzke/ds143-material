// Comparação entre ABB, AVL e árvore rubro-negra: altura obtida e número de
// rotações executadas, nas mesmas sequências de inserção.
//
// Execute com:  go run comparacao.go

package main

import (
	"fmt"
	"math/rand"
)

// ---------------------------------------------------------------- ABB

type NoABB struct {
	Info int
	Esq  *NoABB
	Dir  *NoABB
}

func insereABB(a *NoABB, v int) *NoABB {
	if a == nil {
		return &NoABB{Info: v}
	}
	if v < a.Info {
		a.Esq = insereABB(a.Esq, v)
	} else if v > a.Info {
		a.Dir = insereABB(a.Dir, v)
	}
	return a
}

func alturaABB(a *NoABB) int {
	if a == nil {
		return -1
	}
	return 1 + max(alturaABB(a.Esq), alturaABB(a.Dir))
}

// ---------------------------------------------------------------- AVL

type NoAVL struct {
	Info int
	Alt  int
	Esq  *NoAVL
	Dir  *NoAVL
}

// rotacoesAVL conta as rotações simples executadas. Uma rotação dupla soma 2.
var rotacoesAVL int

func alturaAVL(a *NoAVL) int {
	if a == nil {
		return -1
	}
	return a.Alt
}

func atualizaAlturaAVL(a *NoAVL) {
	a.Alt = 1 + max(alturaAVL(a.Esq), alturaAVL(a.Dir))
}

func fator(a *NoAVL) int {
	return alturaAVL(a.Esq) - alturaAVL(a.Dir)
}

func rotacaoDireitaAVL(a *NoAVL) *NoAVL {
	rotacoesAVL++
	b := a.Esq
	a.Esq = b.Dir
	b.Dir = a
	atualizaAlturaAVL(a)
	atualizaAlturaAVL(b)
	return b
}

func rotacaoEsquerdaAVL(a *NoAVL) *NoAVL {
	rotacoesAVL++
	b := a.Dir
	a.Dir = b.Esq
	b.Esq = a
	atualizaAlturaAVL(a)
	atualizaAlturaAVL(b)
	return b
}

func rebalanceia(a *NoAVL) *NoAVL {
	fb := fator(a)
	if fb > 1 {
		if fator(a.Esq) < 0 {
			a.Esq = rotacaoEsquerdaAVL(a.Esq)
		}
		return rotacaoDireitaAVL(a)
	}
	if fb < -1 {
		if fator(a.Dir) > 0 {
			a.Dir = rotacaoDireitaAVL(a.Dir)
		}
		return rotacaoEsquerdaAVL(a)
	}
	return a
}

func insereAVL(a *NoAVL, v int) *NoAVL {
	if a == nil {
		return &NoAVL{Info: v}
	}
	if v < a.Info {
		a.Esq = insereAVL(a.Esq, v)
	} else if v > a.Info {
		a.Dir = insereAVL(a.Dir, v)
	} else {
		return a
	}
	atualizaAlturaAVL(a)
	return rebalanceia(a)
}

// ------------------------------------------------------- Rubro-negra

type NoRN struct {
	Info     int
	Vermelho bool
	Esq      *NoRN
	Dir      *NoRN
}

// rotacoesRN e recoloracoesRN separam os dois tipos de conserto: a rotação
// mexe em ponteiros, a recoloração só troca cores.
var (
	rotacoesRN     int
	recoloracoesRN int
)

func ehVermelho(a *NoRN) bool {
	return a != nil && a.Vermelho
}

func rotacionaEsquerdaRN(a *NoRN) *NoRN {
	rotacoesRN++
	b := a.Dir
	a.Dir = b.Esq
	b.Esq = a
	b.Vermelho = a.Vermelho
	a.Vermelho = true
	return b
}

func rotacionaDireitaRN(a *NoRN) *NoRN {
	rotacoesRN++
	b := a.Esq
	a.Esq = b.Dir
	b.Dir = a
	b.Vermelho = a.Vermelho
	a.Vermelho = true
	return b
}

func inverteCoresRN(a *NoRN) {
	recoloracoesRN++
	a.Vermelho = true
	a.Esq.Vermelho = false
	a.Dir.Vermelho = false
}

func insereRN(a *NoRN, v int) *NoRN {
	if a == nil {
		return &NoRN{Info: v, Vermelho: true}
	}
	if v < a.Info {
		a.Esq = insereRN(a.Esq, v)
	} else if v > a.Info {
		a.Dir = insereRN(a.Dir, v)
	}
	if ehVermelho(a.Dir) && !ehVermelho(a.Esq) {
		a = rotacionaEsquerdaRN(a)
	}
	if ehVermelho(a.Esq) && ehVermelho(a.Esq.Esq) {
		a = rotacionaDireitaRN(a)
	}
	if ehVermelho(a.Esq) && ehVermelho(a.Dir) {
		inverteCoresRN(a)
	}
	return a
}

func insereRaizRN(a *NoRN, v int) *NoRN {
	a = insereRN(a, v)
	a.Vermelho = false
	return a
}

func alturaRN(a *NoRN) int {
	if a == nil {
		return -1
	}
	return 1 + max(alturaRN(a.Esq), alturaRN(a.Dir))
}

// ---------------------------------------------------------------- main

// constroi insere a sequência nas três estruturas e devolve as três alturas
// e as contagens de conserto.
func constroi(valores []int) (hABB, hAVL, hRN, rotAVL, rotRN, recRN int) {
	rotacoesAVL, rotacoesRN, recoloracoesRN = 0, 0, 0

	var abb *NoABB
	var avl *NoAVL
	var rn *NoRN
	for _, v := range valores {
		abb = insereABB(abb, v)
		avl = insereAVL(avl, v)
		rn = insereRaizRN(rn, v)
	}
	return alturaABB(abb), alturaAVL(avl), alturaRN(rn),
		rotacoesAVL, rotacoesRN, recoloracoesRN
}

func crescente(n int) []int {
	v := make([]int, n)
	for i := range v {
		v[i] = i + 1
	}
	return v
}

func aleatoria(n int, r *rand.Rand) []int {
	v := crescente(n)
	r.Shuffle(n, func(i, j int) { v[i], v[j] = v[j], v[i] })
	return v
}

func main() {
	r := rand.New(rand.NewSource(42))
	tamanhos := []int{1000, 10000, 100000}

	fmt.Println("Inserções em ordem crescente")
	fmt.Println()
	fmt.Println("       n | alt. ABB | alt. AVL | alt. RN | rot. AVL | rot. RN | recolor. RN")
	fmt.Println("---------+----------+----------+---------+----------+---------+------------")
	for _, n := range tamanhos {
		hABB, hAVL, hRN, rotAVL, rotRN, recRN := constroi(crescente(n))
		fmt.Printf("%8d | %8d | %8d | %7d | %8d | %7d | %11d\n",
			n, hABB, hAVL, hRN, rotAVL, rotRN, recRN)
	}

	fmt.Println()
	fmt.Println("Inserções em ordem aleatória")
	fmt.Println()
	fmt.Println("       n | alt. ABB | alt. AVL | alt. RN | rot. AVL | rot. RN | recolor. RN")
	fmt.Println("---------+----------+----------+---------+----------+---------+------------")
	for _, n := range tamanhos {
		hABB, hAVL, hRN, rotAVL, rotRN, recRN := constroi(aleatoria(n, r))
		fmt.Printf("%8d | %8d | %8d | %7d | %8d | %7d | %11d\n",
			n, hABB, hAVL, hRN, rotAVL, rotRN, recRN)
	}

	fmt.Println()
	fmt.Println("Limites teóricos de altura, para conferir as colunas acima:")
	for _, n := range tamanhos {
		fmt.Printf("  n = %6d:  AVL < %.1f      rubro-negra <= %.1f\n",
			n, 1.44*log2(float64(n)+2), 2*log2(float64(n)+1))
	}
}

func log2(x float64) float64 {
	e := 0.0
	for x > 1 {
		x /= 2
		e++
	}
	return e
}
