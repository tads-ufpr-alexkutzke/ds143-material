// Package unionfind reúne as implementações de conectividade dinâmica
// estudadas na disciplina, todas expostas pela mesma interface.
//
// Este arquivo já está pronto e não deve ser alterado: ele define o contrato
// que as quatro implementações precisam cumprir.
package unionfind

// UF é o contrato comum às implementações de conectividade dinâmica sobre
// n objetos numerados de 0 a n-1.
//
// O modelo de custo usado na análise é o número de acessos ao slice interno,
// de leitura ou de escrita.
type UF interface {
	// Union conecta os objetos p e q. Se já estiverem na mesma componente,
	// não faz nada.
	Union(p, q int)

	// Find devolve o identificador da componente de p. Dois objetos estão
	// na mesma componente se, e somente se, têm o mesmo identificador.
	// O identificador é um valor interno da implementação: não há garantia
	// de que seja o mesmo em duas implementações diferentes, nem de que
	// permaneça o mesmo depois de novas chamadas a Union.
	Find(p int) int

	// Connected informa se p e q estão na mesma componente.
	Connected(p, q int) bool

	// Count devolve o número de componentes. Vale n logo após a criação da
	// estrutura e diminui uma unidade a cada Union que conecta componentes
	// distintas.
	Count() int
}
