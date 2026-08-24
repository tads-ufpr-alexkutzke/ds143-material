file = arv_balanceadas.tex
pdflatex = lualatex

all: pdf

pdf:
	latexmk -pdflatex=${pdflatex} -pdf -shell-escape  ${file}

clean: dist
	latexmk -pdflatex=${pdflatex} -C ${file}

dist:
	rm -rf *.auxlock *.nav *.bbl *.snm *.vrb *.aux *.blg *.fdb_latexmk *.fls *.log *.out *.toc *-figure* _minted*
