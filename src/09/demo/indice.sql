-- Demonstração ao vivo: o índice de um banco de dados é uma árvore B.
--
-- Execute com:  sqlite3 aluno.db < indice.sql
-- (apague o aluno.db antes de repetir a demonstração)

.timer on

-- Dois milhões de linhas, sem índice nenhum além do rowid.
CREATE TABLE aluno(grr INTEGER, nome TEXT);
WITH RECURSIVE seq(n) AS (
  SELECT 1 UNION ALL SELECT n + 1 FROM seq WHERE n < 2000000
)
INSERT INTO aluno(grr, nome) SELECT n, 'aluno ' || n FROM seq;

-- Sem índice, o banco lê a tabela inteira: SCAN.
EXPLAIN QUERY PLAN SELECT nome FROM aluno WHERE grr = 1999999;
SELECT nome FROM aluno WHERE grr = 1999999;

-- O índice é uma árvore B construída sobre a coluna grr.
CREATE INDEX idx_grr ON aluno(grr);

-- Com índice, o banco desce a árvore: SEARCH USING INDEX.
EXPLAIN QUERY PLAN SELECT nome FROM aluno WHERE grr = 1999999;
SELECT nome FROM aluno WHERE grr = 1999999;

-- O tamanho do nó da árvore é o tamanho da página do arquivo, em bytes.
PRAGMA page_size;

-- Busca por faixa: o índice serve para intervalo, porque as chaves estão
-- em ordem dentro de cada nó e entre os nós.
EXPLAIN QUERY PLAN SELECT count(*) FROM aluno WHERE grr BETWEEN 1000 AND 2000;
