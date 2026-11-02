// Heap de mínimo de inteiros, escrito à mão. É a estrutura que o pacote
// container/heap mantém por baixo da fila de prioridade do Prim.
//
// O heap é um slice visto como árvore binária completa: os filhos do índice i
// estão em 2i+1 e 2i+2, e o pai em (i-1)/2. Cada elemento é menor ou igual aos
// seus filhos, então o menor de todos está no índice 0.
//
// Execute com:  go run heap.go

package main

import "fmt"

type HeapMin struct {
	v []int
}

func (h *HeapMin) Len() int { return len(h.v) }

// Insere põe x no fim do slice e o faz subir até o lugar certo. O(log n).
func (h *HeapMin) Insere(x int) {
	h.v = append(h.v, x)
	h.sobe(len(h.v) - 1)
}

// RetiraMin devolve o menor elemento, que está na raiz. O último elemento vai
// para a raiz e desce até o lugar certo. O(log n).
func (h *HeapMin) RetiraMin() int {
	menor := h.v[0]
	ultimo := len(h.v) - 1
	h.v[0] = h.v[ultimo]
	h.v = h.v[:ultimo]
	if len(h.v) > 0 {
		h.desce(0)
	}
	return menor
}

// sobe troca o elemento do índice i com o pai enquanto ele for menor que o pai.
func (h *HeapMin) sobe(i int) {
	for i > 0 {
		pai := (i - 1) / 2
		if h.v[pai] <= h.v[i] {
			return
		}
		h.v[pai], h.v[i] = h.v[i], h.v[pai]
		i = pai
	}
}

// desce troca o elemento do índice i com o menor dos filhos enquanto esse
// filho for menor que ele.
func (h *HeapMin) desce(i int) {
	n := len(h.v)
	for {
		menor := i
		esq, dir := 2*i+1, 2*i+2
		if esq < n && h.v[esq] < h.v[menor] {
			menor = esq
		}
		if dir < n && h.v[dir] < h.v[menor] {
			menor = dir
		}
		if menor == i {
			return // nenhum filho é menor: o elemento está no lugar
		}
		h.v[menor], h.v[i] = h.v[i], h.v[menor]
		i = menor
	}
}

func main() {
	h := &HeapMin{}
	// Nesta ordem, nenhuma inserção provoca troca.
	for _, x := range []int{2, 9, 4, 12, 10, 6, 8} {
		h.Insere(x)
	}
	fmt.Println("heap inicial:  ", h.v)

	h.Insere(5)
	fmt.Println("insere 5:      ", h.v)

	fmt.Println("retira o menor:", h.RetiraMin())
	fmt.Println("heap:          ", h.v)

	// Retirar todos esvazia o heap em ordem crescente.
	fmt.Print("retira todos:   ")
	for h.Len() > 0 {
		fmt.Print(h.RetiraMin(), " ")
	}
	fmt.Println()
}
