Qual é o problema da busca de substrings? Defina o que é o "padrão" (pattern) e o "texto" (text).

<details>
  <sumary>Resposta</sumary>

O problema da busca de substrings consiste em encontrar uma ou todas as ocorrências de uma string chamada "padrão" (P) dentro de uma string maior chamada "texto" (T). Por exemplo, encontrar o padrão "ANA" no texto "BANANA".

</details>

---

Descreva o funcionamento e a complexidade de pior caso do algoritmo de busca de substrings de força bruta. Dê um exemplo de um padrão e um texto que causem o pior caso.

<details>
  <sumary>Resposta</sumary>

O algoritmo de força bruta funciona deslizando o padrão sobre o texto, caractere por caractere. Para cada posição $i$ no texto, ele tenta "casar" o padrão $P[0..m-1]$ com o texto $T[i..i+m-1]$.

  * **Funcionamento:** Um laço externo percorre as posições $i$ de $0$ a $n-m$ (onde $n$ é o tamanho do texto e $m$ o do padrão). Um laço interno compara $P[j]$ com $T[i+j]$. Se uma diferença for > encontrada, o laço interno é quebrado e o laço externo avança (incrementa $i$).
  * **Pior Caso:** A complexidade de pior caso é $O(n \cdot m)$.
  * **Exemplo:** Isso ocorre quando, para cada posição do texto, quase todo o padrão casa antes de falhar no último caractere.
      * Texto (T): "AAAAAAAAAAAAAAAAAB"
      * Padrão (P): "AAAAB"

</details>

---

Como o algoritmo KMP (Knuth-Morris-Pratt) consegue realizar uma busca mais rápida? Qual o principal artifício utilizado no algoritmo para isso?

<details>
  <sumary>Resposta</sumary>

O algoritmo KMP consegue uma busca mais rápida porque ele elimina re-comparações desnecessárias no texto.

Ao contrário do algoritmo de força bruta, que após uma falha de correspondência retrocede o ponteiro do texto e avança o padrão em apenas uma posição, o KMP nunca retrocede o ponteiro do texto.

O principal artifício para conseguir isso é o pré-processamento do padrão para construir uma estrutura de dados (um array lps ou um Autômato Finito Determinístico - DFA) que armazena informações sobre as sobreposições internas do próprio padrão. 

</details>

--- 

Qual é a diferença fundamental entre compressão *lossless* (sem perdas) e *lossy* (com perdas)? Em que categoria o algoritmo de Huffman se encaixa e por quê?

<details>
  <sumary>Resposta</sumary>

  * **Compressão Lossless (Sem Perdas):** Reduz o tamanho do arquivo de forma que o arquivo original possa ser perfeitamente reconstruído a partir do arquivo comprimido. Nenhum dado é perdido. É usada para textos, código-fonte e dados científicos.
  * **Compressão Lossy (Com Perdas):** Reduz o tamanho do arquivo descartando permanentemente informações consideradas "menos importantes". O arquivo original não pode ser recuperado. É usada para mídia, como imagens (JPEG), áudio (MP3) e vídeo.
  * **Huffman:** O algoritmo de Huffman é **lossless**. Ele funciona atribuindo códigos binários de tamanho variável aos caracteres (mais frequentes recebem códigos mais curtos), mas a correspondência é exata, permitindo a descompressão perfeita.

</details>

---

O que é uma "codificação livre de prefixo" (ou *prefix-free code*) e por que essa propriedade é essencial para o funcionamento correto da descompressão no algoritmo de Huffman?

<details>
  <sumary>Resposta</sumary>

Uma codificação de prefixo é um conjunto de códigos (binários, no caso de Huffman) onde nenhum código é o prefixo de outro código.

  * **Exemplo:** Se "A" é `01`, "B" não pode ser `010` (pois `01` é prefixo). Um conjunto válido seria: A=`0`, B=`10`, C=`11`.
  * **Importância:** Essa propriedade é essencial para a descompressão, pois permite que o descompressor leia o fluxo de bits comprimido de forma inequívoca. Ao ler os bits, assim que uma sequência corresponde a um código na árvore de Huffman, o descompressor sabe que aquele símbolo foi encontrado. Não há ambiguidade; ele não precisa "olhar para frente" para ver se bits adicionais poderiam formar um código diferente (e mais longo).

</details>

---

Dada a string "A A A B B C D A", construa a árvore de Huffman correspondente. Em seguida, mostre o código binário para cada caractere (A, B, C, D) e o resultado da compressão da string inteira.

