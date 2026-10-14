# Atividade Avaliativa 3: tabela de dispersão própria contra o `map`

**Aula prática presencial, 14/10/2026.** Traga o notebook com Go instalado, ou
use uma das máquinas do laboratório.

| | |
|---|---|
| **Peso** | 5% da nota final |
| **Formato** | individual ou em dupla |
| **Prazo** | 18/10/2026 (domingo), 23h59, na UFPR Virtual |

Feita em sala, com consulta livre ao material da disciplina. O conteúdo é o da
[aula de tabelas de dispersão](aula_10.md), e cai na Prova 2, em 25/11.

São três funções a escrever, nas Partes 1 e 2, e uma série de medidas a fazer e
comentar, na Parte 3. Os últimos 15 minutos do encontro são da apresentação do
Trabalho Prático; o que faltar da atividade pode ser terminado até domingo,
18/10, às 23h59.

## Dupla

A atividade pode ser feita em dupla. Nesse caso, **as duas pessoas enviam o
mesmo `.zip` na UFPR Virtual**, cada uma no seu próprio envio, e o
`respostas.md` traz os dois nomes e os dois GRR. Quem não enviar fica sem nota,
mesmo que o nome esteja no arquivo enviado pela outra pessoa.

## Objetivos

Ao final desta atividade você deve ser capaz de:

1. Escrever a inserção em uma tabela com sondagem linear, mantendo o fator de
   carga limitado pelo redimensionamento;
2. Escrever a remoção sem lápide, decidindo pela posição de origem de cada chave
   se ela desce para a posição esvaziada;
3. Medir o desempenho de uma estrutura própria contra o `map` de Go, e explicar
   as diferenças pela posição que as chaves ocupam na tabela.

## Antes de começar

