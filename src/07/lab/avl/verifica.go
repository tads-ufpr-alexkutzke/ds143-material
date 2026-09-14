package avl

// AlturaRecalculada percorre a árvore inteira e devolve a altura real,
// contada em arestas, sem olhar para o campo Alt de nenhum nó. A árvore
// vazia tem altura -1.
//
// Esta função está pronta. Ela é a Altura da Aula 05, e é o que permite
// desconfiar do campo Alt: se os dois valores divergem em algum nó, a
// árvore está inconsistente.
func AlturaRecalculada(a *No) int {
	if a == nil {
		return -1
	}
	return 1 + maior(AlturaRecalculada(a.Esq), AlturaRecalculada(a.Dir))
}

// VerificaAVL confere, em cada nó da árvore, duas condições:
//
//  1. o campo Alt guarda a altura real da subárvore que começa no nó;
//  2. o fator de balanceamento do nó, calculado a partir das alturas
//     reais, está em {-1, 0, 1}.
//
// Devolve true e nil quando a árvore inteira passa. Quando alguma
// condição falha, devolve false e o ponteiro para um nó em que a
// verificação falhou, para que quem chamou possa imprimi-lo.
//
// A árvore vazia passa.
//
// Sugestão: a condição 1 é uma comparação entre a.Alt e
// AlturaRecalculada(a). Escrita assim, de forma direta, a verificação
// custa O(n²); para esta atividade isso basta, e as árvores dos testes
// são pequenas. Se quiser a versão O(n), faça a recursão devolver a
// altura recalculada junto com o resultado.
func VerificaAVL(a *No) (bool, *No) {
	if a == nil {
		return true, nil
	}
	panic("TODO: implementar VerificaAVL")
}
