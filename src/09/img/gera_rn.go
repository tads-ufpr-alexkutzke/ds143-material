// Gera as figuras de árvore rubro-negra left-leaning do roteiro da aula, no
// desenho de Sedgewick e Wayne (Algorithms, 4a ed., seção 3.3): a cor é a da
// ligação que chega ao nó, e a ligação vermelha é o traço grosso vermelho.
//
// A inserção é a mesma do codes/rn.go, com uma chamada a mais depois de cada
// conserto, que fotografa a árvore inteira para a figura passo a passo.
//
// Execute com:  go run gera_rn.go

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type Cor bool

const (
	Vermelho Cor = true
	Preto    Cor = false
)

type No struct {
	Info int
	Cor  Cor
	Esq  *No
	Dir  *No
}

func vermelho(a *No) bool {
	return a != nil && a.Cor == Vermelho
}

func rotacionaEsquerda(a *No) *No {
	b := a.Dir
	a.Dir = b.Esq
	b.Esq = a
	b.Cor = a.Cor
	a.Cor = Vermelho
	return b
}

func rotacionaDireita(a *No) *No {
	b := a.Esq
	a.Esq = b.Dir
	b.Dir = a
	b.Cor = a.Cor
	a.Cor = Vermelho
	return b
}

func inverteCores(a *No) {
	a.Cor = Vermelho
	a.Esq.Cor = Preto
	a.Dir.Cor = Preto
}

// Arvore guarda a raiz e, durante uma inserção, as chaves do caminho
// percorrido. Se foto não for nil, é chamada depois de cada mudança.
type Arvore struct {
	Raiz    *No
	caminho map[int]bool
	foto    func(legenda string, destaque int)
}

func (t *Arvore) nota(legenda string, destaque int) {
	if t.foto != nil {
		t.foto(legenda, destaque)
	}
}

// Insere acrescenta v. O parâmetro liga de insere pendura o nó devolvido no
// pai antes da foto, para que a foto mostre a árvore inteira já consertada.
func (t *Arvore) Insere(v int) {
	t.caminho = map[int]bool{}
	t.Raiz = t.insere(t.Raiz, v, func(n *No) { t.Raiz = n })
	if t.Raiz.Cor == Vermelho {
		t.Raiz.Cor = Preto
		t.nota("a raiz volta a ser preta", t.Raiz.Info)
	}
}

func (t *Arvore) insere(a *No, v int, liga func(*No)) *No {
	if a == nil {
		n := &No{Info: v, Cor: Vermelho}
		liga(n)
		t.caminho[v] = true
		t.nota(rotulo(v)+" entra como folha vermelha", v)
		return n
	}
	t.caminho[a.Info] = true

	if v < a.Info {
		a.Esq = t.insere(a.Esq, v, func(n *No) { a.Esq = n })
	} else if v > a.Info {
		a.Dir = t.insere(a.Dir, v, func(n *No) { a.Dir = n })
	}

	if vermelho(a.Dir) && !vermelho(a.Esq) {
		k := a.Info
		a = rotacionaEsquerda(a)
		liga(a)
		t.nota("filho direito vermelho: rotação à esquerda em "+rotulo(k), k)
	}
	if vermelho(a.Esq) && vermelho(a.Esq.Esq) {
		k := a.Info
		a = rotacionaDireita(a)
		liga(a)
		t.nota("dois vermelhos seguidos à esquerda: rotação à direita em "+rotulo(k), k)
	}
	if vermelho(a.Esq) && vermelho(a.Dir) {
		inverteCores(a)
		t.nota("dois filhos vermelhos: inversão de cores em "+rotulo(a.Info), a.Info)
	}
	return a
}

// As figuras com letras guardam o código da letra em Info.
func rotulo(v int) string {
	return string(rune(v))
}

func copia(a *No) *No {
	if a == nil {
		return nil
	}
	return &No{Info: a.Info, Cor: a.Cor, Esq: copia(a.Esq), Dir: copia(a.Dir)}
}

// ---------------------------------------------------------------- desenho

const (
	corVermelha = "#c0392b"
	corCinza    = "#b8b8b8"
	corDestaque = "#f8c471"
	fonte       = `font-family="Helvetica, Arial, sans-serif"`
)

// posicoes dá a cada chave a coluna (ordem in-ordem) e a linha (profundidade).
func posicoes(a *No, prof int, col *int, px, py map[int]int) {
	if a == nil {
		return
	}
	posicoes(a.Esq, prof+1, col, px, py)
	px[a.Info] = *col
	py[a.Info] = prof
	*col++
	posicoes(a.Dir, prof+1, col, px, py)
}

func altura(a *No) int {
	if a == nil {
		return -1
	}
	return 1 + max(altura(a.Esq), altura(a.Dir))
}

