// Pacote avl: árvore binária de busca com rebalanceamento.
//
// O código é o da Aula 06, com duas diferenças, ambas tratadas na Parte
// 2 da atividade:
//
//  1. RotacaoEsquerda está por escrever;
//  2. RotacaoDireita tem um defeito de uma linha.
package avl

import "fmt"

// No guarda, além da informação e das duas subárvores, a altura da
// subárvore que começa nele.
type No struct {
	Info int
	Alt  int
	Esq  *No
	Dir  *No
}

// altura devolve a altura armazenada no campo Alt, tratando a árvore
// vazia como -1. Não recalcula nada.
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
// armazenadas nos filhos. Depende, portanto, de os filhos estarem com o
// campo Alt correto no momento da chamada.
func atualizaAltura(a *No) {
	a.Alt = 1 + maior(altura(a.Esq), altura(a.Dir))
}

// Fator é o fator de balanceamento: altura da subárvore esquerda menos a
// altura da direita. Em uma árvore AVL, todo nó tem fator -1, 0 ou 1.
func Fator(a *No) int {
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
// ATENÇÃO: esta função tem um defeito de uma linha. As três atribuições
// de ponteiro estão corretas, e a árvore resultante tem a forma certa;
// o que sai errado é o campo Alt. Encontre o defeito com a sua
// VerificaAVL e corrija-o.
func RotacaoDireita(a *No) *No {
	b := a.Esq
	a.Esq = b.Dir
	b.Dir = a

	atualizaAltura(b)
	atualizaAltura(a)
	return b
}

// RotacaoEsquerda é a operação simétrica da rotação à direita.
//
//	  a                  b
//	 / \                / \
//	x   b      -->     a   z
//	   / \            / \
//	  y   z          x   y
//
// Escreva-a espelhando a rotação à direita, trocando Esq por Dir. Cuide
// da ordem das atribuições e da ordem das atualizações de altura, que é
// justamente onde está o defeito da outra.
func RotacaoEsquerda(a *No) *No {
	panic("TODO: implementar RotacaoEsquerda")
}

// rebalanceia devolve a raiz da subárvore já corrigida. São os quatro
// casos da Aula 06, distinguidos pelo sinal do fator do nó e pelo sinal
// do fator do filho do lado mais pesado.
func rebalanceia(a *No) *No {
	fb := Fator(a)

	if fb > 1 { // pesada à esquerda
		if Fator(a.Esq) < 0 { // caso esquerda-direita: endireita o filho antes
			a.Esq = RotacaoEsquerda(a.Esq)
		}
		return RotacaoDireita(a) // caso esquerda-esquerda
	}

	if fb < -1 { // pesada à direita
		if Fator(a.Dir) > 0 { // caso direita-esquerda
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

// Remove é a remoção da ABB seguida de atualização de altura e
// rebalanceamento em cada nó do caminho de volta.
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

// Busca procura um valor descendo por um único caminho.
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

// EmOrdem devolve as chaves em ordem crescente.
func EmOrdem(a *No) []int {
	if a == nil {
		return nil
	}
	valores := EmOrdem(a.Esq)
	valores = append(valores, a.Info)
	return append(valores, EmOrdem(a.Dir)...)
}

// AlturaArmazenada devolve o valor do campo Alt da raiz, sem
// recalcular nada.
func AlturaArmazenada(a *No) int {
	return altura(a)
}

// Imprime desenha a árvore deitada, com a raiz à esquerda, mostrando ao
// lado de cada nó o campo Alt e o fator de balanceamento.
func Imprime(a *No, profundidade int) {
	if a == nil {
		return
	}
	Imprime(a.Dir, profundidade+1)
	for i := 0; i < profundidade; i++ {
		fmt.Print("    ")
	}
	fmt.Printf("%d (Alt %d, fb %d)\n", a.Info, a.Alt, Fator(a))
	Imprime(a.Esq, profundidade+1)
}
