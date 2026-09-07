// Remoção em árvore binária de busca, com os três casos separados.
//
// Execute com:  go run remocao.go

package main

import "fmt"

type No struct {
	Info int
	Esq  *No
	Dir  *No
}

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

// Remove tira um valor da árvore e devolve a raiz da árvore resultante.
// A primeira parte da função é a mesma busca da inserção: descer até
// encontrar o nó. A remoção em si tem três casos.
func Remove(a *No, v int) *No {
	if a == nil {
		return nil
	}

	if v < a.Info {
		a.Esq = Remove(a.Esq, v)
		return a
	}
	if v > a.Info {
		a.Dir = Remove(a.Dir, v)
		return a
	}

	// Caso 1: o nó é uma folha. Some, e o pai passa a apontar para nil.
	if a.Esq == nil && a.Dir == nil {
		return nil
	}

	// Caso 2: o nó tem uma subárvore só. O filho sobe para o lugar dele.
	if a.Esq == nil {
		return a.Dir
	}
	if a.Dir == nil {
		return a.Esq
	}

	// Caso 3: o nó tem as duas subárvores. Copiamos para ele o maior
	// valor da subárvore da esquerda, o antecessor, e removemos esse
	// valor da subárvore da esquerda. O antecessor não tem filho à
	// direita, então essa segunda remoção cai no caso 1 ou no caso 2.
	antecessor := a.Esq
	for antecessor.Dir != nil {
		antecessor = antecessor.Dir
	}
	a.Info = antecessor.Info
	a.Esq = Remove(a.Esq, antecessor.Info)
	return a
}

func InOrdem(a *No) {
	if a == nil {
		return
	}
	InOrdem(a.Esq)
	fmt.Print(a.Info, " ")
	InOrdem(a.Dir)
}

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

	fmt.Println("Árvore inicial:")
	Imprime(a, 0)

	for _, caso := range []struct {
		valor     int
		descricao string
	}{
		{45, "caso 1: folha"},
		{20, "caso 2: uma subárvore só"},
		{50, "caso 3: duas subárvores (a raiz)"},
	} {
		fmt.Printf("\nRemove(a, %d), %s\n", caso.valor, caso.descricao)
		a = Remove(a, caso.valor)
		Imprime(a, 0)
		fmt.Print("in-ordem: ")
		InOrdem(a)
		fmt.Println()
	}
}
