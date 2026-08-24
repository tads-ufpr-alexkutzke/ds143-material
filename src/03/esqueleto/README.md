# Esqueleto da Atividade Avaliativa 1 (DS143)

Ponto de partida da atividade descrita em [aula_03.md](../../aula_03.md).

## Estrutura

```
go.mod                                    módulo ds143/atividade1
unionfind/uf.go                           interface UF (pronta, não alterar)
unionfind/quickfind.go                    pronta, serve de modelo
unionfind/quickunion.go                   pronta
unionfind/weightedquickunion.go           pronta, base da que falta
unionfind/weightedpathcompression.go      A COMPLETAR
unionfind/uf_test.go                      testes fornecidos (não alterar)
unionfind/testdata/                       entradas tinyUF e mediumUF
cmd/benchmark/main.go                     completar a função medir
```

Escreva seu teste em `unionfind/meus_testes_test.go`, um arquivo novo. O
`uf_test.go` fornecido é o mesmo usado na correção.

## Comandos

```bash
go vet ./...                  # erros que o compilador não pega
go test ./...                 # roda os testes
go test -v ./unionfind        # mostra cada subteste, um por implementação
go test -run TestCadeia ./...  # roda só um teste

go build -o benchmark ./cmd/benchmark
./benchmark
```

Enquanto os métodos de `WeightedPathCompression` estiverem com
`panic("TODO: ...")`, o pacote compila e os testes falham. É o estado esperado
no início.
