# Aula 02 — Revisão: Análise de Algoritmos, Recursão, Busca e Ordenação

## Roteiro da aula

1. Análise de Algoritmos (notação Big-O)
2. Union-Find: análise comparativa das implementações (revisita o desafio da Aula 1)
3. Recursão
4. Busca
5. Ordenação (MergeSort e QuickSort)

Esta aula é uma **revisão** de conteúdos normalmente já vistos em disciplinas anteriores (Estruturas de Dados I / Algoritmos e Programação). O objetivo aqui é formalizar, com a notação Big-O, o que torna uma solução mais eficiente que outra — usando como primeiro exemplo prático o próprio desafio de Union-Find proposto na Aula 1.

Todo o código desta aula está em Go. Cada arquivo é um programa independente (`package main`); execute com `go run <arquivo>.go`.

---

## 1. Análise de Algoritmos (Big-O)

- [Slides](02/analise_alg/analise_alg.pdf)

## 2. Union-Find — análise comparativa

Retomando o desafio da Aula 1: existem várias formas de implementar `Find` e `Union`. Compare a complexidade de cada uma:

- [Quick-Find](02/codes/unionfind/quick_find.go) — `Find`/`Connected` em O(1), mas `Union` em O(n)
- [Quick-Union](02/codes/unionfind/quick_union.go) — `Union` mais barato, mas árvores podem degenerar (O(n) no pior caso)
- [Weighted Quick-Union](02/codes/unionfind/weighted_quick_union.go) — balanceamento por tamanho garante O(log n)

Referências:

- [Sedgewick - Union-find (em inglês)](https://algs4.cs.princeton.edu/15uf/)
- [Sedgewick - Union-find Slides (em inglês)](https://algs4.cs.princeton.edu/lectures/keynote/15UnionFind.pdf)

## 3. Recursão

- [Slides](02/recursao/recursion.pdf)

Exemplos em Go (vários com uma versão "print", que imprime a árvore de chamadas recursivas indentada por profundidade — útil para visualizar como a recursão se desenrola):

- MDC pelo Algoritmo de Euclides: [euclid.go](02/codes/recursao/euclid.go) / [euclid_print.go](02/codes/recursao/euclid_print.go)
- Fatorial (recursivo e iterativo): [fatorial.go](02/codes/recursao/fatorial.go)
- Fibonacci:
  - [fibonacci_recursivo.go](02/codes/recursao/fibonacci_recursivo.go) — versão ingênua, conta o número de chamadas
  - [fibonacci_recursivo_print.go](02/codes/recursao/fibonacci_recursivo_print.go) — mostra a árvore de chamadas repetidas
  - [fibonacci_iterativo.go](02/codes/recursao/fibonacci_iterativo.go) — O(n), sem pilha de recursão
  - [fibonacci_dinamico.go](02/codes/recursao/fibonacci_dinamico.go) — recursivo com memoização, também O(n)
- Lista encadeada com `Tamanho` e `Maximo` recursivos: [lista.go](02/codes/recursao/lista.go)
- Parser de expressões em notação prefixa (ex.: `+ 3 4` = 7): [prefix_parser.go](02/codes/recursao/prefix_parser.go) / [prefix_parser_print.go](02/codes/recursao/prefix_parser_print.go)
- "Puzzle" da Conjectura de Collatz: [puzzle.go](02/codes/recursao/puzzle.go) / [puzzle_print.go](02/codes/recursao/puzzle_print.go)

## 4. Busca

- [Slides](02/busca/busca.pdf)
- [Análise da busca binária](02/busca/pdf/analise-busca-binaria.pdf)
- [Busca sequencial, binária iterativa e binária recursiva em Go](02/codes/busca/busca.go) — as três implementações no mesmo programa, comparando o número de comparações realizado por cada uma em um vetor de 1.000.000 de elementos

## 5. Ordenação

### MergeSort

- [Slides](02/ordenacao/merge/mergesort.pdf)
- [Merge de dois vetores já ordenados](02/codes/ordenacao/merge.go)
- [Merge "in place" (truque clássico de Sedgewick)](02/codes/ordenacao/merge_in_place.go)

### QuickSort

- [Slides](02/ordenacao/quicksort/quicksort.pdf)
- [QuickSort em Go (particionamento de Hoare/Sedgewick)](02/codes/ordenacao/quicksort.go)

### Bônus: comparação com ordenações O(n²)

- [Bubble sort, Insertion sort e Selection sort](02/codes/ordenacao/sort_basico.go) — `go run sort_basico.go <selection|insertion|bubble> <n> [print]`
