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
3. Escrever a rotação à esquerda a partir da rotação à direita, cuidando da
   ordem das atribuições e das atualizações de altura;
4. Escrever um verificador de invariante e usá-lo para localizar um defeito que
   os testes fornecidos não pegam;
5. Explicar por que uma sequência de inserção expõe um defeito de
   rebalanceamento e outra não.

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

São dois pacotes. O `abb` tem a árvore binária de busca da Aula 06, pronta, e o
`avl` tem a árvore balanceada, com duas lacunas. Leia o código pronto de cada
um antes de escrever qualquer coisa: ele define como o restante deve ser
escrito. Não altere `abb/abb_test.go` nem `avl/avl_test.go`, que são os
arquivos usados na correção.

## Parte 1: duas funções sobre a ABB (30 min)

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

## Parte 2: a AVL e um defeito (35 min)

**Escreva a `RotacaoEsquerda`** em [avl/avl.go](07/lab/avl/avl.go), espelhando a
rotação à direita que está no mesmo arquivo. O teste fornecido confere a forma
da árvore resultante e também o campo `Alt` dos dois nós envolvidos.

**Escreva a `VerificaAVL`** em [avl/verifica.go](07/lab/avl/verifica.go). Ela
percorre a árvore e confere, em cada nó, que o campo `Alt` guarda a altura real
da subárvore e que o fator de balanceamento está em $\{-1, 0, +1\}$. Devolve
`true` e `nil` quando tudo passa, ou `false` e o nó em que a verificação falhou.
A `AlturaRecalculada`, que mede a altura sem olhar para o campo `Alt`, já vem
pronta no arquivo.

**Encontre e corrija o defeito.** A `RotacaoDireita` do esqueleto tem um defeito
de uma linha. As três atribuições de ponteiro estão corretas, e a árvore
resultante tem a forma certa; o que sai errado é o campo `Alt`, e o erro se
propaga para o fator de balanceamento dos ancestrais, provocando rotação a mais
ou a menos nas inserções seguintes.

Os testes fornecidos passam mesmo com o defeito no lugar, porque nenhum deles
confere o campo `Alt` das árvores construídas por `Insere`. Quem localiza o
defeito é a sua `VerificaAVL`, rodada pelo programa da Parte 3. Um `PASS` do
`go test` não é prova de que a estrutura está consistente, e esse é o ponto da
Parte 2.

## Parte 3: medir e explicar (15 min)

O programa [cmd/alturas/main.go](07/lab/cmd/alturas/main.go) já está pronto. Ele
constrói a ABB e a AVL sobre as mesmas sequências de inserção, em três ordens
(crescente, decrescente e aleatória), e imprime duas colunas para a AVL: a
altura guardada no campo `Alt` da raiz e a altura obtida percorrendo a árvore.
Em uma AVL consistente as duas são iguais. No fim, ele roda a sua `VerificaAVL`
sobre as três árvores de mil nós.

Rode **antes** de corrigir o defeito da Parte 2 e guarde a saída:

```bash
go build -o alturas ./cmd/alturas
./alturas | tee antes.txt
```

Corrija o defeito e rode de novo:

```bash
go build -o alturas ./cmd/alturas
./alturas | tee depois.txt
```

Preencha o `respostas.md` com as duas saídas e as perguntas 3 a 6.

## Entrega

Um `.zip` na UFPR Virtual, nomeado `ds143_atividade2_<seu_nome>.zip`, com o
projeto Go completo (`go.mod`, os pacotes `abb` e `avl`, `cmd/alturas`), o
`antes.txt`, o `depois.txt` e o `respostas.md` preenchido. Sem binários
compilados: apague o `alturas` antes de compactar. Rode `go test ./...` uma
última vez na pasta que será enviada.

Em dupla, as duas pessoas enviam o mesmo arquivo, cada uma no seu envio.

## Critérios de avaliação

| Critério | Valor |
|---|---|
| Compila, `go vet` limpo, `gofmt` sem pendências, testes fornecidos passam | 20 |
| `EhABB` correta, rejeitando a violação distante | 20 |
| `Sucessor` correto, inclusive com `v` fora da árvore e sem sucessor | 20 |
| `RotacaoEsquerda` correta, com as alturas atualizadas na ordem certa | 8 |
| `VerificaAVL` detecta as duas condições e aponta o nó | 12 |
| Defeito corrigido, com as perguntas 3 e 4 respondidas | 5 |
| Saídas antes e depois, com as perguntas 5 e 6 respondidas | 15 |
| **Total** | **100** |

Entrega que não compila perde o primeiro item inteiro; os demais são avaliados
pelo que for verificável no código.

## Problemas comuns

**`panic: TODO: implementar ...`**: a função ainda está por escrever. Os `panic`
do esqueleto marcam o que falta, e desaparecem conforme você preenche.

**`go: cannot find main module`**: rode os comandos de dentro da pasta `lab`,
onde está o `go.mod`.

**A `VerificaAVL` acusa um nó com o campo `Alt` correto**: leia o fator dele. As
duas condições são independentes, e um nó pode ter a altura certa e ainda assim
estar desbalanceado, por causa de um `Alt` errado que já foi corrigido mais
abaixo.

**`./alturas` não roda**: ele usa a `RotacaoEsquerda` e a `VerificaAVL`. Termine
a Parte 2 antes.
