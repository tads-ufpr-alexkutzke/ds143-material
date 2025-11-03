**1)** Qual é o problema da busca de substrings? Defina o que é o "padrão" (pattern) e o "texto" (text).

<details>
  <sumary>Resposta</sumary>
O problema da busca de substrings consiste em encontrar uma ou todas as ocorrências de uma string chamada "padrão" (P) dentro de uma string maior chamada "texto" (T). Por exemplo, encontrar o padrão "ANA" no texto "BANANA".
</details>

**2)** Descreva o funcionamento e a complexidade de pior caso do algoritmo de busca de substrings de força bruta. Dê um exemplo de um padrão e um texto que causem o pior caso.

<details>
  <sumary>Resposta</sumary>
O algoritmo de força bruta funciona deslizando o padrão sobre o texto, caractere por caractere. Para cada posição $i$ no texto, ele tenta "casar" o padrão $P[0..m-1]$ com o texto $T[i..i+m-1]$.

  * **Funcionamento:** Um laço externo percorre as posições $i$ de $0$ a $n-m$ (onde $n$ é o tamanho do texto e $m$ o do padrão). Um laço interno compara $P[j]$ com $T[i+j]$. Se uma diferença for > encontrada, o laço interno é quebrado e o laço externo avança (incrementa $i$).
  * **Pior Caso:** A complexidade de pior caso é $O(n \cdot m)$.
  * **Exemplo:** Isso ocorre quando, para cada posição do texto, quase todo o padrão casa antes de falhar no último caractere.
      * Texto (T): "AAAAAAAAAAAAAAAAAB"
      * Padrão (P): "AAAAB"
</details>

**3)** Como o algoritmo KMP (Knuth-Morris-Pratt) consegue realizar uma busca mais rápida? Qual o principal artifício utilizado no algoritmo para isso?

<details>
  <sumary>Resposta</sumary>
O algoritmo KMP consegue uma busca mais rápida porque ele elimina re-comparações desnecessárias no texto.

Ao contrário do algoritmo de força bruta, que após uma falha de correspondência retrocede o ponteiro do texto e avança o padrão em apenas uma posição, o KMP nunca retrocede o ponteiro do texto.

O principal artifício para conseguir isso é o pré-processamento do padrão para construir uma estrutura de dados (um array lps ou um Autômato Finito Determinístico - DFA) que armazena informações sobre as sobreposições internas do próprio padrão. 
</details>

**4)** Qual é a diferença fundamental entre compressão *lossless* (sem perdas) e *lossy* (com perdas)? Em que categoria o algoritmo de Huffman se encaixa e por quê?

<details>
  <sumary>Resposta</sumary>
  * **Compressão Lossless (Sem Perdas):** Reduz o tamanho do arquivo de forma que o arquivo original possa ser perfeitamente reconstruído a partir do arquivo comprimido. Nenhum dado é perdido. É usada para textos, código-fonte e dados científicos.
  * **Compressão Lossy (Com Perdas):** Reduz o tamanho do arquivo descartando permanentemente informações consideradas "menos importantes". O arquivo original não pode ser recuperado. É usada para mídia, como imagens (JPEG), áudio (MP3) e vídeo.
  * **Huffman:** O algoritmo de Huffman é **lossless**. Ele funciona atribuindo códigos binários de tamanho variável aos caracteres (mais frequentes recebem códigos mais curtos), mas a correspondência é exata, permitindo a descompressão perfeita.
</details>

**5)** O que é uma "codificação livre de prefixo" (ou *prefix-free code*) e por que essa propriedade é essencial para o funcionamento correto da descompressão no algoritmo de Huffman?

<details>
  <sumary>Resposta</sumary>
Uma codificação de prefixo é um conjunto de códigos (binários, no caso de Huffman) onde nenhum código é o prefixo de outro código.

  * **Exemplo:** Se "A" é `01`, "B" não pode ser `010` (pois `01` é prefixo). Um conjunto válido seria: A=`0`, B=`10`, C=`11`.
  * **Importância:** Essa propriedade é essencial para a descompressão, pois permite que o descompressor leia o fluxo de bits comprimido de forma inequívoca. Ao ler os bits, assim que uma sequência corresponde a um código na árvore de Huffman, o descompressor sabe que aquele símbolo foi encontrado. Não há ambiguidade; ele não precisa "olhar para frente" para ver se bits adicionais poderiam formar um código diferente (e mais longo).
</details>


**6)** Dada a string "A A A B B C D A", construa a árvore de Huffman correspondente. Em seguida, mostre o código binário para cada caractere (A, B, C, D) e o resultado da compressão da string inteira.

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

