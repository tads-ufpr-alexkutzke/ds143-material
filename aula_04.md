# Aula 04: Aula prática de fundamentos

**Presencial, 26/08/2026.** Traga o notebook com Go instalado, ou use uma das
máquinas do laboratório.

| | |
|---|---|
| **Peso** | sem nota |
| **Formato** | individual, com discussão em sala |
| **Entrega** | `respostas.md` na UFPR Virtual, até o final da aula |

O conteúdo é o do [autoestudo de recursão, busca e
ordenação](aula_02_autoestudo.md) e o teste de duplicação da [aula de Análise
de Algoritmos](aula_02.md), seção 9. Tudo isso cai na Prova 1.

## Objetivos

Ao final da aula você deve ser capaz de:

1. Localizar e corrigir defeitos em implementações de busca binária, merge e
   particionamento, a partir da mensagem de um teste que falha;
2. Classificar um algoritmo pela ordem de crescimento usando apenas as razões
   entre tempos medidos;
3. Explicar por que o mesmo algoritmo muda de custo conforme a entrada, e
   decidir, com número na mão, quando compensa ordenar antes de buscar.

## Antes de começar

Baixe o projeto da pasta [04/lab](04/lab), também disponível como `.zip` na
UFPR Virtual, e confirme que ele compila:

```bash
cd lab
go vet ./...     # sem erros
go test ./algoritmos    # FAIL, em quatro funções
```

O `algoritmos/algoritmos_test.go` é o mesmo arquivo usado na devolutiva. Não o
altere.

## Parte 1: corrigir as quatro funções (35 min)

Quatro funções de `algoritmos/` estão incorretas:

- `BuscaBinaria`, em [busca.go](04/lab/algoritmos/busca.go);
- `InsertionSort`, `Merge` e `Particiona`, em
  [ordenacao.go](04/lab/algoritmos/ordenacao.go).

`BuscaSequencial`, `SelectionSort`, `MergeSort` e `QuickSort` estão corretas.
As duas últimas chamam `Merge` e `Particiona`, então erram enquanto essas duas
estiverem erradas.

Cada correção tem entre uma e quatro linhas. Comece rodando `go test -v
./algoritmos` e leia a mensagem: ela diz qual entrada quebrou e o que saiu no
lugar do esperado. O `TestParticiona` é o mais informativo, porque verifica a
pós-condição do particionamento em vez do resultado final da ordenação.

Ao terminar, `go test ./algoritmos` passa e `gofmt -l .` não lista nada.
Anote em `respostas.md` qual era o defeito de cada função, em uma linha.

## Parte 2: medir (40 min)

Com as funções corretas, compile e rode o programa de medição:

```bash
go build -o benchmark ./cmd/benchmark
./benchmark
```

Cada linha é um algoritmo, cada par de colunas é um tamanho `n` com o tempo e a
razão em relação ao tamanho anterior. `ERRO` na linha indica que aquele
algoritmo devolveu vetor não ordenado, ou seja, que a Parte 1 ainda não está
completa.

Rode depois sobre um vetor que já chega ordenado:

```bash
./benchmark -ordenada
```

Cole as duas saídas em `respostas.md` e responda às perguntas 1 e 2.

## Parte 3: buscar muitas vezes (20 min)

Um vetor de 200.000 elementos e `k` consultas a responder. Duas estratégias:
buscar sequencialmente `k` vezes no vetor como ele está, ou ordenar uma vez e
fazer `k` buscas binárias. O programa mede as duas:

```bash
go build -o consultas ./cmd/consultas
./consultas
```

A coluna `B: ord+bin` inclui o tempo de ordenar. Ache na tabela a faixa de `k`
em que a segunda estratégia passa a ganhar e responda à pergunta 3.

## Entrega

`respostas.md` preenchido, enviado na UFPR Virtual até o final da aula. Não
vale nota e não precisa do código; a devolutiva é coletiva, no início da Aula
05.

## Problemas comuns

**`go: cannot find main module`**: rode os comandos de dentro da pasta `lab`,
onde está o `go.mod`.

**O `benchmark` demora demais**: o `SelectionSort` e o `InsertionSort` são
O(n²) e o último tamanho da Tabela 1 é 40.000. Cerca de 10 segundos no total é
o esperado.

**Razões estranhas nos tempos pequenos**: tempos abaixo de 0,0001 s aparecem
como `0.0000` e suas razões não significam nada. Compare as razões nas linhas
em que o tempo é grande o bastante para ser medido.
