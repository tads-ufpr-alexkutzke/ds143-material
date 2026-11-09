# Gera um SVG por página de figuras.pdf, para o roteiro em Markdown.
file = figuras

all: svg

pdf:
	latexmk -pdflatex=lualatex -pdf $(file).tex

svg: pdf
	pdftocairo -svg -f 1 -l 1 $(file).pdf arvore_caminhos.svg
	pdftocairo -svg -f 2 -l 2 $(file).pdf relaxacao.svg
	pdftocairo -svg -f 3 -l 3 $(file).pdf agm_caminhos.svg
	pdftocairo -svg -f 4 -l 4 $(file).pdf peso_negativo.svg

dist:
	latexmk -C $(file).tex
	rm -f $(file).pdf
