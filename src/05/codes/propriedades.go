// Altura, contagem de nós e a diferença entre a árvore cheia e a
// árvore degenerada.
//
// Execute com:  go run propriedades.go

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

// ContaNos devolve o número de nós da árvore. Um nó, mais o que houver
// nas duas subárvores.
func ContaNos(a *No) int {
	if a == nil {
		return 0
	}
	return 1 + ContaNos(a.Esq) + ContaNos(a.Dir)
}

// ContaFolhas devolve o número de nós sem nenhum filho.
func ContaFolhas(a *No) int {
	if a == nil {
		return 0
	}
	if a.Esq == nil && a.Dir == nil {
		return 1
	}
	return ContaFolhas(a.Esq) + ContaFolhas(a.Dir)
}

// Altura devolve o comprimento do caminho mais longo da raiz até uma
// folha, contado em arestas. A árvore com um único nó tem altura 0, e a
// árvore vazia tem altura -1, valor escolhido para que a conta
// 1 + max(...) funcione também para a folha.
func Altura(a *No) int {
	if a == nil {
		return -1
	}

	esq := Altura(a.Esq)
	dir := Altura(a.Dir)

	if esq > dir {
		return 1 + esq
	}
	return 1 + dir
}

// ArvoreCheia monta uma árvore em que todo nó interno tem duas
// subárvores e todas as folhas estão no último nível.
func ArvoreCheia(altura int) *No {
	if altura < 0 {
		return nil
	}
	return Constroi(0, ArvoreCheia(altura-1), ArvoreCheia(altura-1))
}

// ArvoreDegenerada monta uma árvore com um único nó por nível, sempre
// à esquerda. É a mesma forma de uma lista encadeada.
func ArvoreDegenerada(n int) *No {
	if n <= 0 {
		return nil
	}
	return Constroi(0, ArvoreDegenerada(n-1), nil)
}

func main() {
	a := ArvoreExemplo()

	fmt.Println("Árvore de exemplo")
	fmt.Println("  nós    :", ContaNos(a))
	fmt.Println("  folhas :", ContaFolhas(a))
	fmt.Println("  altura :", Altura(a))

	fmt.Println()
	fmt.Println("Mesmo número de nós, duas formas:")
	fmt.Println()
	fmt.Printf("%8s | %8s | %10s\n", "nós", "cheia", "degenerada")
	fmt.Printf("%8s-+-%8s-+-%10s\n", "--------", "--------", "----------")

	for h := 3; h <= 9; h += 2 {
		cheia := ArvoreCheia(h)
		n := ContaNos(cheia)
		degenerada := ArvoreDegenerada(n)

		fmt.Printf("%8d | %8d | %10d\n", n, Altura(cheia), Altura(degenerada))
	}
}
