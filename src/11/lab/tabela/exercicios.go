package tabela

// Insere associa o valor v à chave k. Se a chave já está na tabela, só troca
// o valor, e o número de pares não muda.
//
// Se a chave é nova, ela vai para a primeira posição vazia a partir de h(k),
// andando uma posição por vez e dando a volta no fim do vetor, como na Busca.
// Depois de guardar uma chave nova, se o fator de carga passou de 0,7, a
// tabela cresce para proximoPrimo(2*m) posições. Em inteiros, a condição é
// 10*n > 7*m.
func (t *Tabela) Insere(k, v int) {
	panic("TODO: implementar Insere")
}

// redimensiona troca o vetor por um de m posições e reinsere todos os pares.
// A posição de cada chave muda, porque h(k) = k mod m depende de m.
//
// Sugestão: crie uma tabela com Nova(m), insira nela os pares das posições
// ocupadas e copie o resultado com *t = *nova, como no encadeamento.go da
// aula de tabelas de dispersão.
func (t *Tabela) redimensiona(m int) {
	panic("TODO: implementar redimensiona")
}

// Remove tira a chave k da tabela e devolve true, ou devolve false se ela
// não estava lá.
//
// Esvaziar a posição da chave não basta: uma chave que passou por ela ao ser
// inserida ficaria inalcançável, porque a Busca para na primeira posição
// vazia. A remoção percorre as posições seguintes ao buraco até achar uma
// vazia. Cada chave encontrada no caminho que pode ocupar o buraco desce para
// ele, e o buraco passa para a posição que ela deixou.
//
// A chave na posição j, com h(k) = h, pode descer para o buraco i quando o
// buraco está no caminho que a sondagem fez de h até j. Use a distancia de
// tabela.go para escrever essa condição sem tratar a volta no fim do vetor
// como caso à parte.
func (t *Tabela) Remove(k int) bool {
	panic("TODO: implementar Remove")
}