<details>
  <sumary>Resposta</sumary>
  
 1.  **Frequências:**
       * A: 4
       * B: 2
       * C: 1
       * D: 1
 2.  **Construção da Árvore (Fila de Prioridade):**
       * Início: `(C:1), (D:1), (B:2), (A:4)`
       * Combina os dois menores (C e D): `(CD:2)`
       * Fila: `(CD:2), (B:2), (A:4)`
       * Combina os dois menores (CD e B). (A ordem entre B e CD não importa, o resultado será o mesmo. Vamos pegar CD e B): `(BCD:4)`
       * Fila: `(A:4), (BCD:4)`
       * Combina os dois últimos: `(ABCD:8)`
 3.  **Árvore Final (assumindo '0' para esquerda e '1' para direita):**
     ```
           (ABCD:8)
          /        \
        (A:4)     (BCD:4)
                  /     \
             (B:2)   (CD:2)
                     /    \
                  (C:1)  (D:1)
     ```
 4.  **Códigos Binários:**
       * A: `0`
       * B: `10`
       * C: `110`
       * D: `111`
 5.  **String Comprimida ("A A A B B C D A"):**
     `0 0 0 10 10 110 111 0`
     (Resultado: `00010101101110`)

</details>

--- 

Um grafo pode ser utilizado para mostrar relacionamentos entre pessoas. Por exemplo, dada a seguinte lista de pessoas pertencentes a um mesmo curso e suas relações de amizades abaixo:

```
Pessoas = {Jorge, Jaime, José, Francisco, Frederico, João, Suzana}

Amizades = {(Jorge, Jaime), (Francisco, Frederico), (Jorge, João),
            (Jaime, Frederico), (Jaime, Francisco), (Jaime, Suzana), 
            (Suzana, Francisco)}
```

Determine:

a. todos os amigos de João;
b. todos os amigos de Suzana;
c. todos os amigos de Jaime;
d. O grau de cada vértice.

<details>
  <sumary>Resposta</sumary>

a. **Todos os amigos de João:**
   * Jorge

b. **Todos os amigos de Suzana:**
   * Jaime
   * Francisco

c. **Todos os amigos de Jaime:**
   * Jorge
   * Frederico
   * Francisco
   * Suzana

d. **O grau de cada vértice:**
   * Grau(Jorge) = 2
   * Grau(Jaime) = 4
   * Grau(José) = 0 (José é um vértice isolado)
   * Grau(Francisco) = 3
   * Grau(Frederico) = 2
   * Grau(João) = 1
   * Grau(Suzana) = 2
</details>


---

Compare a Busca em Largura (BFS) e a Busca em Profundidade (DFS). Qual estrutura de dados auxiliar cada uma utiliza e qual é uma aplicação comum para cada uma?

<details>
  <sumary>Resposta</sumary>

  * **BFS (Busca em Largura):** Explora o grafo "em camadas", visitando primeiro todos os vizinhos de um nó antes de se aprofundar.
      * **Estrutura de Dados:** Fila (Queue), para garantir a ordem FIFO (First-In, First-Out).
      * **Aplicação Comum:** Encontrar o caminho mais curto (em número de arestas) em um grafo não-ponderado.
  * **DFS (Busca em Profundidade):** Explora o grafo "indo o mais fundo possível" por um caminho antes de retroceder (backtracking).
      * **Estrutura de Dados:** Pilha (Stack) (ou recursão, que usa a pilha de chamadas).
      * **Aplicação Comum:** Detecção de ciclos, ordenação topológica (em grafos direcionados), encontrar componentes conexas.

</details>

---

Dado um grafo não-direcionado, como você pode usar BFS ou DFS para determinar o número de componentes conexas desse grafo? Descreva o algoritmo.

<details>
  <sumary>Resposta</sumary>

Uma componente conexa é um subgrafo onde existe um caminho entre quaisquer dois vértices.
O algoritmo funciona da seguinte forma:

1.  Inicialize um contador de componentes `count = 0`.
2.  Mantenha um array (ou conjunto) de "visitados" para todos os vértices do grafo, inicializado como falso.
3.  Itere por todos os vértices $v$ do grafo (de $0$ a $V-1$).
4.  Para cada vértice $v$:
      * Se $v$ **não** foi visitado:
          * Incremente `count` (pois encontramos o início de uma nova componente).
          * Inicie uma busca (DFS ou BFS) a partir de $v$. A busca irá percorrer e marcar como "visitados" todos os vértices que são alcançáveis a partir de $v$ (ou seja, todos os vértices daquela componente).
5.  Ao final da iteração, `count` conterá o número total de componentes conexas.

</details>

---

Em um grafo não-direcionado, a BFS encontra o caminho mais curto entre dois nós $s$ e $t$ (em número de arestas). Por que a DFS não garante encontrar o caminho mais curto?

