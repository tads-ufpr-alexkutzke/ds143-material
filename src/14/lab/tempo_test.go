package main

import (
	"math/rand"
	"runtime"
	"sync"
	"testing"
	"time"
)

// Os limites abaixo são muito maiores que os tempos de uma solução com os
// custos exigidos, para não reprovar uma solução correta em uma máquina lenta
// ou ocupada. Por isso, estes testes reprovam custos bem acima dos exigidos,
// mas não demonstram sozinhos que um custo é O(V) ou O(1).
//
// Os testes deste arquivo rodam depois dos de linha_test.go. Cada um depende de
// o anterior ter limitado o custo, e pula a medida quando ele reprovou.
const (
	nGrande = 1000000

	// Um vetor com um vértice por posição, percorrido uma vez, leva menos de um
	// segundo. Percorrer o vetor inteiro para cada vértice leva minutos.
	limitePreprocessa = 10 * time.Second

	// Uma consulta O(1) leva menos de um microssegundo. Percorrer a linha leva
	// mais de um milissegundo.
	limitePorConsulta = 10 * time.Microsecond

	// A primeira consulta não deve percorrer a linha.
	limitePrimeiraConsulta = time.Millisecond

	// Um vetor com um elemento por vértice cabe com folga neste limite, e uma
	// tabela com um elemento por par de vértices, não.
	nMemoria        = 8000
	bytesPorVertice = 2048
)

var (
	// preprocessaLento e consultaLenta registram que um custo reprovou. Os
	// testes seguintes pulam a medida, pois levariam minutos.
	preprocessaLento bool
	consultaLenta    bool

	// A mesma linha grande serve a todos os testes, para montá-la uma vez só.
	linhaGrandeUnica *linhaDeTeste
	montaLinhaGrande sync.Once
)

func linhaGrande() *linhaDeTeste {
	montaLinhaGrande.Do(func() {
		linhaGrandeUnica = linhaSorteada(nGrande, 31)
	})
	return linhaGrandeUnica
}

// cronometra roda Preprocessa com o limite de tempo. Se o limite estourar, o
// teste termina com erro, e a execução interrompida só para no fim do go test.
func cronometra(t *testing.T, g *GrafoP) (time.Duration, *Preprocessamento) {
	t.Helper()
	if preprocessaLento {
		t.Skip("Preprocessa já reprovou em um teste de custo anterior")
	}
	type resultado struct {
		tempo time.Duration
		p     *Preprocessamento
	}
	pronto := make(chan resultado, 1)
	go func() {
		inicio := time.Now()
		p := Preprocessa(g)
		pronto <- resultado{time.Since(inicio), p}
	}()
	select {
	case r := <-pronto:
		return r.tempo, r.p
	case <-time.After(limitePreprocessa):
		preprocessaLento = true
		t.Fatalf("Preprocessa levou mais de %v com %d vértices, muito acima do esperado para O(V)",
			limitePreprocessa, g.V())
		return 0, nil
	}
}

// A memória é conferida antes dos tempos para que uma solução com uma tabela de
// V^2 posições reprove aqui, antes de tentar alocar essa tabela com nGrande.
func TestPreprocessaUsaMemoriaLinear(t *testing.T) {
	g := linhaSorteada(nMemoria, 32).g

	var antes, depois runtime.MemStats
	runtime.ReadMemStats(&antes)
	p := Preprocessa(g)
	runtime.ReadMemStats(&depois)
	runtime.KeepAlive(p)

	if alocado := depois.TotalAlloc - antes.TotalAlloc; alocado > bytesPorVertice*nMemoria {
		preprocessaLento = true
		t.Fatalf("Preprocessa alocou %d bytes para %d vértices, mais de %d bytes por vértice. Isso indica um custo acima de O(V)",
			alocado, nMemoria, bytesPorVertice)
	}
}

func TestPreprocessaEmTempoLinear(t *testing.T) {
	tempo, _ := cronometra(t, linhaGrande().g)
	t.Logf("Preprocessa com %d vértices: %v", nGrande, tempo)
}

// tempoPorConsulta repete consultas sorteadas até somar 20 ms, para medir bem
// tanto uma consulta de nanossegundos quanto uma de milissegundos.
func tempoPorConsulta(p *Preprocessamento, semente int64) time.Duration {
	r := rand.New(rand.NewSource(semente))
	pares := make([][2]int, 1<<16)
	for i := range pares {
		pares[i] = [2]int{r.Intn(nGrande), r.Intn(nGrande)}
	}

	soma := 0
	feitas := 0
	inicio := time.Now()
	for lote := 1; ; lote *= 2 {
		for i := 0; i < lote; i++ {
			par := pares[(feitas+i)&(len(pares)-1)]
			soma += p.Distancia(par[0], par[1])
		}
		feitas += lote
		if decorrido := time.Since(inicio); decorrido >= 20*time.Millisecond {
			runtime.KeepAlive(soma)
			return decorrido / time.Duration(feitas)
		}
	}
}

func TestDistanciaEmTempoConstante(t *testing.T) {
	_, p := cronometra(t, linhaGrande().g)

	melhor := tempoPorConsulta(p, 33)
	for semente := int64(34); semente <= 35; semente++ {
		if tempo := tempoPorConsulta(p, semente); tempo < melhor {
			melhor = tempo
		}
	}

	t.Logf("Distancia com %d vértices: %v por consulta", nGrande, melhor)
	if melhor > limitePorConsulta {
		consultaLenta = true
		t.Fatalf("cada consulta levou %v com %d vértices (limite: %v), muito acima do esperado para O(1)",
			melhor, nGrande, limitePorConsulta)
	}
}

// Preprocessa deve deixar tudo pronto: a primeira consulta custa o mesmo que as
// outras, sem adiar para ela o percurso pela linha.
func TestPrimeiraConsultaEmTempoConstante(t *testing.T) {
	g := linhaGrande().g
	r := rand.New(rand.NewSource(36))

	melhor := time.Hour
	for i := 0; i < 3; i++ {
		_, p := cronometra(t, g)
		v, w := r.Intn(nGrande), r.Intn(nGrande)
		inicio := time.Now()
		d := p.Distancia(v, w)
		if tempo := time.Since(inicio); tempo < melhor {
			melhor = tempo
		}
		runtime.KeepAlive(d)
	}

	t.Logf("primeira consulta com %d vértices: %v", nGrande, melhor)
	if melhor > limitePrimeiraConsulta {
		t.Fatalf("a primeira consulta levou %v com %d vértices (limite: %v). Isso indica que o trabalho de Preprocessa ficou para Distancia",
			melhor, nGrande, limitePrimeiraConsulta)
	}
}

// Fica por último porque consulta a linha grande: uma solução com custo acima
// do exigido já reprovou antes de chegar aqui.
func TestLinhaLonga(t *testing.T) {
	if consultaLenta {
		t.Skip("Distancia já reprovou em um teste de custo anterior")
	}
	l := linhaGrande()
	_, p := cronometra(t, l.g)
	r := rand.New(rand.NewSource(37))
	for i := 0; i < 100000; i++ {
		v, w := r.Intn(nGrande), r.Intn(nGrande)
		if d := p.Distancia(v, w); d != l.esperada(v, w) {
			t.Fatalf("Distancia(%d, %d) = %d, esperado %d", v, w, d, l.esperada(v, w))
		}
	}
}
