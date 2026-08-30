# Laboratório da Aula 04 (DS143)

Projeto de apoio à aula prática descrita em [aula_04.md](../../aula_04.md).

## Estrutura

```
go.mod                              módulo ds143/aula04
algoritmos/busca.go                 busca sequencial e busca binária
algoritmos/ordenacao.go             selection, insertion, merge e quicksort
algoritmos/algoritmos_test.go       testes fornecidos (não alterar)
cmd/benchmark/main.go               medição das ordenações (pronto)
cmd/consultas/main.go               busca sequencial contra ordenar e buscar (pronto)
```

Quatro funções de `algoritmos/` estão incorretas e são o trabalho da Parte 1:
`BuscaBinaria`, `InsertionSort`, `Merge` e `Particiona`. As demais estão
corretas e servem de referência de estilo. Os dois programas de `cmd/`
dependem dessas correções e acusam o problema quando encontram resultado
errado.

## Comandos

```bash
go vet ./...                        # erros que o compilador não pega
go test ./algoritmos                # roda todos os testes
go test -v -run TestMerge ./algoritmos   # roda um teste só, com detalhe

go build -o benchmark ./cmd/benchmark
./benchmark                         # entrada aleatória
./benchmark -ordenada               # entrada já ordenada

go build -o consultas ./cmd/consultas
./consultas
```

O `benchmark` leva entre 7 e 10 segundos, quase todos gastos pelo
`SelectionSort` e pelo `InsertionSort` no maior tamanho. Com `-ordenada` cai
para 3 segundos, porque a Tabela 2 não roda. Tempos abaixo de
0,0001 s aparecem como `0.0000`.
