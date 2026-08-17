# Atividade Avaliativa 1: Union-Find em Go

**Atividade assíncrona, no lugar da aula presencial de 19/08/2026.** Estarei em
viagem nesta data, disponível por e-mail institucional e Microsoft Teams das 19h
às 21h para dúvidas pontuais.

| | |
|---|---|
| **Peso** | 5% da nota final |
| **Formato** | individual |
| **Prazo** | 26/08/2026, 23h59, na UFPR Virtual |
| **Esforço estimado** | 2h30 |

O conteúdo é o das seções 9 a 15 da [aula de Análise de Algoritmos e
Conectividade Dinâmica](aula_02.md).

## Objetivos

Ao final desta atividade você deve ser capaz de:

1. Organizar implementações alternativas de uma estrutura de dados em um pacote
   Go, atrás de uma interface comum;
2. Escrever testes com o pacote `testing` e dizer que defeito cada um detecta;
3. Implementar a compressão de caminho sobre o Weighted Quick-Union;
4. Executar o teste de duplicação e confrontar as razões medidas com a ordem de
   crescimento prevista pela análise.

## Antes de começar

Confirme o Go instalado com `go version` (a partir da 1.21). Se ainda não fez as
seções *Basics* e *Methods and interfaces* do [Tour de
Go](https://go.dev/tour/welcome/1), faça agora: a atividade gira em torno de uma
interface. Baixe o esqueleto do projeto na pasta [03/esqueleto](03/esqueleto),
também disponível como `.zip` na UFPR Virtual.

O esqueleto compila desde o início e os testes falham. É o estado esperado:

```bash
cd esqueleto
go vet ./...     # sem erros
go test ./...    # FAIL, com panic em NewQuickFind
```

## Parte 1: o pacote `unionfind`

[unionfind/uf.go](03/esqueleto/unionfind/uf.go) define a interface que as quatro
implementações precisam cumprir, com `Union`, `Find`, `Connected` e `Count`.
Cada implementação tem seu arquivo, com campos e métodos a completar:

| Arquivo | Tipo | Construtor |
|---|---|---|
| `quickfind.go` | `QuickFind` | `NewQuickFind(n int) *QuickFind` |
| `quickunion.go` | `QuickUnion` | `NewQuickUnion(n int) *QuickUnion` |
| `weightedquickunion.go` | `WeightedQuickUnion` | `NewWeightedQuickUnion(n int) *WeightedQuickUnion` |
| `weightedpathcompression.go` | `WeightedPathCompression` | `NewWeightedPathCompression(n int) *WeightedPathCompression` |

As três primeiras já existem no material, cada uma como programa independente em
[02/codes/unionfind](02/codes/unionfind): parta delas. Ao migrar, o tipo passa a
ter o nome da estratégia, porque os quatro convivem no mesmo pacote, e a leitura
da entrada sai do arquivo da estrutura. A estrutura não lê arquivo nem imprime.

A quarta implementação não tem código pronto: o `WeightedPathCompression` é o
Weighted Quick-Union com compressão de caminho, e a seção 15 da aula anterior
mostra apenas o laço do `Find`.

A linha `var _ UF = (*QuickFind)(nil)` faz o compilador recusar o programa
enquanto faltar método ou a assinatura estiver errada. Não a remova, e não
altere `uf.go` nem `uf_test.go`, que são os arquivos usados na correção. Ao
final, `go vet ./...`, `go test ./...` e `gofmt -l .` devem passar limpos.

## Parte 2: seus próprios testes

Os testes fornecidos verificam o contrato da interface. Acrescente três testes
seus em um arquivo novo, `unionfind/meus_testes_test.go`:

1. **Oráculo cruzado**: processe a mesma sequência de pares aleatórios, com
   semente fixa, nas quatro implementações, e verifique que todas concordam em
   `Count` e `Connected`. Um erro em uma delas aparece contra as outras três.
2. **Casos limite**: `n = 0`, `n = 1` e `Union(p, p)`. Decida o que é correto em
   cada caso e escreva o teste que o exige.
3. **Um teste à sua escolha**, cobrindo algo que os anteriores não cobrem.

Cada teste leva um comentário de uma linha dizendo que defeito ele detecta. Um
teste que passa em qualquer implementação, correta ou não, não vale nota.

## Parte 3: medição empírica

Complete [cmd/benchmark/main.go](03/esqueleto/cmd/benchmark/main.go). Os pares
são gerados em memória com semente fixa, então todas as implementações recebem a
mesma sequência e o tempo medido não inclui leitura de arquivo. Meça o binário
compilado, e não `go run`, que recompila a cada execução:

```bash
go build -o benchmark ./cmd/benchmark
./benchmark | tee resultados.txt
```

**Tabela 1**: as quatro implementações, com `n` em 10.000, 20.000, 40.000,
80.000 e 160.000, e `2n` pares. Informe o tempo e a razão em relação ao tamanho
anterior. A execução leva cerca de um minuto, quase todo gasto pelo Quick-Find e
pelo Quick-Union nos dois maiores tamanhos.

**Tabela 2**: só o `WeightedQuickUnion` e o `WeightedPathCompression`, com `n` em
1.000.000, 2.000.000 e 4.000.000. Nos tamanhos da Tabela 1 essas duas respondem
em poucos milissegundos e a medição vira ruído. O último tamanho ocupa cerca de
250 MB de memória: se a máquina não der conta, pare em 2.000.000 e diga isso no
relatório.

## Parte 4: relatório

Um `relatorio.md` de no máximo duas páginas, com o ambiente de medição
(processador, memória, sistema operacional, versão do Go), as duas tabelas em
Markdown e três respostas de até cinco linhas cada:

1. As razões da Tabela 1 confirmam a ordem de crescimento prevista na seção 14
   da aula anterior? Cite os números que você mediu.
2. A compressão de caminho melhorou o tempo? Compare a Tabela 2 com o custo
   $n + M \lg^{*} n$ da seção 15 e explique por que a diferença observada não é
   maior.
3. Usando apenas a Tabela 1 e a hipótese de duplicação, preveja o tempo do
   Quick-Find para `n = 1.000.000`, mostrando a conta. Não execute essa medição.

O relatório termina com uma seção `## Uso de IA`, informando se você usou
ferramentas de IA generativa e em que partes. **A declaração não altera a nota
aqui**: serve para criar o hábito antes do Trabalho Prático, onde ela é
irrevogável e define a forma de avaliação. Declaração ausente custa 0,5%.

## Entrega

Um `.zip` na UFPR Virtual, nomeado `ds143_atividade1_<seu_nome>.zip`, com o
projeto Go completo (`go.mod`, pacote `unionfind` incluindo
`meus_testes_test.go`, e `cmd/benchmark/main.go`), o `resultados.txt` e o
`relatorio.md`. Sem binários compilados. Rode `go test ./...` uma última vez na
pasta que será enviada.

## Critérios de avaliação

| Critério | Valor |
|---|---|
| Compila, `go vet` limpo, testes fornecidos passam e os três testes próprios passam | 2,0% |
| Implementação correta das quatro variantes, com a compressão de caminho funcionando | 1,5% |
| Medição empírica com as duas tabelas e as razões calculadas | 1,0% |
| Relatório fundamentado nos números medidos, com declaração de uso de IA | 0,5% |
| **Total** | **5,0%** |

Entrega que não compila perde o primeiro item inteiro; os demais são avaliados
pelo que for verificável no código.

## Problemas comuns

**`cannot use ... as UF value: missing method Union`**: método declarado com
receptor por valor. Use receptor por ponteiro em todos.

**Teste de arquivo falhando com `no such file or directory`**: rode `go test
./...` a partir da raiz do projeto, sem mover a pasta `testdata`.
