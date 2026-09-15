# Atividade Avaliativa 2: árvores de busca e AVL em Go

**Aula prática presencial, 16/09/2026.** Traga o notebook com Go instalado, ou
use uma das máquinas do laboratório.

| | |
|---|---|
| **Peso** | 5% da nota final |
| **Formato** | individual ou em dupla |
| **Prazo** | 16/09/2026, 23h59, na UFPR Virtual |

Feita em sala, com consulta livre ao material da disciplina. O conteúdo é o das
aulas [05](aula_05.md) e [06](aula_06.md), e cai na Prova 1, em 23/09.

São três funções a escrever, uma em cada uma das duas partes abaixo.

## Dupla

A atividade pode ser feita em dupla. Nesse caso, **as duas pessoas enviam o
mesmo `.zip` na UFPR Virtual**, cada uma no seu próprio envio, e o
`respostas.md` traz os dois nomes e os dois GRR. Quem não enviar fica sem nota,
mesmo que o nome esteja no arquivo enviado pela outra pessoa.

## Objetivos

Ao final desta atividade você deve ser capaz de:

1. Escrever a verificação da propriedade de busca com limites propagados, e
   dizer por que a comparação entre pai e filhos não basta;
2. Localizar a menor chave maior que um valor dado, descendo por um único
   caminho;
3. Escrever o verificador do invariante da AVL, separando a conferência do campo
   `Alt` da conferência do fator de balanceamento.

## Antes de começar

Confirme o Go instalado com `go version` (a partir da 1.21). Baixe o esqueleto
do projeto na pasta
[07/lab](https://github.com/tads-ufpr-alexkutzke/ds143-material/tree/main/src/07/lab),
também disponível como `.zip` na UFPR Virtual.

O esqueleto compila desde o início e os testes falham. É o estado esperado:

```bash
cd lab
go vet ./...     # sem erros
go test ./...    # FAIL nos dois pacotes, com panic nas funções por escrever
```

São dois pacotes. O `abb` tem a árvore binária de busca da Aula 06 e o `avl` tem
a árvore balanceada, as duas prontas. Leia o código de cada uma antes de escrever
qualquer coisa: ele define como o restante deve ser escrito. Não altere
`abb/abb_test.go` nem `avl/avl_test.go`, que são os arquivos usados na correção.

## Parte 1: duas funções sobre a ABB (35 min)

Complete as duas funções de
[abb/exercicios.go](07/lab/abb/exercicios.go). São os exercícios 5 e 6 da Aula
06.

**`EhABB(a *No) bool`** verifica se uma árvore binária qualquer satisfaz a
propriedade de busca. O teste fornecido inclui a árvore da seção 3 da Aula 06,
que tem o 60 na subárvore esquerda do 50: o 60 é maior que o pai 30, então a
comparação entre pai e filhos aceita a árvore, e a sua função precisa rejeitá-la.

**`Sucessor(a *No, v int) (int, bool)`** devolve a menor chave estritamente
maior que `v`, e um booleano dizendo se ela existe. O valor `v` não precisa
estar na árvore. Desça por um único caminho, sem percorrer a árvore inteira.

Ao terminar, `go test ./abb` passa.

## Parte 2: o verificador da AVL (30 min)

Complete a **`VerificaAVL(a *No) (bool, *No)`** em
[avl/verifica.go](07/lab/avl/verifica.go). Ela percorre a árvore e confere, em
cada nó, duas condições independentes:

1. o campo `Alt` guarda a altura real da subárvore que começa no nó;
2. o fator de balanceamento do nó está em $\{-1, 0, +1\}$.

Devolve `true` e `nil` quando a árvore inteira passa, ou `false` e o nó em que a
verificação falhou. A `AlturaRecalculada`, que mede a altura percorrendo a
árvore sem olhar para o campo `Alt`, já vem pronta no arquivo.

As duas condições são independentes, e a segunda pergunta do `respostas.md`
cobra justamente isso: uma árvore pode ter o campo `Alt` correto em todos os nós
e ainda assim não ser AVL.

O pacote `avl` está completo, com as rotações, a inserção e a remoção da Aula
06. A `VerificaAVL` é o que permite conferir que essas operações mantêm o
invariante, em vez de supor que mantêm. Os testes fornecidos a rodam sobre
árvores montadas à mão, com um único defeito cada.

Ao terminar, `go test ./...` passa inteiro.

## Entrega

Um `.zip` na UFPR Virtual, nomeado `ds143_atividade2_<seu_nome>.zip`, com o
projeto Go completo (`go.mod`, os pacotes `abb` e `avl`) e o `respostas.md`
preenchido. Sem binários compilados. Rode `go test ./...` uma última vez na
pasta que será enviada e cole a saída no `respostas.md`.

Em dupla, as duas pessoas enviam o mesmo arquivo, cada uma no seu envio.

## Critérios de avaliação

| Critério | Valor |
|---|---|
| Compila, `go vet` limpo, `gofmt` sem pendências, testes fornecidos passam | 20 |
| `EhABB` correta, rejeitando a violação distante | 25 |
| `Sucessor` correto, inclusive com `v` fora da árvore e sem sucessor | 25 |
| `VerificaAVL` confere as duas condições e aponta o nó | 25 |
| `respostas.md` com as três perguntas e a declaração de uso de IA | 5 |
| **Total** | **100** |

Entrega que não compila perde o primeiro item inteiro; os demais são avaliados
pelo que for verificável no código.

## Problemas comuns

**`panic: TODO: implementar ...`**: a função ainda está por escrever. Os `panic`
do esqueleto marcam o que falta, e desaparecem conforme você preenche.

**`go: cannot find main module`**: rode os comandos de dentro da pasta `lab`,
onde está o `go.mod`.

**A `EhABB` aceita a árvore do 60**: você está comparando cada nó apenas com os
seus dois filhos. A verificação precisa carregar, na descida, o intervalo de
valores que aquela subárvore inteira pode conter.

**A `VerificaAVL` acusa um nó com o campo `Alt` correto**: leia o fator dele. As
duas condições são independentes, e é disso que trata a pergunta 3.
