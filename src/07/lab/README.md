# Esqueleto da Atividade Avaliativa 2 (DS143)

Ponto de partida da atividade descrita em [aula_07.md](../../aula_07.md).

## Estrutura

```
go.mod                   módulo ds143/atividade2
abb/abb.go               ABB pronta (busca, inserção, remoção, in-ordem, altura)
abb/exercicios.go        A COMPLETAR: EhABB e Sucessor
abb/abb_test.go          testes fornecidos (não alterar)
avl/avl.go               AVL: RotacaoEsquerda A COMPLETAR, RotacaoDireita com defeito
avl/verifica.go          AlturaRecalculada pronta, VerificaAVL A COMPLETAR
avl/avl_test.go          testes fornecidos (não alterar)
cmd/alturas/main.go      pronto: tabelas de altura da ABB e da AVL
respostas.md             modelo a preencher
```

## Comandos

```bash
go vet ./...            # erros que o compilador não pega
go test ./...           # roda os testes dos dois pacotes
go test -v ./abb        # mostra cada teste da Parte 1
go test -v ./avl        # mostra cada teste da Parte 2
gofmt -l .              # lista arquivos fora do padrão de formatação

go build -o alturas ./cmd/alturas
./alturas | tee antes.txt
```

Enquanto as funções estiverem com `panic("TODO: ...")`, o pacote compila e os
testes falham. É o estado esperado no início.

Os testes fornecidos **não** conferem o campo `Alt` das árvores construídas por
`Insere`. Passar em `go test ./...` não é prova de que a AVL está consistente:
essa conferência é a `VerificaAVL` que você vai escrever.
