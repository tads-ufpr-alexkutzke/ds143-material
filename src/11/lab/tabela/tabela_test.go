package tabela

import (
	"math/rand"
	"testing"
)

// exemploDaAula monta a tabela da sondagem linear da aula de tabelas de
// dispersão: m = 11 e as chaves 22, 33, 5, 16, 27, 38 e 44, nessa ordem. O
// valor de cada chave é a chave vezes 10.
func exemploDaAula() *Tabela {
	t := Nova(11)
	for _, k := range []int{22, 33, 5, 16, 27, 38, 44} {
		t.Insere(k, 10*k)
	}
	return t
}

// confereDesenho compara as posições das chaves com o desenho esperado. Os
// desenhos usam h(k) = k mod m; com misturar = true, a conferência é pulada,
// e o restante do teste continua valendo.
func confereDesenho(t *testing.T, tab *Tabela, esperado string) {
	t.Helper()
	if misturar {
		return
	}
	if tab.String() != esperado {
		t.Fatalf("tabela diferente da esperada.\nobtida:\n%s\nesperada:\n%s", tab, esperado)
	}
}

func TestBuscaEmTabelaVazia(t *testing.T) {
	tab := Nova(11)
	if _, achou := tab.Busca(7); achou {
		t.Fatal("Busca(7) achou a chave em uma tabela vazia")
	}
}

func TestInsereExemploDaAula(t *testing.T) {
	tab := exemploDaAula()
	confereDesenho(t, tab,
		"   0   1   2   3   4   5   6   7   8   9  10\n"+
			"  22  33  44   .   .   5  16  27  38   .   .")
	if tab.Len() != 7 {
		t.Fatalf("Len() = %d, esperado 7", tab.Len())
	}
	for _, k := range []int{22, 33, 5, 16, 27, 38, 44} {
		if v, achou := tab.Busca(k); !achou || v != 10*k {
			t.Fatalf("Busca(%d) = %d, %v; esperado %d, true", k, v, achou, 10*k)
		}
	}
}

func TestInsereTrocaValor(t *testing.T) {
	tab := exemploDaAula()
	tab.Insere(38, 1)
	if v, achou := tab.Busca(38); !achou || v != 1 {
		t.Fatalf("depois de Insere(38, 1), Busca(38) = %d, %v; esperado 1, true", v, achou)
	}
	if tab.Len() != 7 {
		t.Fatalf("trocar o valor de uma chave mudou Len() para %d; esperado 7", tab.Len())
	}
}

func TestInsereChavesNegativas(t *testing.T) {
	tab := Nova(11)
	for _, k := range []int{-1, -12, -23, 10} {
		tab.Insere(k, k)
	}
	for _, k := range []int{-1, -12, -23, 10} {
		if v, achou := tab.Busca(k); !achou || v != k {
			t.Fatalf("Busca(%d) = %d, %v; esperado %d, true", k, v, achou, k)
		}
	}
}

func TestInsereCresce(t *testing.T) {
	tab := Nova(11)
	for k := 1; k <= 7; k++ {
		tab.Insere(k, k)
	}
	if tab.m() != 11 {
		t.Fatalf("com 7 chaves, m = %d; esperado 11 (7/11 = 0,64 não passa de 0,7)", tab.m())
	}
	tab.Insere(8, 8)
	if tab.m() != 23 {
		t.Fatalf("com 8 chaves, m = %d; esperado 23 (8/11 = 0,73 passa de 0,7, e 23 é o primeiro primo depois de 22)", tab.m())
	}
	for k := 9; k <= 5000; k++ {
		tab.Insere(k, k)
		if 10*tab.Len() > 7*tab.m() {
			t.Fatalf("com %d chaves e m = %d, o fator de carga passou de 0,7", tab.Len(), tab.m())
		}
	}
	for k := 1; k <= 5000; k++ {
		if v, achou := tab.Busca(k); !achou || v != k {
			t.Fatalf("depois de crescer, Busca(%d) = %d, %v; esperado %d, true", k, v, achou, k)
		}
	}
}

func TestInsereMuitasChaves(t *testing.T) {
	gerador := rand.New(rand.NewSource(1))
	tab := Nova(11)
	esperado := make(map[int]int)
	for i := 0; i < 50000; i++ {
		k := gerador.Intn(2000000) - 1000000
		tab.Insere(k, i)
		esperado[k] = i
	}
	if tab.Len() != len(esperado) {
		t.Fatalf("Len() = %d, esperado %d", tab.Len(), len(esperado))
	}
	for k, v := range esperado {
		if obtido, achou := tab.Busca(k); !achou || obtido != v {
			t.Fatalf("Busca(%d) = %d, %v; esperado %d, true", k, obtido, achou, v)
		}
	}
}

