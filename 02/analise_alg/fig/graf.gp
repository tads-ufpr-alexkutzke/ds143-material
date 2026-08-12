# Crescimento comparado das classes de custo, usado no slide
# "Crescimento comparado das classes" da aula 02.
#
# Regerar:  gnuplot graf.gp
#
# A variavel independente se chama n no material da aula, entao a legenda
# usa n e nao o x padrao do gnuplot.

# Tela larga e fonte grande porque a figura e projetada em slide, ocupando
# cerca de 6 cm de altura.
set terminal pngcairo enhanced size 1500,850 font "sans,30" linewidth 3
set output "graf.png"

set xrange [0:40]
set yrange [0:1000]
# Uma amostra por inteiro, para os marcadores caírem sobre valores de n.
set samples 41

set key right top spacing 1.3
set tics nomirror
set border 3

set xlabel "n"

f1(x) = x
f2(x) = x*x
f3(x) = x*x*x
# gamma(n+1) = n! e aceita argumento real, ao contrario do operador !
f4(x) = gamma(x+1)

plot f1(x) with linespoints title "n", \
     f2(x) with linespoints title "n^2", \
     f3(x) with linespoints title "n^3", \
     f4(x) with linespoints title "n!"
