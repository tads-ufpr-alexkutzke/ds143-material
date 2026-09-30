# Gera um SVG por página de figuras.pdf, para o roteiro em Markdown.
file = figuras

all: svg rn slides

# Figuras de árvore rubro-negra desenhadas a partir da própria inserção.
rn:
	go run gera_rn.go

pdf:
	latexmk -pdflatex=lualatex -pdf $(file).tex

svg: pdf
	pdftocairo -svg -f 1 -l 1 $(file).pdf rn_exemplo.svg
	pdftocairo -svg -f 2 -l 2 $(file).pdf rn_23.svg
	pdftocairo -svg -f 3 -l 3 $(file).pdf rn_exercicio1.svg
	pdftocairo -svg -f 4 -l 4 $(file).pdf b_divisao.svg
	pdftocairo -svg -f 5 -l 5 $(file).pdf b_propagacao.svg
	pdftocairo -svg -f 6 -l 6 $(file).pdf bmais.svg
	pdftocairo -svg -f 7 -l 7 $(file).pdf trie.svg
	pdftocairo -svg -f 8 -l 8 $(file).pdf arv23_a.svg
	pdftocairo -svg -f 9 -l 9 $(file).pdf arv23_b.svg
	pdftocairo -svg -f 10 -l 10 $(file).pdf arv23_c.svg
	pdftocairo -svg -f 11 -l 11 $(file).pdf arv23_d.svg

# Recortes em PDF para os slides, que não leem SVG: um painel da inserção de P
# por arquivo, sem a legenda (o slide traz a legenda), e a árvore B+ inteira.
# Os limites verticais de cada painel vêm das coordenadas do rn_insercao_P.svg.
PAINEIS = 1:38:144 2:230:144 3:422:182 4:652:144 5:844:144 6:1036:144

slides:
	mkdir -p slides
	for p in $(PAINEIS); do \
	  n=$${p%%:*}; r=$${p#*:}; y=$${r%%:*}; h=$${r#*:}; \
	  sed 's/width="440" height="1202" viewBox="0 0 440 1202"/width="266" height="'$$h'" viewBox="4 '$$y' 266 '$$h'"/' \
	    rn_insercao_P.svg | rsvg-convert -f pdf -o slides/rn_insercao_P_$$n.pdf; \
	done
	rsvg-convert -f pdf -o slides/bmais.pdf bmais.svg

dist:
	latexmk -C $(file).tex
	rm -f $(file).pdf
