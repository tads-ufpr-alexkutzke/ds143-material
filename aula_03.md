# Atividade Avaliativa 1: Union-Find em Go

**Atividade assíncrona, no lugar da aula presencial de 19/08/2026.** Estarei em
viagem nesta data, disponível por e-mail institucional e Microsoft Teams das 19h
às 21h para dúvidas pontuais.

| | |
|---|---|
| **Peso** | 5% da nota final |
| **Formato** | individual |
| **Prazo** | 25/08/2026, 23h59, na UFPR Virtual |

O conteúdo é o das seções 9 a 15 da [aula de Análise de Algoritmos e
Conectividade Dinâmica](aula_02.md).

## Objetivos

Ao final desta atividade você deve ser capaz de:

1. Implementar a compressão de caminho sobre o Weighted Quick-Union, dentro de
   um pacote Go organizado atrás de uma interface;
2. Escrever um teste com o pacote `testing` e dizer que defeito ele detecta;
3. Executar o teste de duplicação e confrontar as razões medidas com a ordem de
   crescimento prevista pela análise.

## Antes de começar

Confirme o Go instalado com `go version` (a partir da 1.21). Baixe o esqueleto
do projeto na pasta [03/esqueleto](03/esqueleto), também disponível como `.zip`
na UFPR Virtual. O pacote `unionfind` já vem com `QuickFind`, `QuickUnion` e
`WeightedQuickUnion` implementados: leia os três antes de escrever qualquer
coisa, porque eles definem como o restante deve ser escrito.

O esqueleto compila desde o início e os testes falham. É o estado esperado:

```bash
cd esqueleto
go vet ./...     # sem erros
go test ./...    # FAIL, com panic em NewWeightedPathCompression
```

## Parte 1: compressão de caminho

Complete
[unionfind/weightedpathcompression.go](03/esqueleto/unionfind/weightedpathcompression.go).
O `WeightedPathCompression` é o Weighted Quick-Union com compressão de caminho:
a cada subida até a raiz, o `Find` aproxima da raiz os nós visitados. A seção 15
da aula anterior mostra o laço do `Find`, e o restante sai do
`weightedquickunion.go`, no mesmo pacote.

A linha `var _ UF = (*WeightedPathCompression)(nil)` faz o compilador recusar o
programa enquanto faltar método ou a assinatura estiver errada. Não a remova, e
não altere `uf.go` nem `uf_test.go`, que são os arquivos usados na correção. Ao
final, `go vet ./...`, `go test ./...` e `gofmt -l .` devem passar sem problemas.

## Parte 2: um teste seu

Os testes fornecidos verificam o contrato da interface. Acrescente um teste seu
em um arquivo novo, `unionfind/meus_testes_test.go`: um **oráculo cruzado**, que
processa a mesma sequência de pares aleatórios, com semente fixa, nas quatro
implementações e verifica que todas concordam em `Count` e em `Connected`. Assim,
caso sua implementação tenha um erro, ele aparecerá contra os resultados das outras três.

O teste tem um comentário de uma linha dizendo que defeito ele detecta. Um
teste que passa em qualquer implementação, correta ou não, não vale nota.

## Parte 3: medição empírica

No programa [cmd/benchmark/main.go](03/esqueleto/cmd/benchmark/main.go), só a
função `medir` está por escrever: o resto, incluindo a geração dos pares e a
impressão da tabela, já está pronto. Os pares são gerados em memória com semente
fixa, então todas as implementações recebem a mesma sequência e o tempo medido
não inclui leitura de arquivo.

Meça o binário compilado, e não `go run`, que recompila a cada execução:

```bash
go build -o benchmark ./cmd/benchmark
./benchmark | tee resultados.txt
```

A tabela sai com as quatro implementações, `n` em 10.000, 20.000, 40.000, 80.000
e 160.000, e `2n` pares em cada tamanho. A execução leva cerca de um minuto,
quase todo gasto pelo Quick-Find e pelo Quick-Union nos dois maiores tamanhos.

## Parte 4: relatório

Um `relatorio.md` de no máximo uma página, com o ambiente de medição
(processador, memória, sistema operacional, versão do Go), a tabela em Markdown
e duas respostas de até cinco linhas cada para as seguintes perguntas:

1. As razões medidas confirmam a ordem de crescimento prevista na seção 14 da
   aula anterior para cada implementação? Cite os números que você obteve.
2. O `Weighted` e o `Weighted+PC` marcam tempos praticamente iguais nesses
   tamanhos. Isso contraria a proposição da seção 15? Diga o que precisaria
   mudar na medição para a diferença entre os dois aparecer.

O relatório termina com uma seção `## Uso de IA`, informando se você usou
ferramentas de IA generativa e em que partes (não influencia na nota). 

## Entrega

Um `.zip` na UFPR Virtual, nomeado `ds143_atividade1_<seu_nome>.zip`, com o
projeto Go completo (`go.mod`, o pacote `unionfind` incluindo
`meus_testes_test.go`, e `cmd/benchmark/main.go`), o `resultados.txt` e o
`relatorio.md`. Sem binários compilados. Rode `go test ./...` uma última vez na
pasta que será enviada.

## Critérios de avaliação

| Critério | Valor |
|---|---|
| Compila, `go vet` limpo, testes fornecidos e teste próprio passam | 40 |
| `WeightedPathCompression` correto, com a compressão encurtando o caminho de fato | 30 |
| Tabela de medição com tempos e razões                                            | 20 |
| Relatório com as duas respostas e declaração de uso de IA                        | 10 |
| **Total** | **100** |

Entrega que não compila perde o primeiro item inteiro; os demais são avaliados
pelo que for verificável no código.

## Problemas comuns

**`cannot use ... as UF value: missing method Union`**: método declarado com
receptor por valor. Use receptor por ponteiro em todos.

**Teste de arquivo falhando com `no such file or directory`**: rode `go test
./...` a partir da raiz do projeto, sem mover a pasta `testdata`.
