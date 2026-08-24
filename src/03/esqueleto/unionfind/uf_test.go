package unionfind

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"
)

// implementacoes reúne os construtores das quatro variantes. Todo teste deste
// arquivo roda sobre as quatro, sem saber qual é qual: o que está sendo
// verificado é o contrato da interface UF, que independe da estratégia usada.
var implementacoes = map[string]func(int) UF{
	"QuickFind":               func(n int) UF { return NewQuickFind(n) },
	"QuickUnion":              func(n int) UF { return NewQuickUnion(n) },
	"WeightedQuickUnion":      func(n int) UF { return NewWeightedQuickUnion(n) },
	"WeightedPathCompression": func(n int) UF { return NewWeightedPathCompression(n) },
}

func TestEstadoInicial(t *testing.T) {
	for nome, novo := range implementacoes {
		t.Run(nome, func(t *testing.T) {
			uf := novo(10)
			if uf.Count() != 10 {
				t.Errorf("Count() = %d, esperado 10", uf.Count())
			}
			for i := 0; i < 10; i++ {
				if uf.Find(i) != i {
					t.Errorf("Find(%d) = %d, esperado %d", i, uf.Find(i), i)
				}
			}
			if uf.Connected(3, 7) {
				t.Error("Connected(3, 7) = true em estrutura recém-criada")
			}
			if !uf.Connected(3, 3) {
				t.Error("Connected(3, 3) = false; todo objeto está conectado a si mesmo")
			}
		})
	}
}

func TestUniaoConecta(t *testing.T) {
	for nome, novo := range implementacoes {
		t.Run(nome, func(t *testing.T) {
			uf := novo(10)
			uf.Union(4, 3)
			if !uf.Connected(4, 3) {
				t.Error("Connected(4, 3) = false depois de Union(4, 3)")
			}
			if !uf.Connected(3, 4) {
				t.Error("Connected(3, 4) = false; a relação é simétrica")
			}
			if uf.Count() != 9 {
				t.Errorf("Count() = %d depois de uma união, esperado 9", uf.Count())
			}
			if uf.Connected(4, 5) {
				t.Error("Connected(4, 5) = true sem união entre eles")
			}
		})
	}
}

func TestUniaoRepetidaNaoAlteraContagem(t *testing.T) {
	for nome, novo := range implementacoes {
		t.Run(nome, func(t *testing.T) {
			uf := novo(10)
			uf.Union(4, 3)
			uf.Union(4, 3)
			uf.Union(3, 4)
			if uf.Count() != 9 {
				t.Errorf("Count() = %d depois de unir o mesmo par três vezes, esperado 9", uf.Count())
			}
		})
	}
}

func TestTransitividade(t *testing.T) {
	for nome, novo := range implementacoes {
		t.Run(nome, func(t *testing.T) {
			uf := novo(10)
			uf.Union(0, 1)
			uf.Union(2, 3)
			if uf.Connected(1, 2) {
				t.Error("Connected(1, 2) = true antes de unir as duas componentes")
			}
			uf.Union(1, 2)
			if !uf.Connected(0, 3) {
				t.Error("Connected(0, 3) = false; a relação é transitiva")
			}
			if uf.Count() != 7 {
				t.Errorf("Count() = %d, esperado 7", uf.Count())
			}
		})
	}
}

// TestFindConcordaComConnected verifica a propriedade que liga os dois
// métodos: dois objetos estão conectados se, e somente se, Find devolve o
// mesmo identificador para ambos.
func TestFindConcordaComConnected(t *testing.T) {
	for nome, novo := range implementacoes {
		t.Run(nome, func(t *testing.T) {
			uf := novo(8)
			uf.Union(0, 2)
			uf.Union(2, 4)
			uf.Union(1, 3)
			for p := 0; p < 8; p++ {
				for q := 0; q < 8; q++ {
					mesmoID := uf.Find(p) == uf.Find(q)
					if mesmoID != uf.Connected(p, q) {
						t.Errorf("p=%d q=%d: Find igual = %v, Connected = %v",
							p, q, mesmoID, uf.Connected(p, q))
					}
				}
			}
		})
	}
}

// TestCadeia une os objetos em sequência, caso em que o Quick-Union sem
// balanceamento produz uma árvore de altura máxima.
func TestCadeia(t *testing.T) {
	const n = 200
	for nome, novo := range implementacoes {
		t.Run(nome, func(t *testing.T) {
			uf := novo(n)
			for i := 0; i < n-1; i++ {
				uf.Union(i, i+1)
			}
			if uf.Count() != 1 {
				t.Errorf("Count() = %d depois de unir todos em cadeia, esperado 1", uf.Count())
			}
			if !uf.Connected(0, n-1) {
				t.Errorf("Connected(0, %d) = false", n-1)
			}
		})
	}
}

func TestArquivos(t *testing.T) {
	casos := []struct {
		arquivo     string
		componentes int
	}{
		{"testdata/tinyUF.txt", 2},
		{"testdata/mediumUF.txt", 3},
	}

	for _, caso := range casos {
		for nome, novo := range implementacoes {
			t.Run(caso.arquivo+"/"+nome, func(t *testing.T) {
				n, pares := lerEntrada(t, caso.arquivo)
				uf := novo(n)
				for _, par := range pares {
					uf.Union(par[0], par[1])
				}
				if uf.Count() != caso.componentes {
					t.Errorf("Count() = %d, esperado %d", uf.Count(), caso.componentes)
				}
			})
		}
	}
}

// lerEntrada carrega um arquivo no formato usado pela disciplina: a primeira
// linha traz o número de objetos e cada linha seguinte traz um par. A leitura
// para no fim do arquivo ou no primeiro par com número negativo.
func lerEntrada(t *testing.T, caminho string) (int, [][2]int) {
	t.Helper()

	arquivo, err := os.Open(caminho)
	if err != nil {
		t.Fatalf("abrindo %s: %v", caminho, err)
	}
	defer arquivo.Close()

	scanner := bufio.NewScanner(arquivo)
	if !scanner.Scan() {
		t.Fatalf("%s: arquivo vazio", caminho)
	}
	n, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil {
		t.Fatalf("%s: primeira linha inválida: %v", caminho, err)
	}

	var pares [][2]int
	for scanner.Scan() {
		campos := strings.Fields(scanner.Text())
		if len(campos) < 2 {
			continue
		}
		p, err1 := strconv.Atoi(campos[0])
		q, err2 := strconv.Atoi(campos[1])
		if err1 != nil || err2 != nil {
			t.Fatalf("%s: par inválido em %q", caminho, scanner.Text())
		}
		if p < 0 || q < 0 {
			break
		}
		pares = append(pares, [2]int{p, q})
	}
	return n, pares
}