<details>
  <sumary>Resposta</sumary>

A DFS não garante o caminho mais curto porque ela não explora o grafo em camadas de distância. A natureza da DFS é seguir um caminho até o fim (o mais profundo possível) antes de retroceder (backtracking).

A DFS pode encontrar um caminho válido entre $s$ e $t$, mas é provável que seja o primeiro caminho encontrado, que pode ser um caminho "longo" e "tortuoso". Por exemplo, $s$ pode estar conectado a $t$ por uma aresta direta ($s-t$, distância 1), mas também por um caminho longo ($s-a-b-c-t$, distância 4). Se a DFS decidir explorar o vértice $a$ antes do vértice $t$, ela encontrará o caminho de distância 4 primeiro e o retornará (se o objetivo for apenas encontrar *um* caminho), sem nunca ter "visto" o caminho de distância 1.

A BFS, por outro lado, usa uma Fila. Ela explora primeiro os vizinhos de $s$ (distância 1), depois os vizinhos dos vizinhos (distância 2), e assim por diante. Ao explorar camada por camada, ela garante que o primeiro caminho encontrado para $t$ será, por definição, o caminho com o menor número de arestas.

</details>

---

Qual a relação entre o somatório dos graus dos vértices de um grafo e o número de arestas do grafo?

<details>
  <sumary>Resposta</sumary>

Em um grafo não-direcionado, o somatório dos graus de todos os vértices é exatamente o **dobro** do número de arestas.

Justificativa: Cada aresta $(u, v)$ conecta dois vértices. Ao somar os graus, a aresta é contada duas vezes: uma vez para o grau do vértice $u$ e outra vez para o grau do vértice $v$.

Fórmula: $\sum_{v \in V} \text{grau}(v) = 2 \cdot |E|$

</details>

---

Por que a busca em largura, quando realizada em um grafo não direcionado e sem pesos, é capaz de determinar o menor caminho entre um vértice inicial e todos os demais? Por que isso não é verdade para grafos direcionados com pesos?

<details>
  <sumary>Resposta</sumary>

a) Por que a BFS funciona para grafos não-direcionados e sem pesos:

A Busca em Largura (BFS) explora o grafo em "camadas" a partir do vértice inicial, usando uma Fila (Queue) para garantir essa ordem.
1.  Primeiro, visita todos os vizinhos diretos (distância 1).
2.  Depois, visita todos os vizinhos dos vizinhos (distância 2).
3.  E assim por diante.

Como a BFS avança camada por camada, ela garante que o primeiro caminho encontrado para qualquer vértice $v$ será, por definição, o caminho com o **menor número de arestas**. Em um grafo sem pesos, o "menor caminho" é exatamente aquele com o menor número de arestas.

b) Por que não funciona para grafos direcionados com pesos:

Em um grafo com pesos, o "menor caminho" não é o que tem menos arestas, mas o que tem a **menor soma de pesos**. A BFS ignora completamente os pesos das arestas.

Um caminho $s \rightarrow a$ pode ter 1 aresta com peso 10, enquanto um caminho $s \rightarrow b \rightarrow c \rightarrow a$ pode ter 3 arestas com peso 1 cada (custo total 3). A BFS, otimizando pelo número de arestas, encontraria $s \rightarrow a$ (distância 1) e o consideraria incorretamente "mais curto" que $s \rightarrow b \rightarrow c \rightarrow a$ (distância 3), falhando em encontrar o caminho de menor custo (custo 3). Para isso, é necessário um algoritmo como o de Dijkstra, que usa os pesos para priorizar a exploração.

</details>

---

Prove que um grafo não-direcionado com n vértices contém no máximo $n(n-1)/2$ arestas.

<details>
  <sumary>Resposta</sumary>

Um grafo atinge seu número máximo de arestas quando é um **grafo cheio (completo)**, ou seja, todo vértice está conectado a todo outro vértice (assumindo que não há auto-loops ou arestas paralelas).

A prova pode ser feita por uma contagem combinatória:

1.  O problema de encontrar o número máximo de arestas é equivalente a encontrar o número de maneiras de **escolher 2 vértices distintos** de um conjunto de $n$ vértices.
2.  A ordem da escolha não importa (uma aresta de $u$ para $v$ é a mesma que de $v$ para $u$).
3.  Este é um problema clássico de **Combinação de $n$ elementos, tomados 2 a 2**, ou $\binom{n}{2}$.
4.  A fórmula da combinação é:
    $\binom{n}{2} = \frac{n!}{2!(n-2)!}$
