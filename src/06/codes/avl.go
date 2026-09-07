// Árvore AVL: inserção com rotações, altura garantida em O(log n).
//
// Execute com:  go run avl.go

package main

import "fmt"

// No guarda, além da informação e das duas subárvores, a altura da
// subárvore que começa nele. Guardar a altura evita recalculá-la a cada
// verificação de balanceamento, que passa a custar tempo constante.
type No struct {
	Info int
	Alt  int
	Esq  *No
	Dir  *No
}

// altura devolve a altura armazenada, tratando a árvore vazia como -1.
func altura(a *No) int {
	if a == nil {
		return -1
	}
	return a.Alt
}

func maior(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// atualizaAltura recalcula a altura de um nó a partir das alturas já
// armazenadas nos filhos.
func atualizaAltura(a *No) {
	a.Alt = 1 + maior(altura(a.Esq), altura(a.Dir))
}

// fator é o fator de balanceamento: altura da subárvore esquerda menos
// a altura da direita. Em uma árvore AVL, todo nó tem fator -1, 0 ou 1.
func fator(a *No) int {
	if a == nil {
		return 0
	}
	return altura(a.Esq) - altura(a.Dir)
}

// RotacaoDireita sobe o filho da esquerda e desce a raiz para a direita.
//
//	    a              b
//	   / \            / \
//	  b   z    -->   x   a
//	 / \                / \
//	x   y              y   z
//
// A ordem in-ordem x b y a z é a mesma antes e depois: a rotação muda a
// forma da árvore sem violar a propriedade de busca.
func RotacaoDireita(a *No) *No {
	b := a.Esq
	a.Esq = b.Dir
	b.Dir = a

	atualizaAltura(a)
	atualizaAltura(b)
	return b
}

// RotacaoEsquerda é a operação simétrica.
//
//	  a                  b
//	 / \                / \
//	x   b      -->     a   z
//	   / \            / \
//	  y   z          x   y
func RotacaoEsquerda(a *No) *No {
	b := a.Dir
	a.Dir = b.Esq
	b.Esq = a

	atualizaAltura(a)
	atualizaAltura(b)
	return b
}

// rebalanceia devolve a raiz da subárvore já corrigida. São quatro
// casos, distinguidos pelo sinal do fator do nó e do sinal do fator do
// filho do lado mais pesado.
func rebalanceia(a *No) *No {
	fb := fator(a)

	if fb > 1 { // pesada à esquerda
		if fator(a.Esq) < 0 { // caso esquerda-direita: endireita o filho antes
			a.Esq = RotacaoEsquerda(a.Esq)
		}
		return RotacaoDireita(a) // caso esquerda-esquerda
	}

	if fb < -1 { // pesada à direita
		if fator(a.Dir) > 0 { // caso direita-esquerda
			a.Dir = RotacaoDireita(a.Dir)
		}
		return RotacaoEsquerda(a) // caso direita-direita
	}

	return a
}

// Insere é a inserção da ABB com duas linhas a mais: ao voltar da
// recursão, cada nó do caminho tem sua altura atualizada e é
// rebalanceado se preciso.
func Insere(a *No, v int) *No {
	if a == nil {
		return &No{Info: v}
	}

	if v < a.Info {
		a.Esq = Insere(a.Esq, v)
	} else if v > a.Info {
		a.Dir = Insere(a.Dir, v)
	} else {
		return a
	}

	atualizaAltura(a)
	return rebalanceia(a)
}

// Remove segue a mesma ideia: a remoção da ABB, seguida de atualização
// de altura e rebalanceamento em cada nó do caminho de volta.
func Remove(a *No, v int) *No {
	if a == nil {
		return nil
	}

	if v < a.Info {
		a.Esq = Remove(a.Esq, v)
	} else if v > a.Info {
		a.Dir = Remove(a.Dir, v)
	} else {
		if a.Esq == nil {
			return a.Dir
		}
		if a.Dir == nil {
			return a.Esq
		}
		antecessor := a.Esq
		for antecessor.Dir != nil {
			antecessor = antecessor.Dir
		}
		a.Info = antecessor.Info
		a.Esq = Remove(a.Esq, antecessor.Info)
	}

	atualizaAltura(a)
	return rebalanceia(a)
}

// insereSemBalancear é a inserção da ABB, sem rotações. Serve para
// mostrar, lado a lado, a forma que a árvore teria sem o
// rebalanceamento.
func insereSemBalancear(a *No, v int) *No {
	if a == nil {
		return &No{Info: v}
	}
	if v < a.Info {
		a.Esq = insereSemBalancear(a.Esq, v)
	} else if v > a.Info {
		a.Dir = insereSemBalancear(a.Dir, v)
	}
	atualizaAltura(a)
	return a
}

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

func Imprime(a *No, profundidade int) {
	if a == nil {
		return
	}
	Imprime(a.Dir, profundidade+1)
	for i := 0; i < profundidade; i++ {
		fmt.Print("    ")
	}
	fmt.Printf("%d (fb %d)\n", a.Info, fator(a))
	Imprime(a.Esq, profundidade+1)
}

func InOrdem(a *No) {
	if a == nil {
		return
	}
	InOrdem(a.Esq)
	fmt.Print(a.Info, " ")
	InOrdem(a.Dir)
}

func main() {
	fmt.Println("Os quatro casos, cada um com a sua sequência de inserção:")
	casos := []struct {
		nome      string
		sequencia []int
	}{
		{"esquerda-esquerda (rotação simples à direita)", []int{30, 20, 10}},
		{"direita-direita (rotação simples à esquerda)", []int{10, 20, 30}},
		{"esquerda-direita (rotação dupla)", []int{30, 10, 20}},
		{"direita-esquerda (rotação dupla)", []int{10, 30, 20}},
	}
	for _, caso := range casos {
		var semBalanceamento, comBalanceamento *No
		for _, v := range caso.sequencia {
			semBalanceamento = insereSemBalancear(semBalanceamento, v)
			comBalanceamento = Insere(comBalanceamento, v)
		}
		fmt.Printf("\n%v: %s\n", caso.sequencia, caso.nome)
		fmt.Printf("  sem rebalanceamento (altura %d):\n", altura(semBalanceamento))
		Imprime(semBalanceamento, 1)
		fmt.Printf("  depois da rotação (altura %d):\n", altura(comBalanceamento))
		Imprime(comBalanceamento, 1)
	}

	fmt.Println("\nA árvore de exemplo da aula, inserida na mesma ordem:")
	var c *No
	for _, v := range []int{50, 30, 90, 20, 40, 95, 10, 35, 45} {
		c = Insere(c, v)
	}
	Imprime(c, 0)
	fmt.Print("in-ordem: ")
	InOrdem(c)
	fmt.Println("\naltura:", altura(c))

	fmt.Println("\nInserindo 1, 2, 3, ..., n em ordem crescente:")
	fmt.Println("       n | altura na ABB | altura na AVL")
	fmt.Println("---------+---------------+--------------")
	for _, n := range []int{10, 100, 1000, 10000} {
		var avl *No
		for v := 1; v <= n; v++ {
			avl = Insere(avl, v)
		}
		fmt.Printf("%8d | %13d | %13d\n", n, n-1, altura(avl))
	}
}
