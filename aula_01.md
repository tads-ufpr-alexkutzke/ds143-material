# Aula 01 - Estudo de caso - Union-Find

## O problema da Conectividade Dinâmica

**Conectividade dinâmica**: a entrada é uma sequência de pares de inteiros,
onde cada inteiro representa um objeto de algum tipo e nós interpretamos
o par `p q` como "`p` está conectado a `q`". Assumimos que "estar conectado
a" é uma relação de equivalência:

- *Simétrica*: se `p` está conectado a `q`, então `q` está conectado a `p`;
- *Transitiva*: se `p` está conectado a `q` e `q` está conectado a `r`, então `p` está conectado a `r`;
- *Reflexiva*: `p` está conectado a `p`.

Uma relação de equivalência particiona objetos em classes de equivalência ou
componente conexas (ou componentes conectadas ou, apenas, componentes).
Nesse caso, dois itens estão em uma mesma componente se, e somente se, eles estão
conectados.

Podemos identificar tanto itens quanto componente por números inteiros entre
0 e N-1.

Inicialmente, exitem N componentes independentes (nenhuma conexão), com cada
item em sua própria componente. O identificador de uma componente é um dos
itens que a compõem. Dois itens tem o mesmo identificador de componente se
e somente se eles fazem parte de uma mesma componente.

Nosso objetivo é escrever um programa para filtrar pares externos em uma sequência:
quando o programa ler um par `p q` da entrada padrão, ele deve escrever o próprio par
lido na saída padrão apenas se os pares lidos até o momento não implicam que `p` está
conectado a `q`. Se os pares lidos até o momento implicam que `p` está conectado a `q`,
então o programa não deve imprimir nada e proceder para o próximo par.

Abaixo está a API Union-Find. Ela encapsula as operações básicas que precisamos:

```c
   typedef struct {
      int * id;  // array para armazenar as compoentes de cada elemento;
      int n;     // número de elementos
      int count; // contagem de componentes atual;
   } UF

   UF *  init_UF(int n)                      // inicializa N itens com nomes de inteiros (0 até N-1)
   void  union_UF(UF * uf, int p, int q)     // adiciona conexão entre p e q
    int  find_UF(UF * uf, int p)             // identifica a componente de p (0 até N-1)
    int  connected_UF(UF * uf, int p, int q) // retorna true se p e q estão na mesma componente
    int  count_UF(UF * uf)                   // retorna o número de componentes
```

O identificador de uma componente pode mudar apenas a partir
de uma chamada ao método `union`. Este identificador **não** pode
ser alterado pelos métodos `find`, `connected` ou `count`.

![Exemplo de execução](01/dynamic-connectivity-tiny.png)

Para testar a API acima, o arquivo [uf.c](01/uf.c) resolve parcialmente
o problema da conectividade dinâmica. O programa é encerrado quando
o usuário entra com ao menos um número negativo.

Dados para testar o programa também estão disponíveis: 
- O arquivo [tinyUF.txt](01/tinyUF.txt) contém 11 conexões indicadas na figura acima;
- O arquivo [mediumUF.txt](01/mediumUF.txt) contém 900 conexões e;
- O arquivo [largeUF.txt](01/largeUF.txt) contém milhões de conexões; =O

Considere um vetor como a estrutura de dados básica para armazenar as informações
sobre cada item. Por exemplo `id[0]` armazena a qual componente o item 0 pertence.
Existem diferentes implementações para os métodos `union` e `find`. Qual seria a mais
eficiente que você consegue pensar?

Tente resolver essa questão sem procurar por alternativas na internet. Se desafie!
Essa tarefa está no moodle.

## Referências (consulte apenas depois de fazer sua própria solução)
* [Sedgewick - Union-find (em inglês)](https://algs4.cs.princeton.edu/15uf/);
* [Sedgewick - Union-find Slides (em inglês)](https://algs4.cs.princeton.edu/lectures/keynote/15UnionFind.pdf);
