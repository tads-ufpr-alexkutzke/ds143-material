# Gera um SVG por página de figuras.pdf, para o roteiro em Markdown.
file = figuras

all: svg

pdf:
	latexmk -pdflatex=lualatex -pdf $(file).tex

svg: pdf
	pdftocairo -svg -f 1 -l 1 $(file).pdf grafo_g.svg
	pdftocairo -svg -f 2 -l 2 $(file).pdf grafo_g_bfs.svg
	pdftocairo -svg -f 3 -l 3 $(file).pdf grafo_g_dfs.svg
	pdftocairo -svg -f 4 -l 4 $(file).pdf componentes.svg
	pdftocairo -svg -f 5 -l 5 $(file).pdf pacotes.svg
	pdftocairo -svg -f 6 -l 6 $(file).pdf pacotes_ciclo.svg
	pdftocairo -svg -f 7 -l 7 $(file).pdf pacotes_ordem.svg
	pdftocairo -svg -f 8 -l 8 $(file).pdf fortemente_conexas.svg

dist:
	latexmk -C $(file).tex
	rm -f $(file).pdf