5.  Expandindo o fatorial:
    $\binom{n}{2} = \frac{n \cdot (n-1) \cdot (n-2)!}{2 \cdot 1 \cdot (n-2)!}$
6.  Simplificando (cortando o $(n-2)!$):
    $\binom{n}{2} = \frac{n(n-1)}{2}$

Portanto, um grafo não-direcionado com $n$ vértices contém no máximo $\frac{n(n-1)}{2}$ arestas.

</details>

---

Considere um grafo G não-direcionado, conexo e acíclico com n vértices. Quantas arestas G possui?
<details>
  <sumary>Resposta</sumary>

Um grafo não-direcionado, conexo e acíclico é a definição de uma **Árvore (Tree)**.

Por definição, uma árvore com $n$ vértices possui exatamente **$n-1$** arestas.

</details>


---

Qual problema o algoritmo de Dijkstra resolve? Qual é a principal restrição para que o algoritmo funcione corretamente?

<details>
  <sumary>Resposta</sumary>

O algoritmo de Dijkstra resolve o problema do **caminho mais curto de fonte única** (*single-source shortest path*). Dado um vértice de origem $s$ em um grafo ponderado, ele encontra o caminho de menor custo (menor soma de pesos das arestas) de $s$ para todos os outros vértices do grafo.

  * **Principal Restrição:** O algoritmo de Dijkstra **não funciona corretamente se o grafo contiver arestas com pesos negativos**.

</details>

---

O que significa "relaxar" uma aresta `(u, v)` no contexto do algoritmo de Dijkstra? Descreva a operação de relaxamento.

<details>
  <sumary>Resposta</sumary>

"Relaxar" uma aresta $(u, v)$ com peso $w$ é a operação central do algoritmo. Significa verificar se o caminho para $v$ pode ser melhorado (encurtado) passando por $u$.

Seja $dist[x]$ a distância mais curta conhecida da fonte $s$ até um vértice $x$. A operação de relaxamento para a aresta $(u, v)$ é:

```
se dist[u] + peso(u, v) < dist[v]:
    dist[v] = dist[u] + peso(u, v)
    // Atualiza v na fila de prioridade
```

Isso significa: "Se a distância conhecida até $u$, somada ao peso da aresta de $u$ para $v$, for menor do que a distância que conhecemos atualmente para $v$, então encontramos um novo caminho mais curto para $v$ (passando por $u$)".

</details>

---

Por que o algoritmo de Dijkstra falha se o grafo (direcionado ou não) contiver arestas de peso negativo?

<details>
  <sumary>Resposta</sumary>

A falha ocorre porque a premissa "gulosa" (greedy) de Dijkstra é violada. Dijkstra funciona assumindo que, uma vez que um vértice $u$ é extraído da fila de prioridade (ou seja, ele tem a menor distância $dist[u]$ entre todos os não visitados), essa distância é **final e absoluta**. O algoritmo assume que nunca encontrará um caminho mais curto para $u$ no futuro.

Se houver arestas negativas, essa premissa é quebrada. Um vértice $u$ pode ser "finalizado" com $dist[u]=5$. No entanto, mais tarde no algoritmo, podemos encontrar um caminho para outro vértice $v$ ($dist[v]=10$) que tem uma aresta para $u$ com peso $-6$. O caminho $s \rightarrow v \rightarrow u$ teria um custo total de $10 + (-6) = 4$, que é menor que 5. Como Dijkstra já finalizou $u$, ele não revisitará $u$ para corrigir essa distância, resultando em uma resposta incorreta.

</details>

---

Simule a execução do algoritmo de Dijkstra para o grafo direcionado com as arestas abaixo (o grafo contém 8 vértices), iniciando pelo vértice 1:

```
origem -> destino (peso)

0 -> 1 (5) 
0 -> 7 (8) 
0 -> 4 (9) 
1 -> 7 (4) 
1 -> 3 (15) 
1 -> 2 (12) 
2 -> 3 (3) 
2 -> 6 (11) 
3 -> 6 (9)
4 -> 7 (5)
4 -> 5 (4)
4 -> 6 (20)
5 -> 2 (1)
5 -> 6 (13)
7 -> 5 (6)
7 -> 2 (7)
```

<details>
  <sumary>Resposta</sumary>

Para a simulação, usaremos `dist[]` para armazenar as distâncias (iniciando com $\infty$ para todos, exceto `dist[1] = 0`) e uma Fila de Prioridade (FP) para selecionar o próximo vértice de menor custo.

