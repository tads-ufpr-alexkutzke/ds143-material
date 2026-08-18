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
	"fmt"
	"math/rand"
	"time"

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
// mesma sequência. Função pronta.
func gerarPares(n, quantidade int, semente int64) [][2]int {
	r := rand.New(rand.NewSource(semente))
	pares := make([][2]int, quantidade)
	for i := range pares {
		pares[i] = [2]int{r.Intn(n), r.Intn(n)}
	}
	return pares
}

// medir cria uma estrutura com n objetos usando o construtor recebido,
// processa todos os pares e devolve o tempo gasto no processamento.
//
// TODO: implementar.
//
//  1. crie a estrutura com novo(n);
//  2. marque o instante inicial com time.Now(), depois do passo 1: o tempo de
//     criação não entra na medição;
//  3. para cada par, chame Connected e, se os objetos ainda não estiverem
//     conectados, chame Union. É o que fazem os programas da Aula 02;
//  4. devolva time.Since(inicio).
func medir(novo func(int) unionfind.UF, n int, pares [][2]int) time.Duration {
	panic("TODO: implementar medir")
}

// main percorre os tamanhos, chama medir para cada implementação e imprime a
// tabela com os tempos e as razões entre tamanhos consecutivos. Função pronta.
func main() {
	anterior := make([]time.Duration, len(implementacoes))

	fmt.Printf("%9s", "n")
	for _, impl := range implementacoes {
		fmt.Printf("%14s%8s", impl.nome, "razao")
	}
	fmt.Println()

	for _, n := range tamanhos {
		pares := gerarPares(n, 2*n, 42)
		fmt.Printf("%9d", n)
		for i, impl := range implementacoes {
			gasto := medir(impl.novo, n, pares)
			fmt.Printf("%14.3f", gasto.Seconds())
			if anterior[i] > 0 {
				fmt.Printf("%8.1f", gasto.Seconds()/anterior[i].Seconds())
			} else {
				fmt.Printf("%8s", "")
			}
			anterior[i] = gasto
		}
		fmt.Println()
	}
}
