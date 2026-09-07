# Árvores binárias de busca e AVL

[Slides desta aula (PDF)](06/00_abb_avl/abb_avl.pdf)

## Bibliografia recomendada para o tema

* SEDGEWICK, R.; WAYNE, K. **Algorithms**, 4th ed., seção 3.2 (*Binary Search
  Trees*), inteira, incluindo *Deletion* e *Analysis*;
* CORMEN, T. et al. **Introduction to Algorithms**, 3rd ed., cap. 12 (*Binary
  Search Trees*) e problema 13-3 (*AVL trees*);
* [algs4 - Binary Search Trees](https://algs4.cs.princeton.edu/32bst/);
* [Visualização interativa de ABB e AVL (USF)](https://www.cs.usfca.edu/~galles/visualization/AVLtree.html),
  útil para conferir à mão o resultado das rotações.

## Objetivos da aula

Ao final desta aula você deve ser capaz de:

1. Enunciar a propriedade da árvore binária de busca e verificar se uma árvore
   dada a satisfaz;
2. Implementar busca, inserção, mínimo e máximo em Go, e justificar por que o
   custo das quatro operações é a altura da árvore;
3. Executar à mão a remoção nos três casos, incluindo a substituição pelo
   antecessor;
4. Explicar por que a ordem de inserção determina a forma da árvore, e dizer
   qual sequência produz o pior caso;
5. Calcular o fator de balanceamento de cada nó e identificar qual dos quatro
   casos de desbalanceamento ocorreu;
6. Aplicar rotações simples e duplas, e justificar por que a rotação não viola a
   propriedade de busca;
7. Comparar os custos de ABB e AVL no melhor caso, no caso médio e no pior caso.

Este conteúdo cai na Prova 1.

## 1. Três formas de guardar um conjunto de chaves

O problema desta aula é o de manter um conjunto de chaves com duas operações:
inserir uma chave nova e responder se uma chave está no conjunto. Com o que você
já viu na disciplina, há duas soluções imediatas, e as duas têm um lado ruim.

O programa [06/codes/motivacao.go](06/codes/motivacao.go) mede as duas, e mais a
estrutura desta aula, inserindo 200 mil chaves e fazendo 10 mil buscas:

```bash
cd 06/codes
go run motivacao.go
```

```
200000 chaves inseridas, 10000 buscas

estrutura        | inserção | busca
-----------------+----------+---------
slice sem ordem  |      3ms | 325ms
slice ordenado   |   2.334s | 1ms
árvore de busca  |     54ms | 2ms
```

> Os tempos variam de máquina para máquina e entre execuções. O que interessa
> não são os números, e sim a diferença de ordem de grandeza entre as linhas.

No slice sem ordem, inserir é `append`, que custa tempo constante em média, e
buscar é varrer o slice inteiro, $\Theta(n)$ no pior caso. No slice ordenado a
troca é exata: a busca binária custa $\Theta(\log n)$, e inserir uma chave no
meio obriga a empurrar para a direita, uma a uma, todas as chaves que vêm depois
dela, o que custa $\Theta(n)$.

A árvore binária de busca fica perto do melhor das duas colunas. É a estrutura
que consegue isso porque não precisa manter as chaves contíguas na memória: a
ordem é dada pela posição de cada nó na árvore, e mudar essa posição custa
alterar um ponteiro.

## 2. A busca que a aula anterior deixou em aberto

Na aula de árvores e percursos, a função `Pertence` procurava um valor visitando
todos os nós, porque nada na árvore dizia de que lado continuar:

```go
func Pertence(a *No, procurado int) bool {
    if a == nil {
        return false
    }
    if a.Info == procurado {
        return true
    }
    return Pertence(a.Esq, procurado) || Pertence(a.Dir, procurado)
}
```

Naquela aula, o percurso in-ordem da árvore de exemplo saiu em ordem crescente,
`10 20 30 35 40 45 50 90 95`. Isso não era acidente nem propriedade de toda
árvore binária: a árvore tinha sido montada com todos os valores menores à
esquerda de cada nó e todos os maiores à direita. Esta aula parte desse fato e o
transforma em exigência.

## 3. Definição

Uma **árvore binária de busca** (ABB) é uma árvore binária em que, para todo nó
$x$:

* todas as chaves da subárvore esquerda de $x$ são menores que a chave de $x$;
* todas as chaves da subárvore direita de $x$ são maiores que a chave de $x$.

A exigência vale para **todos** os nós, e não só para a raiz. Esta árvore
satisfaz a condição na raiz e mesmo assim não é uma ABB, porque o 60 está na
subárvore esquerda do 50:

```
      50
     /  \
   30    90
     \
      60
```

A definição, como a da aula anterior, é recursiva: uma árvore é ABB quando as
duas subárvores são ABB e a chave da raiz separa as duas. O exercício 6 pede a
função que verifica isso, e o erro clássico ali é comparar cada nó apenas com os
seus dois filhos, que é exatamente o teste que a árvore acima passa.

Nesta aula as chaves são distintas. Repetições exigem uma decisão a mais
(ignorar, contar ocorrências, ou fixar um dos lados para os iguais), e a
implementação da seção 5 ignora o valor repetido.

Duas consequências da definição:

* **o percurso in-ordem devolve as chaves em ordem crescente.** É a garantia que
  a árvore da aula anterior tinha por construção e que agora é obrigatória;
* **a busca nunca precisa voltar atrás.** Em cada nó, a comparação descarta uma
  subárvore inteira, e o caminho percorrido é único.

## 4. Busca

A comparação em cada nó tem três resultados, e cada um decide o passo seguinte:

```go
func Busca(a *No, procurado int) bool {
    if a == nil {
        return false
    }
    if procurado < a.Info {
        return Busca(a.Esq, procurado)
    }
    if procurado > a.Info {
        return Busca(a.Dir, procurado)
    }
    return true
}
```

As duas chamadas recursivas estão na posição final da função, e nenhuma delas
precisa do resultado da outra. Uma recursão assim se troca por um laço sem
esforço:

```go
func BuscaIterativa(a *No, procurado int) bool {
    for a != nil {
        if procurado < a.Info {
            a = a.Esq
        } else if procurado > a.Info {
            a = a.Dir
        } else {
            return true
        }
    }
    return false
}
```

Compare com `Pertence`, da seção 2, que precisa das duas subárvores e por isso
não vira laço com a mesma facilidade.

O número de comparações é o número de nós no caminho da raiz até onde a busca
termina, ou seja, no máximo a altura da árvore mais um. Esse fato vale para
todas as operações desta aula, e a seção 9 volta a ele.

## 5. Inserção

O nó novo entra sempre como folha, na posição em que a busca por aquela chave
terminaria. A função devolve a raiz da árvore resultante, e o pai reatribui o
próprio ponteiro com o valor devolvido:

```go
func Insere(a *No, v int) *No {
    if a == nil {
        return &No{Info: v}
    }
    if v < a.Info {
        a.Esq = Insere(a.Esq, v)
    } else if v > a.Info {
        a.Dir = Insere(a.Dir, v)
    }
    return a
}
```

A linha `a.Esq = Insere(a.Esq, v)` reatribui o ponteiro em todos os nós do
caminho, embora só um deles mude de fato. É o preço de não guardar o ponteiro
para o pai dentro do nó, e é o mesmo padrão que a remoção e o rebalanceamento
vão usar. Quem chama escreve `a = Insere(a, v)`, o que também cobre o caso da
árvore vazia, em que a raiz passa a existir.

Inserir os valores `50 30 90 20 40 95 10 35 45`, nessa ordem, reconstrói
exatamente a árvore de exemplo da aula anterior:

```bash
go run abb.go
```

```
Árvore construída por inserções (raiz à esquerda):
        95
    90
50
            45
        40
            35
    30
        20
            10

in-ordem: 10 20 30 35 40 45 50 90 95 

altura: 3
mínimo: 10  máximo: 95

Busca(a, 35) = true
Busca(a, 60) = false
Busca(a, 95) = true
```

## 6. Mínimo, máximo e listagem em ordem

O menor valor está no nó mais à esquerda, e o maior no mais à direita. Nenhuma
comparação é necessária, basta descer sempre para o mesmo lado:

```go
func Minimo(a *No) (int, bool) {
    if a == nil {
        return 0, false
    }
    for a.Esq != nil {
        a = a.Esq
    }
    return a.Info, true
}
```

O segundo valor devolvido diz se a resposta existe, e é como Go trata a
ausência de resultado sem recorrer a um valor especial como $-1$. A árvore vazia
não tem mínimo, e devolver 0 sozinho seria confundi-la com uma árvore que contém
o 0.

Listar as chaves em ordem crescente é o percurso in-ordem da aula anterior, sem
nenhuma alteração. O código completo desta seção e das duas anteriores está em
[06/codes/abb.go](06/codes/abb.go).

## 7. A ordem de inserção determina a forma

Um mesmo conjunto de chaves produz árvores diferentes conforme a ordem em que as
chaves são inseridas. O caso extremo é inserir já em ordem crescente: cada chave
nova é maior que todas as anteriores, desce sempre para a direita e vira folha
no fim do único ramo.

```bash
go run forma.go
```

```
Inserindo 50 30 90 20 40 95 10 35 45:
        95
    90
50
            45
        40
            35
    30
        20
            10
altura: 3

Inserindo os mesmos valores em ordem crescente:
                                95
                            90
                        50
                    45
                40
            35
        30
    20
10
altura: 8
```

A segunda árvore tem os mesmos nove valores e é uma lista encadeada com passos
extras. A busca nela custa o mesmo que a varredura do slice sem ordem da seção 1,
e a inserção também.

Dados de entrada em ordem crescente não são um caso raro. Chaves lidas de um
arquivo já ordenado, códigos sequenciais e carimbos de tempo produzem essa
situação justamente nos casos em que a árvore seria mais útil.

Entre os dois extremos está o caso médio. Inserindo em ordem aleatória, a árvore
fica muito mais perto da altura mínima do que da degenerada:

```
Inserções em ordem aleatória (média de 100 árvores, 5 para n >= 100 mil):
       n | mínima | aleatória | degenerada | comparações médias
---------+--------+-----------+------------+-------------------
    1000 |      9 |      21.1 |        999 |               12.0
   10000 |     13 |      30.6 |       9999 |               16.6
  100000 |     16 |      40.6 |      99999 |               21.0
 1000000 |     19 |      46.8 |     999999 |               25.3
```

A última coluna é o número médio de comparações de uma busca bem sucedida.
Sedgewick demonstra que esse valor tende a $2 \ln n \approx 1{,}39 \log_2 n$,
cerca de 39% acima do que daria uma árvore perfeitamente balanceada, e a coluna
medida acompanha essa conta. A altura média também é $\Theta(\log n)$, com uma
constante maior.

O caso médio é bom, e ainda assim insuficiente: ele supõe que a ordem de chegada
das chaves seja aleatória, e quem escreve a estrutura não controla essa ordem.

## 8. Remoção

A remoção começa como uma busca, e o que fazer ao encontrar o nó depende de
quantas subárvores ele tem.

**Caso 1, folha.** O nó some, e o pai passa a apontar para `nil`.

**Caso 2, uma subárvore só.** O filho sobe para o lugar do nó removido. A
propriedade de busca continua valendo, porque toda a subárvore que subiu já
estava do lado certo em relação ao pai.

**Caso 3, duas subárvores.** Nenhum dos dois filhos pode simplesmente subir: o
lugar do nó removido comporta uma chave que seja maior que tudo o que está à
esquerda e menor que tudo o que está à direita, e existem exatamente duas chaves
assim na árvore. Uma delas é o **antecessor**, o maior valor da subárvore
esquerda; a outra é o **sucessor**, o menor valor da subárvore direita. Copiamos
uma das duas para o nó e removemos o nó de onde ela veio.

Essa segunda remoção termina: o antecessor é o nó mais à direita da subárvore
esquerda, e por isso não tem filho à direita, o que a joga no caso 1 ou no caso
2. Nunca no caso 3.

```go
func Remove(a *No, v int) *No {
    if a == nil {
        return nil
    }

    if v < a.Info {
        a.Esq = Remove(a.Esq, v)
        return a
    }
    if v > a.Info {
        a.Dir = Remove(a.Dir, v)
        return a
    }

    if a.Esq == nil && a.Dir == nil {
        return nil
    }
    if a.Esq == nil {
        return a.Dir
    }
    if a.Dir == nil {
        return a.Esq
    }

    antecessor := a.Esq
    for antecessor.Dir != nil {
        antecessor = antecessor.Dir
    }
    a.Info = antecessor.Info
    a.Esq = Remove(a.Esq, antecessor.Info)
    return a
}
```

Os dois primeiros `if` do bloco final tratam o caso 1 e o caso 2 juntos: um nó
sem os dois filhos tem `a.Esq == nil` ou `a.Dir == nil`, e devolver o outro
ponteiro funciona tanto quando ele aponta para um filho quanto quando ele é
`nil`. O teste explícito da folha, que aparece primeiro aqui, é redundante e
está no código apenas para deixar os três casos visíveis.

```bash
go run remocao.go
```

```
Remove(a, 45), caso 1: folha
        95
    90
50
        40
            35
    30
        20
            10
in-ordem: 10 20 30 35 40 50 90 95 

Remove(a, 20), caso 2: uma subárvore só
        95
    90
50
        40
            35
    30
        10
in-ordem: 10 30 35 40 50 90 95 

Remove(a, 50), caso 3: duas subárvores (a raiz)
        95
    90
40
        35
    30
        10
in-ordem: 10 30 35 40 90 95 
```

Na última remoção, o 50 saiu e o 40, maior valor da subárvore esquerda, tomou o
lugar dele na raiz. Em todas as saídas o in-ordem continua crescente, que é o
teste mais rápido para saber se uma operação preservou a propriedade de busca.

## 9. O custo de todas as operações é a altura

Busca, inserção, remoção, mínimo e máximo percorrem um único caminho da raiz para
baixo. O custo de todas elas é $O(h)$, com $h$ a altura da árvore, e a seção 7
mostrou que $h$ depende da ordem de inserção:

| Situação | Altura | Custo das operações |
|---|---|---|
| Árvore cheia (melhor caso) | $\log_2(n+1) - 1$ | $\Theta(\log n)$ |
| Inserções em ordem aleatória (caso médio) | $\Theta(\log n)$ | $\Theta(\log n)$ |
| Inserções em ordem crescente ou decrescente (pior caso) | $n - 1$ | $\Theta(n)$ |

O pior caso é a entrada ordenada, uma das sequências mais comuns na prática. O
resto da aula trata de eliminá-lo.

## 10. O critério de balanceamento da AVL

Manter a árvore perfeitamente balanceada, com todos os níveis cheios menos o
último, custaria caro demais: uma única inserção pode obrigar a reconstruir
quase tudo. A solução das árvores AVL, propostas por Adelson-Velsky e Landis em
1962, é exigir menos.

Uma árvore binária de busca é **AVL** quando, para todo nó, as alturas das duas
subárvores diferem em no máximo 1.

Chamamos de **fator de balanceamento** de um nó a diferença entre a altura da sua
subárvore esquerda e a altura da sua subárvore direita:

$$\mathrm{fb}(x) = \mathrm{altura}(x.\mathrm{Esq}) - \mathrm{altura}(x.\mathrm{Dir})$$

Em uma AVL, todo nó tem fb igual a $-1$, 0 ou $+1$. Um nó com fb $= +2$ está
pesado à esquerda; com fb $= -2$, pesado à direita. Vale a convenção da aula
anterior: a árvore vazia tem altura $-1$.

A árvore de exemplo desta aula é AVL. A degenerada da seção 7 não é: a raiz tem
fb $= -8$.

Calcular a altura a cada verificação custaria $\Theta(n)$, o que jogaria fora a
economia. Por isso cada nó guarda a própria altura em um campo, atualizado
quando os filhos mudam:

```go
type No struct {
    Info int
    Alt  int
    Esq  *No
    Dir  *No
}

func altura(a *No) int {
    if a == nil {
        return -1
    }
    return a.Alt
}

func atualizaAltura(a *No) {
    a.Alt = 1 + maior(altura(a.Esq), altura(a.Dir))
}
```

`maior` devolve o maior de dois inteiros e está definida no mesmo arquivo.

Com a altura armazenada, calcular o fator de balanceamento de um nó custa tempo
constante.

## 11. Rotação

A operação que corrige o desbalanceamento é a **rotação**, que muda a forma da
árvore mexendo em três ponteiros e preserva a ordem das chaves.

```
       a                     b
      / \                   / \
     b   z    ------->     x   a
    / \      rotação          / \
   x   y     à direita       y   z
```

Percorra as duas árvores em in-ordem. À esquerda: $x$, $b$, $y$, $a$, $z$. À
direita: $x$, $b$, $y$, $a$, $z$. As duas sequências são iguais, e é por isso
que a rotação nunca viola a propriedade de busca. O que muda é a altura: $b$
sobe um nível com toda a sua subárvore $x$, e $a$ desce um nível com $z$.

```go
func RotacaoDireita(a *No) *No {
    b := a.Esq
    a.Esq = b.Dir
    b.Dir = a

    atualizaAltura(a)
    atualizaAltura(b)
    return b
}
```

São três atribuições de ponteiro, e a ordem entre elas importa: `b.Dir` é lido
para dentro de `a.Esq` antes de ser sobrescrito com `a`. As alturas são
atualizadas de baixo para cima, primeiro a de `a`, que agora é filho, depois a
de `b`. A função devolve a nova raiz da subárvore, e quem chamou reatribui o
ponteiro, no mesmo padrão da inserção.

`RotacaoEsquerda` é a imagem espelhada, trocando `Esq` por `Dir`.

## 12. Os quatro casos

Depois de uma inserção, o primeiro nó desbalanceado encontrado na volta da
recursão tem fb $= +2$ ou fb $= -2$. Qual rotação aplicar depende também do
lado em que o filho mais pesado cresceu, o que dá quatro casos.

| Caso | fb do nó | fb do filho pesado | Correção |
|---|---|---|---|
| Esquerda-esquerda | $+2$ | $+1$ | rotação simples à direita |
| Direita-direita | $-2$ | $-1$ | rotação simples à esquerda |
| Esquerda-direita | $+2$ | $-1$ | rotação à esquerda no filho, depois à direita no nó |
| Direita-esquerda | $-2$ | $+1$ | rotação à direita no filho, depois à esquerda no nó |

Nos dois primeiros casos, o nó e o filho pendem para o mesmo lado, e uma rotação
resolve. Nos dois últimos, o desequilíbrio faz um zigue-zague, e uma rotação
simples apenas o inverteria: a primeira rotação alinha o filho com o avô, e a
segunda corrige o conjunto.

O caso esquerda-direita, com a sequência de inserção `30 10 20`:

```
     30              30            20
    /               /             /  \
  10       -->    20      -->   10    30
    \            /
     20        10
```

O programa [06/codes/avl.go](06/codes/avl.go) monta os quatro casos e mostra a
árvore antes e depois, com o fator de balanceamento ao lado de cada nó:

```bash
go run avl.go
```

```
[30 20 10]: esquerda-esquerda (rotação simples à direita)
  sem rebalanceamento (altura 2):
    30 (fb 2)
        20 (fb 1)
            10 (fb 0)
  depois da rotação (altura 1):
        30 (fb 0)
    20 (fb 0)
        10 (fb 0)

[30 10 20]: esquerda-direita (rotação dupla)
  sem rebalanceamento (altura 2):
    30 (fb 2)
            20 (fb 0)
        10 (fb -1)
  depois da rotação (altura 1):
        30 (fb 0)
    20 (fb 0)
        10 (fb 0)
```

## 13. Inserção em AVL

A inserção da AVL é a da seção 5 com duas linhas a mais no fim, executadas na
volta da recursão, em cada nó do caminho da folha nova até a raiz:

```go
func Insere(a *No, v int) *No {
    if a == nil {
        return &No{Info: v}
    }

    if v < a.Info {
        a.Esq = Insere(a.Esq, v)
    } else if v > a.Info {
        a.Dir = Insere(a.Dir, v)
    } else {
        return a
    }

    atualizaAltura(a)
    return rebalanceia(a)
}

func rebalanceia(a *No) *No {
    fb := fator(a)

    if fb > 1 {
        if fator(a.Esq) < 0 {
            a.Esq = RotacaoEsquerda(a.Esq)
        }
        return RotacaoDireita(a)
    }

    if fb < -1 {
        if fator(a.Dir) > 0 {
            a.Dir = RotacaoDireita(a.Dir)
        }
        return RotacaoEsquerda(a)
    }

    return a
}
```

Os quatro casos da tabela da seção 12 estão nos dois blocos: o `if` interno
distingue o caso simples do duplo, e a rotação que fecha cada bloco é comum aos
dois.

Uma inserção estraga o balanceamento de no máximo um nó, e uma rotação, simples
ou dupla, devolve à subárvore a altura que ela tinha antes da inserção. Como a
altura da subárvore não mudou, nenhum nó acima dela precisa ser tocado: a
inserção em AVL faz **no máximo uma** rotação. A remoção não tem essa
propriedade, e pode exigir rotações em todos os nós do caminho até a raiz.

A remoção da AVL, também escrita como a da ABB acrescida de `atualizaAltura` e
`rebalanceia`, está em `avl.go`.

## 14. A altura da AVL

Resta mostrar que o critério da seção 10, que é bem mais fraco que o
balanceamento perfeito, basta para garantir altura logarítmica.

Seja $N(h)$ o **menor** número de nós de uma AVL de altura $h$. A árvore mais
esguia possível com altura $h$ tem uma raiz, uma subárvore de altura $h-1$, e a
outra tão pequena quanto o critério permite, ou seja, de altura $h-2$:

$$N(h) = N(h-1) + N(h-2) + 1, \quad N(0) = 1, \quad N(1) = 2$$

| $h$ | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10 |
|---|---|---|---|---|---|---|---|---|---|---|---|
| $N(h)$ | 1 | 2 | 4 | 7 | 12 | 20 | 33 | 54 | 88 | 143 | 232 |

A recorrência é a de Fibonacci deslocada, $N(h) = F(h+3) - 1$, e cresce
exponencialmente com $h$. Invertendo a relação, uma AVL com $n$ nós tem

$$h < 1{,}44 \log_2 (n + 2)$$

o que é $O(\log n)$, com uma constante 44% acima da altura mínima. Uma AVL com
10 mil chaves tem altura no máximo 17, porque a menor AVL de altura 18 já
precisaria de 10.945 nós.

Como busca, inserção e remoção custam $O(h)$, as três passam a custar
$O(\log n)$ **no pior caso**, e não mais só em média.

## 15. ABB e AVL na entrada ordenada

O último bloco de `avl.go` insere `1, 2, 3, ..., n`, a sequência que degenera a
ABB, nas duas estruturas:

```
       n | altura na ABB | altura na AVL
---------+---------------+--------------
      10 |             9 |             3
     100 |            99 |             6
    1000 |           999 |             9
   10000 |          9999 |            13
```

Com 10 mil chaves em ordem crescente, a busca na ABB percorre até 10 mil nós, e
na AVL até 14. A altura 13 obtida é a menor possível para 10 mil nós, porque
essa sequência de inserção leva a AVL a preencher os níveis quase por completo.

O que se paga por isso: um campo a mais em cada nó, o custo de atualizar a
altura em cada nó do caminho, e as rotações. As três parcelas são $O(1)$ por nó
visitado, e não mudam a ordem de grandeza das operações.

| | ABB | AVL |
|---|---|---|
| Busca, inserção e remoção no caso médio | $\Theta(\log n)$ | $\Theta(\log n)$ |
| Busca, inserção e remoção no pior caso | $\Theta(n)$ | $\Theta(\log n)$ |
| Memória por nó | chave e 2 ponteiros | chave, 2 ponteiros e altura |
| Rotações por inserção | nenhuma | no máximo 1 |
| Rotações por remoção | nenhuma | até $O(\log n)$ |

A AVL não é a única resposta ao problema do balanceamento. As árvores
rubro-negras, que veremos na segunda metade da disciplina, admitem árvores um
pouco mais desbalanceadas e fazem menos rotações na remoção, o que as torna a
escolha mais comum em bibliotecas de uso geral. O critério de decisão entre as
duas é a proporção entre buscas e alterações na aplicação.

## 16. De volta à comparação da seção 1

A tabela da abertura pode ser lida de novo, agora com a última linha explicada.
O slice ordenado é rápido na busca porque as chaves estão em ordem, e lento na
inserção porque manter a ordem exige mover as chaves. A árvore de busca guarda a
mesma informação de ordem na forma da árvore, e por isso inserir custa apenas
descer até uma folha e pendurar um nó.

A linha da árvore naquela medição foi obtida com chaves em ordem aleatória, que
é o caso médio da seção 7. Com as mesmas 200 mil chaves em ordem crescente, a
árvore da seção 5 daria uma coluna de 200 mil níveis e as duas medições ficariam
piores que as do slice. A AVL é o que remove essa dependência da ordem de
chegada.

## 17. Exercícios

**Estes exercícios não valem nota** e não têm entrega. Eles preparam a aula
prática de árvores, essa sim avaliada, e o conteúdo cai na Prova 1. Resolva
antes da próxima aula.

Para os exercícios de código, parta de [06/codes/abb.go](06/codes/abb.go) ou de
[06/codes/avl.go](06/codes/avl.go).

**1.** Desenhe a ABB resultante de inserir, nesta ordem, `40 20 60 10 30 50 70`.
Depois desenhe a que resulta de inserir os mesmos valores na ordem
`10 20 30 40 50 60 70`. Dê a altura das duas e o número de comparações que a
busca por 70 faz em cada uma.

**2.** Sobre a primeira árvore do exercício 1, execute à mão, em sequência, as
remoções de 10, de 60 e de 40. Diga em qual caso da seção 8 cada uma cai e
desenhe a árvore depois de cada remoção.

**3.** A remoção da seção 8 usa o antecessor. Reescreva o caso 3 usando o
sucessor, o menor valor da subárvore direita, e verifique que o resultado do
exercício 2 muda de forma mas continua sendo uma ABB.

**4.** Escreva `Conta(a *No, x, y int) int`, que devolve quantas chaves da
árvore estão no intervalo $[x, y]$, sem visitar as subárvores que não podem
conter nenhuma chave do intervalo. Qual é o custo da sua função em uma árvore de
altura $h$ com $k$ chaves no intervalo?

**5.** Escreva `Sucessor(a *No, v int) (int, bool)`, que devolve a menor chave da
árvore que é maior que `v`, sem supor que `v` esteja na árvore.

**6.** Escreva `EhABB(a *No) bool`, que verifica se uma árvore binária qualquer
satisfaz a propriedade de busca. Teste a sua função com a árvore da seção 3, que
tem o 60 na subárvore esquerda do 50, e explique por que a versão que compara
cada nó só com os seus dois filhos aceita essa árvore.

**7.** Sem rodar código, dê o fator de balanceamento de cada nó da árvore
resultante de inserir `50 30 90 20 40 95 10 35 45`, e conclua se ela é AVL.
Confira depois com `go run avl.go`.

**8.** Insira, em uma AVL vazia, os valores `10 20 30 40 50 25`, um a um,
desenhando a árvore após cada inserção e nomeando o caso da seção 12 sempre que
houver rotação.

**9.** Escreva `AlturaRecalculada(a *No) int`, que calcula a altura percorrendo a
árvore, e uma função que verifique se o campo `Alt` de todos os nós de uma AVL
está correto. Use-a depois de uma sequência de inserções e remoções em `avl.go`.

**10.** Altere `forma.go` para medir também a altura de uma AVL construída com as
mesmas permutações aleatórias. Compare as três colunas e diga quanto o
rebalanceamento ganha no caso médio, e não só no pior caso.

### Desafio

Uma ABB pode ser reequilibrada sem rotação nenhuma: percorra a árvore em
in-ordem, guarde as chaves em um slice, que sai ordenado, e reconstrua a árvore
tomando o elemento do meio como raiz, o meio de cada metade como filho, e assim
por diante. Implemente `Reequilibra(a *No) *No` com essa ideia, mostre que a
árvore resultante tem altura mínima, e diga por que essa estratégia não
substitui a AVL em uma estrutura que recebe inserções ao longo do tempo.

---

## Resumo

* Uma árvore binária de busca exige, em todo nó, que a subárvore esquerda tenha
  apenas chaves menores e a direita apenas chaves maiores. A exigência vale para
  todos os nós, e não só para a raiz.
* O percurso in-ordem de uma ABB devolve as chaves em ordem crescente, e é o
  teste mais rápido para verificar se uma operação preservou a propriedade.
* Busca, inserção, remoção, mínimo e máximo percorrem um único caminho da raiz
  para baixo e custam $O(h)$.
* A remoção tem três casos, e o de duas subárvores se resolve copiando o
  antecessor ou o sucessor para o nó e removendo aquele nó, o que sempre recai
  em um dos dois casos mais simples.
* A ordem de inserção determina a forma da árvore. Em ordem aleatória a altura
  fica em $\Theta(\log n)$; em ordem crescente ou decrescente, a árvore degenera
  para $n - 1$.
* Uma árvore AVL é uma ABB em que todo nó tem fator de balanceamento $-1$, 0 ou
  $+1$. O critério garante altura menor que $1{,}44 \log_2(n+2)$.
* A rotação muda a forma da árvore mexendo em três ponteiros e preserva a ordem
  in-ordem. São quatro casos de desbalanceamento, dois resolvidos por rotação
  simples e dois por rotação dupla.
* A inserção em AVL é a da ABB seguida de atualização de altura e
  rebalanceamento na volta da recursão, e faz no máximo uma rotação.