**7)** Compare a Busca em Largura (BFS) e a Busca em Profundidade (DFS). Qual estrutura de dados auxiliar cada uma utiliza e qual é uma aplicação comum para cada uma?

<details>
  <sumary>Resposta</sumary>
  * **BFS (Busca em Largura):** Explora o grafo "em camadas", visitando primeiro todos os vizinhos de um nó antes de se aprofundar.
      * **Estrutura de Dados:** Fila (Queue), para garantir a ordem FIFO (First-In, First-Out).
      * **Aplicação Comum:** Encontrar o caminho mais curto (em número de arestas) em um grafo não-ponderado.
  * **DFS (Busca em Profundidade):** Explora o grafo "indo o mais fundo possível" por um caminho antes de retroceder (backtracking).
      * **Estrutura de Dados:** Pilha (Stack) (ou recursão, que usa a pilha de chamadas).
      * **Aplicação Comum:** Detecção de ciclos, ordenação topológica (em grafos direcionados), encontrar componentes conexas.
</details>


**8)** Dado um grafo não-direcionado, como você pode usar BFS ou DFS para determinar o número de componentes conexas desse grafo? Descreva o algoritmo.

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


**8)** Em um grafo não-direcionado, a BFS encontra o caminho mais curto entre dois nós $s$ e $t$ (em número de arestas). Por que a DFS não garante encontrar o caminho mais curto?

<details>
  <sumary>Resposta</sumary>
A DFS não garante o caminho mais curto porque ela não explora o grafo em camadas de distância. A natureza da DFS é seguir um caminho até o fim (o mais profundo possível) antes de retroceder (backtracking).

A DFS pode encontrar um caminho válido entre $s$ e $t$, mas é provável que seja o primeiro caminho encontrado, que pode ser um caminho "longo" e "tortuoso". Por exemplo, $s$ pode estar conectado a $t$ por uma aresta direta ($s-t$, distância 1), mas também por um caminho longo ($s-a-b-c-t$, distância 4). Se a DFS decidir explorar o vértice $a$ antes do vértice $t$, ela encontrará o caminho de distância 4 primeiro e o retornará (se o objetivo for apenas encontrar *um* caminho), sem nunca ter "visto" o caminho de distância 1.

A BFS, por outro lado, usa uma Fila. Ela explora primeiro os vizinhos de $s$ (distância 1), depois os vizinhos dos vizinhos (distância 2), e assim por diante. Ao explorar camada por camada, ela garante que o primeiro caminho encontrado para $t$ será, por definição, o caminho com o menor número de arestas.
</details>


**9)** Qual problema o algoritmo de Dijkstra resolve? Qual é a principal restrição para que o algoritmo funcione corretamente?

<details>
  <sumary>Resposta</sumary>
O algoritmo de Dijkstra resolve o problema do **caminho mais curto de fonte única** (*single-source shortest path*). Dado um vértice de origem $s$ em um grafo ponderado, ele encontra o caminho de menor custo (menor soma de pesos das arestas) de $s$ para todos os outros vértices do grafo.

  * **Principal Restrição:** O algoritmo de Dijkstra **não funciona corretamente se o grafo contiver arestas com pesos negativos**.
</details>


**10)** O que significa "relaxar" uma aresta `(u, v)` no contexto do algoritmo de Dijkstra? Descreva a operação de relaxamento.

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

**11)** Por que o algoritmo de Dijkstra falha se o grafo (direcionado ou não) contiver arestas de peso negativo?

<details>
  <sumary>Resposta</sumary>
A falha ocorre porque a premissa "gulosa" (greedy) de Dijkstra é violada. Dijkstra funciona assumindo que, uma vez que um vértice $u$ é extraído da fila de prioridade (ou seja, ele tem a menor distância $dist[u]$ entre todos os não visitados), essa distância é **final e absoluta**. O algoritmo assume que nunca encontrará um caminho mais curto para $u$ no futuro.

Se houver arestas negativas, essa premissa é quebrada. Um vértice $u$ pode ser "finalizado" com $dist[u]=5$. No entanto, mais tarde no algoritmo, podemos encontrar um caminho para outro vértice $v$ ($dist[v]=10$) que tem uma aresta para $u$ com peso $-6$. O caminho $s \rightarrow v \rightarrow u$ teria um custo total de $10 + (-6) = 4$, que é menor que 5. Como Dijkstra já finalizou $u$, ele não revisitará $u$ para corrigir essa distância, resultando em uma resposta incorreta.
</details>
