// Tabela de dispersão com encadeamento: cada posição da tabela guarda a lista
// das chaves que caíram nela.
//
// Execute com:  go run encadeamento.go

package main

import (
	"fmt"
	"strings"
)

// Par associa uma chave a um valor. Aqui, uma palavra ao número de vezes que
// ela aparece no texto.
type Par struct {
	Chave string
	Valor int
}

// Tabela guarda m listas de pares. A lista de cada posição é um slice de Go:
// o encadeamento clássico usa lista encadeada, e o comportamento é o mesmo.
type Tabela struct {
	posicoes [][]Par
	n        int // número de pares guardados
}

func NovaTabela(m int) *Tabela {
	return &Tabela{posicoes: make([][]Par, m)}
}

func (t *Tabela) m() int {
	return len(t.posicoes)
}

// hash é a função polinomial de base 31, com o resto tirado a cada passo.
func (t *Tabela) hash(chave string) int {
	h := 0
	for i := 0; i < len(chave); i++ {
		h = (31*h + int(chave[i])) % t.m()
	}
	return h
}

// Busca percorre apenas a lista da posição da chave.
func (t *Tabela) Busca(chave string) (int, bool) {
	for _, p := range t.posicoes[t.hash(chave)] {
		if p.Chave == chave {
			return p.Valor, true
		}
	}
	return 0, false
}

// Insere troca o valor se a chave já existe; senão, acrescenta o par ao fim
// da lista da posição. Quando o número de pares passa do número de posições,
// a tabela cresce.
func (t *Tabela) Insere(chave string, valor int) {
	i := t.hash(chave)
	for j := range t.posicoes[i] {
		if t.posicoes[i][j].Chave == chave {
			t.posicoes[i][j].Valor = valor
			return
		}
	}
	t.posicoes[i] = append(t.posicoes[i], Par{chave, valor})
	t.n++

	if t.n > t.m() {
		antigo := t.m()
		t.redimensiona(proximoPrimo(2 * t.m()))
		fmt.Printf("inserção de %-8q n = %2d > m = %2d, a tabela passa a ter %d posições\n",
			chave, t.n, antigo, t.m())
	}
}

// Remove tira o par da lista da posição, se ele estiver lá.
func (t *Tabela) Remove(chave string) bool {
	i := t.hash(chave)
	for j, p := range t.posicoes[i] {
		if p.Chave == chave {
			t.posicoes[i] = append(t.posicoes[i][:j], t.posicoes[i][j+1:]...)
			t.n--
			return true
		}
	}
	return false
}

// redimensiona cria uma tabela com m posições e reinsere todos os pares. A
// posição de cada chave muda, porque a função de dispersão depende de m.
func (t *Tabela) redimensiona(m int) {
	nova := NovaTabela(m)
	for _, lista := range t.posicoes {
		for _, p := range lista {
			i := nova.hash(p.Chave)
			nova.posicoes[i] = append(nova.posicoes[i], p)
			nova.n++
		}
	}
	*t = *nova
}

// proximoPrimo devolve o menor primo maior ou igual a k.
func proximoPrimo(k int) int {
	for {
		primo := k >= 2
		for d := 2; d*d <= k; d++ {
			if k%d == 0 {
				primo = false
				break
			}
		}
		if primo {
			return k
		}
		k++
	}
}

func (t *Tabela) Imprime() {
	for i, lista := range t.posicoes {
		fmt.Printf("%3d:", i)
		for _, p := range lista {
			fmt.Printf(" [%s %d]", p.Chave, p.Valor)
		}
		fmt.Println()
	}
}

func main() {
	texto := "o rato roeu a roupa do rei de roma e a rainha de raiva " +
		"roeu o resto da roupa do rato"

	t := NovaTabela(7)
	for _, palavra := range strings.Fields(texto) {
		quantas, _ := t.Busca(palavra)
		t.Insere(palavra, quantas+1)
	}

	fmt.Printf("\n%d palavras distintas em %d posições (fator de carga %.2f)\n\n",
		t.n, t.m(), float64(t.n)/float64(t.m()))
	t.Imprime()

	fmt.Println()
	for _, p := range []string{"roupa", "rei", "rainha", "reino"} {
		v, achou := t.Busca(p)
		fmt.Printf("Busca(%q) = %d, %v\n", p, v, achou)
	}

	fmt.Println()
	fmt.Printf("Remove(%q) = %v\n", "rei", t.Remove("rei"))
	v, achou := t.Busca("rei")
	fmt.Printf("Busca(%q) = %d, %v\n", "rei", v, achou)
}
