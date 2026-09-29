// Trie (árvore de prefixos): inserção, busca exata e listagem de todas as
// palavras que começam com um prefixo dado.
//
// Execute com:  go run trie.go

package main

import (
	"fmt"
	"sort"
	"strings"
)

// No é um nó da trie. A chave não fica guardada no nó: ela é o caminho da
// raiz até ele, uma letra por aresta. FimDePalavra marca os caminhos que
// correspondem a uma palavra inteira, e não só a um prefixo dela.
type No struct {
	Filhos       map[rune]*No
	FimDePalavra bool
}

func NovoNo() *No {
	return &No{Filhos: make(map[rune]*No)}
}

// Insere acrescenta uma palavra, criando um nó por letra ainda não vista.
func Insere(raiz *No, palavra string) {
	atual := raiz
	for _, letra := range palavra {
		proximo, existe := atual.Filhos[letra]
		if !existe {
			proximo = NovoNo()
			atual.Filhos[letra] = proximo
		}
		atual = proximo
	}
	atual.FimDePalavra = true
}

// desce segue o caminho das letras e devolve o nó em que ele termina, ou
// nil se o caminho não existe.
func desce(raiz *No, texto string) *No {
	atual := raiz
	for _, letra := range texto {
		proximo, existe := atual.Filhos[letra]
		if !existe {
			return nil
		}
		atual = proximo
	}
	return atual
}

// Contem responde se a palavra exata está na trie. O custo é o comprimento
// da palavra, e não depende de quantas palavras a trie guarda.
func Contem(raiz *No, palavra string) bool {
	n := desce(raiz, palavra)
	return n != nil && n.FimDePalavra
}

// EhPrefixo responde se alguma palavra da trie começa com o texto dado.
func EhPrefixo(raiz *No, texto string) bool {
	return desce(raiz, texto) != nil
}

// ComPrefixo devolve, em ordem alfabética, todas as palavras que começam
// com o prefixo: desce até o nó do prefixo e coleta a subárvore inteira.
func ComPrefixo(raiz *No, prefixo string) []string {
	n := desce(raiz, prefixo)
	if n == nil {
		return nil
	}
	var encontradas []string
	coleta(n, prefixo, &encontradas)
	sort.Strings(encontradas)
	return encontradas
}

func coleta(n *No, caminho string, encontradas *[]string) {
	if n.FimDePalavra {
		*encontradas = append(*encontradas, caminho)
	}
	for letra, filho := range n.Filhos {
		coleta(filho, caminho+string(letra), encontradas)
	}
}

// ContaNos devolve quantos nós a trie tem, sem contar a raiz. É a medida do
// espaço ocupado, e mostra quanto os prefixos comuns economizam.
func ContaNos(n *No) int {
	total := 0
	for _, filho := range n.Filhos {
		total += 1 + ContaNos(filho)
	}
	return total
}

func main() {
	palavras := []string{
		"casa", "casaco", "casal", "caso", "cassino",
		"carro", "carta", "cartaz", "cartão",
		"dado", "dados", "data", "datar",
		"grafo", "grafos", "grau", "grade",
	}

	raiz := NovoNo()
	for _, p := range palavras {
		Insere(raiz, p)
	}

	letras := 0
	for _, p := range palavras {
		letras += len([]rune(p))
	}
	fmt.Printf("%d palavras, %d letras no total, %d nós na trie\n",
		len(palavras), letras, ContaNos(raiz))

	fmt.Println()
	for _, p := range []string{"casa", "cas", "cartão", "carroça"} {
		fmt.Printf("Contem(%q) = %v, EhPrefixo(%q) = %v\n",
			p, Contem(raiz, p), p, EhPrefixo(raiz, p))
	}

	fmt.Println()
	for _, prefixo := range []string{"car", "cas", "gra", "z"} {
		encontradas := ComPrefixo(raiz, prefixo)
		if len(encontradas) == 0 {
			fmt.Printf("ComPrefixo(%q) = nenhuma palavra\n", prefixo)
			continue
		}
		fmt.Printf("ComPrefixo(%q) = %s\n", prefixo, strings.Join(encontradas, ", "))
	}
}
