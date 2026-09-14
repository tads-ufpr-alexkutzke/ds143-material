package abb

// Testes fornecidos. NÃO ALTERE ESTE ARQUIVO: ele é o mesmo usado na
// correção.

import (
	"reflect"
	"testing"
)

// exemplo devolve a árvore da Aula 06, construída por inserções.
func exemplo() *No {
	var a *No
	for _, v := range []int{50, 30, 90, 20, 40, 95, 10, 35, 45} {
		a = Insere(a, v)
	}
	return a
}

func TestEmOrdemSaiCrescente(t *testing.T) {
	esperado := []int{10, 20, 30, 35, 40, 45, 50, 90, 95}
	if obtido := EmOrdem(exemplo()); !reflect.DeepEqual(obtido, esperado) {
		t.Errorf("EmOrdem = %v, esperado %v", obtido, esperado)
	}
}

func TestEhABBAceita(t *testing.T) {
	casos := []struct {
		nome   string
		arvore *No
	}{
		{"vazia", nil},
		{"um nó", &No{Info: 42}},
		{"exemplo da aula", exemplo()},
		{"degenerada à direita", Insere(Insere(Insere(nil, 1), 2), 3)},
	}
	for _, caso := range casos {
		if !EhABB(caso.arvore) {
			t.Errorf("EhABB(%s) = false, esperado true", caso.nome)
		}
	}
}

func TestEhABBRejeitaViolacaoDistante(t *testing.T) {
	// O 60 é maior que o pai 30, então a comparação local não acusa
	// nada. Ele está, porém, na subárvore esquerda do 50.
	//
	//	   50
	//	  /  \
	//	30    90
	//	  \
	//	   60
	a := &No{
		Info: 50,
		Esq:  &No{Info: 30, Dir: &No{Info: 60}},
		Dir:  &No{Info: 90},
	}
	if EhABB(a) {
		t.Error("EhABB = true para a árvore com o 60 à esquerda do 50, esperado false")
	}
}

func TestEhABBRejeitaViolacaoLocal(t *testing.T) {
	// Aqui basta comparar o 40 com o filho 45.
	a := &No{
		Info: 40,
		Esq:  &No{Info: 45},
		Dir:  &No{Info: 90},
	}
	if EhABB(a) {
		t.Error("EhABB = true para a árvore com o 45 à esquerda do 40, esperado false")
	}
}

func TestSucessor(t *testing.T) {
	a := exemplo()
	casos := []struct {
		v        int
		esperado int
		existe   bool
	}{
		{35, 40, true},
		{36, 40, true},  // v não está na árvore
		{5, 10, true},   // v é menor que todas as chaves
		{50, 90, true},  // v está na raiz
		{90, 95, true},  // sucessor está à direita
		{95, 0, false},  // v é o máximo
		{100, 0, false}, // v é maior que todas as chaves
	}
	for _, caso := range casos {
		obtido, existe := Sucessor(a, caso.v)
		if existe != caso.existe {
			t.Errorf("Sucessor(a, %d): existe = %v, esperado %v", caso.v, existe, caso.existe)
			continue
		}
		if existe && obtido != caso.esperado {
			t.Errorf("Sucessor(a, %d) = %d, esperado %d", caso.v, obtido, caso.esperado)
		}
	}
}

func TestSucessorArvoreVazia(t *testing.T) {
	if _, existe := Sucessor(nil, 10); existe {
		t.Error("Sucessor(nil, 10): existe = true, esperado false")
	}
}
