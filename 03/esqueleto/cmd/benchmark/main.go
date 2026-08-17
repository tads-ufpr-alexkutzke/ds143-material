// Programa de medição empírica das implementações de Union-Find.
//
// A medição é feita sobre pares gerados em memória, com semente fixa, para que
// todas as implementações recebam exatamente a mesma sequência e para que o
// tempo medido não inclua leitura de arquivo.
//
// Uso:
//
//	go build -o benchmark ./cmd/benchmark
//	./benchmark
package main

import (
	"math/rand"

	"ds143/atividade1/unionfind"
)

// tamanhos são os valores de n do teste de duplicação. Ajuste a lista se a sua
// máquina levar tempo demais nos dois últimos: o que importa é ter ao menos
// quatro tamanhos consecutivos, cada um o dobro do anterior, com tempos acima
// de um décimo de segundo.
var tamanhos = []int{10000, 20000, 40000, 80000, 160000}

// implementacoes define a ordem das colunas da tabela.
var implementacoes = []struct {
	nome string
	novo func(int) unionfind.UF
}{
	{"Quick-Find", func(n int) unionfind.UF { return unionfind.NewQuickFind(n) }},
	{"Quick-Union", func(n int) unionfind.UF { return unionfind.NewQuickUnion(n) }},
	{"Weighted", func(n int) unionfind.UF { return unionfind.NewWeightedQuickUnion(n) }},
	{"Weighted+PC", func(n int) unionfind.UF { return unionfind.NewWeightedPathCompression(n) }},
}

// gerarPares produz quantidade pares aleatórios de objetos entre 0 e n-1.
// A semente é fixa, então duas execuções com os mesmos argumentos geram a
// mesma sequência. Esta função já está pronta.
func gerarPares(n, quantidade int, semente int64) [][2]int {
	r := rand.New(rand.NewSource(semente))
	pares := make([][2]int, quantidade)
	for i := range pares {
		pares[i] = [2]int{r.Intn(n), r.Intn(n)}
	}
	return pares
}

func main() {
	// TODO: para cada n em tamanhos:
	//   1. gere 2*n pares com gerarPares(n, 2*n, 42);
	//   2. para cada implementação, crie a estrutura com n objetos, marque o
	//      instante inicial com time.Now(), processe todos os pares (chamando
	//      Connected antes de Union, como fazem os programas da Aula 02) e
	//      obtenha o tempo decorrido com time.Since;
	//   3. imprima uma linha da tabela com o tempo de cada implementação e a
	//      razão entre o tempo atual e o tempo do tamanho anterior.
	//
	// Cuidados registrados na aula: meça o binário compilado, não `go run`, e
	// desconsidere tempos abaixo de um décimo de segundo, que são dominados
	// por ruído do sistema operacional.
	_ = tamanhos
	_ = implementacoes
	panic("TODO: implementar a medição")
}