func TestRemoveExemploDaAula(t *testing.T) {
	tab := exemploDaAula()
	if !tab.Remove(16) {
		t.Fatal("Remove(16) = false; o 16 está na tabela")
	}
	// O 27 e o 38 (h = 5) descem uma posição cada, e a posição 9, vazia,
	// encerra a remoção.
	confereDesenho(t, tab,
		"   0   1   2   3   4   5   6   7   8   9  10\n"+
			"  22  33  44   .   .   5  27  38   .   .   .")
	if tab.Len() != 6 {
		t.Fatalf("Len() = %d, esperado 6", tab.Len())
	}
	if v, achou := tab.Busca(38); !achou || v != 380 {
		t.Fatalf("Busca(38) = %d, %v; esperado 380, true", v, achou)
	}
	if _, achou := tab.Busca(16); achou {
		t.Fatal("Busca(16) achou a chave removida")
	}
}

func TestRemoveVoltaCircular(t *testing.T) {
	tab := Nova(11)
	for _, k := range []int{9, 20, 31, 1, 10} {
		tab.Insere(k, k)
	}
	confereDesenho(t, tab,
		"   0   1   2   3   4   5   6   7   8   9  10\n"+
			"  31   1  10   .   .   .   .   .   .   9  20")
	if !tab.Remove(20) {
		t.Fatal("Remove(20) = false; o 20 está na tabela")
	}
	// O 31 (h = 9) desce da posição 0 para a 10, dando a volta no vetor. O 1
	// (h = 1) já está na posição de origem e fica onde está. O 10 (h = 10)
	// passou pela posição 0 ao ser inserido, e desce da 2 para a 0.
	confereDesenho(t, tab,
		"   0   1   2   3   4   5   6   7   8   9  10\n"+
			"  10   1   .   .   .   .   .   .   .   9  31")
	for _, k := range []int{9, 31, 1, 10} {
		if v, achou := tab.Busca(k); !achou || v != k {
			t.Fatalf("Busca(%d) = %d, %v; esperado %d, true", k, v, achou, k)
		}
	}
}

func TestRemoveAusente(t *testing.T) {
	tab := Nova(11)
	if tab.Remove(5) {
		t.Fatal("Remove(5) = true em uma tabela vazia")
	}
	tab = exemploDaAula()
	if tab.Remove(49) {
		t.Fatal("Remove(49) = true; o 49 não está na tabela")
	}
	if tab.Len() != 7 {
		t.Fatalf("remover uma chave ausente mudou Len() para %d", tab.Len())
	}
}

func TestRemoveAleatorio(t *testing.T) {
	// Chaves de um intervalo pequeno, para que as inserções e remoções caiam
	// sempre nos mesmos agrupamentos.
	gerador := rand.New(rand.NewSource(2))
	tab := Nova(11)
	esperado := make(map[int]int)
	for i := 0; i < 200000; i++ {
		k := gerador.Intn(500)
		switch gerador.Intn(3) {
		case 0:
			tab.Insere(k, i)
			esperado[k] = i
		case 1:
			_, estava := esperado[k]
			if tab.Remove(k) != estava {
				t.Fatalf("operação %d: Remove(%d) = %v; esperado %v", i, k, !estava, estava)
			}
			delete(esperado, k)
		default:
			v, achou := tab.Busca(k)
			vEsperado, estava := esperado[k]
			if achou != estava || v != vEsperado {
				t.Fatalf("operação %d: Busca(%d) = %d, %v; esperado %d, %v", i, k, v, achou, vEsperado, estava)
			}
		}
	}
	if tab.Len() != len(esperado) {
		t.Fatalf("Len() = %d, esperado %d", tab.Len(), len(esperado))
	}
}

func TestJanelaDeslizante(t *testing.T) {
	// Insere as chaves 0, 1, 2, ... e remove cada uma 1000 inserções depois,
	// como pedidos que chegam e saem em ordem. Sem lápides, a tabela não
	// precisa crescer além do que as 1000 chaves vivas exigem.
	const janela = 1000
	tab := Nova(11)
	for k := 0; k < 100000; k++ {
		tab.Insere(k, k)
		if k >= janela && !tab.Remove(k-janela) {
			t.Fatalf("Remove(%d) = false; a chave foi inserida %d inserções antes", k-janela, janela)
		}
	}
	if tab.Len() != janela {
		t.Fatalf("Len() = %d, esperado %d", tab.Len(), janela)
	}
	if tab.m() > 4*janela {
		t.Fatalf("m = %d para %d chaves vivas", tab.m(), janela)
	}
	for k := 100000 - janela; k < 100000; k++ {
		if _, achou := tab.Busca(k); !achou {
			t.Fatalf("Busca(%d) não achou uma chave viva", k)
		}
	}
}
