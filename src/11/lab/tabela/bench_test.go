package tabela

// Medidas de desempenho da Tabela contra o map de Go. Rode com:
//
//	go test -run '^$' -bench . ./tabela
//
// O -run '^$' pula os testes e roda só as medidas. Cada linha da saída traz,
// na coluna ns/op, o tempo médio de uma operação em nanossegundos.

import (
	"math/rand"
	"testing"
)

const nMedida = 100000 // chaves em cada tabela medida

// chavesAleatorias devolve n chaves distintas sorteadas entre 0 e 2^40.
func chavesAleatorias(n int) []int {
	gerador := rand.New(rand.NewSource(3))
	vistas := make(map[int]bool, n)
	chaves := make([]int, 0, n)
	for len(chaves) < n {
		k := gerador.Intn(1 << 40)
		if !vistas[k] {
			vistas[k] = true
			chaves = append(chaves, k)
		}
	}
	return chaves
}

// chavesSequenciais devolve 0, 1, 2, ..., n-1, como os identificadores que
// um sistema atribui em ordem.
func chavesSequenciais(n int) []int {
	chaves := make([]int, n)
	for i := range chaves {
		chaves[i] = i
	}
	return chaves
}

// ausentes devolve n chaves que não estão em chaves, sorteadas entre 0 e
// 2^40.
func ausentes(chaves []int, n int) []int {
	presentes := make(map[int]bool, len(chaves))
	for _, k := range chaves {
		presentes[k] = true
	}
	gerador := rand.New(rand.NewSource(4))
	resultado := make([]int, 0, n)
	for len(resultado) < n {
		k := gerador.Intn(1 << 40)
		if !presentes[k] {
			resultado = append(resultado, k)
		}
	}
	return resultado
}

// BenchmarkInsere mede uma inserção, começando de uma tabela pequena que
// cresce até nMedida chaves.
func BenchmarkInsere(b *testing.B) {
	for _, caso := range []struct {
		nome   string
		chaves []int
	}{
		{"aleatorias", chavesAleatorias(nMedida)},
		{"sequenciais", chavesSequenciais(nMedida)},
	} {
		b.Run("Tabela/"+caso.nome, func(b *testing.B) {
			tab := Nova(11)
			for i := 0; i < b.N; i++ {
				if i%nMedida == 0 {
					tab = Nova(11)
				}
				tab.Insere(caso.chaves[i%nMedida], i)
			}
		})
		b.Run("map/"+caso.nome, func(b *testing.B) {
			m := make(map[int]int)
			for i := 0; i < b.N; i++ {
				if i%nMedida == 0 {
					m = make(map[int]int)
				}
				m[caso.chaves[i%nMedida]] = i
			}
		})
	}
}

// BenchmarkBusca mede uma busca em uma tabela com nMedida chaves: por chaves
// presentes, nos dois casos, e por chaves ausentes, no caso das aleatórias.
func BenchmarkBusca(b *testing.B) {
	for _, caso := range []struct {
		nome   string
		chaves []int
	}{
		{"aleatorias", chavesAleatorias(nMedida)},
		{"sequenciais", chavesSequenciais(nMedida)},
	} {
		tab := Nova(11)
		m := make(map[int]int)
		for i, k := range caso.chaves {
			tab.Insere(k, i)
			m[k] = i
		}
		embaralhadas := append([]int(nil), caso.chaves...)
		rand.New(rand.NewSource(5)).Shuffle(len(embaralhadas), func(i, j int) {
			embaralhadas[i], embaralhadas[j] = embaralhadas[j], embaralhadas[i]
		})
		consultas := []struct {
			nome   string
			chaves []int
		}{{"presentes", embaralhadas}}
		if caso.nome == "aleatorias" {
			consultas = append(consultas, struct {
				nome   string
				chaves []int
			}{"ausentes", ausentes(caso.chaves, nMedida)})
		}

		for _, consulta := range consultas {
			b.Run("Tabela/"+caso.nome+"/"+consulta.nome, func(b *testing.B) {
				soma := 0
				for i := 0; i < b.N; i++ {
					v, _ := tab.Busca(consulta.chaves[i%nMedida])
					soma += v
				}
				resultado = soma
			})
			b.Run("map/"+caso.nome+"/"+consulta.nome, func(b *testing.B) {
				soma := 0
				for i := 0; i < b.N; i++ {
					soma += m[consulta.chaves[i%nMedida]]
				}
				resultado = soma
			})
		}
	}
}

// BenchmarkJanela mede um passo de uma janela deslizante: insere a chave i,
// remove a chave i-1000 e busca uma chave viva. A tabela fica sempre com
// cerca de 1000 chaves, como uma lista de pedidos que chegam e saem em ordem.
func BenchmarkJanela(b *testing.B) {
	const janela = 1000
	b.Run("Tabela", func(b *testing.B) {
		tab := Nova(11)
		soma := 0
		for i := 0; i < b.N; i++ {
			tab.Insere(i, i)
			if i >= janela {
				tab.Remove(i - janela)
				v, _ := tab.Busca(i - janela/2)
				soma += v
			}
		}
		resultado = soma
	})
	b.Run("map", func(b *testing.B) {
		m := make(map[int]int)
		soma := 0
		for i := 0; i < b.N; i++ {
			m[i] = i
			if i >= janela {
				delete(m, i-janela)
				soma += m[i-janela/2]
			}
		}
		resultado = soma
	})
}

// resultado guarda a soma dos valores buscados, para o compilador não
// descartar as buscas como código sem efeito.
var resultado int
