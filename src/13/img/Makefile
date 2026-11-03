# Gera um SVG por página de figuras.pdf, para o roteiro em Markdown.
file = figuras

all: svg

pdf:
	latexmk -pdflatex=lualatex -pdf $(file).tex

svg: pdf
	pdftocairo -svg -f 1 -l 1 $(file).pdf grafo_gp.svg
	pdftocairo -svg -f 2 -l 2 $(file).pdf grafo_gp_bfs.svg
	pdftocairo -svg -f 3 -l 3 $(file).pdf grafo_gp_agm.svg
	pdftocairo -svg -f 4 -l 4 $(file).pdf corte.svg
	pdftocairo -svg -f 5 -l 5 $(file).pdf mapa_escolas.svg
	pdftocairo -svg -f 6 -l 6 $(file).pdf heap.svg
	pdftocairo -svg -f 7 -l 7 $(file).pdf heap_insere.svg
	pdftocairo -svg -f 8 -l 8 $(file).pdf heap_retira.svg
	pdftocairo -svg -f 9 -l 9 $(file).pdf prova_corte.svg

dist:
	latexmk -C $(file).tex
	rm -f $(file).pdf
