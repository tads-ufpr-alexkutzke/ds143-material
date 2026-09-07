// Árvore binária de busca: busca, inserção, mínimo, máximo e listagem
// em ordem crescente.
//
// Execute com:  go run abb.go

package main

import "fmt"

// No é um nó da árvore binária de busca. A estrutura é a mesma da aula
// de árvores e percursos: o que muda é a regra que os valores obedecem.
type No struct {
	Info int
	Esq  *No
	Dir  *No
}

// Busca procura um valor descendo por um único caminho: em cada nó, a
// comparação descarta uma das duas subárvores inteiras.
func Busca(a *No, procurado int) bool {
	if a == nil {
		return false
	}
	if procurado < a.Info {
		return Busca(a.Esq, procurado)
	}
	if procurado > a.Info {
		return Busca(a.Dir, procurado)
	}
	return true
}

// BuscaIterativa faz o mesmo percurso sem recursão. A busca em uma ABB
// nunca volta atrás, então um laço basta.
func BuscaIterativa(a *No, procurado int) bool {
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

// Insere acrescenta um valor na árvore e devolve a raiz da árvore
// resultante. O nó novo entra sempre como folha, na posição em que a
// busca por esse valor terminaria. Valores repetidos são ignorados.
func Insere(a *No, v int) *No {
	if a == nil {
		return &No{Info: v}
	}
	if v < a.Info {
		a.Esq = Insere(a.Esq, v)
	} else if v > a.Info {
		a.Dir = Insere(a.Dir, v)
	}
	return a
}

// Minimo devolve o menor valor da árvore, que está no nó mais à
// esquerda. Devolve 0 e false para a árvore vazia.
func Minimo(a *No) (int, bool) {
	if a == nil {
		return 0, false
	}
	for a.Esq != nil {
		a = a.Esq
	}
	return a.Info, true
}

// Maximo devolve o maior valor da árvore, que está no nó mais à direita.
func Maximo(a *No) (int, bool) {
	if a == nil {
		return 0, false
	}
	for a.Dir != nil {
		a = a.Dir
	}
	return a.Info, true
}

// InOrdem imprime os valores em ordem crescente.
func InOrdem(a *No) {
	if a == nil {
		return
	}
	InOrdem(a.Esq)
	fmt.Print(a.Info, " ")
	InOrdem(a.Dir)
}

// Altura devolve o comprimento do caminho mais longo até uma folha,
// contado em arestas. A árvore vazia tem altura -1.
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

// Imprime desenha a árvore deitada, com a raiz à esquerda.
func Imprime(a *No, profundidade int) {
	if a == nil {
		return
	}
	Imprime(a.Dir, profundidade+1)
	for i := 0; i < profundidade; i++ {
		fmt.Print("    ")
	}
	fmt.Println(a.Info)
	Imprime(a.Esq, profundidade+1)
}

func main() {
	var a *No
	for _, v := range []int{50, 30, 90, 20, 40, 95, 10, 35, 45} {
		a = Insere(a, v)
	}

	fmt.Println("Árvore construída por inserções (raiz à esquerda):")
	Imprime(a, 0)

	fmt.Print("\nin-ordem: ")
	InOrdem(a)
	fmt.Println()

	fmt.Println("\naltura:", Altura(a))

	menor, _ := Minimo(a)
	maior, _ := Maximo(a)
	fmt.Println("mínimo:", menor, " máximo:", maior)

	fmt.Println()
	for _, v := range []int{35, 60, 95} {
		fmt.Printf("Busca(a, %d) = %v\n", v, Busca(a, v))
	}
}
