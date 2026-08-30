// Os quatro percursos de uma árvore binária: pré-ordem, in-ordem,
// pós-ordem e em largura.
//
// Execute com:  go run percursos.go

package main

import "fmt"

type No struct {
	Info int
	Esq  *No
	Dir  *No
}

func Constroi(info int, esq, dir *No) *No {
	return &No{Info: info, Esq: esq, Dir: dir}
}

func Folha(info int) *No {
	return Constroi(info, nil, nil)
}

func ArvoreExemplo() *No {
	return Constroi(50,
		Constroi(30,
			Constroi(20, Folha(10), nil),
			Constroi(40, Folha(35), Folha(45))),
		Constroi(90,
			nil,
			Folha(95)))
}

// PreOrdem visita a raiz antes das duas subárvores (R, E, D).
func PreOrdem(a *No) {
	if a == nil {
		return
	}
	fmt.Print(a.Info, " ")
	PreOrdem(a.Esq)
	PreOrdem(a.Dir)
}

// InOrdem visita a raiz entre as duas subárvores (E, R, D).
func InOrdem(a *No) {
	if a == nil {
		return
	}
	InOrdem(a.Esq)
	fmt.Print(a.Info, " ")
	InOrdem(a.Dir)
}

// PosOrdem visita a raiz depois das duas subárvores (E, D, R).
func PosOrdem(a *No) {
	if a == nil {
		return
	}
	PosOrdem(a.Esq)
	PosOrdem(a.Dir)
	fmt.Print(a.Info, " ")
}

// EmLargura visita os nós nível a nível, da esquerda para a direita.
// Não é recursivo: usa uma fila com os nós ainda por visitar. A cada
// nó retirado da frente da fila, os seus filhos entram no fim.
func EmLargura(a *No) {
	if a == nil {
		return
	}

	fila := []*No{a}

	for len(fila) > 0 {
		atual := fila[0]
		fila = fila[1:]

		fmt.Print(atual.Info, " ")

		if atual.Esq != nil {
			fila = append(fila, atual.Esq)
		}
		if atual.Dir != nil {
			fila = append(fila, atual.Dir)
		}
	}
}

// PreOrdemComPilha faz o mesmo que PreOrdem, sem recursão. A pilha
// explícita ocupa o lugar da pilha de chamadas do programa. O filho da
// direita é empilhado primeiro para que o da esquerda saia antes.
func PreOrdemComPilha(a *No) {
	if a == nil {
		return
	}

	pilha := []*No{a}

	for len(pilha) > 0 {
		topo := pilha[len(pilha)-1]
		pilha = pilha[:len(pilha)-1]

		fmt.Print(topo.Info, " ")

		if topo.Dir != nil {
			pilha = append(pilha, topo.Dir)
		}
		if topo.Esq != nil {
			pilha = append(pilha, topo.Esq)
		}
	}
}

func main() {
	a := ArvoreExemplo()

	fmt.Print("pré-ordem  (R,E,D): ")
	PreOrdem(a)
	fmt.Println()

	fmt.Print("in-ordem   (E,R,D): ")
	InOrdem(a)
	fmt.Println()

	fmt.Print("pós-ordem  (E,D,R): ")
	PosOrdem(a)
	fmt.Println()

	fmt.Print("em largura        : ")
	EmLargura(a)
	fmt.Println()

	fmt.Print("pré-ordem c/ pilha: ")
	PreOrdemComPilha(a)
	fmt.Println()
}
