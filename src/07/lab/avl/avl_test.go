package avl

// Testes fornecidos. NÃO ALTERE ESTE ARQUIVO: ele é o mesmo usado na
// correção.
//
// Estes testes não conferem o campo Alt das árvores construídas por
// Insere. Passar aqui, portanto, não é prova de que a AVL está
// consistente: essa conferência é a VerificaAVL que você vai escrever.

import (
	"reflect"
	"sort"
	"testing"
)

func TestRotacaoEsquerda(t *testing.T) {
	// Árvore com as alturas corretas antes da rotação:
	//
	//	   10 (Alt 2)
	//	  /  \
	//	 5    20 (Alt 1)
	//	     /  \
	//	   15    30
	a := &No{
		Info: 10, Alt: 2,
		Esq: &No{Info: 5, Alt: 0},
		Dir: &No{
			Info: 20, Alt: 1,
			Esq: &No{Info: 15, Alt: 0},
			Dir: &No{Info: 30, Alt: 0},
		},
	}

	raiz := RotacaoEsquerda(a)

	if raiz == nil || raiz.Info != 20 {
		t.Fatalf("a raiz depois da rotação deveria ser o 20")
	}
	if raiz.Esq == nil || raiz.Esq.Info != 10 {
		t.Fatalf("o 10 deveria ter descido para a esquerda do 20")
	}
	if raiz.Esq.Dir == nil || raiz.Esq.Dir.Info != 15 {
		t.Errorf("o 15 deveria ter virado filho direito do 10")
	}
	esperado := []int{5, 10, 15, 20, 30}
	if obtido := EmOrdem(raiz); !reflect.DeepEqual(obtido, esperado) {
		t.Errorf("in-ordem depois da rotação = %v, esperado %v", obtido, esperado)
	}
	if raiz.Esq.Alt != 1 {
		t.Errorf("Alt do 10 depois da rotação = %d, esperado 1", raiz.Esq.Alt)
	}
	if raiz.Alt != 2 {
		t.Errorf("Alt do 20 depois da rotação = %d, esperado 2", raiz.Alt)
	}
}

func TestInsereMantemOrdem(t *testing.T) {
	var a *No
	for _, v := range []int{50, 30, 90, 20, 40, 95, 10, 35, 45} {
		a = Insere(a, v)
	}
	esperado := []int{10, 20, 30, 35, 40, 45, 50, 90, 95}
	if obtido := EmOrdem(a); !reflect.DeepEqual(obtido, esperado) {
		t.Errorf("EmOrdem = %v, esperado %v", obtido, esperado)
	}
	for _, v := range esperado {
		if !Busca(a, v) {
			t.Errorf("Busca(a, %d) = false, esperado true", v)
		}
	}
	if Busca(a, 60) {
		t.Error("Busca(a, 60) = true, esperado false")
	}
}

func TestInsereCrescenteMantemOrdem(t *testing.T) {
	var a *No
	var esperado []int
	for v := 1; v <= 200; v++ {
		a = Insere(a, v)
		esperado = append(esperado, v)
	}
	if obtido := EmOrdem(a); !reflect.DeepEqual(obtido, esperado) {
		t.Errorf("in-ordem de 1..200 saiu diferente da sequência inserida")
	}
}

func TestRemoveMantemOrdem(t *testing.T) {
	var a *No
	valores := []int{50, 30, 90, 20, 40, 95, 10, 35, 45}
	for _, v := range valores {
		a = Insere(a, v)
	}
	for _, v := range []int{10, 90, 50} {
		a = Remove(a, v)
	}
	esperado := []int{20, 30, 35, 40, 45, 95}
	if obtido := EmOrdem(a); !reflect.DeepEqual(obtido, esperado) {
		t.Errorf("EmOrdem depois das remoções = %v, esperado %v", obtido, esperado)
	}
	if !sort.IntsAreSorted(EmOrdem(a)) {
		t.Error("o in-ordem deixou de sair crescente")
	}
}

// avlCorreta devolve uma AVL montada à mão, com todos os campos Alt
// corretos.
//
//	   20 (Alt 2)
//	  /  \
//	10    30 (Alt 1)
//	        \
//	         40
func avlCorreta() *No {
	return &No{
		Info: 20, Alt: 2,
		Esq: &No{Info: 10, Alt: 0},
		Dir: &No{
			Info: 30, Alt: 1,
			Dir: &No{Info: 40, Alt: 0},
		},
	}
}

func TestVerificaAVLAceita(t *testing.T) {
	if ok, culpado := VerificaAVL(nil); !ok {
		t.Errorf("VerificaAVL(nil) = false (culpado %v), esperado true", culpado)
	}
	if ok, culpado := VerificaAVL(&No{Info: 7, Alt: 0}); !ok {
		t.Errorf("VerificaAVL de um nó só = false (culpado %v), esperado true", culpado)
	}
	if ok, culpado := VerificaAVL(avlCorreta()); !ok {
		t.Errorf("VerificaAVL da árvore correta = false (culpado %v), esperado true", culpado)
	}
}

func TestVerificaAVLPegaAlturaErrada(t *testing.T) {
	a := avlCorreta()
	a.Alt = 5 // só a raiz fica com o campo Alt errado

	ok, culpado := VerificaAVL(a)
	if ok {
		t.Fatal("VerificaAVL = true com a raiz guardando Alt 5 em vez de 2")
	}
	if culpado == nil || culpado.Info != 20 {
		t.Errorf("nó acusado = %v, esperado o 20", culpado)
	}
}

func TestVerificaAVLPegaDesbalanceamento(t *testing.T) {
	// Alturas corretas, mas a raiz tem fator 2.
	//
	//	     30 (Alt 2)
	//	    /
	//	  20 (Alt 1)
	//	  /
	//	10
	a := &No{
		Info: 30, Alt: 2,
		Esq: &No{
			Info: 20, Alt: 1,
			Esq: &No{Info: 10, Alt: 0},
		},
	}

	ok, culpado := VerificaAVL(a)
	if ok {
		t.Fatal("VerificaAVL = true para a árvore com fator 2 na raiz")
	}
	if culpado == nil || culpado.Info != 30 {
		t.Errorf("nó acusado = %v, esperado o 30", culpado)
	}
}
