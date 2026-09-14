package abb

// EhABB verifica se uma árvore binária qualquer satisfaz a propriedade
// de busca: em todo nó, as chaves da subárvore esquerda são menores que
// a do nó, e as da direita são maiores.
//
// A exigência vale para todos os nós, e não só para cada nó em relação
// aos seus dois filhos. A árvore
//
//	   50
//	  /  \
//	30    90
//	  \
//	   60
//
// tem o 60 maior que o pai 30 e mesmo assim não é ABB, porque o 60 está
// na subárvore esquerda do 50. Uma verificação que só compare pai e
// filhos aceita essa árvore, e o teste fornecido cobra a rejeição.
//
// Sugestão: escreva uma função auxiliar recursiva que receba, além do
// nó, os limites inferior e superior que a chave dele precisa respeitar,
// e que desça trocando o limite correspondente pela chave do nó.
// A árvore vazia é ABB.
func EhABB(a *No) bool {
	panic("TODO: implementar EhABB")
}

// Sucessor devolve a menor chave da árvore que é estritamente maior que
// v, e um booleano dizendo se ela existe. O valor v não precisa estar na
// árvore.
//
// Exemplos, na árvore construída com 50 30 90 20 40 95 10 35 45:
//
//	Sucessor(a, 35) = 40, true
//	Sucessor(a, 36) = 40, true   (v não está na árvore)
//	Sucessor(a, 95) =  0, false  (v é o máximo)
//
// Sugestão: desça como na busca, guardando o último nó em que você foi
// para a esquerda. Não percorra a árvore inteira: o custo esperado é
// O(h).
func Sucessor(a *No, v int) (int, bool) {
	panic("TODO: implementar Sucessor")
}
