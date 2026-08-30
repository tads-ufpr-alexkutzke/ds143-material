# Árvores e percursos

[Slides desta aula (PDF)](05/00_arvores/arvores.pdf)

## Bibliografia recomendada para o tema

* SEDGEWICK, R.; WAYNE, K. **Algorithms**, 4th ed., seção 3.2 (*Binary Search
  Trees*), até a subseção *Analysis*;
* CELES, W. **Introdução a estruturas de dados**, cap. 15 (*Árvores*);
* PEREIRA, S. L. **Estruturas de dados fundamentais**, cap. 7;
* [algs4 - Binary Search Trees](https://algs4.cs.princeton.edu/32bst/).

## Objetivos da aula

Ao final desta aula você deve ser capaz de:

1. Enunciar a definição recursiva de árvore e aplicar a terminologia (raiz,
   folha, nó interno, grau, nível, altura, subárvore);
2. Representar uma árvore binária em Go com uma `struct` e dois ponteiros, e
   construir uma árvore concreta a partir dessa representação;
3. Escrever os percursos em pré-ordem, in-ordem e pós-ordem, e produzir à mão a
   sequência de visitas de cada um sobre uma árvore dada;
4. Escolher o percurso adequado a uma tarefa, justificando pela ordem em que a
   informação fica disponível;
5. Implementar o percurso em largura com uma fila e dizer em que ele difere dos
   três anteriores;
6. Relacionar altura, número de nós e custo de alcançar um nó, e distinguir a
   árvore cheia da degenerada.

Este conteúdo cai na Prova 1.

## 1. Uma hierarquia que você usa todo dia

Antes de qualquer definição, olhe para uma árvore que já está na sua máquina. Na
pasta do material, entre em `src` e rode:

```bash
tree 05/demo
```

```
05/demo
├── curso
│   ├── ds122
│   │   └── aula_01.txt
│   └── ds143
│       ├── aula_01.txt
│       └── aula_02.txt
└── leia-me.txt

4 directories, 4 files
```

Um diretório contém arquivos e outros diretórios, que por sua vez contêm
arquivos e outros diretórios. A regra que descreve `05/demo` é a mesma que
descreve `05/demo/curso`, e a mesma que descreve `05/demo/curso/ds143`. É essa
repetição da mesma regra em escalas diferentes que vamos formalizar hoje.

Agora peça ao sistema o tamanho de cada item, em bytes:

```bash
du -ab 05/demo
```

```
30	05/demo/curso/ds143/aula_01.txt
48	05/demo/curso/ds143/aula_02.txt
78	05/demo/curso/ds143
32	05/demo/curso/ds122/aula_01.txt
32	05/demo/curso/ds122
110	05/demo/curso
46	05/demo/leia-me.txt
156	05/demo
```

Duas coisas nessa saída importam para a aula inteira.

A primeira é a **ordem das linhas**. `05/demo/curso/ds143` aparece depois dos
dois arquivos que estão dentro dele, e `05/demo` aparece por último. O `du` não
tinha alternativa: o tamanho de um diretório é a soma dos tamanhos do que está
dentro, e essa soma só pode ser calculada depois que as partes forem conhecidas.

A segunda é a **aritmética**: 30 + 48 = 78, e 78 + 32 = 110, e 110 + 46 = 156.
Cada linha de diretório é a soma das linhas que a precedem no seu ramo.

> Os tamanhos exatos dependem do sistema de arquivos. Em alguns deles cada
> diretório ocupa alguns bytes por si só, e as somas ficam maiores do que as
> mostradas acima. A ordem das linhas, que é o ponto aqui, não muda.

Essas duas observações voltam com nome na seção 14: o `du` percorre o sistema de
arquivos em **pós-ordem**, e o `tree` o percorre em **pré-ordem**.

## 2. Árvores que você já construiu nesta disciplina

Não é a primeira vez que uma árvore aparece no semestre. Ela apareceu duas
vezes, sem esse nome.

Na aula de análise de algoritmos, o Quick-Union representava as componentes
conexas por um slice `id`, em que `id[i]` guardava o **pai** do elemento `i`.
Cada componente era uma árvore, e o conjunto de todas elas uma floresta. O
`Find` subia da folha até a raiz, e a raiz era o identificador da componente. Foi
justamente a **altura** dessas árvores que fez a diferença entre o Quick-Union e
o Weighted Quick-Union: pendurar a árvore menor sob a maior mantém a altura em
$O(\log n)$, enquanto pendurar sem critério permite que ela chegue a $O(n)$.

Na aula de recursão, o parser de expressões prefixadas
([prefix_parser.go](02/codes/recursao/prefix_parser.go)) lia `* 2 + 3 4` e
devolvia 14. Cada chamada de `eval` tratava uma subexpressão, e as chamadas
aninhadas formavam uma árvore. Diferente do Quick-Union, ali a árvore nunca foi
construída na memória: ela existia apenas como o desenho das chamadas
recursivas. Na seção 11 vamos construí-la de verdade.

## 3. Definição

Uma árvore é um conjunto finito de $n$ elementos chamados **nós**, tal que, se
$n > 0$:

* existe um nó distinguido $r$, chamado **raiz** da árvore;
* os $n - 1$ nós restantes são divididos em $m \geq 0$ conjuntos disjuntos, cada
  um deles também uma árvore, chamados **subárvores** de $r$.

A definição é recursiva, e é assim que vamos programar todas as operações desta
aula. O caso base é $n = 0$: a **árvore vazia**, que não tem raiz.

Duas consequências valem a pena registrar agora, porque voltam mais tarde:

* **existe exatamente um caminho da raiz até cada nó.** Se houvesse dois, algum
  nó teria dois pais, e os conjuntos da definição deixariam de ser disjuntos. A
  estrutura ainda seria útil, mas seria um grafo, assunto do final do semestre;
* **a definição não distingue árvore de subárvore.** Toda subárvore é uma árvore
  completa por direito próprio. É por isso que uma função que recebe uma árvore
  pode ser chamada sobre uma subárvore sem nenhuma adaptação.

Na computação, árvores são desenhadas com a raiz em cima e as folhas embaixo. As
ligações apontam sempre do pai para os filhos, e por isso não se costuma desenhar
as setas.

## 4. Terminologia

Sobre a árvore que vamos usar o resto da aula:

```
            50
          /    \
        30      90
       /  \       \
     20    40      95
    /     /  \
  10    35    45
```

| Termo | Definição | Na árvore acima |
|---|---|---|
| Raiz | o único nó sem pai | 50 |
| Pai de $x$ | o nó imediatamente acima de $x$ | o pai de 35 é 40 |
| Filho de $x$ | nó imediatamente abaixo de $x$ | 20 e 40 são filhos de 30 |
| Grau de um nó | número de subárvores do nó | grau de 30 é 2, de 20 é 1, de 10 é 0 |
| Grau da árvore | maior grau entre todos os nós | 2 |
| Folha (nó externo) | nó de grau 0 | 10, 35, 45, 95 |
| Nó interno | nó com pelo menos um filho | 50, 30, 20, 40, 90 |
| Nível (profundidade) | número de arestas da raiz até o nó | 50 está no nível 0, 35 no nível 3 |
| Altura da árvore | comprimento do caminho mais longo da raiz até uma folha | 3 |
| Floresta | conjunto de árvores disjuntas | as componentes do Quick-Union |

Duas convenções que causam confusão em prova, e que valem para esta disciplina:

* a **altura é contada em arestas**, não em nós. A árvore com um único nó tem
  altura 0;
* a **árvore vazia tem altura -1**. O valor parece arbitrário, mas é o que faz a
  fórmula `altura(a) = 1 + max(altura(esq), altura(dir))` devolver 0 para uma
  folha, sem tratar a folha como caso especial. A seção 13 mostra o código.

## 5. Árvores binárias

Uma **árvore binária** é uma árvore de grau máximo 2, com uma exigência a mais:
cada subárvore é identificada como sendo a da **esquerda** ou a da **direita**.
Qualquer uma das duas pode ser vazia.

A distinção importa. Estas duas árvores são iguais enquanto árvores de grau 2, e
diferentes enquanto árvores binárias:

```
    a          a
   /            \
  b              b
```

Essa é uma restrição, e restringir é o que torna a estrutura útil. A partir da
próxima aula, a posição do filho vai carregar informação: em uma árvore binária
de busca, tudo o que está à esquerda de um nó é menor que ele, e tudo à direita é
maior. Sem a distinção entre esquerda e direita, essa regra não teria como ser
enunciada.

## 6. Representação em Go

Um nó guarda uma informação e dois ponteiros:

```go
type No struct {
    Info int
    Esq  *No
    Dir  *No
}
```

O tipo `No` se refere a si mesmo, o que só é possível através de ponteiros: um
campo `Esq No` daria um tipo de tamanho infinito, e o compilador recusa
(`invalid recursive type No`).

A árvore inteira é representada pelo **ponteiro para o nó raiz**. Um `*No` é,
ao mesmo tempo, um nó e a subárvore que começa nele, e a árvore vazia é o
ponteiro `nil`. Essa coincidência entre a definição matemática e o tipo da
linguagem é o que faz as funções recursivas desta aula ficarem tão curtas: onde
a definição diz "cada subárvore também é uma árvore", o código diz `Percorre(a.Esq)`.

Não existe função para criar árvore vazia nem para liberar a árvore. Em C seria
preciso escrever `cria_arv_vazia` devolvendo `NULL` e uma `arv_libera` recursiva
chamando `free` em pós-ordem; em Go, a árvore vazia é `nil` e o coletor de lixo
recolhe os nós que ficam sem referência.

Duas funções de construção bastam para montar qualquer árvore:

```go
func Constroi(info int, esq, dir *No) *No {
    return &No{Info: info, Esq: esq, Dir: dir}
}

func Folha(info int) *No {
    return Constroi(info, nil, nil)
}
```

`&No{...}` cria o nó e devolve o endereço dele. Em Go isso é seguro mesmo com a
variável saindo de escopo: o compilador percebe que o endereço escapa da função e
aloca o valor onde ele sobreviva.

## 7. Construindo a árvore de exemplo

Com `Constroi` e `Folha`, a árvore da seção 4 sai em uma única expressão, cujo
recuo reproduz o desenho:

```go
func ArvoreExemplo() *No {
    return Constroi(50,
        Constroi(30,
            Constroi(20,
                Folha(10),
                nil),
            Constroi(40,
                Folha(35),
                Folha(45))),
        Constroi(90,
            nil,
            Folha(95)))
}
```

O código completo está em [05/codes/arvore.go](05/codes/arvore.go), junto com
uma função `Imprime` que desenha a árvore deitada e uma `Pertence` que procura um
valor. Rode:

```bash
cd 05/codes
go run arvore.go
```

```
Árvore de exemplo (raiz à esquerda):
        95
    90
50
            45
        40
            35
    30
        20
            10

Pertence(a, 35) = true
Pertence(a, 60) = false
```

A `Pertence` visita todos os nós no pior caso, e não há como fazer melhor:
nada nesta árvore diz de que lado procurar. Essa é exatamente a lacuna que a
árvore binária de busca preenche na próxima aula.

## 8. Percurso

**Percorrer** uma árvore é visitar sistematicamente cada um dos seus nós, uma
vez cada. **Visitar** um nó é fazer alguma coisa com a informação dele: imprimir,
somar, alterar, comparar. Durante um percurso é comum *passar* por um nó várias
vezes sem *visitá-lo*, e a diferença entre as duas palavras é o que separa os
percursos que veremos.

Para muitas tarefas a ordem é indiferente. Para aplicar um reajuste de 5% em
todos os salários guardados na árvore, qualquer ordem serve, desde que nenhum nó
fique de fora e nenhum seja contado duas vezes. Para outras tarefas a ordem é o
problema inteiro, como o `du` da seção 1 mostrou.

Os três primeiros percursos são **em profundidade**: descem por um ramo até o
fim antes de tentar o ramo seguinte. Eles diferem apenas em **quando a raiz é
visitada**, em relação às duas subárvores:

| Percurso | Ordem | Sigla |
|---|---|---|
| Pré-ordem | raiz, esquerda, direita | R, E, D |
| In-ordem | esquerda, raiz, direita | E, R, D |
| Pós-ordem | esquerda, direita, raiz | E, D, R |

Os três códigos têm as mesmas três linhas, em ordens diferentes.

## 9. Pré-ordem, in-ordem e pós-ordem

```go
func PreOrdem(a *No) {
    if a == nil {
        return
    }
    fmt.Print(a.Info, " ")
    PreOrdem(a.Esq)
    PreOrdem(a.Dir)
}

func InOrdem(a *No) {
    if a == nil {
        return
    }
    InOrdem(a.Esq)
    fmt.Print(a.Info, " ")
    InOrdem(a.Dir)
}

func PosOrdem(a *No) {
    if a == nil {
        return
    }
    PosOrdem(a.Esq)
    PosOrdem(a.Dir)
    fmt.Print(a.Info, " ")
}
```

O `if a == nil` é o caso base, e corresponde à árvore vazia da definição. Sem
ele o programa quebra na primeira folha.

Sobre a árvore de exemplo, os três produzem:

```bash
go run percursos.go
```

```
pré-ordem  (R,E,D): 50 30 20 10 40 35 45 90 95
in-ordem   (E,R,D): 10 20 30 35 40 45 50 90 95
pós-ordem  (E,D,R): 10 20 35 45 40 30 95 90 50
```

Acompanhe a pré-ordem à mão sobre o desenho da seção 4: 50 é visitado primeiro,
depois desce-se toda a subárvore da esquerda (30, e dentro dela 20 e 10, depois
40 com 35 e 45), e só então a da direita (90 e 95). A pós-ordem é a que aparece
no `du`: 10 e 20 antes de 30, e a raiz 50 no fim de tudo.

O resultado do in-ordem sobre **esta** árvore veio em ordem crescente. Isso não é
propriedade de toda árvore binária: é propriedade desta árvore, que foi montada
com todos os valores menores à esquerda e maiores à direita. A próxima aula parte
exatamente daí.

Custo dos três percursos: cada nó é visitado uma vez e cada ponteiro é
examinado uma vez, então o tempo é $\Theta(n)$ para uma árvore com $n$ nós. O
espaço é o da pilha de chamadas, proporcional à **altura** da árvore, o que dá
$O(\log n)$ para uma árvore cheia e $O(n)$ para uma degenerada.

## 10. Escolhendo o percurso pela tarefa

A escolha não é estilística. Cada percurso deixa a informação disponível em um
momento diferente, e a tarefa decide qual momento serve.

| Tarefa | Percurso | Por quê |
|---|---|---|
| Copiar ou serializar a árvore preservando a forma | pré-ordem | a raiz precisa existir antes que os filhos sejam pendurados nela |
| Listar os valores de uma árvore binária de busca em ordem | in-ordem | é a ordem que a organização da ABB garante |
| Somar tamanhos, calcular altura, liberar memória em C | pós-ordem | o resultado do nó depende dos resultados dos filhos |
| Encontrar o nó mais próximo da raiz que satisfaça uma condição | em largura | os níveis são visitados em ordem crescente de profundidade |

O caso da pós-ordem é o mais rígido dos quatro. Em `Altura`, na seção 13, as
chamadas às subárvores **têm** que preceder a conta do nó, porque a conta usa os
dois valores devolvidos. Trocar a ordem das linhas não muda o resultado por
gosto: quebra a função.

## 11. Árvore de expressão

O parser da aula de recursão volta aqui, agora com a árvore construída de fato.
A expressão $2 \times (3 + 4)$ vira:

```
    *
   / \
  2   +
     / \
    3   4
```

Operadores nos nós internos, números nas folhas. Os três percursos sobre essa
árvore produzem as três notações aritméticas:

```bash
go run expressao.go
```

```
pré-ordem (prefixa)       : * 2 + 3 4
in-ordem  (sem parênteses): 2 * 3 + 4
in-ordem  (com parênteses): (2 * (3 + 4))
pós-ordem (pós-fixa)      : 2 3 4 + *
valor                     : 14
```

A linha da pré-ordem é exatamente a entrada que o
[prefix_parser.go](02/codes/recursao/prefix_parser.go) da aula de recursão lê:

```bash
echo "* 2 + 3 4" | go run 02/codes/recursao/prefix_parser.go
```

```
14
```

A notação prefixa e a pós-fixa dispensam parênteses, porque a posição do
operador já determina a estrutura. O in-ordem puro, sem parênteses, perde essa
informação: `2 * 3 + 4` lido com a precedência usual vale 10, e não 14. Os
parênteses da terceira linha repõem o que o percurso deixou cair.

A avaliação é um percurso em pós-ordem, pela razão da seção anterior:

```go
func Avalia(a *No) int {
    if a.Esq == nil && a.Dir == nil {
        valor, _ := strconv.Atoi(a.Info)
        return valor
    }

    esq := Avalia(a.Esq)
    dir := Avalia(a.Dir)

    switch a.Info {
    case "+":
        return esq + dir
    case "*":
        return esq * dir
    }
    return 0
}
```

## 12. Percurso em largura

Os três percursos anteriores descem até o fim de um ramo antes de olhar o
seguinte. O **percurso em largura** faz o oposto: visita todos os nós do nível 0,
depois todos os do nível 1, e assim por diante, sempre da esquerda para a direita.

Sobre a árvore de exemplo, a saída é:

```
em largura        : 50 30 90 20 40 95 10 35 45
```

Compare com a pré-ordem, `50 30 20 10 40 35 45 90 95`. A pré-ordem alcança o 10,
que está no nível 3, antes de alcançar o 90, que está no nível 1.

Esse percurso não sai de uma função recursiva de três linhas. A recursão dá
acesso natural aos filhos de quem está sendo visitado, e o que se precisa aqui é
dos **irmãos**, que estão em outro ramo. A solução é guardar explicitamente os
nós ainda por visitar, em uma **fila**: retira-se um nó da frente, visita-se, e
seus filhos entram no fim.

```go
func EmLargura(a *No) {
    if a == nil {
        return
    }

    fila := []*No{a}

    for len(fila) > 0 {
        atual := fila[0]
        fila = fila[1:]

        fmt.Print(atual.Info, " ")

        if atual.Esq != nil {
            fila = append(fila, atual.Esq)
        }
        if atual.Dir != nil {
            fila = append(fila, atual.Dir)
        }
    }
}
```

A fila é um slice usado com `fila[0]` para ler a frente, `fila[1:]` para
descartá-la e `append` para acrescentar no fim. Custo de tempo $\Theta(n)$, como
os outros. O espaço é diferente: a fila chega a conter um nível inteiro de uma
vez, o que dá até $n/2$ nós no último nível de uma árvore cheia, contra a altura
da árvore no caso dos percursos em profundidade.

Este percurso é o mesmo BFS que aparecerá em grafos, no final do semestre, com a
mesma fila e a mesma estrutura de laço.

### Percurso em profundidade sem recursão

Trocar a fila por uma **pilha** transforma o percurso em largura em pré-ordem.
É a mesma ideia com outra disciplina de retirada, e é o que a recursão vinha
fazendo por baixo dos panos, na pilha de chamadas do programa:

```go
func PreOrdemComPilha(a *No) {
    if a == nil {
        return
    }

    pilha := []*No{a}

    for len(pilha) > 0 {
        topo := pilha[len(pilha)-1]
        pilha = pilha[:len(pilha)-1]

        fmt.Print(topo.Info, " ")

        if topo.Dir != nil {
            pilha = append(pilha, topo.Dir)
        }
        if topo.Esq != nil {
            pilha = append(pilha, topo.Esq)
        }
    }
}
```

O filho da direita é empilhado primeiro para que o da esquerda, empilhado
depois, saia antes. A função está no mesmo `percursos.go` e produz a mesma saída
que `PreOrdem`.

## 13. Altura, número de nós e desempenho

A altura da árvore mede o esforço para alcançar o nó mais distante. Em código:

```go
func Altura(a *No) int {
    if a == nil {
        return -1
    }

    esq := Altura(a.Esq)
    dir := Altura(a.Dir)

    if esq > dir {
        return 1 + esq
    }
    return 1 + dir
}
```

O `-1` da árvore vazia é o que faz uma folha devolver `1 + (-1) = 0`, sem caso
especial.

Uma árvore binária **cheia** é aquela em que todo nó interno tem duas subárvores
e todas as folhas estão no último nível. Nela há $2^0 = 1$ nó no nível 0, $2^1$
nós no nível 1, $2^2$ no nível 2, e $2^k$ no nível $k$. Somando todos os níveis
de uma árvore cheia de altura $h$:

$$n = 2^0 + 2^1 + \cdots + 2^h = 2^{h+1} - 1$$

Isolando $h$: uma árvore cheia com $n$ nós tem altura $h = \log_2(n+1) - 1$, ou
seja, $\Theta(\log n)$.

No extremo oposto, uma árvore **degenerada** tem um único filho por nó interno,
um nó por nível. Com $n$ nós, a altura é $n - 1$. A forma é a de uma lista
encadeada, e o desempenho também.

```bash
go run propriedades.go
```

```
Árvore de exemplo
  nós    : 9
  folhas : 4
  altura : 3

Mesmo número de nós, duas formas:

     nós |    cheia | degenerada
---------+----------+-----------
      15 |        3 |         14
      63 |        5 |         62
     255 |        7 |        254
    1023 |        9 |       1022
```

Com 1023 nós, a árvore cheia tem altura 9 e a degenerada 1022: uma diferença de
mais de cem vezes no caminho até o nó mais distante, com a mesma quantidade de
informação armazenada.

É a mesma diferença entre o Quick-Union e o Weighted Quick-Union da aula de
análise de algoritmos, e pela mesma razão. Uma árvore binária com $n$ nós tem
altura mínima proporcional a $\log_2 n$ e altura máxima proporcional a $n$.
Garantir que a árvore fique perto do primeiro caso, e não do segundo, é o assunto
das duas próximas aulas.

## 14. De volta ao `tree` e ao `du`

Com o vocabulário construído, a demonstração da seção 1 pode ser lida de novo.

O `tree` imprime cada diretório **antes** do seu conteúdo, portanto percorre o
sistema de arquivos em **pré-ordem**. É o percurso certo para a tarefa dele:
quem lê a saída precisa saber em que diretório está antes de ver os arquivos
listados.

O `du` imprime cada diretório **depois** do seu conteúdo, portanto em
**pós-ordem**. Também é o percurso certo: o número que ele imprime para um
diretório é a soma dos números dos filhos, e não existe antes deles.

E o sistema de arquivos não é uma árvore binária. Um diretório tem quantos
filhos quiser, o que corresponde ao caso geral da definição da seção 3, com $m$
subárvores. Os percursos continuam valendo, com o laço sobre os filhos no lugar
das duas chamadas fixas.

## 15. Exercícios

**Estes exercícios não valem nota** e não têm entrega. Eles são a preparação
para a aula prática de árvores, essa sim avaliada, e o conteúdo cai na Prova 1.
Resolva antes da próxima aula.

Para os exercícios de código, parta de `05/codes/arvore.go`, que já traz a
estrutura e a árvore de exemplo.

**1.** Sobre a árvore abaixo: (a) dê o grau de cada nó; (b) liste as folhas;
(c) dê a altura e o nível do nó `d`; (d) acrescente os nós necessários para
torná-la uma árvore binária cheia, e diga quantos foram.

```
      a
     / \
    b   c
     \  / \
      d e  f
```

**2.** Escreva as três sequências de percurso em profundidade e a sequência em
largura da árvore do exercício 1, à mão, sem rodar código.

**3.** Escreva `ContaNos(a *No) int`, que devolve o número de nós da árvore.
Compare com `ContaFolhas` de `propriedades.go` e diga por que uma delas tem dois
casos base e a outra tem um.

**4.** Escreva `Maximo(a *No) int`, que devolve o maior valor armazenado na
árvore. Decida o que a função faz com a árvore vazia e justifique a decisão.

**5.** Escreva `Espelha(a *No)`, que troca as subárvores esquerda e direita de
todos os nós. Depois responda: qual dos quatro percursos aplicados à árvore
espelhada devolve exatamente o inverso do percurso in-ordem da árvore original?

**6.** Duas árvores são iguais quando têm a mesma forma e os mesmos valores nas
mesmas posições. Escreva `Iguais(a, b *No) bool`. Quantos casos base a função
precisa?

**7.** Use a notação textual em que a árvore vazia é `<>` e a árvore não vazia é
`<raiz esquerda direita>`. A árvore do exercício 1 fica
`<a <b <> <d <> <>>> <c <e <> <>> <f <> <>>>>`. Escreva uma função que imprima
uma árvore nessa notação, e diga qual percurso ela usa.

**8.** Escreva `NoNoNivel(a *No, nivel int)`, que imprime todos os nós de um
nível dado. Depois, use essa função para escrever um percurso em largura
**sem fila**, chamando-a para cada nível de 0 até a altura da árvore. Qual é o
custo desse percurso em largura, comparado ao $\Theta(n)$ da versão com fila?

### Desafio

Escreva `Reconstroi(preOrdem, inOrdem []int) *No`, que recebe as duas sequências
de percurso de uma árvore com valores distintos e devolve a árvore original.
Antes de programar, convença-se de que o par pré-ordem mais in-ordem determina a
árvore de forma única, e mostre com um contraexemplo que pré-ordem mais
pós-ordem não determina.

---

## Resumo

* Uma árvore é definida recursivamente: uma raiz e um conjunto de subárvores
  disjuntas. Não há distinção entre árvore e subárvore, e é isso que faz as
  funções recursivas serem curtas.
* Em Go, a árvore é o ponteiro para o nó raiz, e a árvore vazia é `nil`.
* Os três percursos em profundidade diferem apenas em quando a raiz é visitada:
  R,E,D na pré-ordem; E,R,D no in-ordem; E,D,R na pós-ordem.
* A pós-ordem é obrigatória quando o resultado de um nó depende dos resultados
  dos filhos, como em `Altura` e na avaliação de uma árvore de expressão.
* O percurso em largura usa uma fila e visita nível a nível. Trocar a fila por
  uma pilha devolve a pré-ordem.
* Todos os percursos custam $\Theta(n)$ em tempo. O espaço é a altura, nos
  percursos em profundidade, e a largura do maior nível, no percurso em largura.
* A altura de uma árvore binária com $n$ nós varia entre $\Theta(\log n)$, na
  árvore cheia, e $n - 1$, na degenerada. Manter a árvore perto do primeiro caso
  é o assunto das próximas aulas.