// estilo diz como desenhar uma árvore de letras.
type estilo struct {
	cinza    map[int]bool // nós fora do caminho da inserção
	novo     int          // chave recém-inserida, com a letra em vermelho
	destaque int          // nó em que o conserto foi aplicado, com fundo laranja
}

const (
	dx   = 30.0
	dy   = 38.0
	raio = 11.0
)

// letras desenha a árvore com a raiz no topo e o canto esquerdo em (x0, y0).
func letras(b *strings.Builder, raiz *No, x0, y0 float64, e estilo) {
	px, py := map[int]int{}, map[int]int{}
	col := 0
	posicoes(raiz, 0, &col, px, py)
	xy := func(k int) (float64, float64) {
		return x0 + raio + float64(px[k])*dx, y0 + raio + float64(py[k])*dy
	}

	var ligacoes func(a *No)
	ligacoes = func(a *No) {
		if a == nil {
			return
		}
		for _, f := range []*No{a.Esq, a.Dir} {
			if f == nil {
				continue
			}
			x1, y1 := xy(a.Info)
			x2, y2 := xy(f.Info)
			cor, larg := "black", 1.2
			if f.Cor == Vermelho {
				cor, larg = corVermelha, 4.5
			}
			if e.cinza[f.Info] {
				cor = corCinza
			}
			fmt.Fprintf(b, `<line x1="%.1f" y1="%.1f" x2="%.1f" y2="%.1f" stroke="%s" stroke-width="%.1f" stroke-linecap="round"/>`+"\n",
				x1, y1, x2, y2, cor, larg)
		}
		ligacoes(a.Esq)
		ligacoes(a.Dir)
	}
	ligacoes(raiz)

	var nos func(a *No)
	nos = func(a *No) {
		if a == nil {
			return
		}
		x, y := xy(a.Info)
		fundo, borda, texto := "white", "black", "black"
		if e.cinza[a.Info] {
			borda, texto = corCinza, corCinza
		}
		if a.Info == e.destaque {
			fundo = corDestaque
		}
		if a.Info == e.novo {
			texto = corVermelha
		}
		fmt.Fprintf(b, `<circle cx="%.1f" cy="%.1f" r="%.1f" fill="%s" stroke="%s" stroke-width="1.2"/>`+"\n",
			x, y, raio, fundo, borda)
		fmt.Fprintf(b, `<text x="%.1f" y="%.1f" %s font-size="13" text-anchor="middle" fill="%s">%s</text>`+"\n",
			x, y+4.5, fonte, texto, rotulo(a.Info))
		nos(a.Esq)
		nos(a.Dir)
	}
	nos(raiz)
}