Confirme o Go instalado com `go version` (a partir da 1.21). Baixe o esqueleto
do projeto na pasta
[11/lab](https://github.com/tads-ufpr-alexkutzke/ds143-material/tree/main/src/11/lab),
também disponível como `.zip` na UFPR Virtual.

O esqueleto compila desde o início e os testes falham. É o estado esperado:

```bash
cd lab
go vet ./...     # sem erros
go test ./...    # FAIL, com panic nas funções por escrever
```

## A tabela do esqueleto

O pacote `tabela` tem uma tabela de dispersão de chaves inteiras com
endereçamento aberto e sondagem linear, a mesma da seção 5 da aula de tabelas de
dispersão, agora guardando um valor junto com cada chave. Leia
[tabela/tabela.go](11/lab/tabela/tabela.go) antes de escrever qualquer coisa:

* a `Tabela` guarda três fatias de mesmo comprimento $m$: `chaves`, `valores` e
  `ocupada`, que diz se a posição tem um par;
* a `hash` é o método da divisão, $h(k) = k \bmod m$, enquanto a constante
  `misturar` for `false`. Ela só muda na Parte 3;
* a `Busca` já está pronta, e mostra como percorrer a sequência de sondagem:
  começa em $h(k)$, anda uma posição por vez, dá a volta no fim do vetor e para
  na primeira posição vazia;
* a `distancia` conta os passos da sondagem entre duas posições, já com a volta
  no fim do vetor, e é a peça da Parte 2;
* o `String` desenha a tabela no formato das tabelas da aula. Um
  `fmt.Println(tab)` no meio de um teste ajuda a depurar.

Não altere `tabela/tabela_test.go` nem `tabela/bench_test.go`, que são os
arquivos usados na correção.

## Parte 1: inserção e crescimento (25 min)

Complete a **`Insere(k, v int)`** e a **`redimensiona(m int)`** em
[tabela/exercicios.go](11/lab/tabela/exercicios.go).

A `Insere` procura a chave como a `Busca`. Se a chave já está na tabela, troca o
valor e termina. Se não está, guarda o par na posição vazia em que a procura
parou. Depois de guardar uma chave nova, confere o fator de carga: se
$n / m$ passou de 0,7, chama a `redimensiona` com `proximoPrimo(2*m)`. Em
inteiros, a condição é `10*n > 7*m`, e evita o arredondamento de um `float64`.

A `redimensiona` cria uma tabela com $m$ posições, reinsere nela todos os pares
e toma o lugar da antiga, como no `encadeamento.go` da aula. A posição de cada
chave muda, porque $h(k)$ depende de $m$.

O limite de 0,7 é o que faz a `Busca` terminar: com pelo menos 30% das posições
vazias, toda sequência de sondagem chega a uma posição vazia.

Os testes desta parte refazem o exemplo da aula, com $m = 11$ e as chaves 22,
33, 5, 16, 27, 38 e 44, e conferem a posição de cada uma. Para rodá-los antes de
escrever a `Remove`:

```bash
go test -v -run 'Busca|Insere' ./tabela
```

## Parte 2: remoção sem lápide (30 min)

Complete a **`Remove(k int) bool`** em
[tabela/exercicios.go](11/lab/tabela/exercicios.go). Ela devolve `true` se a
chave estava na tabela, e `false` se não estava.

A remoção é a alternativa à lápide descrita na seção 6 da aula, que só funciona
na sondagem linear. Ache a posição $i$ da chave e esvazie-a. A partir daí,
percorra as posições seguintes, até a primeira vazia. Cada chave encontrada no
caminho, numa posição $j$, desce para o buraco $i$ se puder; quando desce, a
posição $j$ fica vazia e passa a ser o buraco.

A chave da posição $j$ pode descer quando o buraco $i$ está no caminho que a
sondagem fez, ao inseri-la, entre a posição de origem $h(k)$ e $j$. Se o buraco
está antes da posição de origem, a chave nunca passou por ele, e descer a
deixaria fora do alcance da `Busca`. Escreva essa condição com a `distancia`,
comparando duas distâncias até $j$.

O teste `TestRemoveVoltaCircular` monta a tabela abaixo, com $m = 11$:

```
   0   1   2   3   4   5   6   7   8   9  10
  31   1  10   .   .   .   .   .   .   9  20
```

O 9 e o 20 estão nas posições de origem. O 31 ($h = 9$) deu a volta e parou na
0. O 1 ($h = 1$) está na posição de origem. O 10 ($h = 10$) passou pela 10, pela
0 e pela 1 até parar na 2. Depois de `Remove(20)`, a tabela fica assim:

```
   0   1   2   3   4   5   6   7   8   9  10
  10   1   .   .   .   .   .   .   .   9  31
```

O 1 fica onde está, e o 10, que vem depois dele, desce. A remoção não pode parar
na primeira chave que não desce: só uma posição vazia encerra o percurso. A
pergunta 1 do `respostas.md` trata deste exemplo.

Um segundo exemplo mostra por que a comparação de distâncias resolve a volta no
fim do vetor. Com $m = 11$, inserindo 9, 20, 21, 31 e 2:

```
   0   1   2   3   4   5   6   7   8   9  10
  21  31   2   .   .   .   .   .   .   9  20
```

O 21 ($h = 10$) achou a 10 ocupada e parou na 0. O 31 ($h = 9$) passou pela 9,
pela 10 e pela 0 até parar na 1. O 2 ($h = 2$) está na posição de origem.
`Remove(20)` esvazia a posição 10, e o percurso segue a partir da 0:

| $j$ | chave | $h$ | buraco $i$ | desce? |
|---|---|---|---|---|
| 0 | 21 | 10 | 10 | sim |
| 1 | 31 | 9 | 0 | sim |
| 2 | 2 | 2 | 1 | não |
| 3 | vazia | | | fim |

A tabela fica assim, e a `Busca` acha todas as chaves:

```
   0   1   2   3   4   5   6   7   8   9  10
  31   .   2   .   .   .   .   .   .   9  21
```

O 31 mostra por que comparar as posições diretamente não funciona. O caminho
dele foi 9, 10, 0, 1, e o buraco na 0 está nesse caminho, mas o teste
$h \le i < j$ dá $9 \le 0$, falso. Com esse teste, o 31 ficaria na posição 1
com a 0 vazia, e a `Busca(31)` passaria pela 9 e pela 10, pararia na 0 e não o
acharia. Contadas até $j$ com a `distancia`, as posições 9, 10, 0 e 1 ficam a
3, 2, 1 e 0 passos, e a volta no fim do vetor deixa de existir. O buraco está no
caminho quando fica a no máximo tantos passos de $j$ quanto a origem.

O 2 fica onde está. A origem dele vem depois do buraco: até $j$, ela fica a 0
passos, e o buraco a 1. Se o 2 descesse para a posição 1, ficaria antes da
própria origem, e a `Busca(2)`, que começa na 2, não o acharia.

Ao terminar, `go test ./...` passa inteiro.

## Parte 3: medidas contra o `map` (20 min)

O arquivo [tabela/bench_test.go](11/lab/tabela/bench_test.go) traz medidas de
desempenho (*benchmarks*) que fazem as mesmas operações na sua `Tabela` e num
`map[int]int`. As de inserção e de busca usam 100.000 chaves:

* `Insere`: inserções a partir de uma tabela vazia, com chaves aleatórias e com
  chaves sequenciais (0, 1, 2, ...);
* `Busca`: buscas por chaves presentes e, no caso das aleatórias, por chaves
  ausentes;
* `Janela`: começa com a tabela vazia e repete um passo numerado por $i$, com
  $i = 0, 1, 2, ...$ Nos primeiros 1000 passos, o passo só insere a chave $i$.
  Do passo 1000 em diante, ele insere a chave $i$, remove a chave $i - 1000$,
  inserida 1000 passos antes, e busca a chave $i - 500$, que está na tabela. A
  partir daí, a tabela termina cada passo com as 1000 chaves mais recentes, como
  uma lista de pedidos que chegam e saem em ordem. O Go repete o passo dezenas
  de milhares de vezes ou mais, e os primeiros 1000 passos quase não pesam na
  média.

Rode as medidas com:

```bash
go test -run '^$' -bench . ./tabela
```

A execução leva cerca de meio minuto. Cada linha da saída tem o nome da
medida, o número de repetições que o Go escolheu e, na coluna `ns/op`, o tempo
médio de uma operação em nanossegundos. O sufixo `-12` no nome é o número de
processadores que o Go usou:

```
BenchmarkInsere/Tabela/aleatorias-12      10069634       121.3 ns/op
```

Copie a coluna `ns/op` para a tabela do `respostas.md`. Em seguida, troque a
constante `misturar` para `true` no `tabela/tabela.go`, rode as medidas de novo
e preencha as outras duas colunas. Com `true`, a `hash` embaralha os bits da
chave antes de tirar o resto da divisão. Os testes continuam passando, e pulam
só a conferência das posições, que supõe $h(k) = k \bmod m$.

Os números variam de uma máquina para outra, e de uma execução para outra na
mesma máquina. As perguntas 2 e 3 tratam das diferenças grandes, de várias
vezes, que se repetem em qualquer máquina.

**Volte `misturar` para `false` antes de enviar.**

## Relação com o Trabalho Prático

O Trabalho Prático proíbe o `map`. A tabela desta atividade pode ser
reaproveitada nele, como permite o enunciado do trabalho, desde que o código
seja seu.

## Entrega

Um `.zip` na UFPR Virtual, nomeado `ds143_atividade3_<seu_nome>.zip`, com o
projeto Go completo (`go.mod` e o pacote `tabela`) e o `respostas.md`
preenchido. Sem binários compilados. Rode `go test ./...` uma última vez na
pasta que será enviada, com `misturar = false`, e cole a saída no
`respostas.md`.

Em dupla, as duas pessoas enviam o mesmo arquivo, cada uma no seu envio.

## Critérios de avaliação

| Critério | Valor |
|---|---|
| Compila, `go vet` limpo, `gofmt` sem pendências, testes fornecidos passam | 20 |
| `Insere` e `redimensiona` corretas, com o crescimento no limite de 0,7 | 25 |
| `Remove` correta, inclusive com a volta no fim do vetor | 30 |
| Tabela de medidas preenchida com as duas versões da `hash` | 10 |
| `respostas.md` com as três perguntas e a declaração de uso de IA | 15 |
| **Total** | **100** |

Entrega que não compila perde o primeiro item inteiro; os demais são avaliados
pelo que for verificável no código.

## Problemas comuns

**`panic: TODO: implementar ...`**: a função ainda está por escrever. Os `panic`
do esqueleto marcam o que falta, e desaparecem conforme você preenche.

**`go: cannot find main module`**: rode os comandos de dentro da pasta `lab`,
onde está o `go.mod`.

**`go test` não termina**: a tabela encheu, e a `Busca` ou a `Insere` não acham
posição vazia. Confira se a `Insere` chama a `redimensiona` e se a
`redimensiona` reinsere os pares na tabela nova. Para não esperar os 10 minutos
do limite padrão, rode com `go test -timeout 30s ./...`.

**`TestRemoveVoltaCircular` falha com o 1 fora da posição 1**: a sua `Remove`
desce todas as chaves seguintes, sem conferir a posição de origem.

**`TestRemoveVoltaCircular` falha com o 10 na posição 2**: a sua `Remove` para
na primeira chave que não desce, ou compara as posições sem considerar a volta
no fim do vetor.
