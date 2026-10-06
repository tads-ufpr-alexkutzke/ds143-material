// Tabela de dispersão com endereçamento aberto: todas as chaves ficam no
// próprio vetor, e uma colisão é resolvida procurando outra posição livre.
// Três regras de sondagem: linear, quadrática e dispersão dupla.
//
// Execute com:  go run sondagem.go

package main

import (
	"fmt"
	"math"
	"math/rand"
)

type Estado int

const (
	Vazia Estado = iota
	Ocupada
	Removida // lápide: a posição já teve uma chave, e a busca não pode parar nela
)

type Sondagem int

const (
	Linear Sondagem = iota
	Quadratica
	Dupla
)

func (s Sondagem) String() string {
	return [...]string{"linear", "quadrática", "dupla"}[s]
}

type Tabela struct {
	chaves []int
	estado []Estado
	regra  Sondagem
	n      int
}

func NovaTabela(m int, regra Sondagem) *Tabela {
	return &Tabela{chaves: make([]int, m), estado: make([]Estado, m), regra: regra}
}

func (t *Tabela) m() int {
	return len(t.chaves)
}

// posicao devolve a i-ésima posição da sequência de sondagem da chave, com
// i = 0, 1, 2, ... A posição 0 da sequência é h1(k) nas três regras.
func (t *Tabela) posicao(k, i int) int {
	m := t.m()
	h1 := k % m
	switch t.regra {
	case Linear:
		return (h1 + i) % m
	case Quadratica:
		return (h1 + i*i) % m
	default: // Dupla
		h2 := 1 + k%(m-1)
		return (h1 + i*h2) % m
	}
}

// Busca devolve se a chave está na tabela e quantas posições examinou. Ela
// para na primeira posição vazia, e passa por cima das lápides.
func (t *Tabela) Busca(k int) (bool, int) {
	for i := 0; i < t.m(); i++ {
		p := t.posicao(k, i)
		switch {
		case t.estado[p] == Vazia:
			return false, i + 1
		case t.estado[p] == Ocupada && t.chaves[p] == k:
			return true, i + 1
		}
	}
	return false, t.m()
}

// Insere guarda a chave na primeira posição livre da sua sequência, vazia ou
// com lápide, depois de confirmar que ela ainda não está na tabela. Devolve
// quantas posições examinou, ou -1 se a sequência não achou posição livre.
func (t *Tabela) Insere(k int) int {
	livre := -1
	for i := 0; i < t.m(); i++ {
		p := t.posicao(k, i)
		switch {
		case t.estado[p] == Ocupada && t.chaves[p] == k:
			return i + 1
		case t.estado[p] == Removida && livre == -1:
			livre = p
		case t.estado[p] == Vazia:
			if livre == -1 {
				livre = p
			}
			t.chaves[livre] = k
			t.estado[livre] = Ocupada
			t.n++
			return i + 1
		}
	}
	if livre != -1 {
		t.chaves[livre] = k
		t.estado[livre] = Ocupada
		t.n++
		return t.m()
	}
	return -1
}

// Remove marca a posição com uma lápide. Esvaziá-la cortaria a sequência de
// sondagem das chaves que passaram por ela ao serem inseridas.
func (t *Tabela) Remove(k int) bool {
	for i := 0; i < t.m(); i++ {
		p := t.posicao(k, i)
		switch {
		case t.estado[p] == Vazia:
			return false
		case t.estado[p] == Ocupada && t.chaves[p] == k:
			t.estado[p] = Removida
			t.n--
			return true
		}
	}
	return false
}

func (t *Tabela) Imprime() {
	fmt.Print("   ")
	for i := 0; i < t.m(); i++ {
		fmt.Printf("%4d", i)
	}
	fmt.Print("\n   ")
	for i := 0; i < t.m(); i++ {
		switch t.estado[i] {
		case Vazia:
			fmt.Print("   .")
		case Removida:
			fmt.Print("   X")
		default:
			fmt.Printf("%4d", t.chaves[i])
		}
	}
	fmt.Println()
}

