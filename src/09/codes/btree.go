// Árvore B de ordem 5: busca, inserção com divisão de nó cheio e impressão
// da árvore por níveis.
//
// Execute com:  go run btree.go

package main

import (
	"fmt"
	"strings"
)

// ordem é o número máximo de filhos de um nó. Com ordem 5, cada nó guarda
// no máximo 4 chaves.
const ordem = 5
const maxChaves = ordem - 1

// NoB é um nó da árvore B. As chaves ficam em ordem crescente dentro do nó,
// e Filhos tem sempre uma posição a mais que Chaves, ou está vazio.
type NoB struct {
	Chaves []int
	Filhos []*NoB
}

func ehFolha(n *NoB) bool {
	return len(n.Filhos) == 0
}

// posicao devolve o índice da primeira chave do nó que não é menor que v, e
// diz se essa chave é o próprio v. É a busca sequencial dentro do nó.
func posicao(n *NoB, v int) (int, bool) {
	i := 0
	for i < len(n.Chaves) && v > n.Chaves[i] {
		i++
	}
	return i, i < len(n.Chaves) && n.Chaves[i] == v
}

// Busca desce um nível por vez. Dentro de cada nó, a comparação escolhe
// entre len(Chaves)+1 caminhos, e não entre 2 como na árvore binária.
func Busca(n *NoB, v int) bool {
	if n == nil {
		return false
	}
	i, achou := posicao(n, v)
	if achou {
		return true
	}
	if ehFolha(n) {
		return false
	}
	return Busca(n.Filhos[i], v)
}

// divide separa um nó com maxChaves+1 chaves em dois e devolve a chave do
// meio, que sobe para o pai, junto com o nó da direita.
func divide(n *NoB) (int, *NoB) {
	meio := len(n.Chaves) / 2
	subiu := n.Chaves[meio]

	direito := &NoB{Chaves: append([]int{}, n.Chaves[meio+1:]...)}
	if !ehFolha(n) {
		direito.Filhos = append([]*NoB{}, n.Filhos[meio+1:]...)
		n.Filhos = n.Filhos[:meio+1]
	}
	n.Chaves = n.Chaves[:meio]

	return subiu, direito
}

// insereEm insere v na subárvore de n. Quando o nó estoura, devolve a chave
// que sobe, o nó da direita e true; caso contrário, devolve dividiu = false.
func insereEm(n *NoB, v int) (int, *NoB, bool) {
	i, achou := posicao(n, v)
	if achou {
		return 0, nil, false
	}

	if ehFolha(n) {
		insereChave(n, i, v)
	} else {
		subiu, direito, dividiu := insereEm(n.Filhos[i], v)
		if !dividiu {
			return 0, nil, false
		}
		insereChave(n, i, subiu)
		insereFilho(n, i+1, direito)
	}

	if len(n.Chaves) <= maxChaves {
		return 0, nil, false
	}
	subiu, direito := divide(n)
	return subiu, direito, true
}

// Insere acrescenta v à árvore. Quando a raiz se divide, a árvore ganha um
// nível: é o único jeito de uma árvore B crescer em altura.
func Insere(raiz *NoB, v int) *NoB {
	if raiz == nil {
		return &NoB{Chaves: []int{v}}
	}
	subiu, direito, dividiu := insereEm(raiz, v)
	if !dividiu {
		return raiz
	}
	return &NoB{Chaves: []int{subiu}, Filhos: []*NoB{raiz, direito}}
}

func insereChave(n *NoB, i, v int) {
	n.Chaves = append(n.Chaves, 0)
	copy(n.Chaves[i+1:], n.Chaves[i:])
	n.Chaves[i] = v
}

func insereFilho(n *NoB, i int, filho *NoB) {
	n.Filhos = append(n.Filhos, nil)
	copy(n.Filhos[i+1:], n.Filhos[i:])
	n.Filhos[i] = filho
}

// Altura conta os níveis abaixo da raiz. Todas as folhas de uma árvore B
// estão no mesmo nível, então basta descer por um caminho.
func Altura(n *NoB) int {
	if n == nil {
		return -1
	}
	h := 0
	for !ehFolha(n) {
		h++
		n = n.Filhos[0]
	}
	return h
}

// Imprime desenha a árvore por níveis, com um nó por linha.
func Imprime(n *NoB, nivel int) {
	if n == nil {
		return
	}
	partes := make([]string, len(n.Chaves))
	for i, c := range n.Chaves {
		partes[i] = fmt.Sprint(c)
	}
	fmt.Printf("%s[%s]\n", strings.Repeat("    ", nivel), strings.Join(partes, " "))
	for _, f := range n.Filhos {
		Imprime(f, nivel+1)
	}
}

// EmOrdem percorre a árvore alternando filho e chave, o que devolve as
// chaves em ordem crescente.
func EmOrdem(n *NoB, visita func(int)) {
	if n == nil {
		return
	}
	for i, c := range n.Chaves {
		if !ehFolha(n) {
			EmOrdem(n.Filhos[i], visita)
		}
		visita(c)
	}
	if !ehFolha(n) {
		EmOrdem(n.Filhos[len(n.Chaves)], visita)
	}
}

func main() {
	fmt.Printf("Árvore B de ordem %d: até %d chaves e %d filhos por nó\n\n",
		ordem, maxChaves, ordem)

	var raiz *NoB
	for v := 1; v <= 20; v++ {
		antes := Altura(raiz)
		raiz = Insere(raiz, v)
		if antes >= 0 && Altura(raiz) != antes {
			fmt.Printf("inserção do %2d: a raiz se dividiu, altura agora é %d\n",
				v, Altura(raiz))
		}
	}

	fmt.Println()
	fmt.Println("Árvore depois de inserir 1 a 20 em ordem crescente:")
	Imprime(raiz, 0)

	fmt.Print("\nin-ordem: ")
	EmOrdem(raiz, func(v int) { fmt.Print(v, " ") })
	fmt.Println()

	fmt.Printf("\nBusca(raiz, 13) = %v   Busca(raiz, 21) = %v\n",
		Busca(raiz, 13), Busca(raiz, 21))

	fmt.Println()
	fmt.Println("Altura mínima por ordem, com os nós cheios:")
	fmt.Println()
	fmt.Println("  ordem | chaves/nó |   1 mil |   1 mi | 1 bilhão")
	fmt.Println("--------+-----------+---------+--------+---------")
	for _, m := range []int{5, 101, 201, 1001} {
		fmt.Printf("  %5d | %9d | %7d | %6d | %8d\n",
			m, m-1, niveis(m, 1000), niveis(m, 1000000), niveis(m, 1000000000))
	}
	fmt.Println()
	fmt.Println("(número de níveis, ou seja, de acessos a disco por busca)")
}

// niveis devolve quantos níveis uma árvore B de ordem m precisa para guardar
// n chaves, no caso em que todos os nós estão cheios.
func niveis(m, n int) int {
	capacidade := 0
	nos := 1
	for h := 1; ; h++ {
		capacidade += nos * (m - 1)
		if capacidade >= n {
			return h
		}
		nos *= m
	}
}