| Passo | Vértice Extraído (V, Custo) | Estado de `dist[]` {0, 1, 2, 3, 4, 5, 6, 7} | Fila de Prioridade (FP) | Ações (Relaxamento de Arestas) |
| :--- | :--- | :--- | :--- | :--- |
| **0** | - (Início) | {$\infty$, 0, $\infty$, $\infty$, $\infty$, $\infty$, $\infty$, $\infty$} | {(1, 0)} | |
| **1** | **(1, 0)** | {$\infty$, 0, 12, 15, $\infty$, $\infty$, $\infty$, 4} | {(7, 4), (2, 12), (3, 15)} | Relaxa 1->7 (0+4=4). Relaxa 1->3 (0+15=15). Relaxa 1->2 (0+12=12). |
| **2** | **(7, 4)** | {$\infty$, 0, 11, 15, $\infty$, 10, $\infty$, 4} | {(5, 10), (2, 11), (3, 15)} | Relaxa 7->5 (4+6=10). Relaxa 7->2 (4+7=11, melhor que 12). |
| **3** | **(5, 10)** | {$\infty$, 0, 11, 15, $\infty$, 10, 23, 4} | {(2, 11), (3, 15), (6, 23)} | Relaxa 5->2 (10+1=11, igual a 11, sem mudança). Relaxa 5->6 (10+13=23). |
| **4** | **(2, 11)** | {$\infty$, 0, 11, 14, $\infty$, 10, 22, 4} | {(3, 14), (6, 22)} | Relaxa 2->3 (11+3=14, melhor que 15). Relaxa 2->6 (11+11=22, melhor que 23). |
| **5** | **(3, 14)** | {$\infty$, 0, 11, 14, $\infty$, 10, 22, 4} | {(6, 22)} | Relaxa 3->6 (14+9=23, pior que 22, sem mudança). |
| **6** | **(6, 22)** | {$\infty$, 0, 11, 14, $\infty$, 10, 22, 4} | { } | Vértice 6 não tem arestas de saída. |

**Resultado Final (Distâncias a partir do vértice 1):**
* **Vértice 0:** $\infty$ (Inalcançável)
* **Vértice 1:** 0 (Origem)
* **Vértice 2:** 11 (Caminho: 1 $\rightarrow$ 7 $\rightarrow$ 2)
* **Vértice 3:** 14 (Caminho: 1 $\rightarrow$ 7 $\rightarrow$ 2 $\rightarrow$ 3)
* **Vértice 4:** $\infty$ (Inalcançável)
* **Vértice 5:** 10 (Caminho: 1 $\rightarrow$ 7 $\rightarrow$ 5)
* **Vértice 6:** 22 (Caminho: 1 $\rightarrow$ 7 $\rightarrow$ 2 $\rightarrow$ 6)
* **Vértice 7:** 4 (Caminho: 1 $\rightarrow$ 7)

</details>

---

Para o mesmo grafo da questão anterior, determine:

a) Sua representação na forma de matriz de adjacências;
b) Sua representação na forma de lista de adjacências.

<details>
  <sumary>Resposta</sumary>

a. Sua representação na forma de matriz de adjacências:

(Usando $\infty$ para ausência de aresta e 0 para a diagonal principal)

| | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 |
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
| **0** | 0 | 5 | $\infty$ | $\infty$ | 9 | $\infty$ | $\infty$ | 8 |
| **1** | $\infty$ | 0 | 12 | 15 | $\infty$ | $\infty$ | $\infty$ | 4 |
| **2** | $\infty$ | $\infty$ | 0 | 3 | $\infty$ | $\infty$ | 11 | $\infty$ |
| **3** | $\infty$ | $\infty$ | $\infty$ | 0 | $\infty$ | $\infty$ | 9 | $\infty$ |
| **4** | $\infty$ | $\infty$ | $\infty$ | $\infty$ | 0 | 4 | 20 | 5 |
| **5** | $\infty$ | $\infty$ | 1 | $\infty$ | $\infty$ | 0 | 13 | $\infty$ |
| **6** | $\infty$ | $\infty$ | $\infty$ | $\infty$ | $\infty$ | $\infty$ | 0 | $\infty$ |
| **7** | $\infty$ | $\infty$ | 7 | $\infty$ | $\infty$ | 6 | $\infty$ | 0 |

b. Sua representação na forma de lista de adjacências:

(Formato: `Vértice -> (Vizinho, Peso)`)

* **0** -> (1, 5), (7, 8), (4, 9)
* **1** -> (7, 4), (3, 15), (2, 12)
* **2** -> (3, 3), (6, 11)
* **3** -> (6, 9)
* **4** -> (7, 5), (5, 4), (6, 20)
* **5** -> (2, 1), (6, 13)
* **6** -> ( )
* **7** -> (5, 6), (2, 7)

</details>