// exemplo insere poucas chaves em uma tabela de 11 posições e mostra quantas
// posições cada inserção examinou.
func exemplo(regra Sondagem) *Tabela {
	t := NovaTabela(11, regra)
	fmt.Printf("Sondagem %s, m = 11, h1(k) = k %% 11\n", regra)
	for _, k := range []int{22, 33, 5, 16, 27, 38, 44} {
		fmt.Printf("  insere %2d: h1 = %2d, posições examinadas: %d\n",
			k, k%11, t.Insere(k))
	}
	t.Imprime()
	fmt.Println()
	return t
}

// experimento enche uma tabela grande até o fator de carga alfa com chaves
// aleatórias e mede o número médio de posições examinadas por busca.
func experimento(m int, alfa float64, regra Sondagem, gerador *rand.Rand) (sucesso, fracasso float64, falhas int) {
	t := NovaTabela(m, regra)
	alvo := int(alfa * float64(m))
	vistas := make(map[int]bool)
	var inseridas []int
	for len(inseridas) < alvo {
		k := gerador.Intn(1 << 40)
		if vistas[k] {
			continue
		}
		vistas[k] = true
		if t.Insere(k) == -1 {
			falhas++
			continue
		}
		inseridas = append(inseridas, k)
	}

	total := 0
	for _, k := range inseridas {
		_, sondas := t.Busca(k)
		total += sondas
	}
	sucesso = float64(total) / float64(len(inseridas))

	total = 0
	consultas := 0
	for consultas < 100000 {
		k := gerador.Intn(1 << 40)
		if vistas[k] {
			continue
		}
		_, sondas := t.Busca(k)
		total += sondas
		consultas++
	}
	fracasso = float64(total) / float64(consultas)
	return sucesso, fracasso, falhas
}

func main() {
	t := exemplo(Linear)
	exemplo(Quadratica)
	exemplo(Dupla)

	fmt.Println("Remoção com lápide, na tabela da sondagem linear:")
	fmt.Printf("  Remove(16) = %v\n", t.Remove(16))
	achou, sondas := t.Busca(38)
	fmt.Printf("  Busca(38) = %v, posições examinadas: %d\n", achou, sondas)
	t.Imprime()

	const m = 100003 // primo
	gerador := rand.New(rand.NewSource(1))
	cargas := []float64{0.5, 0.75, 0.9, 0.95}
	medidas := make(map[Sondagem][][3]float64)
	for _, regra := range []Sondagem{Linear, Quadratica, Dupla} {
		for _, alfa := range cargas {
			s, f, falhas := experimento(m, alfa, regra, gerador)
			medidas[regra] = append(medidas[regra], [3]float64{s, f, float64(falhas)})
		}
	}

	fmt.Println()
	fmt.Printf("Média de posições examinadas por busca, m = %d, chaves aleatórias\n", m)
	for j, titulo := range []string{"bem-sucedida", "malsucedida"} {
		fmt.Printf("\nBusca %s\n\n", titulo)
		fmt.Println(" alfa | linear | quadrática |  dupla | teoria linear | teoria uniforme")
		fmt.Println("------+--------+------------+--------+---------------+----------------")
		for i, a := range cargas {
			var linear, uniforme float64
			if j == 0 {
				linear = 0.5 * (1 + 1/(1-a))
				uniforme = 1 / a * math.Log(1/(1-a))
			} else {
				linear = 0.5 * (1 + 1/((1-a)*(1-a)))
				uniforme = 1 / (1 - a)
			}
			fmt.Printf(" %.2f | %6.2f | %10.2f | %6.2f | %13.2f | %15.2f\n", a,
				medidas[Linear][i][j], medidas[Quadratica][i][j], medidas[Dupla][i][j],
				linear, uniforme)
		}
	}

	falhas := 0.0
	for _, regra := range []Sondagem{Linear, Quadratica, Dupla} {
		for _, med := range medidas[regra] {
			falhas += med[2]
		}
	}
	fmt.Printf("\nInserções que não acharam posição livre: %.0f\n", falhas)
}
