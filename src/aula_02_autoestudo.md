# Aula 02 — Autoestudo: Recursão, Busca e Ordenação

Conteúdo de sala invertida da Aula 02: revisão de recursão, busca e ordenação.
Não é apresentado em sala — a aula presencial cobre [Análise de Algoritmos e
Union-Find](aula_02.md) — mas cai na Prova 1, junto com o restante da Aula 02.
Estude ao longo da semana; a aula prática seguinte usa este conteúdo em
exercícios de fixação em Go.

Todo o código está em Go. Cada arquivo é um programa independente
(`package main`); execute com `go run <arquivo>.go`.

## Recursão

- [Slides](02/recursao/recursion.pdf)

Exemplos em Go (vários com uma versão "print", que imprime a árvore de
chamadas recursivas indentada por profundidade — útil para visualizar como a
recursão se desenrola):

- MDC pelo Algoritmo de Euclides: [euclid.go](02/codes/recursao/euclid.go) / [euclid_print.go](02/codes/recursao/euclid_print.go)
- Fatorial (recursivo e iterativo): [fatorial.go](02/codes/recursao/fatorial.go)
- Fibonacci:
  - [fibonacci_recursivo.go](02/codes/recursao/fibonacci_recursivo.go) — versão ingênua, conta o número de chamadas
  - [fibonacci_recursivo_print.go](02/codes/recursao/fibonacci_recursivo_print.go) — mostra a árvore de chamadas repetidas
  - [fibonacci_iterativo.go](02/codes/recursao/fibonacci_iterativo.go) — $O(n)$, sem pilha de recursão
  - [fibonacci_dinamico.go](02/codes/recursao/fibonacci_dinamico.go) — recursivo com memoização, também $O(n)$
- Lista encadeada com `Tamanho` e `Maximo` recursivos: [lista.go](02/codes/recursao/lista.go)
- Parser de expressões em notação prefixa (ex.: `+ 3 4` = 7): [prefix_parser.go](02/codes/recursao/prefix_parser.go) / [prefix_parser_print.go](02/codes/recursao/prefix_parser_print.go)
- "Puzzle" da Conjectura de Collatz: [puzzle.go](02/codes/recursao/puzzle.go) / [puzzle_print.go](02/codes/recursao/puzzle_print.go)

## Busca

- [Slides](02/busca/busca.pdf)
- [Análise da busca binária](02/busca/pdf/analise-busca-binaria.pdf)
- [Busca sequencial, binária iterativa e binária recursiva em Go](02/codes/busca/busca.go) — as três implementações no mesmo programa, comparando o número de comparações realizado por cada uma em um vetor de 1.000.000 de elementos

## Ordenação

### MergeSort

- [Slides](02/ordenacao/merge/mergesort.pdf)
- [Merge de dois vetores já ordenados](02/codes/ordenacao/merge.go)
- [Merge "in place" (truque clássico de Sedgewick)](02/codes/ordenacao/merge_in_place.go)

### QuickSort

- [Slides](02/ordenacao/quicksort/quicksort.pdf)
- [QuickSort em Go (particionamento de Hoare/Sedgewick)](02/codes/ordenacao/quicksort.go)

### Bônus: comparação com ordenações $O(n^2)$

- [Bubble sort, Insertion sort e Selection sort](02/codes/ordenacao/sort_basico.go) — `go run sort_basico.go <selection|insertion|bubble> <n> [print]`
