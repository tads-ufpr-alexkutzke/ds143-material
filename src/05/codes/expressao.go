// Árvore de expressão: a mesma árvore percorrida de três maneiras
// produz as três notações aritméticas, e a avaliação é um percurso em
// pós-ordem.
//
// Execute com:  go run expressao.go

package main

import (
	"fmt"
	"strconv"
)

// No guarda um operador ("+", "*") nos nós internos e um número nas
// folhas, sempre como texto.
type No struct {
	Info string
	Esq  *No
	Dir  *No
}

func Constroi(info string, esq, dir *No) *No {
	return &No{Info: info, Esq: esq, Dir: dir}
}

func Numero(n int) *No {
	return Constroi(strconv.Itoa(n), nil, nil)
}

// ExpressaoExemplo monta a árvore de 2 * (3 + 4):
//
//	  *
//	 / \
//	2   +
//	   / \
//	  3   4
func ExpressaoExemplo() *No {
	return Constroi("*",
		Numero(2),
		Constroi("+",
			Numero(3),
			Numero(4)))
}

// Prefixa devolve o percurso em pré-ordem: a notação lida pelo parser
// recursivo da aula de recursão.
func Prefixa(a *No) string {
	if a == nil {
		return ""
	}
	if a.Esq == nil && a.Dir == nil {
		return a.Info
	}
	return a.Info + " " + Prefixa(a.Esq) + " " + Prefixa(a.Dir)
}

// InOrdem devolve o percurso in-ordem puro, sem nenhum acréscimo. O
// resultado é ambíguo: lido da esquerda para a direita com a
// precedência usual, "2 * 3 + 4" vale 10, e não 14.
func InOrdem(a *No) string {
	if a == nil {
		return ""
	}
	if a.Esq == nil && a.Dir == nil {
		return a.Info
	}
	return InOrdem(a.Esq) + " " + a.Info + " " + InOrdem(a.Dir)
}

// Infixa é o mesmo percurso in-ordem com um par de parênteses em volta
// de cada nó interno. Os parênteses recuperam a informação de estrutura
// que o percurso sozinho perde.
func Infixa(a *No) string {
	if a == nil {
		return ""
	}
	if a.Esq == nil && a.Dir == nil {
		return a.Info
	}
	return "(" + Infixa(a.Esq) + " " + a.Info + " " + Infixa(a.Dir) + ")"
}

// PosFixa devolve o percurso em pós-ordem, a notação usada por
// calculadoras de pilha.
func PosFixa(a *No) string {
	if a == nil {
		return ""
	}
	if a.Esq == nil && a.Dir == nil {
		return a.Info
	}
	return PosFixa(a.Esq) + " " + PosFixa(a.Dir) + " " + a.Info
}

// Avalia calcula o valor da expressão. O operador só pode ser aplicado
// depois que os dois operandos são conhecidos, então a ordem das linhas
// aqui não é escolha de estilo: é a única que funciona.
func Avalia(a *No) int {
	if a.Esq == nil && a.Dir == nil {
		valor, _ := strconv.Atoi(a.Info)
		return valor
	}

	esq := Avalia(a.Esq)
	dir := Avalia(a.Dir)

	switch a.Info {
	case "+":
		return esq + dir
	case "*":
		return esq * dir
	}
	return 0
}

func main() {
	e := ExpressaoExemplo()

	fmt.Println("pré-ordem (prefixa)       :", Prefixa(e))
	fmt.Println("in-ordem  (sem parênteses):", InOrdem(e))
	fmt.Println("in-ordem  (com parênteses):", Infixa(e))
	fmt.Println("pós-ordem (pós-fixa)      :", PosFixa(e))
	fmt.Println("valor                     :", Avalia(e))
}
