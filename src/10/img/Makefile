# Gera um SVG por página de figuras.pdf, para o roteiro em Markdown.
file = figuras

all: svg

pdf:
	latexmk -pdflatex=lualatex -pdf $(file).tex

svg: pdf
	pdftocairo -svg -f 1 -l 1 $(file).pdf swiss_grupo.svg

dist:
	latexmk -C $(file).tex
	rm -f $(file).pdf
