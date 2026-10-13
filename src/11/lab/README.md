# Esqueleto da Atividade Avaliativa 3 (DS143)

Ponto de partida da atividade descrita em [aula_11.md](../../aula_11.md).

## Estrutura

```
go.mod                    módulo ds143/atividade3
tabela/tabela.go          Tabela, hash, Busca, distancia e proximoPrimo prontos
tabela/exercicios.go      A COMPLETAR: Insere, redimensiona e Remove
tabela/tabela_test.go     testes fornecidos (não alterar)
tabela/bench_test.go      medidas de desempenho contra o map (não alterar)
respostas.md              modelo a preencher
```

## Comandos

```bash
go vet ./...                              # erros que o compilador não pega
go test ./...                             # roda todos os testes
go test -v -run 'Busca|Insere' ./tabela   # só os testes da Parte 1
go test -v -run Remove ./tabela           # só os testes da remoção
go test -run '^$' -bench . ./tabela       # medidas da Parte 3, sem os testes
gofmt -l .                                # arquivos fora do padrão de formatação
```

São três funções a escrever, todas em `tabela/exercicios.go`. Enquanto elas
estiverem com `panic("TODO: ...")`, o pacote compila e os testes falham. É o
estado esperado no início.
