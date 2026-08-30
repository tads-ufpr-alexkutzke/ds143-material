// Estrutura de árvore binária e construção da árvore usada na aula.
//
// Execute com:  go run arvore.go

package main

import "fmt"

// No é um nó da árvore binária. A árvore inteira é representada pelo
// ponteiro para o seu nó raiz: um *No é, ao mesmo tempo, um nó e a
// subárvore que começa nele. A árvore vazia é o ponteiro nil.
type No struct {
	Info int
	Esq  *No
	Dir  *No
}

// Constroi devolve um nó com a informação e as duas subárvores dadas.
// Passar nil nas duas subárvores cria uma folha.
func Constroi(info int, esq, dir *No) *No {
	return &No{Info: info, Esq: esq, Dir: dir}
}

// Folha é um atalho para Constroi(info, nil, nil).
func Folha(info int) *No {
	return Constroi(info, nil, nil)
}

// ArvoreExemplo monta a árvore usada no restante da aula:
//
//	          50
//	        /    \
//	      30      90
//	     /  \       \
//	   20    40      95
//	  /     /  \
//	10    35    45
func ArvoreExemplo() *No {
	return Constroi(50,
		Constroi(30,
			Constroi(20,
				Folha(10),
				nil),
			Constroi(40,
				Folha(35),
				Folha(45))),
		Constroi(90,
			nil,
			Folha(95)))
}

// Imprime desenha a árvore deitada, com um nível de indentação por
// nível da árvore. A subárvore da direita é impressa antes da raiz e a
// da esquerda depois, de modo que a saída fica com a raiz à esquerda e
// as folhas à direita.
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

// Pertence devolve true se o valor procurado está em algum nó da
// árvore. Sem nenhuma hipótese sobre a organização dos valores, não há
// como evitar visitar todos os nós no pior caso.
func Pertence(a *No, procurado int) bool {
	if a == nil {
		return false
	}
	if a.Info == procurado {
		return true
	}
	return Pertence(a.Esq, procurado) || Pertence(a.Dir, procurado)
}

func main() {
	a := ArvoreExemplo()

	fmt.Println("Árvore de exemplo (raiz à esquerda):")
	Imprime(a, 0)

	fmt.Println()
	fmt.Println("Pertence(a, 35) =", Pertence(a, 35))
	fmt.Println("Pertence(a, 60) =", Pertence(a, 60))
}
