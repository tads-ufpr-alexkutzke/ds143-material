// Funções de dispersão: como a escolha da função e do tamanho da tabela
// decide se as chaves se espalham pelas posições ou se amontoam em poucas.
//
// Execute com:  go run funcoes.go

package main

import (
	"fmt"
	"hash/fnv"
	"math/rand"
)

// modular é a função de dispersão mais simples para chaves inteiras: o resto
// da divisão pelo tamanho da tabela.
func modular(chave, m int) int {
	return chave % m
}

// somaDasLetras soma os códigos das letras. Duas palavras com as mesmas letras
// em outra ordem dão sempre a mesma soma.
func somaDasLetras(s string, m int) int {
	soma := 0
	for i := 0; i < len(s); i++ {
		soma += int(s[i])
	}
	return soma % m
}

// polinomial trata a palavra como um número escrito na base 31, calculado pelo
// método de Horner. A posição de cada letra muda o resultado. O resto é tirado
// a cada passo para o valor não estourar o int.
func polinomial(s string, m int) int {
	h := 0
	for i := 0; i < len(s); i++ {
		h = (31*h + int(s[i])) % m
	}
	return h
}

// fnv1a usa a função FNV-1a da biblioteca padrão (pacote hash/fnv), que
// mistura os bits de cada byte com uma multiplicação e um ou-exclusivo.
func fnv1a(s string, m int) int {
	h := fnv.New64a()
	h.Write([]byte(s))
	return int(h.Sum64() % uint64(m))
}

// sorteio ignora a chave e sorteia a posição. Não serve como função de
// dispersão, porque a mesma chave precisa cair sempre na mesma posição, mas
// mostra como fica a ocupação quando as posições são independentes.
var gerador = rand.New(rand.NewSource(1))

func sorteio(s string, m int) int {
	return gerador.Intn(m)
}

// ocupacao conta quantas posições da tabela receberam ao menos uma chave e
// quantas chaves caíram na posição mais cheia.
func ocupacao(posicoes []int, m int) (usadas, maior int) {
	contagem := make([]int, m)
	for _, p := range posicoes {
		contagem[p]++
	}
	for _, c := range contagem {
		if c > 0 {
			usadas++
		}
		if c > maior {
			maior = c
		}
	}
	return usadas, maior
}

func main() {
	// 1. Chaves inteiras com um padrão: todas múltiplas de 20.
	var multiplos []int
	for k := 20; k <= 20000; k += 20 {
		multiplos = append(multiplos, k)
	}
	fmt.Printf("1. %d chaves múltiplas de 20, com h(k) = k %% m\n\n", len(multiplos))
	fmt.Println("    m | posições usadas | chaves na posição mais cheia")
	fmt.Println("------+-----------------+-----------------------------")
	for _, m := range []int{100, 97} {
		pos := make([]int, len(multiplos))
		for i, k := range multiplos {
			pos[i] = modular(k, m)
		}
		usadas, maior := ocupacao(pos, m)
		fmt.Printf("%5d | %15d | %28d\n", m, usadas, maior)
	}

	// 2. Anagramas.
	fmt.Println()
	fmt.Println("2. Anagramas, com m = 97")
	fmt.Println()
	fmt.Println("palavra | soma das letras | polinomial")
	fmt.Println("--------+-----------------+-----------")
	for _, p := range []string{"amor", "roma", "mora", "ramo", "omar"} {
		fmt.Printf("%-7s | %15d | %10d\n", p, somaDasLetras(p, 97), polinomial(p, 97))
	}

	// 3. Códigos de matrícula: chaves de texto muito parecidas entre si.
	var grrs []string
	for i := 1; i <= 2000; i++ {
		grrs = append(grrs, fmt.Sprintf("GRR2026%04d", i))
	}
	fmt.Println()
	fmt.Printf("3. %d matrículas, de %s a %s, com m = 997\n\n",
		len(grrs), grrs[0], grrs[len(grrs)-1])
	fmt.Println("função          | posições usadas | chaves na posição mais cheia")
	fmt.Println("----------------+-----------------+-----------------------------")
	for _, f := range []struct {
		nome string
		h    func(string, int) int
	}{
		{"soma das letras", somaDasLetras},
		{"polinomial", polinomial},
		{"FNV-1a", fnv1a},
		{"sorteio", sorteio},
	} {
		pos := make([]int, len(grrs))
		for i, g := range grrs {
			pos[i] = f.h(g, 997)
		}
		usadas, maior := ocupacao(pos, 997)
		fmt.Printf("%-15s | %15d | %28d\n", f.nome, usadas, maior)
	}
}
