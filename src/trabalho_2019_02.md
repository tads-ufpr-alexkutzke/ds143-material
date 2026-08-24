# Especificação do Trabalho prático 2018/2 e 2019/2

O trabalho consiste na criação de um programa, em linguagem C, capaz de indexar palavras de um ou mais documentos de texto. Tal programa deverá ser nomeado como `indexer`.

## Arquivos para testes

* http://200.236.3.126:8999/ds141-texts/

## Funcionamento

```
NOME
    indexer - indexa palavras de documentos de texto

SINOPSE
    indexer --freq N ARQUIVO
    indexer --freq-word PALAVRA ARQUIVO
    indexer --search TERMO ARQUIVO [ARQUIVO ...]

DESCRIÇÃO
    O programa **indexer** realiza uma contagem de palavras em documentos de 
    texto. A partir dessa contagem, ele ainda permite uma busca pelo número de 
    ocorrências de uma palavra específica em um documento, ou a seleção de 
    documentos relevantes para um dado termo de busca.
    Quando executado com a opção --freq, o programa **indexer** irá exibir o 
    número de ocorrências das N palavras mais frequentes no documento passado 
    como parâmetro, em ordem decrescente de ocorrências.
    Quando executado com a opção --freq-word, o programa **indexer** exibe a 
    contagem de uma palavra específica no documento passado como parâmetro.
    Por fim, quando executado com a opção --search, o programa **indexer** 
    apresenta uma listagem dos documentos mais relevantes para um dado termo de 
    busca. Para tanto, o programa utiliza o cálculo TF-IDF (Term 
    Frequency-Inverse Document Frequency).

OPÇÕES
  --freq N ARQUIVO
    Exibe o número de ocorrência das N palavras que mais aparecem em ARQUIVO, em
    ordem decrescente de ocorrência.
  --freq-word PALAVRA ARQUIVO
    Exibe o número de ocorrências de PALAVRA em ARQUIVO. 
  --search TERMO ARQUIVO [ARQUIVO ...]
    Exibe uma listagem dos ARQUIVOS mais relevantes encontrados pela busca por 
    TERMO. A listagem é apresentada em ordem descrescente de relevância. 
    TERMO pode conter mais de uma palavra. Neste caso, deve ser indicado entre 
    àspas.
```

## Requisitos

O trabalho deve ser realizado em grupos de até 3 alunos.

O programa deve conter a implementação de uma estrutura de dados eficiente capaz
de atender às funcionalidades descritas acima, mesmo quando executado com 
arquivos grandes (> 1GB). Tal estrutura deve ser implementada pelos próprios 
alunos e deverá ser explicada no dia da defesa do trabalho.

O programa deve ser capaz de ser compilado e executado em um ambiente Linux. 
Isso não deve gerar maiores problemas para aqueles que programam em Windows ou 
outro S.O.. Basta ter cuidado para não utilizar bibliotecas ou outras diretrizes 
específicas desses S.O.'s.

Além disso, o programa deverá ser acompanhado de um Makefile (sugestão sobre o 
tema https://pt.wikibooks.org/wiki/Programar_em_C/Makefiles ).

A entrega será realizada por meio da criação de um repositório no gitlab 
(enviar endereço do repositório por tarefa no moodle). A data da entrega 
será definida no moodle da disciplina.

Basicamente, vou clonar o repositório recebido, executar o 
comando `make`e a partir daí, um programa chamado `indexer` já deve estar 
disponível para que eu inicie testes em minha máquina.

Caso eu não seja capaz de compilar e executar o trabalho na minha máquina, 
a nota máxima será 50.

O trabalho que apresentar o melhor desempenho (medido em segundos) ganhará um 
ponto extra. O trabalho que apresentar o melhor desempenho entre as duas turmas 
(tarde e noite) ganhará outro ponto extra.


## Cálculo de Relevância

O cálculo para determinar a relevância de um texto para um dado termo de busca deve ser realizado utilizando a técnica **Term Frequency-Inverse Document Frequency** (https://pt.wikipedia.org/wiki/Tf–idf e https://en.wikipedia.org/wiki/Tf–idf#Term_frequency).

**Term Frequency (TF)** ou Frequência de um termo `t`, deve ser calculada como o número de vezes que `t` aparece no documento, dividido pelo número total de palavras no documento:

```
TF(t,d) = (Número de vezes que t aparece em d) / (Total de palavras em d)
```

**Inverse Document Frequency (IDF)** ou Frequência inversa de documento para um conjunto de documentos `D` é o logaritmo da divisão do número total de documentos pelo número de documentos que contém o termo `t` em questão:

```
IDF(t,D) = log[ (Número de Documentos) / (Número de documentos em que t está presente) ]
```

Logaritmo de base 10.

Por fim, TFIDF é a multiplicação das duas medidas. Ela indica qual a relevância de um termo `t` para um documento `d` dentre uma coleção de documentos `D`:

```
TFIDF(t,d,D) = TF(t,d) * IDF(t,D)
```

Para termos com mais de uma palavra, deve ser calculado a média entre o TF-IDF de cada palavra.

Mais uma referência: http://www.tfidf.com