func grava(nome string, larg, alt float64, corpo string) {
	svg := fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" width="%.0f" height="%.0f" viewBox="0 0 %.0f %.0f">
<rect width="100%%" height="100%%" fill="white"/>
%s</svg>
`, larg, alt, larg, alt, corpo)
	if err := os.WriteFile(nome, []byte(svg), 0o644); err != nil {
		panic(err)
	}
	fmt.Println("gerado", nome)
}

func texto(b *strings.Builder, x, y float64, tam int, cor, s string) {
	fmt.Fprintf(b, `<text x="%.1f" y="%.1f" %s font-size="%d" fill="%s">%s</text>`+"\n",
		x, y, fonte, tam, cor, s)
}

// ---------------------------------------------------------------- figuras

func arvoreDeLetras(chaves string) *Arvore {
	t := &Arvore{}
	for _, c := range chaves {
		t.Insere(int(c))
	}
	return t
}

// insercaoP mostra cada conserto da inserção de P, um painel por foto.
func insercaoP() {
	type foto struct {
		raiz     *No
		legenda  string
		destaque int
	}
	t := arvoreDeLetras(inicialP)
	var fotos []foto
	t.foto = func(legenda string, destaque int) {
		fotos = append(fotos, foto{copia(t.Raiz), legenda, destaque})
	}
	t.Insere('P')

	var b strings.Builder
	y := 12.0
	for i, f := range fotos {
		alt := float64(altura(f.raiz))*dy + 2*raio
		texto(&b, 12, y+16, 14, "black", fmt.Sprintf("%d. %s", i+1, f.legenda))
		letras(&b, f.raiz, 30, y+30, estilo{novo: 'P', destaque: f.destaque})
		y += 30 + alt + 26
	}
	grava("rn_insercao_P.svg", 440, y, b.String())
}

// sequencia mostra a árvore depois de cada chave de SEARCHXMPL, em duas
// colunas. Nós fora do caminho da inserção ficam em cinza.
func sequencia() {
	const chaves = "SEARCHXMPL"
	const colunas, largCol = 2, 400.0
	t := &Arvore{}
	var b strings.Builder
	y := [colunas]float64{12, 12}
	metade := (len(chaves) + 1) / 2
	for i, c := range chaves {
		t.Insere(int(c))
		cinza := map[int]bool{}
		var marca func(a *No)
		marca = func(a *No) {
			if a == nil {
				return
			}
			if !t.caminho[a.Info] {
				cinza[a.Info] = true
			}
			marca(a.Esq)
			marca(a.Dir)
		}
		marca(t.Raiz)

		k := i / metade
		x := 12 + float64(k)*largCol
		texto(&b, x, y[k]+18, 15, corVermelha, "insere "+string(c))
		letras(&b, t.Raiz, x+80, y[k], estilo{cinza: cinza, novo: int(c)})
		y[k] += float64(altura(t.Raiz))*dy + 2*raio + 26
	}
	grava("rn_sequencia.svg", colunas*largCol, max(y[0], y[1]), b.String())
}

// pontos desenha uma árvore grande, um ponto por nó, com as estatísticas
// no canto superior esquerdo.
func pontos(nome string, chaves []int) {
	t := &Arvore{}
	for _, v := range chaves {
		t.Insere(v)
	}
	const px, py, r = 4.6, 30.0, 2.4
	x0, y0 := 20.0, 110.0
	pos := func(col, prof int) (float64, float64) {
		return x0 + float64(col)*px, y0 + float64(prof)*py
	}
	cx, cy := map[int]int{}, map[int]int{}
	col := 0
	posicoes(t.Raiz, 0, &col, cx, cy)

	var b strings.Builder
	var ligacoes func(a *No)
	vermelhas := 0
	ligacoes = func(a *No) {
		if a == nil {
			return
		}
		for _, f := range []*No{a.Esq, a.Dir} {
			if f == nil {
				continue
			}
			x1, y1 := pos(cx[a.Info], cy[a.Info])
			x2, y2 := pos(cx[f.Info], cy[f.Info])
			cor, larg := "black", 0.8
			if f.Cor == Vermelho {
				cor, larg = corVermelha, 1.8
				vermelhas++
			}
			fmt.Fprintf(&b, `<line x1="%.1f" y1="%.1f" x2="%.1f" y2="%.1f" stroke="%s" stroke-width="%.1f"/>`+"\n",
				x1, y1, x2, y2, cor, larg)
		}
		ligacoes(a.Esq)
		ligacoes(a.Dir)
	}
	ligacoes(t.Raiz)
	for k := range cx {
		x, y := pos(cx[k], cy[k])
		cor := "black"
		if encontra(t.Raiz, k).Cor == Vermelho {
			cor = corVermelha
		}
		fmt.Fprintf(&b, `<circle cx="%.1f" cy="%.1f" r="%.1f" fill="%s"/>`+"\n", x, y, r, cor)
	}

	n := len(chaves)
	soma := 0
	for k := range cy {
		soma += cy[k] + 1
	}
	niveis := altura(t.Raiz) + 1
	texto(&b, 20, 28, 15, "black", fmt.Sprintf("n = %d", n))
	texto(&b, 20, 48, 15, "black", fmt.Sprintf("níveis = %d", niveis))
	texto(&b, 20, 68, 15, "black", "média = "+virgula(float64(soma)/float64(n)))
	texto(&b, 20, 88, 15, "black", "ótimo = "+virgula(otimo(n)))
	fmt.Printf("%s: %d níveis, média %.2f, ótimo %.2f, %d ligações vermelhas\n",
		nome, niveis, float64(soma)/float64(n), otimo(n), vermelhas)

	grava(nome, x0*2+float64(n-1)*px, y0+float64(niveis-1)*py+20, b.String())
}

func encontra(a *No, v int) *No {
	for a.Info != v {
		if v < a.Info {
			a = a.Esq
		} else {
			a = a.Dir
		}
	}
	return a
}

// otimo é o número médio de nós visitados em uma busca na árvore com n nós
// e todos os níveis cheios, exceto o último.
func otimo(n int) float64 {
	soma, nivel, cabem := 0, 1, 1
	for resto := n; resto > 0; nivel++ {
		k := min(cabem, resto)
		soma += k * nivel
		resto -= k
		cabem *= 2
	}
	return float64(soma) / float64(n)
}

func virgula(x float64) string {
	return strings.Replace(fmt.Sprintf("%.1f", x), ".", ",", 1)
}

// inicialP é uma ordem de inserção que produz a árvore de partida da figura
// de Sedgewick: R no topo, com E ligado por vermelho, A e H vermelhos.
const inicialP = "AEHCRSM"

func main() {
	insercaoP()
	sequencia()

	const n = 255
	crescente := make([]int, n)
	decrescente := make([]int, n)
	for i := range n {
		crescente[i] = i + 1
		decrescente[i] = n - i
	}
	aleatoria := rand.New(rand.NewSource(143)).Perm(n)
	pontos("rn_crescente.svg", crescente)
	pontos("rn_decrescente.svg", decrescente)
	pontos("rn_aleatoria.svg", aleatoria)
}
