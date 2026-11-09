# Atividade Avaliativa 4: menor caminho em grafos não direcionados 

**Aula prática presencial, 11/11/2026.** Traga o notebook com Go instalado, ou
use uma das máquinas do laboratório.

| | |
|---|---|
| **Peso** | 5% da nota final |
| **Formato** | individual ou em dupla |
| **Prazo** | 15/11/2026 (domingo), 23h59, na UFPR Virtual |

Feita em sala ou em casa, com consulta livre ao material da disciplina. O conteúdo é o da
[terceira aula de grafos](aula_14.md), e cai na Prova 2, em 25/11.

A atividade envolve a conclusão da implementação de apenas um programa, 
com a escrita das funções que forem julgadas necessárias e a resposta a uma questão discursiva. 

## Dupla

A atividade pode ser feita em dupla. Nesse caso, **as duas pessoas enviam o
mesmo `.zip` na UFPR Virtual**, cada uma no seu próprio envio, e o
`respostas.md` traz os dois nomes e os dois GRR. Quem não enviar fica sem nota,
mesmo que o nome esteja no arquivo enviado pela outra pessoa.

## Objetivos

Ao final desta atividade você deve ser capaz de:

1. Escrever uma solução para um problema específico de menor caminho em grafos não direcionados;
2. Explicar a complexidade computacional da sua solução.

## Antes de começar

Confirme o Go instalado com `go version` (a partir da 1.21). Baixe o esqueleto
do projeto na pasta
[14/lab](https://github.com/tads-ufpr-alexkutzke/ds143-material/tree/main/src/14/lab),
também disponível como `.zip` na UFPR Virtual.

O esqueleto compila desde o início e os testes falham. É o estado esperado:

```bash
cd lab
go vet ./...     # sem erros
go test ./...    # FAIL, com panic nas funções por escrever
```

## A implementação disponível

Não altere os arquivos de teste (`*_test.go`). Eles serão utilizados 
na correção.

## Parte 1: menor caminho em grafo "linha"


Dado um grafo "linha" (grafo não-direcionado conexo com todos os vértices de grau 2, exceto as duas pontas que possuem grau 1), escreva um algoritmo que pré-processa o grafo em tempo linear e é capaz de retornar em tempo constante a distância do menor caminho entre qualquer par de vértices.

Considere o grafo abaixo:

```
Grafo: (5) --- (2) --- (7) --- (4) --- (1) --- (6) --- (0) --- (3)
Pesos:      4       3       9       1       7       2       5
```

O programa disponível já implementa toda a lógica de leitura do grafo e saída do programa. Leia
os comentários iniciais no arquivo `main.go` para mais detalhes.

Um exemplo possível de execução com o grafo acima é o seguinte:

```
8
4 7 9
6 0 2
1 6 7
5 2 4
3 0 5
7 2 3
1 4 1
1 4
1
1 7
10
5 7
7
3 5
31
-1 -1
```

A primeira linha lê a quantidade de vértices (V). Na sequência são esperadas a definição de `V-1` arestas (`vertice vertice peso`). Por fim, o programa aguarda uma lista de pares de vértices e responde, para cada par, a menor distância entre eles. A execução termina com Ctrl+D ou com o par -1 -1.

O arquivo `exercicios.go` possui a estrutura de as funções que devem ser implementadas como resposta a essa atividade. A compreensão do restante do código é parte da atividade. 

Ao terminar, `go test ./...` passa inteiro.

## Parte 2: explicação da complexidade

Explique, no espaço indicado no arquivo `respostas.md`, por que sua implementação
respeita a complexidade computacional exigida pelo enunciado: preprocessamento do grafo
em tempo linear e apresentação da distância entre qualquer par de vértices em tempo constante.

## Entrega

Um `.zip` na UFPR Virtual, nomeado `ds143_atividade4_<seu_nome>.zip`, com o
projeto Go completo (`go.mod` e o pacote `grafo`) e o `respostas.md`
preenchido. Sem binários compilados. 

Em dupla, as duas pessoas enviam o mesmo arquivo, cada uma no seu envio.

## Critérios de avaliação

| Critério                                                                  | Valor   |
|---------------------------------------------------------------------------|---------|
| Compila, `go vet` limpo, `gofmt` sem pendências, testes fornecidos passam | 60      |
| `respostas.md` com a pergunta e a declaração de uso de IA                 | 40      |
| **Total**                                                                 | **100** |

Entrega que não compila perde o primeiro item inteiro; os demais são avaliados
pelo que for verificável no código.

## Problemas comuns

**`panic: TODO: implementar ...`**: a função ainda está por escrever. Os `panic`
do esqueleto marcam o que falta, e desaparecem conforme você preenche.

**`go: cannot find main module`**: rode os comandos de dentro da pasta `lab`,
onde está o `go.mod`.

## Referência

Exercício adaptado de [Sedgewick (2024)](https://algs4.cs.princeton.edu/44sp/).
